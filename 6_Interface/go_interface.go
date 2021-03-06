package __Interface

import (
	"fmt"
	"os"
)

func Interfacer() {
	fmt.Println(os.Getwd())
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func IsNil() {
	// var in interface{}
	// 区别
	//    type People interface {
	//        Show()
	//    }
	// 底层结构如下：
	//
	//   type eface struct {      //空接口
	//       _type *_type         //类型信息
	//       data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
	//   }
	//    type iface struct {      //带有方法的接口
	//        tab  *itab           //存储type信息还有结构实现方法的集合
	//       data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
	//   }
	//   type _type struct {
	//          size       uintptr  //类型大小
	//       ptrdata    uintptr  //前缀持有所有指针的内存大小
	//       hash       uint32   //数据hash值
	//       tflag      tflag
	//       align      uint8    //对齐
	//       fieldalign uint8    //嵌入结构体时的对齐
	//       kind       uint8    //kind 有些枚举值kind等于0是无效的
	//       alg        *typeAlg //函数指针数组，类型实现的所有方法
	//       gcdata    *byte
	//       str       nameOff
	//       ptrToThis typeOff
	//   }
	//   type itab struct {
	//       inter  *interfacetype  //接口类型
	//       _type  *_type          //结构类型
	//       link   *itab
	//       bad    int32
	//       inhash int32
	//       fun    [1]uintptr      //可变大小 方法集合
	//   }
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

// func SwitchType()  {
// 	i := GetValue()
//
// 	switch i.(type) {
// 	case int:
// 		println("int")
// 	case string:
// 		println("string")
// 	case interface{}:
// 		println("interface")
// 	default:
// 		println("unknown")
// 	}
//
// }
//
// func GetValue() int {
// 	return 1
// }

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func InterIsNil() {
	var x *int = nil
	Foo(x)
}
