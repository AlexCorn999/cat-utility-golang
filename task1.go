package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func parser(args *[]string, file1 *string, file2 *string, file3 *string) (fileCount int) {
	array := make([]string, 4)
	copy(array, *args)
	for i := 1; i < len(*args); i++ {
		if i == 1 {
			*file1 = array[i]
			fileCount = 1
		} else if i == 2 {
			*file2 = array[i]
			fileCount = 2
		} else if i == 3 {
			*file3 = array[i]
			fileCount = 3
		}

	}
	return
}

func fileOpen(fileCount int, file1 *string, file2 *string, file3 *string) {
	if fileCount == 1 {

		// 1 файл
		file1data, err := ioutil.ReadFile(*file1)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		fmt.Println(string(file1data))

	} else if fileCount == 2 {

		// 1 файл
		file1data, err := ioutil.ReadFile(*file1)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		fmt.Println(string(file1data))

		// 2 файл
		file2data, err := ioutil.ReadFile(*file2)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		fmt.Println(string(file2data))

	} else if fileCount == 3 {

		// 1 файл
		file1data, err := ioutil.ReadFile(*file1)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}

		// 2 файл
		file2data, err := ioutil.ReadFile(*file2)
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}

		// обработка текста
		file1Text := string(file1data)
		file2Text := string(file2data)

		file1Exit := make([]string, 1)
		file2Exit := make([]string, 1)

		file1Exit[0] = file1Text
		file2Exit[0] = file2Text

		// запись в 3 файл
		f3, err := os.Create(*file3)
		writer := bufio.NewWriter(f3)

		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}

		writer.WriteString(strings.Join(file1Exit, ""))
		writer.WriteString("\n")
		writer.WriteString(strings.Join(file2Exit, ""))
		writer.WriteString("\n")
		writer.Flush()
		f3.Close()
	}
}

func main() {
	var file1 string
	var file2 string
	var file3 string

	fileCount := parser(&os.Args, &file1, &file2, &file3)
	fileOpen(fileCount, &file1, &file2, &file3)

}
