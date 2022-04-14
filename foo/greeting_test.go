package foo

import (
	"testing"
)

func TestHello(t *testing.T) {
	name := "kj"
	want := Hello(name)
	if "hello, kj" != want {
		t.Fatal("TestHello failed...")
	}
}
