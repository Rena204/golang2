package main

import (
	"fmt"
	"io/ioutil"
)

func main2() {
	fileData, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	if len(fileData) == 0 {
		fmt.Println("Файл пуст")
		return
	}

	fmt.Println(string(fileData))
}
