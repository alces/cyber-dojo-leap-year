package leapyear

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test1900ShouldNotBeLeap(t *testing.T) {
    assert.False(t, isLeap(1900))
}

func Test1917ShouldNotBeLeap(t *testing.T) {
    assert.False(t, isLeap(1917))
}

func Test1980ShouldBeLeap(t *testing.T) {
    assert.True(t, isLeap(1980))
}
