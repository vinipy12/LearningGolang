package main

import (
	"errors"
	"fmt"
	"log"
)

type Car struct {
	id string
}

func (c *Car) PreCleaning() error {
	return errors.New("not implemented")
}

func (c *Car) Washing() error {
	return errors.New("not implemented")
}

// processCar handles all the washing steps.

func processCar(car Car) error {
	fmt.Printf("Processing car: %s\n", car.id)

	if err := car.PreCleaning(); err != nil {
		return fmt.Errorf("Error pre-cleaning car: %w", err)
	}

	if err := car.Washing(); err != nil {
		return fmt.Errorf("Error washing car: %w", err)
	}

	return nil
}

func main() {
	cars := []Car{
		{id: "Car A"},
		{id: "Car B"},
		{id: "Car C"},
	}

	for _, car := range cars {
		fmt.Printf("Car %s arrived.\n", car.id)

		/*err := processCar(car)
		if err != nil {
			log.Fatalf("Error processing car: %s", err)
		}*/

		if err := processCar(car); err != nil {
			log.Fatalf("Error processing car: %s", err)
		}
	}
}
