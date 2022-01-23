package slvlib

import (
	"bytes"
	"fmt"

	"github.com/lprao/slv-go-lib/internal/pkg/encoding"
	slvpb "github.com/lprao/slv-proto"
)

// Serverless integer ledger variable struct
type SlvInt struct {
	// Name of the variable
	Name string

	// Integer value of the variable
	Val int

	// Scope of the variable(can be set only once when the variable is created)
	Scope slvpb.VarScope

	// Permissions of the variable(can be set only once when the variable is created)
	Permissions slvpb.VarPermissions

	// SlvVar local reference
	slvVar *slvpb.SlvVar
}

// Integer opertations on Serverless ledger variables
type SlvIntOps interface {
	// Creates new variable of integer type in the ledger
	NewSlvInt(string, int, slvpb.VarScope, slvpb.VarPermissions) (*SlvInt, error)

	// Get the SlvInt object by name of the varible in the ledger
	GetSlvIntByName(string) (*SlvInt, error)

	// Get the value of the variable from the ledger
	Get() (int, error)

	// Set the value of the variable in the ledger
	Set(int) (int, error)

	// Adds the integer parameter to the value in the ledger and stores
	// the out come back into the ledger
	Add(int) (int, error)

	// Subtracts the integer parameter from the value in the ledger and stores
	// the out come back into the ledger
	Sub(int) (int, error)

	// Multiplies the integer parameter to the value in the ledger and stores
	// the out come back into the ledger
	Mul(int) (int, error)

	// Adds the integer parameter to the value in the ledger and stores
	// the out come back into the ledger
	Div(int) (int, error)
}

// Creates new variable of integer type in the ledger
func NewSlvInt(name string, val int, scope slvpb.VarScope, perms slvpb.VarPermissions) (*SlvInt, error) {
	var err error

	i := &SlvInt{
		Name:        name,
		Val:         val,
		Scope:       scope,
		Permissions: perms,
	}

	err = i.createSlvVarObj()
	if err != nil {
		return &SlvInt{}, err
	}

	outVar, err := client.execOp(i.slvVar, slvpb.Operation_NEW, "")
	if err != nil {
		return &SlvInt{}, err
	}

	err = i.updateSlvIntObj(outVar)
	if err != nil {
		return &SlvInt{}, err
	}

	return i, nil
}

// Get the variable by name
func GetSlvIntByName(name string) (*SlvInt, error) {
	i := &SlvInt{}

	i.slvVar = &slvpb.SlvVar{
		Name: i.Name,
		Type: slvpb.VarType_INT32,
	}

	outVar, err := client.execOp(i.slvVar, slvpb.Operation_GET, "")
	if err != nil {
		return &SlvInt{}, err
	}

	err = i.updateSlvIntObj(outVar)
	if err != nil {
		return &SlvInt{}, err
	}

	i.Name = outVar.Name
	i.Scope = outVar.Scope
	i.Permissions = outVar.Permissions

	return i, nil
}

// Create the local slvVar object to be used during GRPC call
func (i *SlvInt) createSlvVarObj() error {
	eVal, err := encoding.EncodeInt(i.Val)
	if err != nil {
		return err
	}

	i.slvVar = &slvpb.SlvVar{
		Name:        i.Name,
		Val:         eVal,
		Type:        slvpb.VarType_INT32,
		Scope:       i.Scope,
		Permissions: i.Permissions,
	}

	return nil
}

// Get the value of the variable from the ledger
func (i *SlvInt) Get() (int, error) {
	return i.execOp(0, slvpb.Operation_SET)
}

// Set the value of the variable in the ledger
func (i *SlvInt) Set(val int) (int, error) {
	return i.execOp(val, slvpb.Operation_SET)
}

// Adds the integer parameter to the value in the ledger and stores
// the out come back into the ledger
func (i *SlvInt) Add(val int) (int, error) {
	return i.execOp(val, slvpb.Operation_ADD)
}

// Subtracts the integer parameter from the value in the ledger and stores
// the out come back into the ledger
func (i *SlvInt) Sub(val int) (int, error) {
	return i.execOp(val, slvpb.Operation_SUB)
}

// Multiplies the integer parameter to the value in the ledger and stores
// the out come back into the ledger
func (i *SlvInt) Mul(val int) (int, error) {
	return i.execOp(val, slvpb.Operation_MUL)
}

// Adds the integer parameter to the value in the ledger and stores
// the out come back into the ledger
func (i *SlvInt) Div(val int) (int, error) {
	if val == 0 {
		return 0, fmt.Errorf(ErrInvalidIntDiv)
	}

	return i.execOp(val, slvpb.Operation_DIV)
}

// Helper routine to execute common tasks for each operation
func (i *SlvInt) execOp(val int, op slvpb.Operation) (int, error) {
	i.Val = val
	err := i.updateSlvVarObj()
	if err != nil {
		return 0, err
	}

	outVar, err := client.execOp(i.slvVar, op, "")
	if err != nil {
		return 0, err
	}

	err = i.updateSlvIntObj(outVar)
	if err != nil {
		return 0, err
	}

	return i.Val, nil
}

// Updates the local SlvVar object with the user provided value.
func (i *SlvInt) updateSlvVarObj() error {
	var err error
	i.slvVar.Val, err = encoding.EncodeInt(i.Val)
	if err != nil {
		return err
	}

	return nil
}

// Updates the local SlvInt object with the value returned by slv-svc
func (i *SlvInt) updateSlvIntObj(slvVar *slvpb.SlvVar) error {
	dVal, err := encoding.DecodeInt(*bytes.NewBuffer(slvVar.Val))
	if err != nil {
		return err
	}

	i.Val = dVal

	return nil
}
