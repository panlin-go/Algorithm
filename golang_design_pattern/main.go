package main

import (
	"pattern"
)

func main() {

	var battery pattern.RechargeableBattery

	battery = pattern.AdapterNonToYes{pattern.NonRechargeableA{}}
	battery.Use()
	battery.Charge()

	battery = pattern.NonRechargeableB{}
	battery.Use()
	battery.Charge()

}
