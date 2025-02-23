package model

import "strings"

// define storage provider include service
var (
	// GatewayService defines the name of gateway service
	GatewayService = strings.ToLower("Gateway")
	// UploaderService defines the name of uploader service
	UploaderService = strings.ToLower("Uploader")
	// DownloaderService defines the name of downloader service
	DownloaderService = strings.ToLower("Downloader")
	// ChallengeService defines the name of challenge service
	ChallengeService = strings.ToLower("Challenge")
	// TaskNodeService defines the name of task node service
	TaskNodeService = strings.ToLower("TaskNode")
	// ReceiverService defines the name of receiver service
	ReceiverService = strings.ToLower("Receiver")
	// SignerService defines the name of signer service
	SignerService = strings.ToLower("Signer")
	// MetadataService defines the name of metadata service
	MetadataService = strings.ToLower("Metadata")
	// BlockSyncerService defines the name of block sync service
	BlockSyncerService = strings.ToLower("BlockSyncer")
	// BlockSyncerServiceBackup defines the name of block sync service
	BlockSyncerServiceBackup = strings.ToLower("BlockSyncerBackup")
	// ManagerService defines the name of manager service
	ManagerService = strings.ToLower("Manager")
	// MetricsService defines the name of metrics service
	MetricsService = strings.ToLower("Metrics")
	// P2PService defines the name of p2p service
	P2PService = strings.ToLower("p2p")
	// AuthService defines the name of auth service
	AuthService = strings.ToLower("auth")
	// PProfService defines the name of pprof service
	PProfService = strings.ToLower("pprof")
	// StopServingService defines the name of stop serving service
	StopServingService = strings.ToLower("StopServing")
)

// SpServiceDesc defines the service description in storage provider
var SpServiceDesc = map[string]string{
	GatewayService:     "Receives the sdk request",
	UploaderService:    "Uploads object payload to greenfield",
	DownloaderService:  "Downloads object from the backend and statistical read traffic",
	ChallengeService:   "Provides the ability to query the integrity hash and piece data",
	TaskNodeService:    "Executes background task",
	ReceiverService:    "Receives data pieces of an object from other storage provider and store",
	SignerService:      "Sign the transaction and broadcast to chain",
	MetadataService:    "Provides the ability to query meta data",
	BlockSyncerService: "Syncs block data to db",
	P2PService:         "Communicates with SPs on p2p protocol",
	AuthService:        "Handles off-chain-auth requests",
	StopServingService: "Discontinue buckets for greenfield testnet",
}

// define storage provider service default listening address
const (
	// GatewayHTTPAddress default HTTP address of gateway
	GatewayHTTPAddress = "localhost:9033"
	// UploaderGRPCAddress default gRPC address of uploader
	UploaderGRPCAddress = "localhost:9133"
	// DownloaderGRPCAddress default gRPC address of downloader
	DownloaderGRPCAddress = "localhost:9233"
	// ChallengeGRPCAddress default gRPC address of challenge
	ChallengeGRPCAddress = "localhost:9333"
	// TaskNodeGRPCAddress default gRPC address of task node
	TaskNodeGRPCAddress = "localhost:9433"
	// ReceiverGRPCAddress default gRPC address of receiver
	ReceiverGRPCAddress = "localhost:9533"
	// SignerGRPCAddress default gRPC address of signer
	SignerGRPCAddress = "localhost:9633"
	// MetadataGRPCAddress default gRPC address of meta data service
	MetadataGRPCAddress = "localhost:9733"
	// MetricsHTTPAddress default HTTP address of metrics service
	MetricsHTTPAddress = "localhost:24036"
	// P2PGRPCAddress default gRPC address of p2p service
	P2PGRPCAddress = "localhost:9833"
	// P2PListenAddress default p2p protocol listen address of p2p node
	P2PListenAddress = "127.0.0.1:9933"
	// AuthGRPCAddress default gRPC address of auth service
	AuthGRPCAddress = "localhost:8933"
	// PProfHTTPAddress default HTTP address of pprof service
	PProfHTTPAddress = "localhost:25341"
)

// define greenfield chain default address
const (
	// GreenfieldAddress default greenfield chain address
	GreenfieldAddress = "localhost:9090"
	// TendermintAddress default Tendermint address
	TendermintAddress = "http://localhost:26750"
	// GreenfieldChainID default greenfield chainID
	GreenfieldChainID = "greenfield_9000-1741"
	// WhiteListCIDR default whitelist CIDR
	WhiteListCIDR = "0.0.0.0/0"
)

