# 🔺 Triforce
> Standard Tessellation Language library for Gophers

An STL file describes a raw, unstructured triangulated surface by the unit normal and vertices (ordered by the right-hand rule) of the triangles using a three-dimensional Cartesian coordinate system.

The `triforce` package allows you to easily work with small and large binary STL files.

```go
package main

import (
  "fmt"
  "github.com/picatz/triforce"
)

// find the maximum X, Y and Z values of an arbitrarily large STL file
func main() {
  var x float32
  var y float32
  var z float32
  for triangle := range triforce.ReadBinaryTriangleStream("binary.stl") {
    for _, vertex := range triangle.Vertices {
      if vertex.X > x {
        x = vertex.X
      }
      if vertex.Y > y {
        y = vertex.Y
      }
      if vertex.Z > z {
        z = vertex.Z
      }
    }
  }
  fmt.Println("max X value: ", x)
  fmt.Println("max Y value: ", y)
  fmt.Println("max Z value: ", z)
}
```
