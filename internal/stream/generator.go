package stream

import (
	"fmt"

	"github.com/tentameneu/cvm-go/internal/config"
)

type StreamGenerator interface {
	Generate() []int
}

type repeatingStreamGenerator struct {
	total    int
	distinct int
}

func (repeating *repeatingStreamGenerator) Generate() []int {
	stream := make([]int, repeating.total)
	for i := 0; i < repeating.total; i++ {
		stream[i] = i % repeating.distinct
	}
	return stream
}

func NewStreamGenerator(conf *config.Config) (StreamGenerator, error) {
	switch conf.GetGenType() {
	case "repeating":
		return &repeatingStreamGenerator{total: conf.GetTotal(), distinct: conf.GetDistinct()}, nil
	default:
		return nil, fmt.Errorf("unknown generator type '%s'", conf.GetGenType())
	}
}
