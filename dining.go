package main

import "fmt"
import "sync"
import "time"

var permission chan bool
var finished chan bool
var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(num int) {
	for i := 0; i < 3; i++ {
		<-permission
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat", num+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("finishing eating", num+1)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		finished <- true
	}
	fmt.Println("philosopher full ", num+1)
	wg.Done()
}

func host(max_eaters int) {
	cur_eaters := 2
	permission <- true
	permission <- true
	for range finished {
		cur_eaters--
		if cur_eaters < max_eaters {
			cur_eaters++
			permission <- true
		}
	}
	fmt.Println("dinner party over")
	wg.Done()
}

func main() {
	permission = make(chan bool, 2)
	finished = make(chan bool)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(i)
	}
	go host(2)
	wg.Wait()
	close(finished)
	wg.Add(1)
	wg.Wait()
}
