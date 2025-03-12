#ifndef MACRO_H
#define MACRO_H

#include "Python.h"

int _go_Py_EnterRecursiveCall(const char *where);
void _go_Py_LeaveRecursiveCall();

int _go_PyType_Check(PyObject *o);
int _go_PyType_CheckExact(PyObject *o);

int _go_PyLong_Check(PyObject *p);
int _go_PyLong_CheckExact(PyObject *p);

int _go_PyBool_Check(PyObject *o);

int _go_PyFloat_Check(PyObject *p);
int _go_PyFloat_CheckExact(PyObject *p);

int _go_PyComplex_Check(PyObject *p);
int _go_PyComplex_CheckExact(PyObject *p);

int _go_PyBytes_Check(PyObject *o);
int _go_PyBytes_CheckExact(PyObject *o);

int _go_PyByteArray_Check(PyObject *o);
int _go_PyByteArray_CheckExact(PyObject *o);

int _go_PyUnicode_Check(PyObject *o);
int _go_PyUnicode_CheckExact(PyObject *o);

int _go_PyTuple_Check(PyObject *p);
int _go_PyTuple_CheckExact(PyObject *p);

int _go_PyList_Check(PyObject *p);
int _go_PyList_CheckExact(PyObject *p);

int _go_PyDict_Check(PyObject *p);
int _go_PyDict_CheckExact(PyObject *p);

int _go_PyModule_Check(PyObject *p);
int _go_PyModule_CheckExact(PyObject *p);

int _go_PyObject_DelAttr(PyObject *o, PyObject *attr_name);
int _go_PyObject_DelAttrString(PyObject *o, const char *attr_name);

int _go_PyObject_TypeCheck(PyObject *o, PyTypeObject *type);

void _go_PyDatetime_Init();
int _go_PyDateTime_Check(PyObject *obj);
int _go_PyDateTime_CheckExact(PyObject *obj);

typedef struct datetime_tpl {
    int year;
    unsigned int microsecond;
    unsigned char month;
    unsigned char day;
    unsigned char hour;
    unsigned char minute;
    unsigned char second;
} datetime_tpl;
PyObject *_go_PyDateTime_FromDateAndTime(datetime_tpl tpl);
datetime_tpl _go_PyDateTime_GetDateTime(PyObject *obj);
#endif
