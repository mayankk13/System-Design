package main

import "fmt"

// Abstraction: defines what a message sender should do
type MessageSender interface {
	Send(message string)
}

// Low-level: EmailSender implements MessageSender
type EmailSender struct{}

func (e *EmailSender) Send(message string) {
	fmt.Println("📧 Email sent:", message)
}

// Low-level: SMSSender also implements MessageSender
type SMSSender struct{}

func (s *SMSSender) Send(message string) {
	fmt.Println("📱 SMS sent:", message)
}

// High-level: now depends on interface, not concrete types
type NotificationService struct {
	sender MessageSender
}

func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{sender: sender}
}

func (n *NotificationService) Notify(message string) {
	n.sender.Send(message)
}

func main() {
	email := &EmailSender{}
	sms := &SMSSender{}

	emailNotifier := NewNotificationService(email)
	smsNotifier := NewNotificationService(sms)

	emailNotifier.Notify("Your account was credited ₹500")
	smsNotifier.Notify("Your OTP is 123456")
}
