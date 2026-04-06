package subject

import "weather-observer/observer"

// WeatherStation is a concrete subject (Publisher) that implements the Subject interface
// It maintains a list of observers and notifies them of any state changes
type WeatherStation struct {
	observers   []observer.Observer
	temperature float64
	humidity    float64
	pressure    float64
}

// Constructor
func NewWeatherStation() *WeatherStation {
	return &WeatherStation{}
}

// RegisterObserver adds an observer to the list
func (w *WeatherStation) RegisterObserver(o observer.Observer) {
	w.observers = append(w.observers, o)
}

// RemoveObserver removes an observer from the list
func (w *WeatherStation) RemoveObserver(o observer.Observer) {
	for i, obs := range w.observers {
		if obs == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers notifies all registered observers of a change in weather data
func (w *WeatherStation) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

// SetMeasurements updates the weather data and notifies observers of the change
func (w *WeatherStation) SetMeasurements(temp, humidity, pressure float64) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.NotifyObservers()
}
