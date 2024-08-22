//go:generate protoc -I=proto -I=$GOPATH/src -I=$GOPATH/src/github.com/DharitriOne/protobuf/protobuf  --gogoslick_out=. dct.proto
package dct

import "math/big"

// New returns a new batch from given buffers
func New() *DCToken {
	return &DCToken{
		Value: big.NewInt(0),
	}
}
