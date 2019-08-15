package main

import (
	"fmt"
	"os"
)

// Check is a function.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func test1() *os.File {
	f, err := os.Open("/tmp/test.txt")
	Check(err)
	return f
}

func test2() {
	f, err := os.Open("/tmp/test.txt") //ISSUE
	Check(err)
	//defer f.Close()
	b := make([]byte, 5)
	n, err := f.Read(b)
	Check(err)
	fmt.Printf("%d bytes: %s\n", n, string(b))
}

func test3() (f *os.File) {
	f, err := os.Open("/tmp/test.txt")
	Check(err)
	return
}

func main() {

	f1, err := os.Open("/tmp/test.txt") //ISSUE
	Check(err)
	//defer f1.Close()
	b1 := make([]byte, 5)
	n1, err := f1.Read(b1)
	Check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	f2, err := os.Open("/tmp/test.txt")
	Check(err)
	defer f2.Close()
	b2 := make([]byte, 5)
	n2, err := f2.Read(b2)
	Check(err)
	fmt.Printf("%d bytes: %s\n", n2, string(b2))
	f2.Close()

	f3, err := os.OpenFile("/tmp/test.txt", os.O_RDWR, 0644) //ISSUE
	Check(err)
	//defer f3.Close()
	b3 := make([]byte, 5)
	n3, err := f3.Read(b3)
	Check(err)
	fmt.Printf("%d bytes: %s\n", n3, string(b3))

}
