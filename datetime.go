package cpy

/*
#include "Python.h"
#include "macro.h"
#include "datetime.h"
*/
import "C"
import (
	"time"
)

// PyDatetime_Init initializes the datetime module. This function must be called before using any other functions in the
// datetime module.

// Reference: https://docs.python.org/3/c-api/datetime.html#c.PyDateTime_IMPORT

func PyDatetime_Init() {
	C._go_PyDatetime_Init()
}

// PyDateTime_Check returns true if its argument is a PyDateTimeObject or a subtype of PyDateTimeObject. This function
// always succeeds.

// Reference: https://docs.python.org/3/c-api/datetime.html#c.PyDateTime_Check

func PyDateTime_Check(obj *PyObject) bool {
	return C._go_PyDateTime_Check((*C.PyObject)(obj)) != 0
}

// PyDateTime_CheckExact returns true if its argument is a PyDateTimeObject, but not a subtype of PyDateTimeObject. This
// function always succeeds.

// Reference: https://docs.python.org/3/c-api/datetime.html#c.PyDateTime_CheckExact

func PyDateTime_CheckExact(obj *PyObject) bool {
	return C._go_PyDateTime_CheckExact((*C.PyObject)(obj)) != 0
}

// PyDateTime_FromGoTime returns a new reference to a PyDateTime object initialized from date and time.

// Reference: https://docs.python.org/3/c-api/datetime.html#c.PyDateTime_FromDateAndTime

func PyDateTime_FromGoTime(timestamp time.Time) *PyObject {
	year, month, day := timestamp.Date()
	hour, minute, second := timestamp.Clock()
	microseconds := timestamp.Nanosecond() / 1_000
	datetime := C.datetime_tpl{
		year:        C.int(year),
		microsecond: C.uint(microseconds),
		month:       C.uchar(month),
		day:         C.uchar(day),
		hour:        C.uchar(hour),
		minute:      C.uchar(minute),
		second:      C.uchar(second),
	}
	return togo(C._go_PyDateTime_FromDateAndTime(datetime))
}

// PyDateTime_AsGoTime returns the date and time representation of the PyDateTime object.

// Reference: https://docs.python.org/3/c-api/datetime.html#c.PyDateTime_DATE_GET

func PyDateTime_AsGoTime(obj *PyObject) time.Time {
	result := C._go_PyDateTime_GetDateTime(toc(obj))
	return time.Date(
		int(result.year),
		time.Month(result.month),
		int(result.day),
		int(result.hour),
		int(result.minute),
		int(result.second),
		int(result.microsecond)*1_000,
		time.Local,
	)
}
