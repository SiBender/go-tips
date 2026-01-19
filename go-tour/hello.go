package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var Global = 5

const response = `{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
} `

func main() {
	fmt.Println("Hello World")

	ShowGlobal()
	ShowGlobal()
	ShowGlobal()

	man := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Println("Man %v", string(jsMan)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}

	qwe, _ := ParceResponce(response)
	fmt.Println(qwe.Header)

	//PrintAllFiles("./go-tour")
	//PrintAllFilesFiltered(".", ".go")
}

func ShowGlobal() {
	defer RestoreGlobal(Global)
	Global++
	fmt.Println(Global)
}

func RestoreGlobal(prevValue int) {
	Global = prevValue
}

func ParceResponce(response string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(response), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil

}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesFiltered(path, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFilesFiltered(filename, filter)
		}
	}
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		if predicate(path) {
			fmt.Println(filename)
		}

		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintFilesWithFuncFilter(filename, predicate)
		}
	}
}

// отложенный вызов - defere
func EvaluationOrder() {
	defer fmt.Println("deferred")
	fmt.Println("evaluated")
}

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		Id         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIds []int  `json:"article_ids"`
		}
	} `json:"data"`
}
