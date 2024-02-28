package main

import (
	"fmt"
	"os"
)

func main3() {
	file, err := os.Create("readonly.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	err = file.Chmod(0400) // Установка режима только для чтения
	if err != nil {
		fmt.Println("Ошибка при установке режима файла:", err)
		return
	}

	file.Close()

	file, err = os.Open("readonly.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println(string(data))
}
