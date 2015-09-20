package py3

// #include <Python.h>
import "C"

type List PyObject

func NewList(size int) (*List) {
    return (*List)(PyObjectToGO(C.PyList_New(C.Py_ssize_t(size))))
}

// Insert a reference to object o at position pos in the list pointed to by self
// Return 0 on success.
// Note This function “steals” a reference to o.
func (self *List) SetItem(pos int, obj *PyObject) (int) {
    ret := C.PyList_SetItem(self.ptr, C.Py_ssize_t(pos), obj.ptr)
    return (int)(ret)
   // return int2Err(ret)
}
