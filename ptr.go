package ptr

import "time"

// StringPtr returns a pointer to the string value passed in.
func StringPtr(s string) *string {
	return &s
}

// UintPtr returns a pointer to the uint value passed in.
func UintPtr(i uint) *uint {
	return &i
}

// Uint8Ptr returns a pointer to the uint8 value passed in.
func Uint8Ptr(i uint8) *uint8 {
	return &i
}

// Uint16Ptr returns a pointer to the uint16 value passed in.
func Uint16Ptr(i uint16) *uint16 {
	return &i
}

// Uint32Ptr returns a pointer to the uint32 value passed in.
func Uint32Ptr(i uint32) *uint32 {
	return &i
}

// Uint64Ptr returns a pointer to the uint64 value passed in.
func Uint64Ptr(i uint64) *uint64 {
	return &i
}

// IntPtr returns a pointer to the int value passed in.
func IntPtr(i int) *int {
	return &i
}

// Int8Ptr returns a pointer to the int8 value passed in.
func Int8Ptr(i int8) *int8 {
	return &i
}

// Int16Ptr returns a pointer to the int16 value passed in.
func Int16Ptr(i int16) *int16 {
	return &i
}

// Int32Ptr returns a pointer to the int32 value passed in.
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr returns a pointer to the int64 value passed in.
func Int64Ptr(i int64) *int64 {
	return &i
}

// BoolPtr returns a pointer to the bool value passed in.
func BoolPtr(b bool) *bool {
	return &b
}

// Float32Ptr returns a pointer to the float32 value passed in.
func Float32Ptr(f float32) *float32 {
	return &f
}

// Float64Ptr returns a pointer to the float64 value passed in.
func Float64Ptr(f float64) *float64 {
	return &f
}

// Complex64Ptr returns a pointer to the complex64 value passed in.
func Complex64Ptr(c complex64) *complex64 {
	return &c
}

// Complex128Ptr returns a pointer to the complex128 value passed in.
func Complex128Ptr(c complex128) *complex128 {
	return &c
}

// TimePtr returns a pointer to the time.Time value passed in.
func TimePtr(t time.Time) *time.Time {
	return &t
}

// DurationPtr returns a pointer to the time.Duration value passed in.
func DurationPtr(d time.Duration) *time.Duration {
	return &d
}

// RunePtr returns a pointer to the rune value passed in.
func RunePtr(r rune) *rune {
	return &r
}

// BytePtr returns a pointer to the byte value passed in.
func BytePtr(b byte) *byte {
	return &b
}

// ---------------

// StringValue returns the value of the string pointer passed in, or an empty string if the pointer is nil.
func StringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// UintValue returns the value of the uint pointer passed in, or 0 if the pointer is nil.
func UintValue(i *uint) uint {
	if i == nil {
		return 0
	}
	return *i
}

// Uint8Value returns the value of the uint8 pointer passed in, or 0 if the pointer is nil.
func Uint8Value(i *uint8) uint8 {
	if i == nil {
		return 0
	}
	return *i
}

// Uint16Value returns the value of the uint16 pointer passed in, or 0 if the pointer is nil.
func Uint16Value(i *uint16) uint16 {
	if i == nil {
		return 0
	}
	return *i
}

// Uint32Value returns the value of the uint32 pointer passed in, or 0 if the pointer is nil.
func Uint32Value(i *uint32) uint32 {
	if i == nil {
		return 0
	}
	return *i
}

// Uint64Value
func Uint64Value(i *uint64) uint64 {
	if i == nil {
		return 0
	}
	return *i
}

// IntValue
func IntValue(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// Int8Value
func Int8Value(i *int8) int8 {
	if i == nil {
		return 0
	}
	return *i
}

// Int16Value
func Int16Value(i *int16) int16 {
	if i == nil {
		return 0
	}
	return *i
}

// Int32Value
func Int32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

// Int64Value
func Int64Value(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

// BoolValue
func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

// Float32Value
func Float32Value(f *float32) float32 {
	if f == nil {
		return 0
	}
	return *f
}

// Float64Value
func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// Complex64Value
func Complex64Value(c *complex64) complex64 {
	if c == nil {
		return 0
	}
	return *c
}

// Complex128Value
func Complex128Value(c *complex128) complex128 {
	if c == nil {
		return 0
	}
	return *c
}

// TimeValue
func TimeValue(t *time.Time) time.Time {
	if t == nil {
		return time.Time{}
	}
	return *t
}

// DurationValue
func DurationValue(d *time.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return *d
}

// RuneValue
func RuneValue(r *rune) rune {
	if r == nil {
		return 0
	}
	return *r
}

// ByteValue
func ByteValue(b *byte) byte {
	if b == nil {
		return 0
	}
	return *b
}

// ---------------

// IsEqualStringPtr
func IsEqualStringPtr(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualUintPtr
func IsEqualUintPtr(a, b *uint) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualUint8Ptr
func IsEqualUint8Ptr(a, b *uint8) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualUint16Ptr
func IsEqualUint16Ptr(a, b *uint16) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualUint32Ptr
func IsEqualUint32Ptr(a, b *uint32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualUint64Ptr
func IsEqualUint64Ptr(a, b *uint64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualIntPtr
func IsEqualIntPtr(a, b *int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualInt8Ptr
func IsEqualInt8Ptr(a, b *int8) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualInt16Ptr
func IsEqualInt16Ptr(a, b *int16) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualInt32Ptr
func IsEqualInt32Ptr(a, b *int32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualInt64Ptr
func IsEqualInt64Ptr(a, b *int64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualBoolPtr
func IsEqualBoolPtr(a, b *bool) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualFloat32Ptr
func IsEqualFloat32Ptr(a, b *float32) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualFloat64Ptr
func IsEqualFloat64Ptr(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualComplex64Ptr
func IsEqualComplex64Ptr(a, b *complex64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualComplex128Ptr
func IsEqualComplex128Ptr(a, b *complex128) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualTimePtr
func IsEqualTimePtr(a, b *time.Time) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equal(*b)
}

// IsEqualDurationPtr
func IsEqualDurationPtr(a, b *time.Duration) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualRunePtr
func IsEqualRunePtr(a, b *rune) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

// IsEqualBytePtr
func IsEqualBytePtr(a, b *byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
