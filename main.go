package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}

		line := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), input)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		return
	}
}
