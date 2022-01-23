package transport

import (
	"reflect"
	"testing"

	slvpb "github.com/lprao/slv-proto"
	"google.golang.org/grpc"
)

func TestNewGrpcClient(t *testing.T) {
	type args struct {
		ep       string
		insecure bool
		loc      string
	}
	tests := []struct {
		name string
		args args
		want *Grpc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGrpcClient(tt.args.ep, tt.args.insecure, tt.args.loc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGrpcClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpc_connect(t *testing.T) {
	type fields struct {
		conn           *grpc.ClientConn
		slvSvcEndpoint string
		insecure       bool
		caCertLoc      string
	}
	type args struct {
		ep       string
		insecure bool
		loc      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *grpc.ClientConn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grpc{
				conn:           tt.fields.conn,
				slvSvcEndpoint: tt.fields.slvSvcEndpoint,
				insecure:       tt.fields.insecure,
				caCertLoc:      tt.fields.caCertLoc,
			}
			got, err := g.connect(tt.args.ep, tt.args.insecure, tt.args.loc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grpc.connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc.connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpc_getSlvSvcClient(t *testing.T) {
	type fields struct {
		conn           *grpc.ClientConn
		slvSvcEndpoint string
		insecure       bool
		caCertLoc      string
	}
	type args struct {
		ep       string
		insecure bool
		loc      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    slvpb.LvSvcClient
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grpc{
				conn:           tt.fields.conn,
				slvSvcEndpoint: tt.fields.slvSvcEndpoint,
				insecure:       tt.fields.insecure,
				caCertLoc:      tt.fields.caCertLoc,
			}
			got, err := g.getSlvSvcClient(tt.args.ep, tt.args.insecure, tt.args.loc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grpc.getSlvSvcClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc.getSlvSvcClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpc_DoGrpc(t *testing.T) {
	type fields struct {
		conn           *grpc.ClientConn
		slvSvcEndpoint string
		insecure       bool
		caCertLoc      string
	}
	type args struct {
		req slvpb.ExecOpReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    slvpb.ExecOpResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grpc{
				conn:           tt.fields.conn,
				slvSvcEndpoint: tt.fields.slvSvcEndpoint,
				insecure:       tt.fields.insecure,
				caCertLoc:      tt.fields.caCertLoc,
			}
			got, err := g.DoGrpc(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Grpc.DoGrpc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc.DoGrpc() = %v, want %v", got, tt.want)
			}
		})
	}
}
