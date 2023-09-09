package main

import (
	"fmt"
	"time"
)

type barberShop struct {
	shopCapacity    int
	hairCutDuration time.Duration
	numberOfBarbers int
	barbersDoneChan chan bool
	clientsChan     chan string
	open            bool
}

func (shop *barberShop) addBarber(barber string) {
	shop.numberOfBarbers++

	go func() {
		isSleeping := false
		fmt.Printf("%s goes to the waiting room to check fo clients.\n", barber)
		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.clientsChan) == 0 {
				fmt.Printf("There is nothing to do, so %s takes a nap.\n", barber)
				isSleeping = true
			}

			client, isShopOpen := <-shop.clientsChan
			if isShopOpen {
				if isSleeping {
					fmt.Printf("%s wakes %s up.\n", client, barber)
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home
				shop.sendBarberHome(barber)
				// closing the goroutine
				return
			}
		}
	}()
}

func (shop *barberShop) cutHair(barber, client string) {
	fmt.Printf("%s is cutting %s's hair.\n", barber, client)
	time.Sleep(shop.hairCutDuration)
	fmt.Printf("%s is finished cutting %s's hair.\n", barber, client)
}

func (shop *barberShop) sendBarberHome(barber string) {
	fmt.Printf("%s is going home.\n", barber)
	shop.barbersDoneChan <- true
}

func (shop *barberShop) closeShopForDay() {
	fmt.Println("Closing shop for the day.")
	close(shop.clientsChan)
	shop.open = false

	//wait until the barbers are done
	for a := 1; a <= shop.numberOfBarbers; a++ {
		// block until every single barber is done
		// received value ignored
		<-shop.barbersDoneChan
	}

	close(shop.barbersDoneChan)

	fmt.Println("-----------------------------")
	fmt.Println("The barbershop is now closed.")
}

func (shop *barberShop) addClient(client string) {
	// print out a message
	fmt.Printf("client %s arrives!", client)

	if shop.open {
		select {
		case shop.clientsChan <- client:
			fmt.Printf("%s takes a seat in the waiting room!\n", client)
		default:
			fmt.Printf("The waiting room is full, so %s leaves!\n", client)
		}
	} else {
		fmt.Printf("The shop is already closed, so %s leaves!\n", client)
	}
}
