package builtin

import (
//  "fmt"
//  "regexp"
  "unsafe"
)

type GenArr struct{
  Ptr    uintptr
  Size   int
  Stride int
}





func  Map_hof(in GenArr, fn func(uintptr, uintptr), ex GenArr){

  // Eventually add a parallelism to this
  for i := 0; i < in.Size; i++ {
    x := in.Ptr + uintptr(in.Stride * i);
    y := ex.Ptr + uintptr(ex.Stride * i);
    fn(x, y);
  }
}



func Fold_hof(in GenArr, carry uintptr, fn func(uintptr, uintptr, uintptr), ex *uintptr){

  for i := 0; i < in.Size; i++ {
    x := in.Ptr + uintptr(in.Stride * i);
    fn(x, carry, carry);
  }
  *ex = carry;
}



func Scan_hof(in GenArr, carry uintptr, fn func(uintptr, uintptr, uintptr), ex GenArr){

  for i := 0; i < in.Size; i++ {
    x := in.Ptr + uintptr(in.Stride * i);
    y := ex.Ptr + uintptr(ex.Stride * i);
    if i == 0 {
      fn(x, carry, y);
    }else{
      fn(x, y- uintptr(ex.Stride), y);
    }
  }
}



func ZipWith_hof(a, b GenArr, fn func(uintptr, uintptr, uintptr),  q GenArr){

  len := 0;
  if a.Size < b.Size {
    len = a.Size
  }else{
    len = b.Size
  }

  for i := 0; i < len; i++ {
    x := a.Ptr + uintptr(a.Stride * i)
    y := b.Ptr + uintptr(b.Stride * i)
    z := q.Ptr + uintptr(q.Stride * i)
    fn(x, y, z);
  }
}



func memcpy(size int, from, to uintptr){
  var dw *int64
  var bw *int8

  i := 0;
  for i = 0; (i+7) < size; i += 8 {
     dw =  (*int64)(unsafe.Pointer(to   + uintptr(i)))
    *dw = *(*int64)(unsafe.Pointer(from + uintptr(i)))
  }

  for j := i;  j < size; j++ {
     bw =  (*int8)(unsafe.Pointer(to   + uintptr(j)))
    *bw = *(*int8)(unsafe.Pointer(from + uintptr(j)))
  }
}
