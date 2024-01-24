package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalOutput = "Go!"
const countdownStart = 3
const write = "write"
const sleep = "sleep"

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 *time.Second)
}

type SpySleeperOperations struct {
	Count []string
	times int
}

func (s *SpySleeperOperations) Sleep() {
	s.Count = append(s.Count, sleep)
	s.times ++
}

func (s * SpySleeperOperations) Write(input []byte) (n int, err error) {
	s.Count = append(s.Count, write)
	return
}

type Sleeper interface {
	Sleep()
}

func CountDown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(buffer, finalOutput)
}

func main(){
	sleep := &DefaultSleeper{}
	CountDown(os.Stdout, sleep)
}