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

package log

import "context"

// FromContext will create a new Entry based on the Global entry with any
// custom fields defined in ctx.
func FromContext(ctx context.Context) *Entry {
	if fields, ok := FieldsFromContext(ctx); ok {
		return entry.WithFields(fields)
	}

	return entry
}

// AddFieldToCtx adds a new field to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
func AddFieldToCtx(ctx context.Context, key string, value interface{}) (context.Context, *Entry) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		fields = Fields{}
	}

	// Add all Fields from the context, then add the new field
	e := entry.WithFields(fields).WithField(key, value)

	// Fetch the Fields from the Entry which will be stored in the new context.
	fields = Fields(e.Entry.Data)

	return ContextWithFields(ctx, fields), e
}

// AddFieldsToCtx adds newFields to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
func AddFieldsToCtx(ctx context.Context, newFields Fields) (context.Context, *Entry) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		fields = Fields{}
	}

	// Add all Fields from the context, then add the new Fields
	e := entry.WithFields(fields).WithFields(newFields)

	// Fetch the Fields from the Entry which will be stored in the new context.
	fields = Fields(e.Entry.Data)

	return ContextWithFields(ctx, fields), e
}

// key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// fieldsKey is the context key for the Fields.
const fieldsKey key = 0

// FieldsFromContext retrieves the Fields from ctx.
func FieldsFromContext(ctx context.Context) (Fields, bool) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	return fields, ok
}

// ContextWithFields set fields in a new context based on ctx, and returns this
// context. Any Fields defined in ctx will be overriden.
func ContextWithFields(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, fieldsKey, fields)
}
