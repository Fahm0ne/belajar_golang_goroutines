package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func TestRunAsychronus(t *testing.T) {

	group := &sync.WaitGroup{}

	for i:=0; i<1000; i++ {
	go	RunAsychronus(group)
	}

	group.Wait()
	fmt.Println("complite")


}

func RunAsychronus(group *sync.WaitGroup) {

	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")


}
