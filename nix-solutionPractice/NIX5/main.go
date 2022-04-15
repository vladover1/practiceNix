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

const (
	path  = "storage/posts/%d.txt"
	count = 10
)

func main() {
	wg := &sync.WaitGroup{}

	for UserId := 1; UserId <= count; UserId++ {
		wg.Add(1)
		go getUserById(UserId, wg)
	}

	wg.Wait()
}

func getUserById(UserId int, wg *sync.WaitGroup) {
	resp, err := http.Get(URL + strconv.Itoa(UserId)) // отправили запрос
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(fmt.Sprintf(path, UserId)) // создаем файлы
	if err != nil {
		fmt.Println(err)
	}

	content, err := ioutil.ReadAll(resp.Body) // получили ответ
	if err != nil {
		fmt.Println(err)
	}
	s := strings.TrimSpace(string(content))

	_ , err = file.WriteString(s) //записали в тхт файл
	if err != nil {
		fmt.Println(err)
	}

	wg.Done()
}
