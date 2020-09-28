package base

import (
	"bytes"
	"io"
	"os"
	"testing"
)

/*
nil在Go语言中只能被赋值给指针和接口。接口在底层的实现有两部分：type和data
（参考上一篇博客：https://blog.csdn.net/weixin_42117918/article/details/90372113）。
在源码中，显示地将nil赋值给接口时，接口的type和data都将为nil。此时，接口与nil值判断是相等的。
但如果将一个带有类型的nil赋值给接口时，只有data为nil，而type不为nil，此时，接口与nil判断将不相等。
*/

type book struct {
	name string
}

func Test_outSome(t *testing.T) {
	var buf *bytes.Buffer
	outSome(buf)
}

func outSome(out io.Writer) {
	var b book
	//所以此处判断为true，因为此时data为nil，但是type不为nil
	if out != nil {
		if &b == nil {
			os.Exit(-2)
		}

		os.Exit(-1)
	}
}
