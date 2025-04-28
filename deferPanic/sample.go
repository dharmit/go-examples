/* This program demonstrates how defer is called upon panic;
point being that panic traces back the pending defer calls, runs them, and then panics
*/

package main

import "fmt"

func main() {
	fmt.Println("main routine")
	funcOne()
}

func funcOne() {
	defer fmt.Println("exiting funcOne")
	fmt.Println("started funcOne")
	funcTwo()
}

func funcTwo() {
	defer fmt.Println("exiting funcTwo")
	fmt.Println("started funcTwo")
	funcThree()
}

func funcThree() {
	defer fmt.Println("exiting funcThree")
	fmt.Println("started funcThree")
	panic("!! PANICKING !!")
}
