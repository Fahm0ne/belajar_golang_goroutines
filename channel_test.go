package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func () {
		time.Sleep(2 *time.Second)
		channel <- "Hello Fahmi "
		fmt.Println("Selesai Mengirim Data Ke channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 *time.Second)

}

func GiveMeRespon (channel chan string) {

	time.Sleep(2 * time.Second)
	channel <- "Muhamad Fahmi Firdaus"

}

func TestChannelasParameter(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go GiveMeRespon(channel)

	data := <- channel
	fmt.Println(data)

}

func OnlyIn (channel chan <- string) {

	time.Sleep(1 * time.Second)
	channel <- "Muhamad Fahmi Firdaus"

}

func OnlyOut (channel <- chan string) {

	data := <- channel
	fmt.Println(data)

}

func TestInOutChannel (t *testing.T) {

channel := make(chan string)
defer close(channel)

go OnlyIn(channel)
go OnlyOut(channel)
time.Sleep(5 * time.Second)

}

func TestBufferedChannel(t *testing.T) {

	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Muhamad"
	channel <- "Fahmi"
	channel <- "Firdaus"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)

	time.Sleep(1 * time.Second)

	fmt.Println("Selesai")

}

func TestChannelRange(t *testing.T) {
	channel := make(chan string)

	go func () {

		for i:=0; i<10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		} 
		close(channel)
	}()

	for data := range channel {
		fmt.Println("channel " , data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make (chan string)
	defer close(channel1)
	defer close (channel2)

go	GiveMeRespon(channel1)
go	GiveMeRespon(channel2)

counter := 0

for {
	select {
	case 		data := <- channel1:
		fmt.Println("Data dari channel 1 = " , data)
		counter++
	case data := <- channel2:
		fmt.Println("Data dari chhanel 2 = ", data)
		counter++
	}
	if counter == 3{
		break
	}
}
}

// func TestSelectDefaultChannel(t *testing.T) {

// 	channel1 := make(chan string)
// 	channel2 := make (chan string)
// 	defer close(channel1)
// 	defer close (channel2)

// go	GiveMeRespon(channel1)
// go	GiveMeRespon(channel2)

// counter := 0

// for {
// 	select {
// 	case 		data := <- channel1:
// 		fmt.Println("Data dari channel 1 = " , data)
// 		counter++
// 	case data := <- channel2:
// 		fmt.Println("Data dari chhanel 2 = ", data)
// 		counter++
// 		default :
// 		fmt.Println("Menunggu Data")
// 	}
// 	if counter == 2{
// 		break
// 	}
// }
// }