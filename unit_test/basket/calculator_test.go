package basket

import (
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	x, y := 3, 3

	c := Calculate{}

	actual := c.Add(x, y)
	expected := 6

	if actual != expected {
		t.Error("Toplam hatalı hesaplandı")
	}
}

func TestSubtract(t *testing.T) {
	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 0},
		{5, 3, 2},
		{5, 4, 1},
	}

	for _, v := range tables {
		actual := c.Subtract(v.x, v.y)
		expected := v.expected

		if actual != expected {
			t.Error("çıkarma hatalı hesaplandı")
		}
	}
}

func TestMultiply(t *testing.T) {

	assert := asrt.New(t)

	c := Calculate{}

	tables := []struct {
		x        int
		y        int
		expected int
	}{
		{2, 2, 4},
		{5, 3, 15},
		{5, 4, 20},
	}

	for _, v := range tables {
		actual := c.Multiply(v.x, v.y)
		expected := v.expected

		assert.Equal(actual, expected)
	}
}
