package gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/mux"

	"github.com/bnb-chain/greenfield-storage-provider/model"
	chainclient "github.com/bnb-chain/greenfield-storage-provider/pkg/greenfield"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/lifecycle"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/log"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/metrics"
	localhttp "github.com/bnb-chain/greenfield-storage-provider/pkg/middleware/http"
	authclient "github.com/bnb-chain/greenfield-storage-provider/service/auth/client"
	challengeclient "github.com/bnb-chain/greenfield-storage-provider/service/challenge/client"
	downloaderclient "github.com/bnb-chain/greenfield-storage-provider/service/downloader/client"
	metadataclient "github.com/bnb-chain/greenfield-storage-provider/service/metadata/client"
	receiverclient "github.com/bnb-chain/greenfield-storage-provider/service/receiver/client"
	signerclient "github.com/bnb-chain/greenfield-storage-provider/service/signer/client"
	uploaderclient "github.com/bnb-chain/greenfield-storage-provider/service/uploader/client"
)

var _ lifecycle.Service = &Gateway{}

// Gateway is the primary entry point of SP
type Gateway struct {
	config     *GatewayConfig
	running    atomic.Value
	httpServer *http.Server

	// chain is the required component, used to check authorization
	chain *chainclient.Greenfield

	// the below components are optional according to the config
	uploader   *uploaderclient.UploaderClient
	downloader *downloaderclient.DownloaderClient
	challenge  *challengeclient.ChallengeClient
	receiver   *receiverclient.ReceiverClient
	signer     *signerclient.SignerClient
	metadata   *metadataclient.MetadataClient
	auth       *authclient.AuthClient
}

// NewGatewayService return the gateway instance
func NewGatewayService(cfg *GatewayConfig) (*Gateway, error) {
	var (
		gateway *Gateway
		err     error
	)

	gateway = &Gateway{
		config: cfg,
	}
	if gateway.chain, err = chainclient.NewGreenfield(cfg.ChainConfig); err != nil {
		log.Errorw("failed to create chain client", "error", err)
		return nil, err
	}
	if cfg.UploaderServiceAddress != "" {
		if gateway.uploader, err = uploaderclient.NewUploaderClient(cfg.UploaderServiceAddress); err != nil {
			log.Errorw("failed to create uploader client", "error", err)
			return nil, err
		}
	}
	if cfg.DownloaderServiceAddress != "" {
		if gateway.downloader, err = downloaderclient.NewDownloaderClient(cfg.DownloaderServiceAddress); err != nil {
			log.Errorw("failed to create downloader client", "error", err)
			return nil, err
		}
	}
	if cfg.ChallengeServiceAddress != "" {
		if gateway.challenge, err = challengeclient.NewChallengeClient(cfg.ChallengeServiceAddress); err != nil {
			log.Errorw("failed to create challenge client", "error", err)
			return nil, err
		}
	}
	if cfg.ReceiverServiceAddress != "" {
		if gateway.receiver, err = receiverclient.NewReceiverClient(cfg.ReceiverServiceAddress); err != nil {
			log.Errorw("failed to create receiver client", "error", err)
			return nil, err
		}
	}
	if cfg.SignerServiceAddress != "" {
		if gateway.signer, err = signerclient.NewSignerClient(cfg.SignerServiceAddress); err != nil {
			log.Errorw("failed to create signer client", "error", err)
			return nil, err
		}
	}
	if cfg.MetadataServiceAddress != "" {
		if gateway.metadata, err = metadataclient.NewMetadataClient(cfg.MetadataServiceAddress); err != nil {
			log.Errorw("failed to create metadata client", "error", err)
			return nil, err
		}
	}
	if cfg.AuthServiceAddress != "" {
		if gateway.auth, err = authclient.NewAuthClient(cfg.AuthServiceAddress); err != nil {
			log.Errorw("failed to create auth client", "error", err)
			return nil, err
		}
	}

	return gateway, nil
}

// Name return the descriptions of gateway service
func (gateway *Gateway) Name() string {
	return model.GatewayService
}

// Start gateway service
func (gateway *Gateway) Start(ctx context.Context) error {
	if gateway.running.Swap(true) == true {
		return errors.New("gateway has started")
	}
	go gateway.serve()
	return nil
}

// Serve starts http server.
func (gateway *Gateway) serve() {
	router := mux.NewRouter().SkipClean(true)
	if metrics.GetMetrics().Enabled() {
		router.Use(metrics.DefaultHTTPServerMetrics.InstrumentationHandler)
	}
	if err := localhttp.NewAPILimiter(gateway.config.APILimiterCfg); err != nil {
		log.Errorw("failed to new api limiter", "err", err)
		return
	}
	if gateway.config.BandwidthLimitCfg.Enable {
		localhttp.NewBandwidthLimiter(gateway.config.BandwidthLimitCfg.R, gateway.config.BandwidthLimitCfg.B)
	}
	gateway.registerHandler(router)
	server := &http.Server{
		Addr:    gateway.config.HTTPAddress,
		Handler: router,
	}
	gateway.httpServer = server
	if err := server.ListenAndServe(); err != nil {
		log.Errorw("failed to listen", "error", err)
		return
	}
}

// Stop gateway service
func (gateway *Gateway) Stop(ctx context.Context) error {
	if gateway.running.Swap(false) == false {
		return errors.New("gateway has stopped")
	}
	var errs []error
	if err := gateway.httpServer.Shutdown(ctx); err != nil {
		errs = append(errs, err)
	}
	if errs != nil {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
