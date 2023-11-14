package main

import (
	"fmt"
	"html"
	"strconv"
	"sync"
	"time"
)

const (
	FOOD   = "0x0001F35C"
	FINISH = "0x0001F44C"
)

type Philosopher struct {
	ID             int
	Name           string
	LeftChopStick  *ChopStick
	RightChopStick *ChopStick
	Host           *Host
}

func (p *Philosopher) Eat(wg *sync.WaitGroup) {
	wg.Add(1)

	p.LeftChopStick.Lock()
	p.RightChopStick.Lock()

	p.Host.requestChannel <- p
	fmt.Println(p.Name + " is eating " + GetEmoticon(FOOD))
	time.Sleep(time.Millisecond)

	p.LeftChopStick.Unlock()
	p.RightChopStick.Unlock()
}

func GetEmoticon(value string) string {
	i, _ := strconv.ParseInt(value, 0, 64)
	foodEmoticon := html.UnescapeString(string(i))
	return foodEmoticon
}
