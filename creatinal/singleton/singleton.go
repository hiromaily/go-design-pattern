package singleton

import (
	"fmt"
	"time"
)

//https://gobyexample.com/stateful-goroutines

type command int

type singleton struct {
	uniqueNumber int
	ticker       *time.Ticker
	Controller   chan command
	Receiver     chan int
}

// command list
const (
	Add command = iota + 1
	Reduce
	ReceiveNumber
	Stop
)

var instance *singleton

const span = 3 * time.Second

func init() {
	instance = new(singleton)
	instance.Controller = make(chan command)
	instance.Receiver = make(chan int)

	startTicker()
}

func startTicker() {
	go func() {
		instance.ticker = time.NewTicker(span)
		// receiver
		for {
			select {
			case <-instance.ticker.C: //ticker
				// add
				instance.add()
			case cmd := <-instance.Controller:
				switch cmd {
				case 0: //when close channel
					//after close channel
					fmt.Println("[Debug]channel was closed.")
					//`return` is better than `break` in this case
					return
				case Add:
					instance.add()
				case Reduce:
					instance.reduce()
				case ReceiveNumber:
					instance.Receiver <- instance.uniqueNumber
				case Stop:
					instance.stop()
				default:
					fmt.Println("[Debug]this code should not be run.")
				}
			}
		}
	}()
}

func Get() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

//add number
func (s *singleton) add() {
	fmt.Println(" -> add")
	s.uniqueNumber++
}

//reduce number
func (s *singleton) reduce() {
	fmt.Println(" -> reduce")
	s.uniqueNumber--
}

//stop timer
func (s *singleton) stop() {
	fmt.Println(" -> stop")
	s.ticker.Stop()
}

func (s *singleton) GetNumber() int {
	s.Controller <- ReceiveNumber //get number
	return <-s.Receiver
}

func (s *singleton) AddNumber() {
	s.Controller <- Add
}

func (s *singleton) ReduceNumber() {
	s.Controller <- Reduce
}
