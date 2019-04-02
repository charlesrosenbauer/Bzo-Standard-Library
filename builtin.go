package builtin

import (
//  "fmt"
//  "regexp"
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



func Fold_hof(in GenArr, carry uintptr, fn func(uintptr, uintptr, uintptr), ex uintptr){

  for i := 0; i < in.Size; i++ {
    x := in.Ptr + uintptr(in.Stride * i);
    fn(x, carry, carry);
  }
  ex = carry;
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
