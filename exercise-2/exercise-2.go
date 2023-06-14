package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder
	builder.WriteRune('a')
	builder.WriteRune('b')
	builder.WriteRune('c')

	builder.WriteString("defg")
	str := builder.String()
	fmt.Println(str)
}
