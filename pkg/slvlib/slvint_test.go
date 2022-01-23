package slvlib

import (
	"reflect"
	"testing"

	slvpb "github.com/lprao/slv-proto"
)

func TestNewSlvInt(t *testing.T) {
	type args struct {
		name  string
		val   int
		scope slvpb.VarScope
		perms slvpb.VarPermissions
	}
	tests := []struct {
		name    string
		args    args
		want    *SlvInt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSlvInt(tt.args.name, tt.args.val, tt.args.scope, tt.args.perms)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSlvInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSlvIntByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *SlvInt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSlvIntByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSlvIntByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSlvIntByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_createSlvVarObj(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			if err := i.createSlvVarObj(); (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.createSlvVarObj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlvInt_Get(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_Set(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Set(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_Add(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Add(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_Sub(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Sub(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Sub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_Mul(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Mul(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_Div(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.Div(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_execOp(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		val int
		op  slvpb.Operation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			got, err := i.execOp(tt.args.val, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.execOp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SlvInt.execOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlvInt_updateSlvVarObj(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			if err := i.updateSlvVarObj(); (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.updateSlvVarObj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSlvInt_updateSlvIntObj(t *testing.T) {
	type fields struct {
		Name        string
		Val         int
		Scope       slvpb.VarScope
		Permissions slvpb.VarPermissions
		slvVar      *slvpb.SlvVar
	}
	type args struct {
		slvVar *slvpb.SlvVar
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &SlvInt{
				Name:        tt.fields.Name,
				Val:         tt.fields.Val,
				Scope:       tt.fields.Scope,
				Permissions: tt.fields.Permissions,
				slvVar:      tt.fields.slvVar,
			}
			if err := i.updateSlvIntObj(tt.args.slvVar); (err != nil) != tt.wantErr {
				t.Errorf("SlvInt.updateSlvIntObj() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
