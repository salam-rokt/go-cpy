#include "macro.h"
#include "datetime.h"

int _go_Py_EnterRecursiveCall(const char *where) {
    return Py_EnterRecursiveCall(where);
}

void _go_Py_LeaveRecursiveCall() {
    Py_LeaveRecursiveCall();
}

int _go_PyType_Check(PyObject *o) {
    return PyType_Check(o);
}

int _go_PyType_CheckExact(PyObject *o) {
    return PyType_CheckExact(o);
}

int _go_PyLong_Check(PyObject *p) {
    return PyLong_Check(p);
}

int _go_PyLong_CheckExact(PyObject *p) {
    return PyLong_CheckExact(p);
}

int _go_PyBool_Check(PyObject *o) {
    return PyBool_Check(o);
}

int _go_PyFloat_Check(PyObject *p) {
    return PyFloat_Check(p);
}

int _go_PyFloat_CheckExact(PyObject *p) {
    return PyFloat_CheckExact(p);
}

int _go_PyComplex_Check(PyObject *p) {
    return PyComplex_Check(p);
}

int _go_PyComplex_CheckExact(PyObject *p) {
    return PyComplex_CheckExact(p);
}

int _go_PyBytes_Check(PyObject *o) {
    return PyBytes_Check(o);
}

int _go_PyBytes_CheckExact(PyObject *o) {
    return PyBytes_CheckExact(o);
}

int _go_PyByteArray_Check(PyObject *o) {
    return PyByteArray_Check(o);
}

int _go_PyByteArray_CheckExact(PyObject *o) {
    return PyByteArray_CheckExact(o);
}

int _go_PyUnicode_Check(PyObject *o) {
    return PyUnicode_Check(o);
}

int _go_PyUnicode_CheckExact(PyObject *o) {
    return PyUnicode_CheckExact(o);
}

int _go_PyTuple_Check(PyObject *p) {
    return PyTuple_Check(p);
}

int _go_PyTuple_CheckExact(PyObject *p) {
    return PyTuple_CheckExact(p);
}

int _go_PyList_Check(PyObject *p) {
    return PyList_Check(p);
}

int _go_PyList_CheckExact(PyObject *p) {
    return PyList_CheckExact(p);
}

int _go_PyDict_Check(PyObject *p) {
    return PyDict_Check(p);
}

int _go_PyDict_CheckExact(PyObject *p) {
    return PyDict_CheckExact(p);
}

int _go_PyModule_Check(PyObject *p) {
    return PyModule_Check(p);
}

int _go_PyModule_CheckExact(PyObject *p) {
    return PyModule_CheckExact(p);
}

int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name) {
    return PyObject_DelAttr(o, attr_name);
}

int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name) {
    return PyObject_DelAttrString(o, attr_name);
}

int _go_PyObject_TypeCheck(PyObject *o, PyTypeObject *type) {
    return PyObject_TypeCheck(o, type);
}

void _go_PyDatetime_Init() {
	PyDateTime_IMPORT;
}

int _go_PyDateTime_Check(PyObject *obj) {
	return PyDateTime_Check(obj);
}

int _go_PyDateTime_CheckExact(PyObject *obj) {
	return PyDateTime_CheckExact(obj);
}

PyObject *_go_PyDateTime_FromDateAndTime(datetime_tpl tpl) {
	return PyDateTime_FromDateAndTime(
		tpl.year,
		tpl.month,
		tpl.day,
		tpl.hour,
		tpl.minute,
		tpl.second,
		tpl.microsecond
	);
}

datetime_tpl _go_PyDateTime_GetDateTime(PyObject *obj) {
	datetime_tpl dt;
	dt.year = PyDateTime_GET_YEAR(obj);
	dt.month = PyDateTime_GET_MONTH(obj);
	dt.day = PyDateTime_GET_DAY(obj);
	dt.hour = PyDateTime_DATE_GET_HOUR(obj);
	dt.minute = PyDateTime_DATE_GET_MINUTE(obj);
	dt.second = PyDateTime_DATE_GET_SECOND(obj);
	dt.microsecond = PyDateTime_DATE_GET_MICROSECOND(obj);
	return dt;
}