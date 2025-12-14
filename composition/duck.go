package main

import (
  "fmt"
)

type Appearence struct {
  Description string
}

func (a *Appearence) DescribeShape(){
  fmt.Printf(a.Description)
}

type LivingTraits struct {}

func (l *LivingTraits) Quack() {
  fmt.Printf("Quack")
}
func (l *LivingTraits) Eat() {
  fmt.Printf("Eat")
}

type RealDuck struct {
  Appearence
  LivingTraits
}

type RubberDuck struct {
  Appearence
}

func main() {
  d := RealDuck{
       Appearence: Appearence{Description: "Can walk"},
    }
  rd := RubberDuck{
        Appearence: Appearence{Description: "Cannot walk"},
    }
  d.Quack()
  d.Eat()
  d.DescribeShape()
  rd.DescribeShape()
}
