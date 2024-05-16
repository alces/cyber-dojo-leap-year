package leapyear

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test1917ShouldNotBeLeap(t *testing.T) {
    assert.False(t, isLeap(1917))
}