// environment constants
const (
	// SpDBUser defines env variable name for sp db user name
	SpDBUser = "SP_DB_USER"
	// SpDBPasswd defines env variable name for sp db user passwd
	SpDBPasswd = "SP_DB_PASSWORD"
	// SpDBAddress defines env variable name for sp db address
	SpDBAddress = "SP_DB_ADDRESS"
	// SpDBDataBase defines env variable name for sp db database
	SpDBDataBase = "SP_DB_DATABASE"

	// BsDBUser defines env variable name for block syncer db user name
	BsDBUser = "BS_DB_USER"
	// BsDBPasswd defines env variable name for block syncer db user passwd
	BsDBPasswd = "BS_DB_PASSWORD"
	// BsDBAddress defines env variable name for block syncer db address
	BsDBAddress = "BS_DB_ADDRESS"
	// BsDBDataBase defines env variable name for block syncer db database
	BsDBDataBase = "BS_DB_DATABASE"
	// BsDBSwitchedUser defines env variable name for switched block syncer db user name
	BsDBSwitchedUser = "BS_DB_SWITCHED_USER"
	// BsDBSwitchedPasswd defines env variable name for switched block syncer db user passwd
	BsDBSwitchedPasswd = "BS_DB_SWITCHED_PASSWORD"
	// BsDBSwitchedAddress defines env variable name for switched block syncer db address
	BsDBSwitchedAddress = "BS_DB_SWITCHED_ADDRESS"
	// BsDBSwitchedDataBase defines env variable name for switched block syncer db database
	BsDBSwitchedDataBase = "BS_DB_SWITCHED_DATABASE"

	// SpOperatorAddress defines env variable name for sp operator address
	SpOperatorAddress = "greenfield-storage-provider"
	// SpSignerAPIKey defines env variable for signer api key
	SpSignerAPIKey = "SIGNER_API_KEY"
	// SpOperatorPrivKey defines env variable name for sp operator priv key
	SpOperatorPrivKey = "SIGNER_OPERATOR_PRIV_KEY"
	// SpFundingPrivKey defines env variable name for sp funding priv key
	SpFundingPrivKey = "SIGNER_FUNDING_PRIV_KEY"
	// SpApprovalPrivKey defines env variable name for sp approval priv key
	SpApprovalPrivKey = "SIGNER_APPROVAL_PRIV_KEY"
	// SpSealPrivKey defines env variable name for sp seal priv key
	SpSealPrivKey = "SIGNER_SEAL_PRIV_KEY"
	// SpGcPrivKey defines env variable name for sp gc priv key
	SpGcPrivKey = "SIGNER_GC_PRIV_KEY"
	// DsnBlockSyncer defines env variable name for block syncer dsn
	DsnBlockSyncer = "BLOCK_SYNCER_DSN"
	// DsnBlockSyncerSwitched defines env variable name for block syncer backup dsn
	DsnBlockSyncerSwitched = "BLOCK_SYNCER_DSN_SWITCHED"
	// P2PPrivateKey defines env variable for p2p protocol private key
	P2PPrivateKey = "P2P_PRIVATE_KEY"
)

// SQLDB default configuration parmas
const (
	// DefaultConnMaxLifetime defines the default max liveness time of connection
	DefaultConnMaxLifetime = 60
	// DefaultConnMaxIdleTime defines the default max idle time of connection
	DefaultConnMaxIdleTime = 30
	// DefaultMaxIdleConns defines the default max number of idle connections
	DefaultMaxIdleConns = 16
	// DefaultMaxOpenConns defines the default max number of open connections
	DefaultMaxOpenConns = 32
)

