package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type income struct {
	source string
	amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: %d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []income{
		{source: "Main job", amount: 500},
		{source: "Gifts", amount: 10},
		{source: "Part time job", amount: 50},
		{source: "Investments", amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, in := range incomes {
		go func(i int, in income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += in.amount
				bankBalance = temp
				balance.Unlock()
				fmt.Printf("One week %d, you earned $%d.00 from %s\n", week, in.amount, in.source)
			}
		}(i, in)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
}
