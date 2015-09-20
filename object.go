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

// PyObject* PyObject_CallObject(PyObject *callable_object, PyObject *args)
// Return value: New reference.
// Call a callable Python object callable_object, with arguments given by the
// tuple args. If no arguments are needed, then args may be NULL. Returns the
// result of the call on success, or NULL on failure. This is the equivalent of
// the Python expression apply(callable_object, args) or callable_object(*args).
func (self *PyObject) CallObject(args *PyObject) *PyObject {
    return PyObjectToGO(C.PyObject_CallObject(self.ptr, args.ptr))
}

func (self *PyObject) DecRef() {
	C.Py_DecRef(self.ptr)
}

func (self *PyObject) GetAttrString(attr_name string) (*PyObject) {
	c_attr_name := C.CString(attr_name)
	defer C.free(unsafe.Pointer(c_attr_name))
	return PyObjectToGO(C.PyObject_GetAttrString(self.ptr, c_attr_name))
}

func GoToPyObject(self *PyObject) (*C.PyObject) {
    return self.ptr
}

func (self *PyObject) IncRef() () {
	C.Py_IncRef(self.ptr)
}

// Converts a *PyObject to a *C.PyObject
func (self *PyObject) Ptr() (*C.PyObject) {
    return self.ptr
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

// Converts a *PyObject to a *C.PyObject, alias of the Ptr getter
func (self *PyObject) ToPy() *C.PyObject {
    return self.ptr
}
