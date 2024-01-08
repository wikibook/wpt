package main

import (
  "fmt"
)

// fmt.Stringer를 만족하는 Dog 구조체
type Dog struct {
  Name string
  Age  int
}

// fmt.Stringer.String 구현
func (d *Dog) String() string {
  return fmt.Sprintf("Dog: %s (%d)", d.Name, d.Age)
}

// fmt.GoStringer는 구현되지 않았으므로 기본 동작이 사용
// func (d *Dog) GoString() string {
//   // TODO: Implmenet fmt.GoStringer
// }

// fmt.Stringer 및 fmt.GoStringer를 충족하는 Cat 구조체
type Cat struct {
  Name string
  Age  int
}

// fmt.Stringer.String 구현
func (c *Cat) String() string {
  return fmt.Sprintf("(=^_^=) Cat: %s(%d)", c.Name, c.Age)
}

// fmt.GoStringer.GoString 구현
func (c *Cat) GoString() string {
  return fmt.Sprintf("(=^_^=) &Cat{%#v, %#v}", c.Name, c.Age)
}

func main() {
  // Dog를 생성
  dog := &Dog{Name:"coco", Age:5}

  // Cat를 생성
  cat := &Cat{Name:"nana", Age:3}

  fmt.Println(dog) // => Dog: coco (5)
  fmt.Println(cat) // => (=^_^=) Cat: nana(3)

  fmt.Printf("%#v\n", dog) // => &main.Dog{Name:"coco", Age:5}
  fmt.Printf("%#v\n", cat) // => (=^_^=) &Cat{"nana", 3}
}
