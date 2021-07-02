package hello_go

import "fmt"

// 打招呼函数
func Greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
