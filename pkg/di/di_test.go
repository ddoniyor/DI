package di

import (
	"testing"
)

func TestNewMsg_NoDependencies(t *testing.T) {
	container := NewContainer()
	err := container.Provide(NewMsg)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}


func TestNewMsg_Dependency(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		NewMsg,
		NewDependency)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}

func TestNewMsg_Dep_Consumer(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		NewMsg,
		NewDependency,
		NewConsumer)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}