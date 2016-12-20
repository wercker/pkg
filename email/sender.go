package email

import "context"

// Sender is a service which can deliver a message.
type Sender interface {
	Send(ctx context.Context, message *Message) error
}

// Message represents a email.
type Message struct {
	From  string
	To    []string
	CC    []string
	BCC   []string
	Reply []string

	Subject string

	HTMLBody string
	TextBody string
}
