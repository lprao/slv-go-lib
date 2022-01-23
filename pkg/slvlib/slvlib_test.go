// Copyright (c) Puneeth Rao Lokapalli 2022. All rights reserved.
// Licensed under the Apache license. See LICENSE file in the project root for full license information.

package slvlib

import (
	"os"
	"reflect"
	"testing"

	"github.com/lprao/slv-go-lib/internal/pkg/transport"
	slvpb "github.com/lprao/slv-proto"
)

func Test_slvClient_execOp(t *testing.T) {
	type fields struct {
		grpc *transport.Grpc
	}
	type args struct {
		slvVar      *slvpb.SlvVar
		op          slvpb.Operation
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *slvpb.SlvVar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	os.Setenv("SLV_SVC_ENDPOINT", "slv-svc.local")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &slvClient{
				grpc: tt.fields.grpc,
			}
			got, err := s.execOp(tt.args.slvVar, tt.args.op, tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("slvClient.execOp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slvClient.execOp() = %v, want %v", got, tt.want)
			}
		})
	}
}
