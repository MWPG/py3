package py3

// #include <Python.h>
import "C"

import (
    "unsafe"
)

// wrapper for a *C.PyObject structure pointer
type PyObject struct {
	ptr *C.PyObject
}

func (self *PyObject) DecRef() {
	C.Py_DecRef(self.ptr)
}

func (self *PyObject) GetAttrString(attr_name string) *PyObject {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return PyObjectToGO(C.PyObject_GetAttrString(self.ptr, c_attr_name))
}

func GoToPyObject(self *PyObject) *C.PyObject {
    return self.ptr
}

func (self *PyObject) IncRef() {
	C.Py_IncRef(self.ptr)
}

// Converts a PyObject from a *C.PyObject.
func PyObjectToGO(obj *C.PyObject) *PyObject {
	if obj == nil {
		return nil
	}
	return &PyObject{ptr: obj}
}

// PyObject_FromVoidPtr converts a PyObject from an unsafe.Pointer.
func PyObject_FromVoidPtr(ptr unsafe.Pointer) *PyObject {
    return PyObjectToGO((*C.PyObject)(ptr))
}

// Converts a *PyObject to a *C.PyObject
func (self *PyObject) ToPy() *C.PyObject {
    return self.ptr
}

