package main

import (
	"fmt"
	"log"
	"time"
)

// https://leetcode.com/problems/time-needed-to-buy-tickets/description/

func timeRequiredToBuy2(tickets []int, k int) int {
	defer duration(track("timeRequiredToBuy n"))
	neededTickets := tickets[k]
	seconds := 0

	for i := 0; i < len(tickets); i++ {
		if neededTickets > tickets[i] {
			seconds += tickets[i]
			tickets[i] = 0
		} else if neededTickets < tickets[i] || neededTickets == tickets[i] && i <= k {
			seconds += neededTickets
			tickets[i] -= neededTickets
		} else {
			seconds += neededTickets - 1
			tickets[i] -= neededTickets
		}
		if i > k && tickets[i] > 0 {
			seconds -= 1
		}
	}

	return seconds
}

func timeRequiredToBuy(tickets []int, k int) int {
	defer duration(track("timeRequiredToBuy n^2"))
	n := len(tickets)
	seconds := 0
	for tickets[k] != 0 {
		for i := 0; i < n; i++ {
			if tickets[i] > 0 {
				tickets[i]--
				seconds++
			}
			if tickets[k] == 0 {
				return seconds
			}
		}
	}

	return seconds
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v ns\n", msg, time.Since(start).Seconds())
}

func main() {
	tickets := make([]int, 100000)
	for i := range tickets {
		tickets[i] = 100000
	}
	fmt.Println(timeRequiredToBuy(tickets, 99999))
	for i := range tickets {
		tickets[i] = 100000
	}
	fmt.Println(timeRequiredToBuy2(tickets, 99999))
}
