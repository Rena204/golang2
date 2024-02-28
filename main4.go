package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main4() {
	input := "продам гараж"
	line := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), input)

	err := ioutil.WriteFile("output.txt", []byte(line), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

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
