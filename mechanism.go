package kafka

import (
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Mechanism string

const (
	MechanismScram = "scram"
	MechanismPlain = "plain"
)

type SASLConfig struct {
	Type     Mechanism
	Username string
	Password string
}

func (s SASLConfig) Mechanism() (sasl.Mechanism, error) {
	if s.Type == MechanismScram {
		return scram.Mechanism(scram.SHA512, s.Username, s.Password)
	}

	return s.plain(), nil
}

func (s SASLConfig) plain() sasl.Mechanism {
	return &plain.Mechanism{
		Username: s.Username,
		Password: s.Password,
	}
}
