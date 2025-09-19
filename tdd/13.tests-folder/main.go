package hellotdd

import (
	"fmt"
)

// test comment in hello word example
func Hello(name string) string {

	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(Hello("Werld"))
}
