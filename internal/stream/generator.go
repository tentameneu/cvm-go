package stream

import (
	"fmt"
)

type StreamGenerator interface {
	Generate() []int
}

type repeatingStreamGenerator struct {
	total    int
	distinct int
}

var invalidParamType error

func (repeating *repeatingStreamGenerator) Generate() []int {
	stream := make([]int, repeating.total)
	for i := 0; i < repeating.total; i++ {
		stream[i] = i % repeating.distinct
	}
	return stream
}

func NewStreamGenerator(genType string, args map[string]interface{}) (StreamGenerator, error) {
	switch genType {
	case "repeating":
		if err := verifyGeneratorArgs(args, "total", "distinct"); err != nil {
			return nil, err
		}

		total, ok := args["total"].(int)
		if !ok {
			return nil, invalidParamType
		}

		distinct, ok := args["distinct"].(int)
		if !ok {
			return nil, invalidParamType
		}

		return &repeatingStreamGenerator{total: total, distinct: distinct}, nil

	default:
		return nil, fmt.Errorf("unknown generator type '%s'", genType)
	}
}

func verifyGeneratorArgs(args map[string]interface{}, required ...string) error {
	for _, requiredParam := range required {
		_, ok := args[requiredParam]

		if !ok {
			return fmt.Errorf("argument '%s' missing", requiredParam)
		}
	}
	return nil
}
