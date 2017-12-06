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
	"context"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// JSONSender implements the email.Sender interface. It will marshal the
// message to JSON and will write a single message on a single line (see
// http://jsonlines.org/). It is intended to be used with file, but should also
// work with os.Stdout.
type JSONSender struct {
	out io.Writer
}

// NewJSONSender creates a new JSONSender.
func NewJSONSender(out io.Writer) *JSONSender {
	return &JSONSender{out: out}
}

var _ Sender = (*JSONSender)(nil)

// Send will write the marshaled message to s.out.
func (s *JSONSender) Send(ctx context.Context, message *Message) error {
	b, err := json.Marshal(message)
	if err != nil {
		// NOTE(bvdberg): not sure if we need this check
		return errors.Wrap(err, "unable to marshal message")
	}

	b = append(b, byte('\n'))

	_, err = s.out.Write(b)
	if err != nil {
		return errors.Wrap(err, "unable to write to out")
	}

	return nil
}
