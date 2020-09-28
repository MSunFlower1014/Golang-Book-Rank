package base

import (
	"fmt"
	"io"
)

func OutSome(out io.Writer) {
	if out != nil {
		fmt.Print("not nil ")
	}
}
