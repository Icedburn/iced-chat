package parallel_reading

import (
	"math/rand"
	"strings"
	"time"
)

func Read(word string) string {
		var ch chan string = make(chan string, 20)
		var teste []string = strings.Split(word, "")
		for _,v := range teste {
			go writeLetter(v, ch)
		}
		var result string = ""
		for i:=0; i< len(teste); i++{
			el, isOpen := <- ch
			if (isOpen) {
				result = result + el
			} else {
				break
			}
		}
		return result
}

func writeLetter(c string, ch chan string) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	ch <- c
}
