package ptr

import "time"

// StringPtr returns a pointer to the string value passed in.
func StringPtr(s string) *string {
	return &s
}

func UintPtr(i uint) *uint {
	return &i
}

func Uint8Ptr(i uint8) *uint8 {
	return &i
}

func Uint16Ptr(i uint16) *uint16 {
	return &i
}

func Uint32Ptr(i uint32) *uint32 {
	return &i
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func IntPtr(i int) *int {
	return &i
}

func Int8Ptr(i int8) *int8 {
	return &i
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func Float32Ptr(f float32) *float32 {
	return &f
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Complex64Ptr(c complex64) *complex64 {
	return &c
}

func Complex128Ptr(c complex128) *complex128 {
	return &c
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func DurationPtr(d time.Duration) *time.Duration {
	return &d
}

func RunePtr(r rune) *rune {
	return &r
}

func BytePtr(b byte) *byte {
	return &b
}

// ---------------

// StringValue Returns the value of the string pointer passed in, or an empty string if the pointer is nil.
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func UintValue(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}

func Uint8Value(i *uint8) uint8 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint16Value(i *uint16) uint16 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint32Value(i *uint32) uint32 {
	if i == nil {
		return 0
	}
	return *i
}

func Uint64Value(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func Int8Value(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

func Int16Value(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

func Int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func Float32Value(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

func Complex64Value(c *complex64) complex64 {
	if c == nil {
		return 0
	}
	return *c
}

func Complex128Value(c *complex128) complex128 {
	if c == nil {
		return 0
	}
	return *c
}

func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

func DurationValue(d *time.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return *d
}

func RuneValue(r *rune) rune {
	if r == nil {
		return 0
	}
	return *r
}

func ByteValue(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

// ---------------

func IsEqualStringPtr(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualUintPtr(a, b *uint) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualUint8Ptr(a, b *uint8) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualUint16Ptr(a, b *uint16) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualUint32Ptr(a, b *uint32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualUint64Ptr(a, b *uint64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualIntPtr(a, b *int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualInt8Ptr(a, b *int8) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualInt16Ptr(a, b *int16) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualInt32Ptr(a, b *int32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualInt64Ptr(a, b *int64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualBoolPtr(a, b *bool) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualFloat32Ptr(a, b *float32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualFloat64Ptr(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualComplex64Ptr(a, b *complex64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualComplex128Ptr(a, b *complex128) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualTimePtr(a, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equal(*b)
}

func IsEqualDurationPtr(a, b *time.Duration) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualRunePtr(a, b *rune) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func IsEqualBytePtr(a, b *byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
