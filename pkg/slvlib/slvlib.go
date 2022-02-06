// Copyright (c) Puneeth Rao Lokapalli 2022. All rights reserved.
// Licensed under the Apache license. See LICENSE file in the project root for full license information.

package slvlib

import (
	"errors"
	"os"

	"github.com/lprao/slv-go-lib/internal/pkg/transport"
	"github.com/lprao/slv-go-lib/pkg/logger"
	slvpb "github.com/lprao/slv-proto"
)

// slvClient struct to manage the client side artifacts.
type slvClient struct {
	// Grpc pkg object
	grpc *transport.Grpc
}

var (
	client *slvClient
	log    logger.Log
)

const (
	SLV_SVC_CA_CERT = "certs/ca/ca.crt"
)

// Send the user requested operation to slv-svc
func (s *slvClient) execOp(slvVar *slvpb.SlvVar, op slvpb.Operation, accessToken string) (*slvpb.SlvVar, error) {
	req := slvpb.ExecOpReq{
		Operation: op,
		Var:       slvVar,
	}

	resp, err := s.grpc.DoGrpc(req, accessToken)
	if err != nil {
		return &slvpb.SlvVar{}, nil
	}

	return resp.Var, nil
}

// Initial setup and validations.
func init() {
	log.InitLogger()

	ep := os.Getenv("SLV_SVC_ENDPOINT")
	if ep == "" {
		log.Fatalf("SLV_SVC_ENDPOINT environment variable not set")
		return
	}

	var insecure bool = true
	_, err := os.Stat(SLV_SVC_CA_CERT)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		log.Warnf("CA cert for slv-svc not provided, using insecure mode")
		insecure = false
	}

	client.grpc = transport.NewGrpcClient(ep, insecure, SLV_SVC_CA_CERT)
}
