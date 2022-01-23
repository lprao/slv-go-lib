package transport

import (
	"context"
	"fmt"
	"time"

	"github.com/lprao/slv-go-lib/pkg/logger"
	slvpb "github.com/lprao/slv-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Grpc provider
type Grpc struct {
	// grpc connection object
	conn *grpc.ClientConn

	// The GRPC endpoint(or a Load Balancer) that will serve requests
	slvSvcEndpoint string

	// CA cert for GRPC endpoint or the loadbalancer is not provided.
	insecure bool

	// CA cert location
	caCertLoc string
}

const (
	GRPC_TIME_OUT = (10 * time.Second)
)

// Get a new Grpc client
func NewGrpcClient(ep string, insecure bool, loc string) *Grpc {
	return &Grpc{
		slvSvcEndpoint: ep,
		insecure:       insecure,
		caCertLoc:      loc,
	}
}

// Connect to the Grpc endpoint
func (g *Grpc) connect(ep string, insecure bool, loc string) (*grpc.ClientConn, error) {
	if g.conn != nil {
		return g.conn, nil
	}

	if !insecure {
		creds, err := credentials.NewClientTLSFromFile(loc, "")
		if err != nil {
			return nil, fmt.Errorf(ErrInvalidCaCert, err.Error())
		}

		g.conn, err = grpc.Dial(ep, grpc.WithTransportCredentials(creds))
		if err != nil {
			return nil, fmt.Errorf(ErrGrpcDialFailed, ep, err.Error())
		}
		return g.conn, nil
	}

	var err error
	g.conn, err = grpc.Dial(ep)
	if err != nil {
		return nil, fmt.Errorf(ErrGrpcDialFailed, ep, err.Error())
	}

	return g.conn, nil
}

// Get the Grpc client object
func (g *Grpc) getSlvSvcClient(ep string, insecure bool, loc string) (slvpb.LvSvcClient, error) {
	c, err := g.connect(ep, insecure, loc)
	if err != nil {
		return nil, fmt.Errorf(ErrGrpcDialFailed, ep, err.Error())
	}

	return slvpb.NewLvSvcClient(c), nil
}

// Helper routine to send the grpc request and return the response from slv-svc
func (g *Grpc) DoGrpc(req slvpb.ExecOpReq) (slvpb.ExecOpResp, error) {
	c, err := g.getSlvSvcClient(g.slvSvcEndpoint, g.insecure, g.caCertLoc)
	if err != nil {
		return slvpb.ExecOpResp{}, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), GRPC_TIME_OUT)
	defer cancel()
	resp, err := c.Op(ctx, &req)
	if err != nil {
		return slvpb.ExecOpResp{}, fmt.Errorf("execOp failed - %s", err.Error())
	}

	return *resp, nil
}

var (
	log logger.Log
)

// Initial setup.
func init() {
	log.InitLogger()
}
