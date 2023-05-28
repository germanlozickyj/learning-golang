package main

import "fmt"

//SMS, MAIL

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type InotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

//SMS
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Enviando notificacion via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}


//Email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando notificacion via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "mailgun"
}

func getNotificationFactory (notificationType string) (InotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f InotificationFactory) {
	f.SendNotification()
}

func getMethod(f InotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}