/*
High-level modules should not depend on low-level modules.
Both should depend on abstractions.

*/

package main

import (
	"fmt"
)

// low-level: EmailSender
type EmailSender struct{}

func (e *EmailSender) Send(message string) {
	fmt.Println("📧 Email sent:", message)
}

// High-level: NotificationService depends on EmailSender directly
type NotificationService struct {
	email *EmailSender
}

func (n *NotificationService) Notify(message string) {
	n.email.Send(message)
}

func main() {
	emailSender := &EmailSender{}
	notifier := NotificationService{email: emailSender}

	notifier.Notify("Your balance is low")
}

/*
What’s Wrong?
	•	NotificationService is locked to EmailSender.
	•	You can’t switch to SMS, push notification, or anything else without changing NotificationService.
*/