// define all kinds of http constants
const (
	// ContentTypeHeader is used to indicate the media type of the resource
	ContentTypeHeader = "Content-Type"
	// ContentLengthHeader indicates the size of the message body, in bytes
	ContentLengthHeader = "Content-Length"
	// ETagHeader is an MD5 digest of the object data
	ETagHeader = "ETag"
	// RangeHeader asks the server to send only a portion of an HTTP message back to a client
	RangeHeader = "Range"
	// ContentRangeHeader response HTTP header indicates where in a full body message a partial message belongs
	ContentRangeHeader = "Content-Range"
	// OctetStream is used to indicate the binary files
	OctetStream = "application/octet-stream"
	// ContentTypeJSONHeaderValue is used to indicate json
	ContentTypeJSONHeaderValue = "application/json"
	// ContentTypeXMLHeaderValue is used to indicate xml
	ContentTypeXMLHeaderValue = "application/xml"
	// ContentDispositionHeader is used to indicate the media disposition of the resource
	ContentDispositionHeader = "Content-Disposition"
	// ContentDispositionAttachmentValue is used to indicate attachment
	ContentDispositionAttachmentValue = "attachment"
	// ContentDispositionInlineValue is used to indicate inline
	ContentDispositionInlineValue = "inline"

	// SignAlgorithm uses secp256k1 with the ECDSA algorithm
	SignAlgorithm = "ECDSA-secp256k1"
	// SignedMsg is the request hash
	SignedMsg = "SignedMsg"
	// Signature is the request signature
	Signature = "Signature"
	// SignTypeV1 is an authentication algorithm, which is used by dapps
	SignTypeV1 = "authTypeV1"
	// SignTypeV2 is an authentication algorithm, which is used by metamask
	SignTypeV2 = "authTypeV2"

	SignTypeOffChain   = "OffChainAuth" // sign type - off-chain-auth
	SignTypePersonal   = "PersonalSign" // sign type -  PersonalSign
	SignAlgorithmEddsa = "EDDSA"

	// GetApprovalPath defines get-approval path style suffix
	GetApprovalPath = "/greenfield/admin/v1/get-approval"
	// ActionQuery defines get-approval's type, currently include create bucket and create object
	ActionQuery = "action"
	// UploadProgressQuery defines upload progress query, which is used to route request
	UploadProgressQuery = "upload-progress"
	// GetBucketReadQuotaQuery defines bucket read quota query, which is used to route request
	GetBucketReadQuotaQuery = "read-quota"
	// GetBucketReadQuotaMonthQuery defines bucket read quota query month
	GetBucketReadQuotaMonthQuery = "year-month"
	// ListBucketReadRecordQuery defines list bucket read record query, which is used to route request
	ListBucketReadRecordQuery = "list-read-record"
	// ListBucketReadRecordMaxRecordsQuery defines list read record max num
	ListBucketReadRecordMaxRecordsQuery = "max-records"
	// ListObjectsMaxKeysQuery defines the maximum number of keys returned to the response
	ListObjectsMaxKeysQuery = "max-keys"
	// ListObjectsStartAfterQuery defines where you want to start listing from
	ListObjectsStartAfterQuery = "start-after"
	// ListObjectsContinuationTokenQuery indicates that the list is being continued on this bucket with a token
	ListObjectsContinuationTokenQuery = "continuation-token"
	// ListObjectsDelimiterQuery defines a character you use to group keys
	ListObjectsDelimiterQuery = "delimiter"
	// ListObjectsPrefixQuery defines limits the response to keys that begin with the specified prefix
	ListObjectsPrefixQuery = "prefix"
	// GetBucketMetaQuery defines get bucket metadata query, which is used to route request
	GetBucketMetaQuery = "bucket-meta"
	// GetObjectMetaQuery defines get object metadata query, which is used to route request
	GetObjectMetaQuery = "object-meta"
	// StartTimestampUs defines start timestamp in microsecond, which is used by list read record, [start_ts,end_ts)
	StartTimestampUs = "start-timestamp"
	// EndTimestampUs defines end timestamp in microsecond, which is used by list read record, [start_ts,end_ts)
	EndTimestampUs = "end-timestamp"
	// ChallengePath defines challenge path style suffix
	ChallengePath = "/greenfield/admin/v1/challenge"
	// ReplicateObjectPiecePath defines replicate-object path style
	ReplicateObjectPiecePath = "/greenfield/receiver/v1/replicate-piece"
	// AuthRequestNoncePath defines path to request auth nonce
	AuthRequestNoncePath = "/auth/request_nonce"
	// AuthUpdateKeyPath defines path to update user public key
	AuthUpdateKeyPath = "/auth/update_key"
	// GnfdRequestIDHeader defines trace-id, trace request in sp
	GnfdRequestIDHeader = "X-Gnfd-Request-ID"
	// GnfdAuthorizationHeader defines authorization, verify signature and check authorization
	GnfdAuthorizationHeader = "Authorization"
	// GnfdObjectIDHeader defines object id
	GnfdObjectIDHeader = "X-Gnfd-Object-ID"
	// GnfdPieceIndexHeader defines piece idx, which is used by challenge
	GnfdPieceIndexHeader = "X-Gnfd-Piece-Index"
	// GnfdRedundancyIndexHeader defines redundancy idx, which is used by challenge and receiver
	GnfdRedundancyIndexHeader = "X-Gnfd-Redundancy-Index"
	// GnfdIntegrityHashHeader defines integrity hash, which is used by challenge and receiver
	GnfdIntegrityHashHeader = "X-Gnfd-Integrity-Hash"
	// GnfdPieceHashHeader defines piece hash list, which is used by challenge
	GnfdPieceHashHeader = "X-Gnfd-Piece-Hash"
	// GnfdUnsignedApprovalMsgHeader defines unsigned msg, which is used by get-approval
	GnfdUnsignedApprovalMsgHeader = "X-Gnfd-Unsigned-Msg"
	// GnfdSignedApprovalMsgHeader defines signed msg, which is used by get-approval
	GnfdSignedApprovalMsgHeader = "X-Gnfd-Signed-Msg"
	// GnfdPieceSizeHeader defines piece size, which is used to split by receiver
	GnfdPieceSizeHeader = "X-Gnfd-Piece-Size"
	// GnfdReplicateApproval defines SP approval that allow to replicate piece data, which is used by receiver
	GnfdReplicateApproval = "X-Gnfd-Replicate-Approval"
	// GnfdIntegrityHashSignatureHeader defines integrity hash signature, which is used by receiver
	GnfdIntegrityHashSignatureHeader = "X-Gnfd-Integrity-Hash-Signature"
	// GnfdUserAddressHeader defines the user address
	GnfdUserAddressHeader = "X-Gnfd-User-Address"
	// GnfdResponseXMLVersion defines the response xml version
	GnfdResponseXMLVersion = "1.0"

	// off-chain-auth headers

	// GnfdOffChainAuthAppDomainHeader defines the app domain from where user is trying to do the EDDSA auth interactions
	GnfdOffChainAuthAppDomainHeader = "X-Gnfd-App-Domain"
	// GnfdOffChainAuthAppRegNonceHeader defines nonce for which user is trying to register his/her EDDSA public key
	GnfdOffChainAuthAppRegNonceHeader = "X-Gnfd-App-Reg-Nonce"
	// GnfdOffChainAuthAppRegPublicKeyHeader defines the EDDSA public key for which user is trying to register
	GnfdOffChainAuthAppRegPublicKeyHeader = "X-Gnfd-App-Reg-Public-Key"
	// GnfdOffChainAuthAppRegExpiryDateHeader defines the Expiry-Date is the ISO 8601 datetime string (e.g. 2021-09-30T16:25:24Z), used to register the EDDSA public key
	GnfdOffChainAuthAppRegExpiryDateHeader = "X-Gnfd-App-Reg-Expiry-Date"
)

