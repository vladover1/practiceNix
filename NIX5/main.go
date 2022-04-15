package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var URL = os.Args[1]

func main() {
	wg := &sync.WaitGroup{}

	for id := 1; id <= 10; id++ {
		wg.Add(1)
		go getUserById(id, wg)
	}

	wg.Wait()

}

func getUserById(id int, wg *sync.WaitGroup) {
	resp, err := http.Get(URL + strconv.Itoa(id)) // отправили запрос
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(fmt.Sprintf("storage/posts/%d.txt", id)) // создаем файлы
	if err != nil {
		fmt.Println(err)
	}

	content, _ := ioutil.ReadAll(resp.Body) // получили ответ
	s := strings.TrimSpace(string(content))

	file.WriteString(s) //записали в тхт файл

	wg.Done()

}
