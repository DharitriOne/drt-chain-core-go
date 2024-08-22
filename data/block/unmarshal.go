package block

import (
	"github.com/DharitriOne/drt-chain-core-go/core/check"
	"github.com/DharitriOne/drt-chain-core-go/data"
	"github.com/DharitriOne/drt-chain-core-go/marshal"
)

// GetHeaderFromBytes will unmarshal the header bytes based on the header type
func GetHeaderFromBytes(marshaller marshal.Marshalizer, creator EmptyBlockCreator, headerBytes []byte) (data.HeaderHandler, error) {
	if check.IfNil(marshaller) {
		return nil, data.ErrNilMarshalizer
	}
	if check.IfNil(creator) {
		return nil, data.ErrNilEmptyBlockCreator
	}

	header := creator.CreateNewHeader()
	err := marshaller.Unmarshal(header, headerBytes)
	if err != nil {
		return nil, err
	}

	return header, nil
}
