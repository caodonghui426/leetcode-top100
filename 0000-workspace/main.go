package main

import "fmt"

type A struct {
	Name string
	Age  int
	Sex  bool
}

func (a *A) Song(name string) (restr string) {
	restr = "非常好"
	fmt.Printf("%v 唱了一首 %v，观众觉得 %v \n", a.Name, name, restr)
	return restr
}
func main() {
	a := A{
		Name: "LittlePaddy",
		Age:  23,
		Sex:  true,
	}
	fmt.Println(a.Song("若是月亮还没来"))
}
