package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/salam-rokt/go-cpy"
)

func TestSequenceContains_Dict(t *testing.T) {
	cpy.Py_Initialize()

	dict := cpy.PyDict_New()
	defer dict.DecRef()

	dict.SetItem(cpy.PyUnicode_FromString("foo"), cpy.PyUnicode_FromString("bar"))

	assert.Equal(t, 1, cpy.PySequence_Contains(dict, cpy.PyUnicode_FromString("foo")))
	assert.Equal(t, 0, cpy.PySequence_Contains(dict, cpy.PyUnicode_FromString("bar")))
}
