package py3

// #include <Python.h>
import "C"

type Tuple PyObject

// NewTuple returns a new *PyObject of the specified size. However the entries
// are all set to NULL, so the tuple should not be shared, especially with
// Python code, until the entries have all been set.
//
// Return value: New Reference.
func NewTuple(size int) (*Tuple) {
    return (*Tuple)(PyObjectToGO(C.PyTuple_New(C.Py_ssize_t(size))))
}

// Insert a reference to object o at position pos of the tuple pointed to by p.
// Return 0 on success.
// Note This function “steals” a reference to o.
func (self *Tuple) SetItem(pos int, obj *PyObject) (int) {
    ret := C.PyTuple_SetItem(self.ptr, C.Py_ssize_t(pos), obj.ptr)
    return (int)(ret)
   // return int2Err(ret)
}
