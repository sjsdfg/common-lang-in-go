package StringUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncateName(t *testing.T) {
	assert.Equal(t, "test", TruncateName("test", 6))
	assert.Equal(t, "这是一个测试...", TruncateName("这是一个测试哈哈哈", 6))
}
