package main

import (
	"fmt"
)

type Truck struct {
	id string
}

func processTruck(truck Truck) {
	fmt.Printf("Processing truck: %s\n", truck.id)
}

func mian() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s arrived.\n", truck.id)
		processTruck(truck)
	}
}
