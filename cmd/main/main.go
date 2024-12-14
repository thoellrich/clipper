package main

import (
	"fmt"
	"github.com/thoellrich/clipper"
)

func main() {
	c := clipper.NewClipboardMonitor()
	fmt.Println(c.Foo())
}
