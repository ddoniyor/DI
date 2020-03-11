package main

import (
	"di/pkg/di"
	"errors"
	"fmt"
	"log"
)

type Msg string

func NewMsg() Msg {
	return "hello msg"
}

// inversion of control - инверсия управления

// rules: для генерации компонентов, должна быть функция конструктор, возвращающая ровно один результат
// rules: компонент может существовать только в единственном экземпляре (singleton)
type dependency struct{
	value string
}

//
// func NewDependency() [Value(*dependency)] {
// [0].Interface()
func NewDependency(message Msg) *dependency {
	log.Print("dependency created")
	return &dependency{string(message)}
}

func (d *dependency) Start() {
	log.Print("dependency started")
}

func (d *dependency) Stop() {
	log.Print("dependency stopped")
}

type consumer struct {
	dep *dependency
}

func NewConsumer(dep *dependency) *consumer {
	if dep == nil {
		panic(errors.New("dependency can't be nil"))
	}
	log.Print("consumer created")
	return &consumer{dep: dep}
}

// компонент - сущность, которая может иметь зависимости и сама выступать чьей-то зависимостью

// dependency lookup - искать (service locator)
// component должен знать, где и когда искать

// dependency injection - ему подставляют зависимость, если он говорит, что она ему нужна
// component ничего не должен знать (di может выполняться вручную)
// он сам не знает, кто ему зависимость подставит

func main() {
	{
		container := di.NewContainer()
		container.Provide(
			NewMsg,
			NewDependency,
		)
	}
	{
		container := di.NewContainer()
		container.Provide(
			NewMsg,
			NewConsumer,
			NewDependency,
		)

		var component *consumer
		container.Component(&component)
		fmt.Print(component.dep)
	}
}
