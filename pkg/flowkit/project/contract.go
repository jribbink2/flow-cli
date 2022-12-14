package project

import (
	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
)

// Contract is a Cadence contract definition for a project.
type Contract struct {
	Name           string
	location       string
	code           []byte
	AccountAddress flow.Address
	AccountName    string
	Args           []cadence.Value
}

func NewContract(
	name string,
	location string,
	code []byte,
	accountAddress flow.Address,
	accountName string,
	args []cadence.Value,
) *Contract {
	return &Contract{
		Name:           name,
		location:       location,
		code:           code,
		AccountAddress: accountAddress,
		AccountName:    accountName,
		Args:           args,
	}
}

func (c *Contract) Code() []byte {
	return c.code
}

func (c *Contract) SetCode(code []byte) {
	c.code = code
}

func (c *Contract) Location() string {
	return c.location
}

type Aliases map[string]string