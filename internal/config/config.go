package config

import "fmt"

type Config struct {
	genType    string
	total      int
	distinct   int
	randomMin  int
	randomMax  int
	bufferSize int
}

type ValidationError struct {
	param string
	msg   string
}

func newValidationError(param, msg string) ValidationError {
	return ValidationError{param: param, msg: msg}
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("invalid parameter '%s': %s", ve.param, ve.msg)
}

func NewConfig(params map[string]any) (*Config, error) {
	genType, ok := params["genType"].(string)
	if !ok {
		return nil, newValidationError("gen-type", "must be a string")
	}

	total, ok := params["total"].(int)
	if !ok {
		return nil, newValidationError("total", "must be an integer")
	}

	distinct, ok := params["distinct"].(int)
	if !ok {
		return nil, newValidationError("distinct", "must be an integer")
	}

	bufferSize, ok := params["bufferSize"].(int)
	if !ok {
		return nil, newValidationError("buffer-size", "must be an integer")
	}

	if total <= 0 {
		return nil, newValidationError("total", "must be a positive integer")
	}

	if distinct <= 0 {
		return nil, newValidationError("distinct", "must be a positive integer")
	}

	if bufferSize <= 0 {
		return nil, newValidationError("buffer-size", "must be a positive integer")
	}

	if total < distinct {
		return nil, newValidationError("total < distinct", "total number of elements can't be smaller than distinct number of elements")
	}

	if genType == "random" {
		randomMin, ok := params["randomMin"].(int)
		if !ok {
			return nil, newValidationError("random-min", "must be an integer")
		}

		randomMax, ok := params["randomMax"].(int)
		if !ok {
			return nil, newValidationError("random-max", "must be an integer")
		}

		if randomMin < 0 {
			return nil, newValidationError("random-min", "must be a positive integer or 0")
		}

		if randomMax < 0 {
			return nil, newValidationError("random-max", "must be a positive integer")
		}

		if randomMax < randomMin {
			return nil, newValidationError("random-max < random-min", "random-max can't be smaller than random-min")
		}

		if randomMax-randomMin < distinct {
			return nil, newValidationError("random-max - random-min < distinct", "(random-max - random-min) can't be smaller than distinct number of elements")
		}

		conf := &Config{
			genType:    genType,
			total:      total,
			distinct:   distinct,
			randomMin:  randomMin,
			randomMax:  randomMax,
			bufferSize: bufferSize,
		}

		return conf, nil
	}

	conf := &Config{
		genType:    genType,
		total:      total,
		distinct:   distinct,
		bufferSize: bufferSize,
	}

	return conf, nil
}

func (conf *Config) GetGenType() string {
	return conf.genType
}

func (conf *Config) GetTotal() int {
	return conf.total
}

func (conf *Config) GetDistinct() int {
	return conf.distinct
}

func (conf *Config) GetRandomMin() int {
	return conf.randomMin
}

func (conf *Config) GetRandomMax() int {
	return conf.randomMax
}

func (conf *Config) GetBufferSize() int {
	return conf.bufferSize
}
