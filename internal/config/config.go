package config

type Config struct {
	genType    string
	total      int
	distinct   int
	bufferSize int
}

var InvalidParameterTypeError error

func NewConfig(params map[string]any) (*Config, error) {
	genType, ok := params["genType"].(string)
	if !ok {
		return nil, InvalidParameterTypeError
	}

	total, ok := params["total"].(int)
	if !ok {
		return nil, InvalidParameterTypeError
	}

	distinct, ok := params["distinct"].(int)
	if !ok {
		return nil, InvalidParameterTypeError
	}

	bufferSize, ok := params["bufferSize"].(int)
	if !ok {
		return nil, InvalidParameterTypeError
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
