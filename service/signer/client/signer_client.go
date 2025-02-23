package client

import (
	"context"

	storagetypes "github.com/bnb-chain/greenfield/x/storage/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bnb-chain/greenfield-storage-provider/pkg/log"
	"github.com/bnb-chain/greenfield-storage-provider/pkg/metrics"
	p2ptpyes "github.com/bnb-chain/greenfield-storage-provider/pkg/p2p/types"
	"github.com/bnb-chain/greenfield-storage-provider/service/signer/types"
	utilgrpc "github.com/bnb-chain/greenfield-storage-provider/util/grpc"
)

type SignerClient struct {
	address string
	signer  types.SignerServiceClient
	conn    *grpc.ClientConn
}

const signerRPCServiceName = "service.signer.types.SignerService"

func NewSignerClient(address string) (*SignerClient, error) {
	options := []grpc.DialOption{}
	if metrics.GetMetrics().Enabled() {
		options = append(options, utilgrpc.GetDefaultClientInterceptor()...)
	}
	retryOption, err := utilgrpc.GetDefaultGRPCRetryPolicy(signerRPCServiceName)
	if err != nil {
		log.Errorw("failed to get signer client retry option", "error", err)
		return nil, err
	}
	options = append(options, retryOption)
	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(context.Background(), address, options...)
	if err != nil {
		log.Errorw("failed to dial signer", "error", err)
		return nil, err
	}
	client := &SignerClient{
		address: address,
		conn:    conn,
		signer:  types.NewSignerServiceClient(conn),
	}
	return client, nil
}

func (client *SignerClient) Close() error {
	return client.conn.Close()
}

func (client *SignerClient) SignBucketApproval(ctx context.Context,
	msg *storagetypes.MsgCreateBucket, opts ...grpc.CallOption) ([]byte, error) {
	resp, err := client.signer.SignBucketApproval(ctx,
		&types.SignBucketApprovalRequest{CreateBucketMsg: msg}, opts...)
	if err != nil {
		return nil, err
	}
	return resp.Signature, nil
}

func (client *SignerClient) VerifyBucketApproval(ctx context.Context,
	msg *storagetypes.MsgCreateBucket, opts ...grpc.CallOption) (bool, error) {
	resp, err := client.signer.VerifyBucketApproval(ctx,
		&types.VerifyBucketApprovalRequest{CreateBucketMsg: msg}, opts...)
	if err != nil {
		return false, err
	}
	return resp.GetResult(), nil
}

func (client *SignerClient) SignObjectApproval(ctx context.Context,
	msg *storagetypes.MsgCreateObject, opts ...grpc.CallOption) ([]byte, error) {
	resp, err := client.signer.SignObjectApproval(ctx,
		&types.SignObjectApprovalRequest{CreateObjectMsg: msg}, opts...)
	if err != nil {
		return nil, err
	}
	return resp.Signature, nil
}

func (client *SignerClient) VerifyObjectApproval(ctx context.Context,
	msg *storagetypes.MsgCreateObject, opts ...grpc.CallOption) (bool, error) {
	resp, err := client.signer.VerifyObjectApproval(ctx,
		&types.VerifyObjectApprovalRequest{CreateObjectMsg: msg}, opts...)
	if err != nil {
		return false, err
	}
	return resp.GetResult(), nil
}

func (client *SignerClient) SignIntegrityHash(ctx context.Context,
	objectID uint64,
	checksum [][]byte, opts ...grpc.CallOption,
) ([]byte, []byte, error) {
	resp, err := client.signer.SignIntegrityHash(ctx,
		&types.SignIntegrityHashRequest{Data: checksum, ObjectId: objectID}, opts...)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	return resp.GetIntegrityHash(), resp.GetSignature(), nil
}

func (client *SignerClient) SealObjectOnChain(ctx context.Context,
	sealObject *storagetypes.MsgSealObject, opts ...grpc.CallOption) ([]byte, error) {
	resp, err := client.signer.SealObjectOnChain(ctx,
		&types.SealObjectOnChainRequest{SealObject: sealObject}, opts...)
	if err != nil {
		return []byte{}, err
	}
	return resp.GetTxHash(), nil
}

func (client *SignerClient) DiscontinueBucketOnChain(ctx context.Context,
	discontinueBucket *storagetypes.MsgDiscontinueBucket, opts ...grpc.CallOption) ([]byte, error) {
	resp, err := client.signer.DiscontinueBucketOnChain(ctx,
		&types.DiscontinueBucketOnChainRequest{DiscontinueBucket: discontinueBucket}, opts...)
	if err != nil {
		return []byte{}, err
	}
	return resp.GetTxHash(), nil
}

func (client *SignerClient) SignPingMsg(ctx context.Context, ping *p2ptpyes.Ping, opts ...grpc.CallOption) (*p2ptpyes.Ping, error) {
	req := &types.SignPingMsgRequest{
		Ping: ping,
	}
	resp, err := client.signer.SignPingMsg(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetPing(), nil
}

func (client *SignerClient) SignPongMsg(ctx context.Context, pong *p2ptpyes.Pong, opts ...grpc.CallOption) (*p2ptpyes.Pong, error) {
	req := &types.SignPongMsgRequest{
		Pong: pong,
	}
	resp, err := client.signer.SignPongMsg(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetPong(), nil
}

func (client *SignerClient) SignReplicateApprovalReqMsg(ctx context.Context, approval *p2ptpyes.GetApprovalRequest, opts ...grpc.CallOption) (*p2ptpyes.GetApprovalRequest, error) {
	req := &types.SignReplicateApprovalReqMsgRequest{
		Approval: approval,
	}
	resp, err := client.signer.SignReplicateApprovalReqMsg(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetApproval(), nil
}

func (client *SignerClient) SignReplicateApprovalRspMsg(ctx context.Context, approval *p2ptpyes.GetApprovalResponse, opts ...grpc.CallOption) (*p2ptpyes.GetApprovalResponse, error) {
	req := &types.SignReplicateApprovalRspMsgRequest{
		Approval: approval,
	}
	resp, err := client.signer.SignReplicateApprovalRspMsg(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return resp.GetApproval(), nil
}
