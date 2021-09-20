package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpression_minuteFieldParser(t *testing.T) {
	t.Run("should return full range given wildcard pattern", func(t *testing.T) {
		pattern := "*"
		expr := &Expression{}
		err := expr.minuteFieldParser(pattern)

		expectedMinutesList := generateRange(0, 59, 1)

		assert.NoError(t, err)
		assert.Equal(t, expr.MinutesList, expectedMinutesList)
	})
}
