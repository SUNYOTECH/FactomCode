package notaryapi

import (
	"encoding"
	"errors"
	
	"math/big"
)

type BinaryMarshallable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	MarshalledSize() uint64
}

func bigIntMarshalBinary(i *big.Int) (data []byte, err error) {
	intd, err := i.GobEncode()
	if err != nil { return }
	
	size := len(intd)
	if size > 255 { return nil, errors.New("Big int is too big") }
	
	data = make([]byte, size+1)
	data[0] = byte(size)
	copy(data[1:], intd)
	return
}

func bigIntMarshalledSize(i *big.Int) uint64 {
	intd, err := i.GobEncode()
	if err != nil { return 0 }
	
	return uint64(1 + len(intd))
}

func bigIntUnmarshalBinary(data []byte) (retd []byte, i *big.Int, err error) {
	size, data := uint8(data[0]), data[1:]
	
	i = new(big.Int)
	err, retd = i.GobDecode(data[:size]), data[size:]
	
	return
}

type simpleData struct {
	data []byte
}

func (d *simpleData) MarshalBinary() ([]byte, error) {
	return d.data, nil
}

func (d *simpleData) MarshalledSize() uint64 {
	return uint64(len(d.data))
}

func (d *simpleData) UnmarshalBinary([]byte) error {
	return errors.New("simpleData cannot be unmarshalled")
}
