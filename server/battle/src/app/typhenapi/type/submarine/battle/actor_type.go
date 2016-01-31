// This file was generated by typhen-api

package battle

import (
	"errors"
)

type ActorType int

func (i ActorType) String() string {
	switch i {
	case ActorType_Submarine:
		return "0"
	case ActorType_Torpedo:
		return "1"
	case ActorType_Decoy:
		return "2"
	case ActorType_Lookout:
		return "3"
	}
	return ""
}

func (i ActorType) Coerce() error {
	switch i {
	case ActorType_Submarine:
		return nil
	case ActorType_Torpedo:
		return nil
	case ActorType_Decoy:
		return nil
	case ActorType_Lookout:
		return nil
	}
	return errors.New("No ActorType")
}

const (
	ActorType_Submarine ActorType = 0
	ActorType_Torpedo   ActorType = 1
	ActorType_Decoy     ActorType = 2
	ActorType_Lookout   ActorType = 3
)