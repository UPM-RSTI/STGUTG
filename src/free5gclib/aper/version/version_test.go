package version_test

import (
	"free5gclib/aper/version"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "2020-03-31-01", version.GetVersion())
}
