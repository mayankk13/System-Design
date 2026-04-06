package main

import (
	"weather-observer/observer"
	"weather-observer/subject"
)

func main() {
	// 1. Create a WeatherStation (Subject)
	weatherStation := subject.NewWeatherStation()

	// 2. Create Observers (Displays/Subscribers)
	currentDisplay := observer.NewCurrentConditionDisplay()
	statisticsDisplay := observer.NewStatisticsDisplay()

	// 3. Register Observers with the WeatherStation
	weatherStation.RegisterObserver(currentDisplay)
	weatherStation.RegisterObserver(statisticsDisplay)

	// 4. Simulate new weather measurements
	weatherStation.SetMeasurements(25.0, 65.0, 1013.0)
	weatherStation.SetMeasurements(27.0, 70.0, 1012.0)
	weatherStation.SetMeasurements(22.0, 60.0, 1014.0)

	// 5. Remove an observer and simulate more measurements
	weatherStation.RemoveObserver(currentDisplay)
	weatherStation.SetMeasurements(24.0, 68.0, 1011.0)
}
