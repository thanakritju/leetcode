package findmedianfromdatastream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	assert.Equal(t, 1.5, obj.FindMedian())
	obj.AddNum(3)
	assert.Equal(t, 2.0, obj.FindMedian())
}

func TestMedianFinder2(t *testing.T) {
	obj := Constructor()
	obj.AddNum(6)
	assert.Equal(t, 6.0, obj.FindMedian())
	obj.AddNum(10)
	assert.Equal(t, 8.0, obj.FindMedian())
	obj.AddNum(2)
	assert.Equal(t, 6.0, obj.FindMedian())
	obj.AddNum(6)
	assert.Equal(t, 6.0, obj.FindMedian())
	obj.AddNum(5)
	assert.Equal(t, 6.0, obj.FindMedian())
	obj.AddNum(0)
	assert.Equal(t, 5.5, obj.FindMedian())
	obj.AddNum(6)
	assert.Equal(t, 6.0, obj.FindMedian())
	obj.AddNum(3)
	assert.Equal(t, 5.5, obj.FindMedian())
	obj.AddNum(1)
	assert.Equal(t, 5.0, obj.FindMedian())
	obj.AddNum(0)
	assert.Equal(t, 4.0, obj.FindMedian())
	obj.AddNum(0)
	assert.Equal(t, 3.0, obj.FindMedian())
}
