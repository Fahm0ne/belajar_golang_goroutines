package belajar_golang_goroutines

// func TestPool(t *testing.T) {

// 	pool := sync.Pool{
// 		New: func() interface{} { //new untuk default value
// 			return "New"
// 		},
// 	}

// 	pool.Put("Muhamad")
// 	pool.Put("Fahmi")
// 	pool.Put("Firdaus")

// 	for i:=0; i<100; i++{
// 		go func () {
// 			data:=pool.Get()
// 			fmt.Println(data)
// 			time.Sleep(1 * time.Second)
// 			pool.Put(data)

// 		}()
// 	}

// 	time.Sleep(3 * time.Second)

// }