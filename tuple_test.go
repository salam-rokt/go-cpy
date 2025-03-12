package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/salam-rokt/go-cpy"
)

func TestTupleCheck(t *testing.T) {
	cpy.Py_Initialize()

	tuple := cpy.PyTuple_New(0)
	defer tuple.DecRef()

	assert.True(t, cpy.PyTuple_Check(tuple))
	assert.True(t, cpy.PyTuple_CheckExact(tuple))
}

func TestTupleNew(t *testing.T) {
	cpy.Py_Initialize()

	tuple := cpy.PyTuple_New(0)
	defer tuple.DecRef()

	assert.NotNil(t, tuple)
}

func TestTupleSize(t *testing.T) {
	cpy.Py_Initialize()

	size := 45

	tuple := cpy.PyTuple_New(size)
	defer tuple.DecRef()

	assert.Equal(t, size, cpy.PyTuple_Size(tuple))
}

func TestTupleGetSetItem(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("test")

	i := cpy.PyLong_FromGoInt(34)

	tuple := cpy.PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, cpy.PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, cpy.PyTuple_SetItem(tuple, 1, i))

	assert.Equal(t, i, cpy.PyTuple_GetItem(tuple, 1))
}

func TestTupleGetSlice(t *testing.T) {
	cpy.Py_Initialize()

	s := cpy.PyUnicode_FromString("test")

	i := cpy.PyLong_FromGoInt(34)

	tuple := cpy.PyTuple_New(2)
	defer tuple.DecRef()

	assert.Zero(t, cpy.PyTuple_SetItem(tuple, 0, s))
	assert.Zero(t, cpy.PyTuple_SetItem(tuple, 1, i))

	slice := cpy.PyTuple_GetSlice(tuple, 0, 1)
	defer slice.DecRef()

	assert.True(t, cpy.PyTuple_Check(slice))
	assert.Equal(t, 1, cpy.PyTuple_Size(slice))
	assert.Equal(t, s, cpy.PyTuple_GetItem(slice, 0))
}
