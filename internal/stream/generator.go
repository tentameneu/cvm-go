package stream

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/logging"
)

var log = logging.Logger

type StreamGenerator interface {
	Generate() ([]int, error)
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

type fileStreamGenerator struct {
	filePath string
}

func (gen *incrementalStreamGenerator) Generate() ([]int, error) {
	stream := make([]int, gen.total)
	for i := 0; i < gen.total; i++ {
		stream[i] = i % gen.distinct
	}
	return stream, nil
}

func (gen *randomStreamGenerator) Generate() ([]int, error) {
	stream := make([]int, gen.total)
	for i := 0; i < gen.total; i++ {
		if i < gen.distinct {
			stream[i] = rand.Intn(gen.max-gen.min+1) + gen.min
		} else {
			stream[i] = stream[i%gen.distinct]
		}
	}
	return stream, nil
}

func (gen *fileStreamGenerator) Generate() ([]int, error) {
	file, err := os.Open(gen.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	total := 0
	stream := make([]int, 0)
	found := map[int]bool{}
	buff := []byte{}
	current := make([]byte, 1)

	processBuffer := func() error {
		n, err := strconv.Atoi(string(buff))
		if err != nil {
			return err
		}
		found[n] = true
		stream = append(stream, n)
		logging.LogDeep("Element appended to stream", "n", n, "stream", stream)
		total++
		buff = []byte{}
		return nil
	}

	log().Info(fmt.Sprintf("Started generating stream from file '%s'", gen.filePath))

	for {
		_, err := file.Read(current)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}

			if rune(current[0]) != ' ' && rune(current[0]) != '\n' {
				if err := processBuffer(); err != nil {
					return nil, err
				}
			}

			break
		}

		if rune(current[0]) != ' ' && rune(current[0]) != '\n' {
			buff = append(buff, current[0])
		} else {
			if err := processBuffer(); err != nil {
				return nil, err
			}
		}
	}
	distinct := len(found)
	config.SetDistinct(distinct)
	log().Info("Stream generated.", "Number of distinct elements", distinct, "Number of total elements", total)

	return stream, nil
}

func NewStreamGenerator() (StreamGenerator, error) {
	switch config.StreamType() {
	case "incremental":
		return &incrementalStreamGenerator{
			total:    config.Total(),
			distinct: config.Distinct(),
		}, nil
	case "random":
		return &randomStreamGenerator{
			total:    config.Total(),
			distinct: config.Distinct(),
			min:      config.RandomMin(),
			max:      config.RandomMax(),
		}, nil
	case "file":
		return &fileStreamGenerator{
			filePath: config.FilePath(),
		}, nil
	default:
		return nil, fmt.Errorf("unknown generator type '%s'", config.StreamType())
	}
}
