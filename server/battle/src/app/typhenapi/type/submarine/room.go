// This file was generated by typhen-api

package submarine

import (
	"app/typhenapi/core"
	"errors"
)

var _ = errors.New

// Room is a kind of TyphenAPI type.
type Room struct {
	Id      int     `codec:"id"`
	Members []*User `codec:"members"`
}

// Coerce the fields.
func (t *Room) Coerce() error {
	if t.Members == nil {
		return errors.New("Members should not be empty")
	}
	return nil
}

// Bytes creates the byte array.
func (t *Room) Bytes(serializer typhenapi.Serializer) ([]byte, error) {
	if err := t.Coerce(); err != nil {
		return nil, err
	}

	data, err := serializer.Serialize(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
