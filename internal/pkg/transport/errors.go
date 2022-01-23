package transport

const (
	ErrInvalidCaCert  = "could not load tls cert; err: [%s]"
	ErrGrpcDialFailed = "failed to connect to grpc endpoint: [%s] err: [%s]"
	ErrExecOpFailed   = "execOp failed; err: [%s]"
)
