package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("washVehicle", func(t *testing.T) {
		t.Run("should pre clean and wash a vehcile", func(t *testing.T) {
			car := &Car{id: "1", clean: false}
			motocycle := &Motocycle{id: "2", clean: false}

			err := washVehicle(car)
			if err != nil {
				t.Fatalf("error washing vehcile %s", car.id)
			}

			err = washVehicle(motocycle)
			if err != nil {
				t.Fatalf("error washing vehicle %s", motocycle.id)
			}

			if !car.clean {
				t.Fatal("car is expected to be clean")
			}

			if !motocycle.clean {
				t.Fatal("motocycle is expected to be clean")
			}
		})
	})
}
