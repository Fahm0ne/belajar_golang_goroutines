package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {

	fmt.Println("Hello World")
}

func TestCreateGoroutines (t *testing.T) {
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

//tes performa gourutines
func DisplayNumber (number int) {
	fmt.Println("Display = " , number)
}

func TestManyGourutines (t *testing.T) {

	for i:=0; i<100000; i++ {
	DisplayNumber(i)
	}

	//time.Sleep(5 * time.Second)
	//dengan gourutines 31.762
	//tanpa gourutines 30.45

}

