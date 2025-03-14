package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/salam-rokt/go-cpy"
)

func TestWarnEx(t *testing.T) {
	cpy.Py_Initialize()

	assert.Equal(t, -1, cpy.PyErr_WarnEx(cpy.PyExc_RuntimeWarning, "test warning", 3))
}

func TestWarnExplicitObject(t *testing.T) {
	cpy.Py_Initialize()

	message := cpy.PyUnicode_FromString("test warning")
	defer message.DecRef()

	filename := cpy.PyUnicode_FromString("test.py")
	defer filename.DecRef()

	module := cpy.PyUnicode_FromString("test_module")
	defer module.DecRef()

	assert.Equal(t, -1, cpy.PyErr_WarnExplicitObject(cpy.PyExc_RuntimeError, message, filename, 4, module, nil))
}

func TestWarnExplicit(t *testing.T) {
	cpy.Py_Initialize()

	assert.Equal(t, -1, cpy.PyErr_WarnExplicit(cpy.PyExc_RuntimeError, "test warning", "test.py", 4, "test_module", nil))
}
