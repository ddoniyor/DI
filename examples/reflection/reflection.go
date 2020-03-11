package main

import (
	"fmt"
	"log"
	"reflect"
)

// my - junior marker (noob)
type myint int

func sample() {
	log.Print("called")
}

func sample2(msg string) {
	log.Print("called " + msg)
}

func sample3(msg string) string {
	return fmt.Sprintf("called %s", msg)
}

func callMe(c interface{}) {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		return
	}
	// we can call it
	if reflectType.NumIn() != 0 {
		return
	}
	// we can call it without params
	reflectValue := reflect.ValueOf(c)
	reflectValue.Call(nil)
}

func callMeAdvanced(c interface{}, args ...interface{}) {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		return
	}
	// we can call it
	reflectValue := reflect.ValueOf(c)
	// create reflect params for it
	valueArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		valueArgs[i] = reflect.ValueOf(arg)
	}
	// we can call it with
	reflectValue.Call(valueArgs)
}

func callMeAdvancedWithResult(c interface{}, args ...interface{}) interface{} {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		panic("we can't do it")
	}
	// we can call it
	reflectValue := reflect.ValueOf(c)
	// create reflect params for it
	valueArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		valueArgs[i] = reflect.ValueOf(arg)
	}
	// we can call it with
	resultValue := reflectValue.Call(valueArgs)[0]
	return resultValue.Interface()
}

func main() {
	// Reflection -> самоанализ
	var count myint = 10

	// Type - описывает тип
	// Value - описывает значение
	// Fabric'а типовые объекты: штампует эти самые объекты из каких-то ресурсов
	// Fabric-функция
	reflectType := reflect.TypeOf(count)
	// Kind() - набор констант
	log.Print(reflectType.Kind() == reflect.Int)

	reflectValue := reflect.ValueOf(&count)
	// Elem() - "берём из указателя элемент"
	reflectElem := reflectValue.Elem()
	log.Print(reflectValue)
	// Value <-> Type
	reflectElem.SetInt(11)
	log.Print(count)

	callMe(sample)
	callMeAdvanced(sample2, "it works!")
	result := callMeAdvancedWithResult(sample3, "Cosmos!!!")
	log.Print(result)
}
