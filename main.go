package main

import (
	"fmt"
	"log"
)

type Vehicle interface {
	PreCleaning() error
	Washing() error
}

type Car struct {
	id string
}

func (c *Car) PreCleaning() error {
	fmt.Printf("Vacuuming interior and removing trash...\n")
	return nil
}

func (c *Car) Washing() error {
	fmt.Printf("Soaping, rinsing and drying exterior...\n")
	return nil
}

type Motocycle struct {
	id string
}

func (m *Motocycle) PreCleaning() error {
	fmt.Printf("Covering exhaust and removing accessories...\n")
	return nil
}

func (m *Motocycle) Washing() error {
	fmt.Printf("Jet spraing, soap cleaning and polishing crome...\n")
	return nil
}

func washVehicle(vehcile Vehicle) error {

	if err := vehcile.PreCleaning(); err != nil {
		return fmt.Errorf("Error pre-cleaning vehicle: %w", err)
	}

	if err := vehcile.Washing(); err != nil {
		return fmt.Errorf("Error washing vehicle: %w", err)
	}

	return nil
}

func main() {
	car := &Car{id: "1"}
	motocycle := &Motocycle{id: "2"}

	err := washVehicle(car)
	if err != nil {
		log.Fatalf("error washing vehcile %s", car.id)
	}

	err = washVehicle(motocycle)
	if err != nil {
		log.Fatalf("error washing vehicle %s", motocycle.id)
	}
}
