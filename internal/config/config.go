package config

import "fmt"

var Conf *Config

type Config struct {
	streamType string
	total      int
	distinct   int
	randomMin  int
	randomMax  int
	bufferSize int
	logLevel   string
	filePath   string
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

func SetConfig(params map[string]any) error {
	streamType, ok := params["streamType"].(string)
	if !ok {
		return newValidationError("stream-type", "must be a string")
	}

	total, ok := params["total"].(int)
	if !ok {
		return newValidationError("total", "must be an integer")
	}

	distinct, ok := params["distinct"].(int)
	if !ok {
		return newValidationError("distinct", "must be an integer")
	}

	bufferSize, ok := params["bufferSize"].(int)
	if !ok {
		return newValidationError("buffer-size", "must be an integer")
	}

	logLevel, ok := params["logLevel"].(string)
	if !ok {
		return newValidationError("log-level", "must be a string")
	}

	filePath, ok := params["filePath"].(string)
	if !ok {
		return newValidationError("file-path", "must be a string")
	}

	Conf = &Config{
		streamType: streamType,
		total:      total,
		distinct:   distinct,
		bufferSize: bufferSize,
		logLevel:   logLevel,
		filePath:   filePath,
	}

	if streamType == "file" {
		return nil
	}

	if total <= 0 {
		return newValidationError("total", "must be a positive integer")
	}

	if distinct <= 0 {
		return newValidationError("distinct", "must be a positive integer")
	}

	if bufferSize <= 0 {
		return newValidationError("buffer-size", "must be a positive integer")
	}

	if total < distinct {
		return newValidationError("total < distinct", "total number of elements can't be smaller than distinct number of elements")
	}

	if streamType == "random" {
		randomMin, ok := params["randomMin"].(int)
		if !ok {
			return newValidationError("random-min", "must be an integer")
		}

		randomMax, ok := params["randomMax"].(int)
		if !ok {
			return newValidationError("random-max", "must be an integer")
		}

		if randomMin < 0 {
			return newValidationError("random-min", "must be a positive integer or 0")
		}

		if randomMax < 0 {
			return newValidationError("random-max", "must be a positive integer")
		}

		if randomMax < randomMin {
			return newValidationError("random-max < random-min", "random-max can't be smaller than random-min")
		}

		if randomMax-randomMin < distinct {
			return newValidationError("random-max - random-min < distinct", "(random-max - random-min) can't be smaller than distinct number of elements")
		}

		Conf.randomMin = randomMin
		Conf.randomMax = randomMax
	}

	return nil
}

func StreamType() string {
	return Conf.streamType
}

func Total() int {
	return Conf.total
}

func Distinct() int {
	return Conf.distinct
}

func SetDistinct(n int) {
	if Conf.streamType == "file" {
		Conf.distinct = n
	}
}

func RandomMin() int {
	return Conf.randomMin
}

func RandomMax() int {
	return Conf.randomMax
}

func BufferSize() int {
	return Conf.bufferSize
}

func LogLevel() string {
	return Conf.logLevel
}

func FilePath() string {
	return Conf.filePath
}
