package main 

import (
	"fmt"
	"gopkg.in/qml.v1"
	"os"
	"strconv"
)

const FIZZ = 3
const BUZZ = 5
func main() {	
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}


type FizzBuzz struct {
	Number int
	Status string
}

func (f *FizzBuzz) SetNumber(number string) {
	fmt.Println("Received (string): " + number)
	num , err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Error converting number: ", number)
	}
	f.Status = fizzBuzz(num)
	fmt.Println("f.Status: " + f.Status)
	qml.Changed(f, &f.Status)
}
func run() error {
	engine := qml.NewEngine()

	engine.On("quit", func() { os.Exit(0) })

	number := 1
	context := engine.Context()
	//the var is set as the string
	context.SetVar("fizzbuzz", &FizzBuzz{Number: number, Status: "welcome"})

	component, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}
	window := component.CreateWindow(nil)
	window.Show()
	window.Wait()
	return nil
}

func fizzBuzz(num int) string {
	if (num % FIZZ == 0 && num % BUZZ == 0) {
		return "fizz buzz!"
	}
	if (num % FIZZ == 0) {
		return "fizz!"
	}
	if (num % BUZZ == 0) {
		return "buzz!"
	}
	return strconv.Itoa(num)
		
}