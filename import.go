package py3

// #include <Python.h>
import "C"

import (
    "unsafe"
)

func Import(name *PyObject) *PyObject {
    return PyObjectToGO(C.PyImport_Import(GoToPyObject(name)))
}

func ImportModule(name string) *PyObject {
    c_name := C.CString(name)
    defer C.free(unsafe.Pointer(c_name))
    return PyObjectToGO(C.PyImport_ImportModule(c_name))
}
