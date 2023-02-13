Домашнее задание к курсу 
GOLANG NINJA - Разработка Веб-Приложений на Go

go build main.go 


Example:

unc main() {
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
