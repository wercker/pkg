package log

import "context"

// FromContext will create a new Entry based on the Global entry with any
// custom fields defined in ctx.
func FromContext(ctx context.Context) *Entry {
	if fields, ok := fieldsFromContext(ctx); ok {
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

	return fieldsNewContext(ctx, fields), e
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

	return fieldsNewContext(ctx, fields), e
}

// key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// fieldsKey is the context key for the Fields.
const fieldsKey key = 0

// fieldsFromContext retrieves the Fields from ctx. Currently it is unexported,
// since the main entrypoint should be through FromContext and
// AddField(s)ToCtx.
func fieldsFromContext(ctx context.Context) (Fields, bool) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	return fields, ok
}

// fieldsNewContext set fields in a new context based on ctx, and returns this
// context. Any Fields defined in ctx will be overriden.
func fieldsNewContext(ctx context.Context, fields Fields) context.Context {
	return context.WithValue(ctx, fieldsKey, fields)
}
