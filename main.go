package main

import "fmt"

type controller interface {
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed 		int
	enginePower int
}
type bycyle struct {
	speed  	   int 
	humanPower int
}
func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}

func (b *bycyle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bycyle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}
func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func main() {
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	b := &bycyle{0, 5}
	speedUpAndDown(b)
}


