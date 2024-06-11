package stream

import (
	"fmt"
	"math/rand"

	"github.com/tentameneu/cvm-go/internal/config"
)

type StreamGenerator interface {
	Generate() []int
}

type incrementalStreamGenerator struct {
	total    int
	distinct int
}

type randomStreamGenerator struct {
	min      int
	max      int
	total    int
	distinct int
}

func (inc *incrementalStreamGenerator) Generate() []int {
	stream := make([]int, inc.total)
	for i := 0; i < inc.total; i++ {
		stream[i] = i % inc.distinct
	}
	return stream
}

func (random *randomStreamGenerator) Generate() []int {
	stream := make([]int, random.total)
	for i := 0; i < random.total; i++ {
		if i < random.distinct {
			stream[i] = rand.Intn(random.max-random.min+1) + random.min
		} else {
			stream[i] = stream[i%random.distinct]
		}
	}
	return stream
}

func NewStreamGenerator(conf *config.Config) (StreamGenerator, error) {
	switch conf.GetGenType() {
	case "incremental":
		return &incrementalStreamGenerator{total: conf.GetTotal(), distinct: conf.GetDistinct()}, nil
	case "random":
		return &randomStreamGenerator{total: conf.GetTotal(), distinct: conf.GetDistinct(), min: conf.GetRandomMin(), max: conf.GetRandomMax()}, nil
	default:
		return nil, fmt.Errorf("unknown generator type '%s'", conf.GetGenType())
	}
}
