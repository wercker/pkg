//-----------------------------------------------------------------------------
// Copyright (c) 2017 Oracle and/or its affiliates.  All rights reserved.
// This program is free software: you can modify it and/or redistribute it
// under the terms of:
//
// (i)  the Universal Permissive License v 1.0 or at your option, any
//      later version (http://oss.oracle.com/licenses/upl); and/or
//
// (ii) the Apache License v 2.0. (http://www.apache.org/licenses/LICENSE-2.0)
//-----------------------------------------------------------------------------

package email

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_JSONSender_Send(t *testing.T) {
	tests := []struct {
		Name     string
		Messages []*Message
		Expected []string
	}{
		{
			"NoMessage",
			[]*Message{},
			[]string{},
		},
		{
			"EmptyMessage",
			[]*Message{{}},
			[]string{"{}"},
		},
		{
			"ActualMessage",
			[]*Message{
				{
					From:     "foo+from@example.com",
					To:       []string{"foo+to@example.com"},
					CC:       []string{"foo+cc@example.com"},
					BCC:      []string{"foo+bcc@example.com"},
					Reply:    []string{"foo+reply@example.com"},
					Subject:  "foo subject",
					HTMLBody: "foo HTML body\nSecond line",
					TextBody: "foo Text body\nSecond line",
				},
			},
			[]string{"{\"from\":\"foo+from@example.com\",\"to\":[\"foo+to@example.com\"],\"cc\":[\"foo+cc@example.com\"],\"bcc\":[\"foo+bcc@example.com\"],\"reply\":[\"foo+reply@example.com\"],\"subject\":\"foo subject\",\"htmlBody\":\"foo HTML body\\nSecond line\",\"textBody\":\"foo Text body\\nSecond line\"}"},
		},
		{
			"ActualMessage",
			[]*Message{
				{
					From:     "foo+from@example.com",
					To:       []string{"foo+to@example.com"},
					CC:       []string{"foo+cc@example.com"},
					BCC:      []string{"foo+bcc@example.com"},
					Reply:    []string{"foo+reply@example.com"},
					Subject:  "foo subject",
					HTMLBody: "foo HTML body\nSecond line",
					TextBody: "foo Text body\nSecond line",
				},
				{
					From:     "foo+from2@example.com",
					To:       []string{"foo+to2@example.com"},
					CC:       []string{"foo+cc2@example.com"},
					BCC:      []string{"foo+bcc2@example.com"},
					Reply:    []string{"foo+reply2@example.com"},
					Subject:  "foo subject2",
					HTMLBody: "foo HTML body2\nSecond line",
					TextBody: "foo Text body2\nSecond line",
				},
			},
			[]string{
				"{\"from\":\"foo+from@example.com\",\"to\":[\"foo+to@example.com\"],\"cc\":[\"foo+cc@example.com\"],\"bcc\":[\"foo+bcc@example.com\"],\"reply\":[\"foo+reply@example.com\"],\"subject\":\"foo subject\",\"htmlBody\":\"foo HTML body\\nSecond line\",\"textBody\":\"foo Text body\\nSecond line\"}",
				"{\"from\":\"foo+from2@example.com\",\"to\":[\"foo+to2@example.com\"],\"cc\":[\"foo+cc2@example.com\"],\"bcc\":[\"foo+bcc2@example.com\"],\"reply\":[\"foo+reply2@example.com\"],\"subject\":\"foo subject2\",\"htmlBody\":\"foo HTML body2\\nSecond line\",\"textBody\":\"foo Text body2\\nSecond line\"}",
			},
		},
		{
			"MultipleEmptyMessages",
			[]*Message{{}, {}, {}},
			[]string{"{}", "{}", "{}"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := context.Background()
			sink := new(bytes.Buffer)
			sender := NewJSONSender(sink)

			for _, m := range tt.Messages {
				err := sender.Send(ctx, m)
				require.NoError(t, err)
			}

			actual := sink.String()
			results := strings.Split(actual, "\n")

			if assert.Equal(t, len(tt.Expected), len(results)-1) {
				for i, e := range tt.Expected {
					assert.JSONEq(t, e, results[i])
				}
			}
		})
	}
}
