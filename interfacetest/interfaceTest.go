package interfacetest


import (
	"fmt"
)

type IParamer interface {
	Do()
}

type testStruct int

func (t *testStruct) Do() {
	fmt.Println("Doing!!", (*t))
	(*t)++
}

func CreateObj(i IParamer) {
	var t testStruct
	i = &t
}

// interface 的传递也是值传递，但对interface所指变量的修改，是对原变量数据进行修改。
func RunInterface() {
	var i IParamer
	CreateObj(i)
	i.Do()
}

//
// type IParamer interface {
// 	Do()
// }

// type testStruct int

// func (t *testStruct) Do() {
// 	fmt.Println("Doing!!", (*t))
// 	(*t)++
// }

// func CreateObj(i IParamer) {
// 	// var t testStruct
// 	// (*i) = &t
// 	i.Do()
// }

// func RunInterface() {
// 	var t testStruct
// 	var i IParamer = &t
// 	CreateObj(i)
// 	i.Do()
// }
