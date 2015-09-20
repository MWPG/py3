package py3

// #include <Python.h>
import "C"

type PyFloat PyObject

// PyObject* PyFloat_FromDouble(double v)
// Return value: New reference.
// Create a PyFloatObject object from v, or NULL on failure.
func PyFloat_FromDouble(v float32) *PyFloat {
    return (*PyFloat)(PyObjectToGO(C.PyFloat_FromDouble(C.double(v))))
}
