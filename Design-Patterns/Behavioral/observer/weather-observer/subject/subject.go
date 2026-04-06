package subject

import "weather-observer/observer"

// Subject defines behavior of the publisher
// It maintains a list of observers and notifies them of any state changes
type Subject interface {
	RegisterObserver(o observer.Observer)
	RemoveObserver(o observer.Observer)
	NotifyObservers()
}
