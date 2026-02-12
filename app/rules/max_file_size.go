package rules

import (
	"context"
	"mime/multipart"

	"github.com/goravel/framework/contracts/validation"
	"github.com/spf13/cast"
)

type MaxFileSize struct {
}

// Signature The name of the rule.
func (receiver *MaxFileSize) Signature() string {
	return "max_file_size"
}

// Passes Determine if the validation rule passes.
// Checks that the uploaded file does not exceed the given size in KB.
// Usage: max_file_size:1024 (1024 KB = 1 MB)
func (receiver *MaxFileSize) Passes(ctx context.Context, data validation.Data, val any, options ...any) bool {
	var size int64
	switch v := val.(type) {
	case *multipart.FileHeader:
		size = v.Size
	case multipart.FileHeader:
		size = v.Size
	default:
		return false
	}

	maxKB := int64(1024)
	if len(options) > 0 {
		maxKB = cast.ToInt64(options[0])
		if maxKB <= 0 {
			maxKB = 1024
		}
	}

	return size <= maxKB*1024
}

// Message Get the validation error message.
func (receiver *MaxFileSize) Message(ctx context.Context) string {
	return "El :attribute no puede exceder el tamaño máximo permitido"
}
