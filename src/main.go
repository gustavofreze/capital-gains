package main

import "capital-gains/src/starter"

func main() {
	dependencies := starter.NewDependencies()
	dependencies.CalculateCapitalGain.Handle()
}
