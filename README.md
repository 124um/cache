Домашнее задание к курсу 
GOLANG NINJA - Разработка Веб-Приложений на Go

go build main.go 


Example:

package main

import (
	"fmt"
	"github.com/124um/cache.git"
)

func main() {
	cache
	newCache := cache.New()
	err := newCache.Set("Boniy", 24)
	printErr(err)

	userId, err := newCache.Get("Boniy")
	printErr(err)

	fmt.Println(userId)

	err = newCache.Delete("Boniy")
	printErr(err)

	userId, err = newCache.Get("Boniy")
	printErr(err)

	fmt.Println(userId)
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
