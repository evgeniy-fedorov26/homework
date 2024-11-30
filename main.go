package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите целое число: ")

	n, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %v\n", n)
}
