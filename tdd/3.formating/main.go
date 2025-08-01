package main

import (
	"fmt"
)

func FormatValue(value any, format string, label string) string {
	return fmt.Sprintf("%s: "+format, label, value)
}

func main() {
	fmt.Println(FormatValue(42, "%d", "Decimal"))
	fmt.Println(FormatValue(42, "%b", "Binary"))
	fmt.Println(FormatValue(42, "%x", "Hex"))
	fmt.Println(FormatValue("hello", "%q", "Quoted"))
	fmt.Println(FormatValue("hello", "%T", "Type"))
	fmt.Println(FormatValue(3.14159, "%.2f", "Float"))
}
