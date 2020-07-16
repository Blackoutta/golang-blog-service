package convert_test

import (
	"testing"

	"github.com/Blackoutta/blog-service/pkg/convert"
	"github.com/stretchr/testify/assert"
)

func TestUint32(t *testing.T) {
	result, _ := convert.StrTo("-22").Uint32()
	assert.Equal(t, uint32(0), result)

	result, _ = convert.StrTo("1").Uint32()
	assert.Equal(t, uint32(1), result)
}
