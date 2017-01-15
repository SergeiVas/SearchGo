// SearchGo project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var(
		countAddres int	 // Количество url'ов
		AllUrl string	 // Строка со всеми Url'ами
	//	gorutines int
	//	k = 5
	)
	ch := make(chan string) // Канал, в котором хранятся url
	countUrls := make(chan int) // Канал, в котором хранятся количство строк go на url'е
	fmt.Scanln(&AllUrl)
	go func(){
		countAddres = readData(ch, AllUrl)
	}()
	for num := range ch {
			countStrings(num, countUrls)
	}
	count := 0
	// Считываем количество строк go с каждого url'а
	for i := range countUrls{
			count += i
			countAddres--
			if countAddres <= 0{
				break
			}
		}
	fmt.Print("Total: ")
	fmt.Println(count)

}

//Переходим по url'у и считаем кол - во искомых строк
func countStrings(url string, c chan int) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("You can not go to url")
	} else {
		defer response.Body.Close()
		s, _ := ioutil.ReadAll(response.Body)
		count := strings.Count(string(s), "Go")
		fmt.Print("Count for "+string(url)+":")
		fmt.Println(count)
		c <- count
	}
}

// Разбиваем строку на отдельные url'ы и отправляем их в канал
func readData(c chan string, AllURL string) int{
	s := ""
	countAddres := 0
	// Считываем посимвольно строку пока не встретим "/n" и отправляем в канал
	for i := 0; i < len(AllURL); i++ {
		if AllURL[i] != 92 {
			s += string(AllURL[i])
		} else {
			c <- string(s)
			s = ""
			i++
			countAddres++
		}
	}
	// В конце "/n" не встречается, поэтому просто отправляем url в канал
	if len(s) > 5 {
		c <- string(s)
		close(c)
		countAddres++
	}
	return countAddres
}
