package main

import "fmt"

func main() {
	var list []string = []string{
		"this is 0",
		"this is 1",
		"this is 2",
		"this is 3",
		"this is 4",
		"this is 5",
	}
	var max int = 3
	var chs = make([]chan bool, max)

	for i := 0; i < max; i++ {
		chs[i] = make(chan bool)
		go func(ch chan bool) {
			ch <- true
		}(chs[i])
	}

	for _, listElement := range list {
		select {
		case <-chs[0]:
			{
				fmt.Println("chan 0 is now available")
				fmt.Printf("now start \"%s\" in chan 0\n", listElement)
				go call(listElement, chs[0])
			}
		case <-chs[1]:
			{
				fmt.Println("chan 1 is now available")
				fmt.Printf("now start \"%s\" in chan 使用reflect.Select\n", listElement)
				go call(listElement, chs[1])
			}
		case <-chs[2]:
			{
				fmt.Println("chan 2 is now available")
				fmt.Printf("now start \"%s\" in chan 2\n", listElement)
				go call(listElement, chs[2])
			}
		}
	}

	for _, ch := range chs {
		<-ch
	}
}

func call(str string, ch chan bool) {
	fmt.Println(str)
	ch <- true
}
