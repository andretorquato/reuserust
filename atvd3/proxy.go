package main

import (
	"fmt"
)

// ReactiveVariable é a estrutura que implementa o padrão Proxy para monitorar atualizações de uma variável.
type ReactiveVariable struct {
	value     interface{}
	listeners []func(interface{})
}

// new reactive variable
func NewReactiveVariable(initialValue interface{}) *ReactiveVariable {
	return &ReactiveVariable{
		value:     initialValue,
		listeners: []func(interface{}){},
	}
}

// update and notify
func (rv *ReactiveVariable) Set(newValue interface{}) {
	rv.value = newValue
	for _, listener := range rv.listeners {
		listener(newValue)
	}
}

// get current value
func (rv *ReactiveVariable) Get() interface{} {
	return rv.value
}

// add listener
func (rv *ReactiveVariable) AddListener(listener func(interface{})) {
	rv.listeners = append(rv.listeners, listener)
}

func main() {
	// init reactive variable
	client := NewReactiveVariable("John")

	// add listener to watch for changes
	client.AddListener(func(newValue interface{}) {
		fmt.Printf("hey %v, your order has left\n", newValue)
	})

	client.Set("Andre")
	client.Set("Joao")
	client.Set("Maria")

	// current value
	fmt.Printf("current client: %v\n", client.Get())
}
