package config

import "fmt"

type Config struct {
	genType    string
	total      int
	distinct   int
	bufferSize int
}

type ValidationError struct {
	param string
}

func newValidationError(param string) ValidationError {
	return ValidationError{param: param}
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("parameter '%s' is not valid type", ve.param)
}

func NewConfig(params map[string]any) (*Config, error) {
	genType, ok := params["genType"].(string)
	if !ok {
		return nil, newValidationError("genType")
	}

	total, ok := params["total"].(int)
	if !ok {
		return nil, newValidationError("total")
	}

	distinct, ok := params["distinct"].(int)
	if !ok {
		return nil, newValidationError("distinct")
	}

	bufferSize, ok := params["bufferSize"].(int)
	if !ok {
		return nil, newValidationError("bufferSize")
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

func (conf *Config) GetBufferSize() int {
	return conf.bufferSize
}
