package foo

import (
	"math/rand"
	"time"
)

func RandomHello(name string) string {
	l := []string{"你好", "Bonjour", "こんにちは"}
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(l))
	return l[i] + ", " + name
}
