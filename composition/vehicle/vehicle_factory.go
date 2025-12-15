package main

import (
	"fmt"
)

type Engine struct{}

func (e *Engine) Start() { fmt.Println("Vroom!") }

type Weapon struct{}

func (w *Weapon) Fire() { fmt.Println("Boom!") }

type SportsCar struct {
	Engine
}
type Tank struct {
	Engine
	Weapon
}

func main() {
	sportscar := SportsCar{
		Engine: Engine{},
	}
	sportscar.Start()
	tank := Tank{
		Engine: Engine{},
		Weapon: Weapon{},
	}
	tank.Start()
	tank.Fire()
}
