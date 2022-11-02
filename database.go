package main

import (
	"log"
	_ "os"
)

type Car struct {
	Mark string
	Country  string
	Price    string
	Year     string
}

const size = 1000

var carList = make([]Car, 0, size)

func addCar(mark, country, price, year string) {
	var car Car
	car.Mark = mark
	car.Country = country
	car.Price = price
	car.Year = year
	carList = append(carList, car)
	log.Println("New car ", car)
}

func getCars(mark string) []Car {
	log.Println("carList =  ", carList)
	var result = make([]Car, 0, size)

	for _, car := range carList {
		if car.Mark == mark || mark == "" {
			result = append(result, car)
		}
	}

	return result
}

func getCarAmount() int {
	return len(carList)
}
