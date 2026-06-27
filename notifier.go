package main

import "fmt"

type Notifier interface {
	Notify(msg string) error
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) Notify(msg string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Email, msg)
	return nil
}

type SlackNotifier struct {
	Channel string
}

func (s SlackNotifier) Notify(msg string) error {
	fmt.Printf("Sending Slack message to %s: %s\n", s.Channel, msg)
	return nil
}

func notifyAll(ns []Notifier, msg string) {
	for _, notifier := range ns {
		err := notifier.Notify(msg)

		if err != nil {
			fmt.Println("Failed to notify:", err)
		}
	}
}
