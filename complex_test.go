package cpy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/salam-rokt/go-cpy"
)

func TestComplex(t *testing.T) {
	cpy.Py_Initialize()

	nReal := 2.
	nImaginary := 5.

	nComplex := cpy.PyComplex_FromDoubles(nReal, nImaginary)
	defer nComplex.DecRef()

	assert.True(t, cpy.PyComplex_Check(nComplex))
	assert.True(t, cpy.PyComplex_CheckExact(nComplex))

	assert.InDelta(t, nReal, cpy.PyComplex_RealAsDouble(nComplex), 0.01)
	assert.InDelta(t, nImaginary, cpy.PyComplex_ImagAsDouble(nComplex), 0.01)
}
