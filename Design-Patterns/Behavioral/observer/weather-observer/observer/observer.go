// Observer (Subscriber) interface
package observer

// Observer is implemented by all display/subscriber
type Observer interface {
	Update(temp, humidity, pressure float64)
}
