package godemo

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("[v3]hello %v", name)
}
