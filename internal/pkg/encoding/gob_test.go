package encoding

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeInt(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeInt(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeInt(t *testing.T) {
	type args struct {
		b bytes.Buffer
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeInt(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
