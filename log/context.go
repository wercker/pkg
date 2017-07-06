package log

import "context"

// FromContext will create a new Entry based on the Global entry with any
// custom fields defined in ctx.
func FromContext(ctx context.Context) Logger {
	if fields, ok := FieldsFromContext(ctx); ok {
		return baseLogger.WithFields(fields)
	}

	return baseLogger
}

// AddFieldToCtx adds a new field to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
//
// Deprecated: Use AddFieldToContext instead.
func AddFieldToCtx(ctx context.Context, key string, value interface{}) (context.Context, Logger) {
	return AddFieldToContext(ctx, key, value)
}

// AddFieldToContext adds a new field to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
func AddFieldToContext(ctx context.Context, key string, value interface{}) (context.Context, Logger) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		fields = Fields{}
	}

	// Add all Fields from the context, then add the new field
	l := baseLogger.WithFields(fields).With(key, value)

	return ContextWithFields(ctx, l.Fields()), l
}

// AddFieldsToCtx adds newFields to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
//
// Deprecated: Use AddFieldsToContext instead.
func AddFieldsToCtx(ctx context.Context, newFields Fields) (context.Context, Logger) {
	return AddFieldsToContext(ctx, newFields)
}

// AddFieldsToContext adds newFields to fields in the ctx (creates a new one if
// there is none). It returns the new context and a entry will the current
// Fields objects.
func AddFieldsToContext(ctx context.Context, newFields Fields) (context.Context, Logger) {
	fields, ok := ctx.Value(fieldsKey).(Fields)
	if !ok {
		fields = Fields{}
	}

	// Add all Fields from the context, then add the new Fields
	l := baseLogger.WithFields(fields).WithFields(newFields)

	return ContextWithFields(ctx, l.Fields()), l
}

// contextKey type is unexported to prevent collisions with context keys defined in
// other packages.
type contextKey struct{}

// fieldsKey is the context key for the Fields.
var fieldsKey = contextKey{}

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
