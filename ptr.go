package ptr

import "time"

func StringPtr(s string) *string {
	return &s
}

func UintPtr(i uint) *uint {
	return &i
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func IntPtr(i int) *int {
	return &i
}

func Int64Ptr(i int64) *int64 {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func TimePtr(t time.Time) *time.Time {
	return &t
}

func DurationPtr(d time.Duration) *time.Duration {
	return &d
}

// ---------------

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

func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
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
