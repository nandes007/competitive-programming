package main

import "fmt"

func strReverse(s string) string {
	var temp string

	fmt.Println(len(s))
	// for i := len(s) - 1; i >= 0; i-- {
	// 	fmt.Println(string([]rune(s)[i]))
	// 	temp += string([]rune(s)[i])
	// }

	return temp
}

func strReverseV2(s string) string {
	runes := []rune(s)
	var temp string
	fmt.Println(runes)
	fmt.Println(len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		temp += string(runes[i])
	}

	return temp
}

func fizzBuzz() {
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	// fmt.Println(strReverse("héllo"))
	// fmt.Println(strReverseV2("hello"))
	fmt.Println(strReverseV2("héllo"))
	fizzBuzz()
}
