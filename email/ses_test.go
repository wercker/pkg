package email

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func ExampleSESSender() {
	// Create a SenderSESAPI compatible object, example:
	// client := ses.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	var client SenderSESAPI = nil
	sender := NewSESSender(client)

	msg := &Message{
		From:     "foo@example.com",
		To:       []string{"foofoo@example.com"},
		Subject:  "foofoo",
		HTMLBody: "foofoo",
	}

	sender.Send(context.TODO(), msg)
}

func Test_convertToSendEmailInput(t *testing.T) {
	tests := []struct {
		Name     string
		Input    *Message
		Expected *ses.SendEmailInput
	}{
		{
			"Nil",
			nil,
			&ses.SendEmailInput{},
		},
		{
			"EmptyObject",
			&Message{},
			&ses.SendEmailInput{},
		},
		{
			"AllProperties",
			&Message{
				From:     "from_foo@example.com",
				To:       []string{"to_foo@example.com"},
				CC:       []string{"cc_foo@example.com"},
				BCC:      []string{"bcc_foo@example.com"},
				Reply:    []string{"reply_foo@example.com"},
				Subject:  "subject_foo",
				HTMLBody: "htmlbody_foo",
				TextBody: "textbody_foo",
			},
			&ses.SendEmailInput{
				Source: aws.String("from_foo@example.com"),
				Destination: &ses.Destination{
					ToAddresses:  []*string{aws.String("to_foo@example.com")},
					CcAddresses:  []*string{aws.String("cc_foo@example.com")},
					BccAddresses: []*string{aws.String("bcc_foo@example.com")},
				},
				ReplyToAddresses: []*string{aws.String("reply_foo@example.com")},
				Message: &ses.Message{
					Subject: &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String("subject_foo"),
					},
					Body: &ses.Body{
						Html: &ses.Content{
							Charset: aws.String("UTF-8"),
							Data:    aws.String("htmlbody_foo"),
						},
						Text: &ses.Content{
							Charset: aws.String("UTF-8"),
							Data:    aws.String("textbody_foo"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			actual := convertToSendEmailInput(tt.Input)
			require.Equal(t, tt.Expected, actual)
		})
	}
}

// Test_NewSESSender_RequiresSESAPI checks whether NewSESSender panics when
// passing nil
func Test_NewSESSender_RequiresSESAPI(t *testing.T) {
	require.Panics(t, func() {
		NewSESSender(nil)
	})
}

// Test_SESSender_Send_InvalidMessage checks whether the validation provided by
// ses will be triggered by incorrect input from a message.
// NOTE(bvdberg): Not sure how useful this unit test is, as this mostly
// delgates to the aws library and might change over time.
func Test_SESSender_Send_InvalidMessage(t *testing.T) {
	tests := []struct {
		Name  string
		Input *Message
	}{
		{"NilMessage", nil},
		{"EmptyMessage", &Message{}},
		{
			"MissingFrom",
			&Message{
				To:       []string{"foo@example.com"},
				TextBody: "foofoo",
				Subject:  "foofoo",
			},
		},
		{
			"MissingDestination",
			&Message{
				From:     "foo@example",
				TextBody: "foofoo",
				Subject:  "foofoo",
			},
		},
		{
			"MissingSubject",
			&Message{
				From:     "foo@example",
				To:       []string{"foo@example.com"},
				TextBody: "foofoo",
			},
		},
		{
			"MissingBody",
			&Message{
				From:    "foo@example",
				To:      []string{"foo@example.com"},
				Subject: "foofoo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m := new(MockSenderSESAPI)
			sender := NewSESSender(m)

			err := sender.Send(context.Background(), tt.Input)
			require.Error(t, err)
		})
	}
}

func Test_SESSender_Send_ValidMessage(t *testing.T) {
	input := &Message{
		From:     "foo@example",
		To:       []string{"foo@example.com"},
		Subject:  "foofoo",
		TextBody: "foofoo",
	}

	m := new(MockSenderSESAPI)
	m.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)

	sender := NewSESSender(m)

	err := sender.Send(context.Background(), input)
	require.NoError(t, err)
}

func Test_SESSender_Send_FailingSESAPI(t *testing.T) {
	input := &Message{
		From:     "foo@example",
		To:       []string{"foo@example.com"},
		Subject:  "foofoo",
		TextBody: "foofoo",
	}

	m := new(MockSenderSESAPI)
	m.On("SendEmail", mock.Anything).Return(nil, errors.New("fail from mock"))

	sender := NewSESSender(m)

	err := sender.Send(context.Background(), input)
	require.Error(t, err)
}
