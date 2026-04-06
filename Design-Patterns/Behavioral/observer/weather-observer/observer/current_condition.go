package observer

import "fmt"

// CurrentConditionDisplay is a concrete observer that implements the observer interface
// It dispalys latest temperature & humidity
type CurrentConditionDisplay struct {
	temperature float64
	humidity    float64
}

// constructor
func NewCurrentConditionDisplay() *CurrentConditionDisplay {
	return &CurrentConditionDisplay{}
}

// Update is called by the subject (WeatherStation) when there is a change in weather data
func (c *CurrentConditionDisplay) Update(temp, humidity, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.display()
}

// Internal method to display the current conditions
func (c *CurrentConditionDisplay) display() {
	fmt.Printf("Current condition: %.1f°C and %.1f%% humidity\n", c.temperature, c.humidity)
}
