package note

import "os"

/*
每一个可独立运行的Go程序，必定包含一个package main，
在这个main包中必定包含一个入口函数main，而这个函数既没有参数，也没有返回值。
*/

const Des string = "记录学习之路中的一些收获"

/*
Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。
其中rune是int32的别称，byte是uint8的别称。
浮点数的类型有float32和float64两种（没有float类型），默认是float64。
*/

/*
传指针使得多个函数能操作同一个对象。
传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传递的话,
在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的时候，用指针是一个明智的选择。
Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。
（注：若函数需改变slice的长度，则仍需要取地址传递指针）
*/

/*
函数也可以作为变量，类似 lambda ，函数式编程
*/
type testInt func(int) bool // 声明了一个函数类型

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

/*
Panic

是一个内建函数，可以中断原有的控制流程

Recover

是一个内建的函数，可以让进入panic状态的goroutine恢复过来。recover仅在延迟函数中有效。
*/
var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
