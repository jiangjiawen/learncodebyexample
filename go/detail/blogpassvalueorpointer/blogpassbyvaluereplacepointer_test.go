// cmd: GO111MODULE=auto go test -bench=.
// https://medium.com/better-programming/why-you-should-avoid-pointers-in-go-36724365a2a7
package test

import (
	"fmt"
	"testing"
)

type CoffeeMachine struct {
	UID                 string
	Description         string
	NumberOfCoffeeBeans int
}

func NewCoffeeMachinePointer() *CoffeeMachine {
	return &CoffeeMachine{}
}

func (cm *CoffeeMachine) SetUIDPointer(uid string) {
	cm.UID = uid
}

func (cm *CoffeeMachine) SetDescriptionPointer(description string) {
	cm.Description = description
}

func (cm *CoffeeMachine) SetNumberOfCoffeeBeansPointer(n int) {
	cm.NumberOfCoffeeBeans = n
}

func NewCoffeeMachineValue() CoffeeMachine {
	return CoffeeMachine{}
}

func (cm CoffeeMachine) SetUIDValue(uid string) CoffeeMachine {
	cm.UID = uid
	return cm
}

func (cm CoffeeMachine) SetDescriptionValue(description string) CoffeeMachine {
	cm.Description = description
	return cm
}

func (cm CoffeeMachine) SetNumberOfCoffeeBeansValue(n int) CoffeeMachine {
	cm.NumberOfCoffeeBeans = n
	return cm
}

func BenchmarkWithPointer(b *testing.B) {
	cm := NewCoffeeMachinePointer()
	for i := 0; i < b.N; i++ {
		cm.SetUIDPointer(fmt.Sprintf("random generate uid %d", i))
		cm.SetNumberOfCoffeeBeansPointer(i)
		cm.SetDescriptionPointer(fmt.Sprintf("This is the best coffe machine that is around! This is version %d", i))
	}
}

func BenchmarkWithValue(b *testing.B) {
	cm := NewCoffeeMachineValue()
	for i := 0; i < b.N; i++ {
		cm = cm.SetUIDValue(fmt.Sprintf("random generate uid %d", i))
		cm = cm.SetNumberOfCoffeeBeansValue(i)
		cm = cm.SetDescriptionValue(fmt.Sprintf("This is the best coffe machine that is around! This is version %d", i))
	}
}
