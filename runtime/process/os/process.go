// Package os runs processes locally
package os

import (
	"github.com/stack-labs/stack-rpc/runtime/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
