package belajar_golang_goroutines

//  var Lo = sync.Mutex{}
// var cond = sync.Cond{&Lo}
// var group = sync.WaitGroup{}

// func WaitCondition (value int) {
// 	defer group.Done()
// 	group.Add(1)

// 	cond.Wait()
// 	fmt.Println("Done" , value)
// 	cond.L.Lock()
// 	cond.L.Unlock()
// }

// func TestCond(t *testing.T) {
// for i:=0; i<10; i++{
// 	go WaitCondition(i)
// }

// go func () {
// 	for i:=0; i<10; i++ {
// 		time.Sleep(1* time.Second)
// 		cond.Broadcast()
// 	}
// }()

// group.Wait()

// }