package email

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/pkg/errors"
)

// SESSender is a email.Sender which sends emails using the AWS SES API.
type SESSender struct {
	ses SenderSESAPI
}

var _ Sender = (*SESSender)(nil)

// NewSESSender creates a new SESSender.
func NewSESSender(ses SenderSESAPI) *SESSender {
	if ses == nil {
		panic(errors.New("NewSESSender requires SenderSESAPI"))
	}

	return &SESSender{ses: ses}
}

// SenderSESAPI specifies the requirements from the sesiface.SESAPI. It only
// uses a subset of sesiface.SESAPI and thus ses.SES can still be used.
type SenderSESAPI interface {
	SendEmail(*ses.SendEmailInput) (*ses.SendEmailOutput, error)
}

// Send an email using SES.
func (s *SESSender) Send(ctx context.Context, message *Message) error {
	input := convertToSendEmailInput(message)

	// Validate using AWS validation rules
	err := input.Validate()
	if err != nil {
		return errors.Wrap(err, "invalid message")
	}

	_, err = s.ses.SendEmail(input)
	if err != nil {
		return errors.Wrap(err, "unable to send email through SES")
	}

	return nil
}

// convertToSendEmailInput takes a generic Message and converts it to a SES
// specific ses.SendEmailInput. If m is nil then it will return an empty
// ses.SendEmailInput (not nil).
func convertToSendEmailInput(m *Message) *ses.SendEmailInput {
	result := &ses.SendEmailInput{}

	if m != nil {
		if m.From != "" {
			result.Source = aws.String(m.From)
		}

		if len(m.To) > 0 || len(m.CC) > 0 || len(m.BCC) > 0 || len(m.Reply) > 0 {
			result.Destination = &ses.Destination{}
			if len(m.To) > 0 {
				result.Destination.ToAddresses = aws.StringSlice(m.To)
			}

			if len(m.CC) > 0 {
				result.Destination.CcAddresses = aws.StringSlice(m.CC)
			}

			if len(m.BCC) > 0 {
				result.Destination.BccAddresses = aws.StringSlice(m.BCC)
			}

			if len(m.Reply) > 0 {
				result.ReplyToAddresses = aws.StringSlice(m.Reply)
			}
		}

		if m.Subject != "" || m.HTMLBody != "" || m.TextBody != "" {
			result.Message = &ses.Message{}
			if m.Subject != "" {
				result.Message.Subject = &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(m.Subject),
				}
			}

			if m.HTMLBody != "" || m.TextBody != "" {
				result.Message.Body = &ses.Body{}
				if m.HTMLBody != "" {
					result.Message.Body.Html = &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(m.HTMLBody),
					}
				}

				if m.TextBody != "" {
					result.Message.Body.Text = &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(m.TextBody),
					}
				}
			}
		}
	}

	return result
}