// define all kinds of size
const (
	// LruCacheLimit defines maximum number of cached items in service trace queue
	LruCacheLimit = 8192
	// MaxCallMsgSize defines gPRC max send or receive msg size
	MaxCallMsgSize = 25 * 1024 * 1024
	// MaxRetryCount defines getting the latest height from the RPC client max retry count
	MaxRetryCount = 50
	// DefaultBlockHeightDiff defines default block height diff of main and backup service
	DefaultBlockHeightDiff = 100
	// DefaultCheckDiffPeriod defines check interval of block height diff
	DefaultCheckDiffPeriod = 1
	// DefaultPingPeriod defines p2p node ping period
	DefaultPingPeriod = 1
	// DefaultSpFreeReadQuotaSize defines sp bucket's default free quota size, the SP can modify it by itself
	DefaultSpFreeReadQuotaSize = 10 * 1024 * 1024 * 1024
	// DefaultStreamBufSize defines gateway stream forward payload buf size
	DefaultStreamBufSize = 64 * 1024
	// DefaultTimeoutHeight defines approval timeout height
	DefaultTimeoutHeight = 100
	// DefaultPartitionSize defines partition size
	DefaultPartitionSize = 10_000
	// DefaultMaxListLimit defines maximum number of the list request
	DefaultMaxListLimit = 1000
)

// PieceType defines the object payload data type
type PieceType int32

const (
	// SegmentPieceType defines the data type of segment piece
	SegmentPieceType PieceType = 0
	// ECPieceType defines the data type of ec piece
	ECPieceType PieceType = 1
)
