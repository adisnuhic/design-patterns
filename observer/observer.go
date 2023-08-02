/*
	Consider a weather station application where multiple displays need to show real-time weather updates based on weather changes.
*/

package main

import "fmt"

// Subject represents the weather station that holds the current weather data and maintains a list of observers.
type Subject struct {
	observers []Observer
	weather   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s.weather)
	}
}

func (s *Subject) SetWeather(weather string) {
	s.weather = weather
	s.NotifyObservers()
}

// Observer is the interface that defines the update method.
type Observer interface {
	Update(weather string)
}

// ConcreteObserver represents a display that shows weather updates.
type ConcreteObserver struct {
	displayName string
}

func NewConcreteObserver(displayName string) *ConcreteObserver {
	return &ConcreteObserver{
		displayName: displayName,
	}
}

func (o *ConcreteObserver) Update(weather string) {
	fmt.Printf("Display %s received weather update: %s\n", o.displayName, weather)
}

func main() {
	// Create the weather station subject and observers (displays).
	weatherStation := NewSubject()
	display1 := NewConcreteObserver("Display 1")
	display2 := NewConcreteObserver("Display 2")
	display3 := NewConcreteObserver("Display 3")

	// Register the observers with the weather station.
	weatherStation.RegisterObserver(display1)
	weatherStation.RegisterObserver(display2)
	weatherStation.RegisterObserver(display3)

	// Simulate weather updates and notify observers.
	weatherStation.SetWeather("Sunny")
	weatherStation.SetWeather("Rainy")

	// Unregister an observer (display3) and notify remaining observers.
	weatherStation.RemoveObserver(display3)
	weatherStation.SetWeather("Cloudy")
}
