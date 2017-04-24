package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func swap(a, b string) (string, string) {
	return b, a
}

func multiply(a int, b float32) float32 {
	if v := 5; v > a {
		return float32(0) * b
	}
	return float32(a) * b

}

func get(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

}

func deferF() {
	defer fmt.Println("Cool!")
	fmt.Println("Uncool")
}

func modifyglob(pointer *int) {
	fmt.Println(&pointer)
	*pointer = 50
}

func deferedLoop(length int) {
	for i := 0; i < length; i++ {
		// defer fmt.Println(i)
	}
}

func printer(name string) string {
	fmt.Println("Your name is", name, "and you are ", rand.Intn(10), "years old")
	return name
}

// type Gamer struct {
// 	game  string
// 	level string
// }

// type Person struct {
// 	name string
// 	Gamer
// }

func main() {
	printer("aleksandar")
	fmt.Println(multiply(6, 22.4))

	x, y := swap("c", "d")
	glob := 11
	modifyglob(&glob)

	fmt.Println("Glob is ", glob)

	fmt.Println(x, y)

	deferF()

	deferedLoop(10)

	slice := make([]int, 5)

	for i := 0; i < 5; i++ {
		slice[i] = i * 3
	}

	slice = append(slice, 1)
	slice = append(slice, 2)

	fmt.Println(slice)

	// get("https://jsonplaceholder.typicode.com/users")

	b := 5

	fmt.Printf("%d is the best number ever", b)

}
