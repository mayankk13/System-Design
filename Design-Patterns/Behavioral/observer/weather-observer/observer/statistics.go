package observer

import "fmt"

// StatisticsDisplay is a concrete observer that implements the observer interface
// It calculates and displays the average temperature
type StatisticsDisplay struct {
	temperature []float64
}

// Constructor
func NewStatisticsDisplay() *StatisticsDisplay {
	return &StatisticsDisplay{}
}

// Update is called by the subject (WeatherStation) when there is a change in weather data
func (s *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	s.temperature = append(s.temperature, temp)
}

// Internal method to calculate and display the average temperature
func (s *StatisticsDisplay) display() {
	sum := 0.0
	for _, temp := range s.temperature {
		sum += temp
	}

	// Calculate average temperature
	avg := sum / float64(len(s.temperature))

	fmt.Printf("Avg. Temperature: %.1f°C\n", avg)
}
