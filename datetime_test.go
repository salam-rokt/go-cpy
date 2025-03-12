package cpy_test

import (
	"github.com/salam-rokt/go-cpy"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPyDateTime_Check(t *testing.T) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	defer pyDateTime.DecRef()
	assert.True(t, cpy.PyDateTime_Check(pyDateTime))
	cpy.Py_Finalize()
}

func TestPyDateTime_CheckExact(t *testing.T) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	defer pyDateTime.DecRef()
	assert.True(t, cpy.PyDateTime_CheckExact(pyDateTime))
	cpy.Py_Finalize()
}

func TestPyDateTime_FromGoTime(t *testing.T) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	defer pyDateTime.DecRef()
	returned := cpy.PyDateTime_AsGoTime(pyDateTime)
	assert.Equal(t, now, returned)
	cpy.Py_Finalize()
}

func TestPyDateTime_AsGoTime(t *testing.T) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	defer pyDateTime.DecRef()
	returned := cpy.PyDateTime_AsGoTime(pyDateTime)
	assert.Equal(t, now, returned)
	cpy.Py_Finalize()
}

func BenchmarkPyDateTime_FromGoTime(b *testing.B) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now := time.UnixMilli(int64(1741746096220))
		pyDateTime := cpy.PyDateTime_FromGoTime(now)
		pyDateTime.DecRef()
	}
	cpy.Py_Finalize()
}

func BenchmarkPyDateTime_AsGoTime(b *testing.B) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	defer pyDateTime.DecRef()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = cpy.PyDateTime_AsGoTime(pyDateTime).UnixMilli()
	}
	cpy.Py_Finalize()
}

func BenchmarkImportDatetime_NewDatetime(b *testing.B) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	cpy.PyRun_SimpleString("from datetime import datetime")
	mainModule := cpy.PyImport_ImportModule("__main__")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := cpy.PyLong_FromGoInt(2021)
		month := cpy.PyLong_FromGoInt(1)
		day := cpy.PyLong_FromGoInt(1)
		hour := cpy.PyLong_FromGoInt(1)
		minute := cpy.PyLong_FromGoInt(2)
		second := cpy.PyLong_FromGoInt(3)
		microsecond := cpy.PyLong_FromGoInt(5)
		pyDatetime := mainModule.CallMethodArgs("datetime", year, month, day, hour, minute, second, microsecond)
		pyDatetime.DecRef()
	}
	cpy.Py_Finalize()
}

func BenchmarkImportDatetime_FromTimestamp(b *testing.B) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	datetimeModule := cpy.PyImport_ImportModule("datetime")
	datetimeCls := datetimeModule.GetAttrString("datetime")
	now := float64(time.Now().Truncate(time.Microsecond).UnixMicro()) / 1e6
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pyTimestamp := cpy.PyFloat_FromDouble(now)
		ret := datetimeCls.CallMethodArgs("fromtimestamp", pyTimestamp)
		pyTimestamp.DecRef()
		ret.DecRef()
	}
	datetimeCls.DecRef()
	datetimeModule.DecRef()
	cpy.Py_Finalize()
}

func BenchmarkImportDatetime_DatetimeToGo(b *testing.B) {
	cpy.Py_Initialize()
	cpy.PyDatetime_Init()
	now := time.Now().Truncate(time.Microsecond)
	pyDateTime := cpy.PyDateTime_FromGoTime(now)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := pyDateTime.CallMethodArgs("timestamp")
		ts := cpy.PyFloat_AsDouble(ret)
		_ = time.Unix(int64(ts), 0)
		ret.DecRef()
	}
	pyDateTime.DecRef()
	cpy.Py_Finalize()
}
