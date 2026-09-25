package notify

import (
	"errors"
	"fmt"
	"log/slog"
)

type Sender interface {
	Name() string
	Send(n Notification) error
}

type EmailSender struct {
	Host string
	From string
}

var _ Sender = (*EmailSender)(nil)

func (EmailSender) Name() string {
	return "Email"
}

func (es EmailSender) Send(n Notification) error {
	fmt.Println("TODO: Email Sending Logic")
	return nil
}

type SmsSender struct {
	Gateway string
}

var _ Sender = (*SmsSender)(nil)

func (SmsSender) Name() string {
	return "Sms"
}

func (ss SmsSender) Send(n Notification) error {
	fmt.Println("TODO: Sms Sending Logic")
	return nil
}

type LoggingSender struct {
	Sender
	Logger slog.Logger
}

var _ Sender = (*LoggingSender)(nil)

func NewLoggingSender(s Sender, l slog.Logger) (LoggingSender, error) {
	if s == nil {
		return LoggingSender{}, errors.New("Sender can't be nil")
	}

	return LoggingSender{Sender: s, Logger: l}, nil
}

func (ls LoggingSender) Send(n Notification) error {
	senderName := ls.Sender.Name()
	ls.Logger.Info(fmt.Sprintf("Sending %s notification", senderName), "notification", n)

	if err := ls.Sender.Send(n); err != nil {
		ls.Logger.Error(fmt.Sprintf("Error occurred while sending %s notification", senderName), "error", err)
		return err
	}

	ls.Logger.Info(fmt.Sprintf("%s notification sent successfully", senderName), "notification", n)

	return nil
}
