package gogreeting

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("hello, %v", name)
}
