package cvm

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/logging"
	"github.com/tentameneu/cvm-go/internal/stream"
)

func newTestStreamRunner() *CVMRunner {
	streamGenerator, _ := stream.NewStreamGenerator()
	stream, _ := streamGenerator.Generate()
	return NewCVMRunner(stream)
}

func TestRun(t *testing.T) {
	t.Run("SmallerBuffer", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      1_000_000,
			"distinct":   10_000,
			"bufferSize": 1_000,
			"logLevel":   "info",
			"filePath":   "",
		})
		logging.InitializeLogger(os.Stdout)
		runner := newTestStreamRunner()
		n := runner.Run()
		assert.InDelta(t, 10_000, n, 1_000)
	})

	t.Run("ExactBuffer", func(t *testing.T) {
		config.SetConfig(map[string]any{
			"streamType": "incremental",
			"total":      1_000_000,
			"distinct":   10_000,
			"bufferSize": 10_000,
			"logLevel":   "info",
			"filePath":   "",
		})
		logging.InitializeLogger(os.Stdout)
		runner := newTestStreamRunner()
		n := runner.Run()
		assert.Exactly(t, 10_000, n)
	})

	t.Run("Logging", func(t *testing.T) {
		t.Run("Info", func(t *testing.T) {
			config.SetConfig(map[string]any{
				"streamType": "incremental",
				"total":      50,
				"distinct":   15,
				"bufferSize": 10,
				"logLevel":   "info",
				"filePath":   "",
			})
			writerBuffer := new(bytes.Buffer)
			logging.InitializeLogger(writerBuffer)
			runner := newTestStreamRunner()
			runner.Run()
			lines := strings.Split(writerBuffer.String(), "\n")

			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Starting CVM Algorithm\.\.\.$`, lines[0])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Done estimating number of distinct elements\. N=\d+ precision=100|\d+\.\d+\%$`, lines[1])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Buffer status: Size=\d+ Root=\<Value: \d+, Priority: 0\.\d+>$`, lines[2])
		})

		t.Run("Debug", func(t *testing.T) {
			config.SetConfig(map[string]any{
				"streamType": "incremental",
				"total":      10,
				"distinct":   5,
				"bufferSize": 5,
				"logLevel":   "debug",
				"filePath":   "",
			})
			writerBuffer := new(bytes.Buffer)
			logging.InitializeLogger(writerBuffer)
			runner := newTestStreamRunner()
			runner.Run()
			lines := strings.Split(writerBuffer.String(), "\n")

			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Starting CVM Algorithm\.\.\.$`, lines[0])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 1\/10 p=1 u=0\.\d+ Root=nil$`, lines[1])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 2\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[2])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 3\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[3])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 4\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[4])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 5\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[5])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 6\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[6])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 7\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[7])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 8\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[8])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 9\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[9])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 10\/10 p=1|0\.\d+ u=0\.\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[10])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Done estimating number of distinct elements\. N=\d+ precision=1|0\.\d{3}$`, lines[11])
			assert.Regexp(t, `^\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Buffer status: Size=\d+ Root=\<Value: \d+, Priority: 0\.\d+\>$`, lines[12])
		})

		t.Run("Deep", func(t *testing.T) {
			config.SetConfig(map[string]any{
				"streamType": "incremental",
				"total":      5,
				"distinct":   5,
				"bufferSize": 5,
				"logLevel":   "deep",
				"filePath":   "",
			})
			writerBuffer := new(bytes.Buffer)
			logging.InitializeLogger(writerBuffer)
			runner := newTestStreamRunner()
			runner.Run()

			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Starting CVM Algorithm\.\.\.`, writerBuffer.String())
			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| DEBUG \|\| Starting loop 1\/5 p=1 u=0\.\d+ Root=nil`, writerBuffer.String())
			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| DEEP \|\| Deleting next element from buffer a=0`, writerBuffer.String())
			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| DEEP \|\| |B| < s. Inserting new node a=0 u=0.\d+\>`, writerBuffer.String())
			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Done estimating number of distinct elements\. N=\d+ precision=1|0\.\d{3}`, writerBuffer.String())
			assert.Regexp(t, `\d{2}:\d{2}:\d{2}.\d{3} \|\| INFO \|\| Buffer status: Size=\d+ Root=\<Value: \d+, Priority: 0\.\d+\>`, writerBuffer.String())
		})
	})
}
