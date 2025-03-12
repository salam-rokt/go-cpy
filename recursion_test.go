package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/salam-rokt/go-cpy"
)

func TestRecursion(t *testing.T) {
	cpy.Py_Initialize()

	assert.Zero(t, cpy.Py_EnterRecursiveCall("in test function"))

	cpy.Py_LeaveRecursiveCall()
}
