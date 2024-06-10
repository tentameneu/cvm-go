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
		return nil, newValidationError("genType", "must be a string")
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
		return nil, newValidationError("bufferSize", "must be an integer")
	}

	if total <= 0 {
		return nil, newValidationError("total", "must be a positive integer")
	}

	if distinct <= 0 {
		return nil, newValidationError("distinct", "must be a positive integer")
	}

	if bufferSize <= 0 {
		return nil, newValidationError("bufferSize", "must be a positive integer")
	}

	if total < distinct {
		return nil, newValidationError("total < distinct", "total number of elements can't be smaller than distinct number of elements")
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
