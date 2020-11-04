package Int64Utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatToFinancial(t *testing.T) {
	assert.Equal(t, "123", FormatToFinancial(123))
	assert.Equal(t, "1,123", FormatToFinancial(1123))
	assert.Equal(t, "1,111,223", FormatToFinancial(1111223))
	assert.Equal(t, "-111,111,223", FormatToFinancial(-111111223))
}
