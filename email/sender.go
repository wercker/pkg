package email

import "context"

// Sender is a service which can deliver a message.
type Sender interface {
	Send(ctx context.Context, message *Message) error
}

// Message represents a email.
type Message struct {
	From     string            `json:"from,omitempty"`
	To       []string          `json:"to,omitempty"`
	CC       []string          `json:"cc,omitempty"`
	BCC      []string          `json:"bcc,omitempty"`
	Reply    []string          `json:"reply,omitempty"`
	Subject  string            `json:"subject,omitempty"`
	HTMLBody string            `json:"htmlBody,omitempty"`
	TextBody string            `json:"textBody,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
}
