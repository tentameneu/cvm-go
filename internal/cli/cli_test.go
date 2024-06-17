package cli

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestCLIParse(t *testing.T) {
	t.Run("Defaults", func(t *testing.T) {
		err := Parse()

		assert.Nil(t, err)
	})

	t.Run("InvalidParam", func(t *testing.T) {
		*total = -1
		err := Parse()

		assert.EqualError(t, err, "invalid parameter 'total': must be a positive integer")
	})
}
