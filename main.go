package main

import (
	"fmt"
)

type Vehicle interface {
	PreCleaning() error
	Washing() error
}

type Car struct {
	id    string
	clean bool
}

func (c *Car) PreCleaning() error {
	fmt.Printf("Vacuuming interior and removing trash...\n")
	return nil
}

func (c *Car) Washing() error {
	fmt.Printf("Soaping, rinsing and drying exterior...\n")
	c.clean = true
	return nil
}

type Motocycle struct {
	id    string
	clean bool
}

func (m *Motocycle) PreCleaning() error {
	fmt.Printf("Covering exhaust and removing accessories...\n")
	return nil
}

func (m *Motocycle) Washing() error {
	fmt.Printf("Jet spraing, soap cleaning and polishing crome...\n")
	m.clean = true
	return nil
}

func washVehicle(vehicle Vehicle) error {

	if err := vehicle.PreCleaning(); err != nil {
		return fmt.Errorf("Error pre-cleaning vehicle: %w", err)
	}

	if err := vehicle.Washing(); err != nil {
		return fmt.Errorf("Error washing vehicle: %w", err)
	}

	return nil
}

func washVehicleConc(vehicles []Vehicle) error {
	return fmt.Errorf("not implemented")
}

func main() {

	vehicles := []Vehicle{
		&Car{id: "1"},
		&Motocycle{id: "2"},
		&Car{id: "3"},
		&Motocycle{id: "4"},
		&Car{id: "5"},
		&Motocycle{id: "6"},
	}

}
