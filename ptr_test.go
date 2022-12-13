package ptr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	t.Run("StringPtr", func(t *testing.T) {
		emptyStr := ""
		nonEmptyStr := "non-empty"

		cases := []struct {
			name     string
			actual   *string
			expected string
		}{
			{
				name:     "empty string",
				actual:   StringPtr(emptyStr),
				expected: emptyStr,
			},
			{
				name:     "non-empty string",
				actual:   StringPtr(nonEmptyStr),
				expected: nonEmptyStr,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("StringValue", func(t *testing.T) {
		emptyStrPtr := StringPtr("")
		nonEmptyStrPtr := StringPtr("non-empty")

		cases := []struct {
			name     string
			actual   string
			expected string
		}{
			{
				name:     "empty string",
				actual:   StringValue(emptyStrPtr),
				expected: "",
			},
			{
				name:     "non-empty string",
				actual:   StringValue(nonEmptyStrPtr),
				expected: "non-empty",
			},
			{
				name:     "nil string",
				actual:   StringValue(nil),
				expected: "",
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IntPtr", func(t *testing.T) {
		zero := 0
		nonZero := 1

		cases := []struct {
			name     string
			actual   *int
			expected int
		}{
			{
				name:     "zero",
				actual:   IntPtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   IntPtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("IntValue", func(t *testing.T) {
		zeroPtr := IntPtr(0)
		nonZeroPtr := IntPtr(1)

		cases := []struct {
			name     string
			actual   int
			expected int
		}{
			{
				name:     "zero",
				actual:   IntValue(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   IntValue(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil int",
				actual:   IntValue(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Int8Ptr", func(t *testing.T) {
		zero := int8(0)
		nonZero := int8(1)

		cases := []struct {
			name     string
			actual   *int8
			expected int8
		}{
			{
				name:     "zero",
				actual:   Int8Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Int8Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Int8Value", func(t *testing.T) {
		zeroPtr := Int8Ptr(0)
		nonZeroPtr := Int8Ptr(1)

		cases := []struct {
			name     string
			actual   int8
			expected int8
		}{
			{
				name:     "zero",
				actual:   Int8Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Int8Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil int8",
				actual:   Int8Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Int16Ptr", func(t *testing.T) {
		zero := int16(0)
		nonZero := int16(1)

		cases := []struct {
			name     string
			actual   *int16
			expected int16
		}{
			{
				name:     "zero",
				actual:   Int16Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Int16Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Int16Value", func(t *testing.T) {
		zeroPtr := Int16Ptr(0)
		nonZeroPtr := Int16Ptr(1)

		cases := []struct {
			name     string
			actual   int16
			expected int16
		}{
			{
				name:     "zero",
				actual:   Int16Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Int16Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil int16",
				actual:   Int16Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Int32Ptr", func(t *testing.T) {
		zero := int32(0)
		nonZero := int32(1)

		cases := []struct {
			name     string
			actual   *int32
			expected int32
		}{
			{
				name:     "zero",
				actual:   Int32Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Int32Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Int32Value", func(t *testing.T) {
		zeroPtr := Int32Ptr(0)
		nonZeroPtr := Int32Ptr(1)

		cases := []struct {
			name     string
			actual   int32
			expected int32
		}{
			{
				name:     "zero",
				actual:   Int32Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Int32Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil int32",
				actual:   Int32Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Int64Ptr", func(t *testing.T) {
		zero := int64(0)
		nonZero := int64(1)

		cases := []struct {
			name     string
			actual   *int64
			expected int64
		}{
			{
				name:     "zero",
				actual:   Int64Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Int64Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Int64Value", func(t *testing.T) {
		zeroPtr := Int64Ptr(0)
		nonZeroPtr := Int64Ptr(1)

		cases := []struct {
			name     string
			actual   int64
			expected int64
		}{
			{
				name:     "zero",
				actual:   Int64Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Int64Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil int64",
				actual:   Int64Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Complex64Ptr", func(t *testing.T) {
		zero := complex64(0)
		nonZero := complex64(1)

		cases := []struct {
			name     string
			actual   *complex64
			expected complex64
		}{
			{
				name:     "zero",
				actual:   Complex64Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Complex64Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Complex64Value", func(t *testing.T) {
		zeroPtr := Complex64Ptr(0)
		nonZeroPtr := Complex64Ptr(1)

		cases := []struct {
			name     string
			actual   complex64
			expected complex64
		}{
			{
				name:     "zero",
				actual:   Complex64Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Complex64Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil complex64",
				actual:   Complex64Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Complex128Ptr", func(t *testing.T) {
		zero := complex128(0)
		nonZero := complex128(1)

		cases := []struct {
			name     string
			actual   *complex128
			expected complex128
		}{
			{
				name:     "zero",
				actual:   Complex128Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Complex128Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Complex128Value", func(t *testing.T) {
		zeroPtr := Complex128Ptr(0)
		nonZeroPtr := Complex128Ptr(1)

		cases := []struct {
			name     string
			actual   complex128
			expected complex128
		}{
			{
				name:     "zero",
				actual:   Complex128Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Complex128Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil complex128",
				actual:   Complex128Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("UintPtr", func(t *testing.T) {
		zero := uint(0)
		nonZero := uint(1)

		cases := []struct {
			name     string
			actual   *uint
			expected uint
		}{
			{
				name:     "zero",
				actual:   UintPtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   UintPtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("UintValue", func(t *testing.T) {
		zeroPtr := UintPtr(0)
		nonZeroPtr := UintPtr(1)

		cases := []struct {
			name     string
			actual   uint
			expected uint
		}{
			{
				name:     "zero",
				actual:   UintValue(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   UintValue(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil uint",
				actual:   UintValue(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Uint8Ptr", func(t *testing.T) {
		zero := uint8(0)
		nonZero := uint8(1)

		cases := []struct {
			name     string
			actual   *uint8
			expected uint8
		}{
			{
				name:     "zero",
				actual:   Uint8Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Uint8Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Uint8Value", func(t *testing.T) {
		zeroPtr := Uint8Ptr(0)
		nonZeroPtr := Uint8Ptr(1)

		cases := []struct {
			name     string
			actual   uint8
			expected uint8
		}{
			{
				name:     "zero",
				actual:   Uint8Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Uint8Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil uint8",
				actual:   Uint8Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Uint16Ptr", func(t *testing.T) {
		zero := uint16(0)
		nonZero := uint16(1)

		cases := []struct {
			name     string
			actual   *uint16
			expected uint16
		}{
			{
				name:     "zero",
				actual:   Uint16Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Uint16Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Uint16Value", func(t *testing.T) {
		zeroPtr := Uint16Ptr(0)
		nonZeroPtr := Uint16Ptr(1)

		cases := []struct {
			name     string
			actual   uint16
			expected uint16
		}{
			{
				name:     "zero",
				actual:   Uint16Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Uint16Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil uint16",
				actual:   Uint16Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Uint32Ptr", func(t *testing.T) {
		zero := uint32(0)
		nonZero := uint32(1)

		cases := []struct {
			name     string
			actual   *uint32
			expected uint32
		}{
			{
				name:     "zero",
				actual:   Uint32Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Uint32Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Uint32Value", func(t *testing.T) {
		zeroPtr := Uint32Ptr(0)
		nonZeroPtr := Uint32Ptr(1)

		cases := []struct {
			name     string
			actual   uint32
			expected uint32
		}{
			{
				name:     "zero",
				actual:   Uint32Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Uint32Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil uint32",
				actual:   Uint32Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Uint64Ptr", func(t *testing.T) {
		zero := uint64(0)
		nonZero := uint64(1)

		cases := []struct {
			name     string
			actual   *uint64
			expected uint64
		}{
			{
				name:     "zero",
				actual:   Uint64Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Uint64Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Uint64Value", func(t *testing.T) {
		zeroPtr := Uint64Ptr(0)
		nonZeroPtr := Uint64Ptr(1)

		cases := []struct {
			name     string
			actual   uint64
			expected uint64
		}{
			{
				name:     "zero",
				actual:   Uint64Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Uint64Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil uint64",
				actual:   Uint64Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Float32Ptr", func(t *testing.T) {
		zero := float32(0)
		nonZero := float32(1)

		cases := []struct {
			name     string
			actual   *float32
			expected float32
		}{
			{
				name:     "zero",
				actual:   Float32Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Float32Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Float32Value", func(t *testing.T) {
		zeroPtr := Float32Ptr(0)
		nonZeroPtr := Float32Ptr(1)

		cases := []struct {
			name     string
			actual   float32
			expected float32
		}{
			{
				name:     "zero",
				actual:   Float32Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Float32Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil float32",
				actual:   Float32Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("Float64Ptr", func(t *testing.T) {
		zero := float64(0)
		nonZero := float64(1)

		cases := []struct {
			name     string
			actual   *float64
			expected float64
		}{
			{
				name:     "zero",
				actual:   Float64Ptr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   Float64Ptr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("Float64Value", func(t *testing.T) {
		zeroPtr := Float64Ptr(0)
		nonZeroPtr := Float64Ptr(1)

		cases := []struct {
			name     string
			actual   float64
			expected float64
		}{
			{
				name:     "zero",
				actual:   Float64Value(zeroPtr),
				expected: 0,
			},
			{
				name:     "non-zero",
				actual:   Float64Value(nonZeroPtr),
				expected: 1,
			},
			{
				name:     "nil float64",
				actual:   Float64Value(nil),
				expected: 0,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("BoolPtr", func(t *testing.T) {
		zero := false
		nonZero := true

		cases := []struct {
			name     string
			actual   *bool
			expected bool
		}{
			{
				name:     "zero",
				actual:   BoolPtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   BoolPtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("BoolValue", func(t *testing.T) {
		zeroPtr := BoolPtr(false)
		nonZeroPtr := BoolPtr(true)

		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "zero",
				actual:   BoolValue(zeroPtr),
				expected: false,
			},
			{
				name:     "non-zero",
				actual:   BoolValue(nonZeroPtr),
				expected: true,
			},
			{
				name:     "nil bool",
				actual:   BoolValue(nil),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("TimePtr", func(t *testing.T) {
		zero := time.Time{}
		nonZero := time.Now()

		cases := []struct {
			name     string
			actual   *time.Time
			expected time.Time
		}{
			{
				name:     "zero",
				actual:   TimePtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   TimePtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("TimeValue", func(t *testing.T) {
		zeroPtr := TimePtr(time.Time{})
		nonZeroPtr := TimePtr(time.Now())

		cases := []struct {
			name     string
			actual   time.Time
			expected time.Time
		}{
			{
				name:     "zero",
				actual:   TimeValue(zeroPtr),
				expected: time.Time{},
			},
			{
				name:     "non-zero",
				actual:   TimeValue(nonZeroPtr),
				expected: *nonZeroPtr,
			},
			{
				name:     "nil time",
				actual:   TimeValue(nil),
				expected: time.Time{},
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("DurationPtr", func(t *testing.T) {
		zero := time.Duration(0)
		nonZero := time.Second

		cases := []struct {
			name     string
			actual   *time.Duration
			expected time.Duration
		}{
			{
				name:     "zero",
				actual:   DurationPtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   DurationPtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("DurationValue", func(t *testing.T) {
		zeroPtr := DurationPtr(time.Duration(0))
		nonZeroPtr := DurationPtr(time.Second)

		cases := []struct {
			name     string
			actual   time.Duration
			expected time.Duration
		}{
			{
				name:     "zero",
				actual:   DurationValue(zeroPtr),
				expected: time.Duration(0),
			},
			{
				name:     "non-zero",
				actual:   DurationValue(nonZeroPtr),
				expected: *nonZeroPtr,
			},
			{
				name:     "nil duration",
				actual:   DurationValue(nil),
				expected: time.Duration(0),
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("RunePtr", func(t *testing.T) {
		zero := rune(0)
		nonZero := rune('a')

		cases := []struct {
			name     string
			actual   *rune
			expected rune
		}{
			{
				name:     "zero",
				actual:   RunePtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   RunePtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("RuneValue", func(t *testing.T) {
		zeroPtr := RunePtr(rune(0))
		nonZeroPtr := RunePtr(rune('a'))

		cases := []struct {
			name     string
			actual   rune
			expected rune
		}{
			{
				name:     "zero",
				actual:   RuneValue(zeroPtr),
				expected: rune(0),
			},
			{
				name:     "non-zero",
				actual:   RuneValue(nonZeroPtr),
				expected: *nonZeroPtr,
			},
			{
				name:     "nil rune",
				actual:   RuneValue(nil),
				expected: rune(0),
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("BytePtr", func(t *testing.T) {
		zero := byte(0)
		nonZero := byte('a')

		cases := []struct {
			name     string
			actual   *byte
			expected byte
		}{
			{
				name:     "zero",
				actual:   BytePtr(zero),
				expected: zero,
			},
			{
				name:     "non-zero",
				actual:   BytePtr(nonZero),
				expected: nonZero,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, *c.actual)
			})
		}
	})

	t.Run("ByteValue", func(t *testing.T) {
		zeroPtr := BytePtr(byte(0))
		nonZeroPtr := BytePtr(byte('a'))

		cases := []struct {
			name     string
			actual   byte
			expected byte
		}{
			{
				name:     "zero",
				actual:   ByteValue(zeroPtr),
				expected: byte(0),
			},
			{
				name:     "non-zero",
				actual:   ByteValue(nonZeroPtr),
				expected: *nonZeroPtr,
			},
			{
				name:     "nil byte",
				actual:   ByteValue(nil),
				expected: byte(0),
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualIntPtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualIntPtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualIntPtr(nil, IntPtr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualIntPtr(IntPtr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualIntPtr(IntPtr(1), IntPtr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualIntPtr(IntPtr(1), IntPtr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualInt8Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualInt8Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualInt8Ptr(nil, Int8Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualInt8Ptr(Int8Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualInt8Ptr(Int8Ptr(1), Int8Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualInt8Ptr(Int8Ptr(1), Int8Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualInt16Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualInt16Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualInt16Ptr(nil, Int16Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualInt16Ptr(Int16Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualInt16Ptr(Int16Ptr(1), Int16Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualInt16Ptr(Int16Ptr(1), Int16Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualInt32Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualInt32Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualInt32Ptr(nil, Int32Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualInt32Ptr(Int32Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualInt32Ptr(Int32Ptr(1), Int32Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualInt32Ptr(Int32Ptr(1), Int32Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})
	t.Run("IsEqualInt64Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualInt64Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualInt64Ptr(nil, Int64Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualInt64Ptr(Int64Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualInt64Ptr(Int64Ptr(1), Int64Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualInt64Ptr(Int64Ptr(1), Int64Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualUintPtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualUintPtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualUintPtr(nil, UintPtr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualUintPtr(UintPtr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualUintPtr(UintPtr(1), UintPtr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualUintPtr(UintPtr(1), UintPtr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualUint8Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualUint8Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualUint8Ptr(nil, Uint8Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualUint8Ptr(Uint8Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualUint8Ptr(Uint8Ptr(1), Uint8Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualUint8Ptr(Uint8Ptr(1), Uint8Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualUint16Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualUint16Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualUint16Ptr(nil, Uint16Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualUint16Ptr(Uint16Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualUint16Ptr(Uint16Ptr(1), Uint16Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualUint16Ptr(Uint16Ptr(1), Uint16Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualUint32Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualUint32Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualUint32Ptr(nil, Uint32Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualUint32Ptr(Uint32Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualUint32Ptr(Uint32Ptr(1), Uint32Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualUint32Ptr(Uint32Ptr(1), Uint32Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualUint64Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualUint64Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualUint64Ptr(nil, Uint64Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualUint64Ptr(Uint64Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualUint64Ptr(Uint64Ptr(1), Uint64Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualUint64Ptr(Uint64Ptr(1), Uint64Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualFloat32Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualFloat32Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualFloat32Ptr(nil, Float32Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualFloat32Ptr(Float32Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualFloat32Ptr(Float32Ptr(1), Float32Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualFloat32Ptr(Float32Ptr(1), Float32Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualFloat64Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualFloat64Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualFloat64Ptr(nil, Float64Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualFloat64Ptr(Float64Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualFloat64Ptr(Float64Ptr(1), Float64Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualFloat64Ptr(Float64Ptr(1), Float64Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualBoolPtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualBoolPtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualBoolPtr(nil, BoolPtr(true)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualBoolPtr(BoolPtr(true), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualBoolPtr(BoolPtr(true), BoolPtr(true)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualBoolPtr(BoolPtr(true), BoolPtr(false)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualStringPtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualStringPtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualStringPtr(nil, StringPtr("foo")),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualStringPtr(StringPtr("foo"), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualStringPtr(StringPtr("foo"), StringPtr("foo")),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualStringPtr(StringPtr("foo"), StringPtr("bar")),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualTimePtr", func(t *testing.T) {
		now := time.Now()

		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualTimePtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualTimePtr(nil, TimePtr(now)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualTimePtr(TimePtr(now), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualTimePtr(TimePtr(now), TimePtr(now)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualTimePtr(TimePtr(now), TimePtr(now.Add(time.Hour))),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualDurationPtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualDurationPtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualDurationPtr(nil, DurationPtr(time.Hour)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualDurationPtr(DurationPtr(time.Hour), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualDurationPtr(DurationPtr(time.Hour), DurationPtr(time.Hour)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualDurationPtr(DurationPtr(time.Hour), DurationPtr(time.Hour*2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualComplex64Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualComplex64Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualComplex64Ptr(nil, Complex64Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualComplex64Ptr(Complex64Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualComplex64Ptr(Complex64Ptr(1), Complex64Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualComplex64Ptr(Complex64Ptr(1), Complex64Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualComplex128Ptr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualComplex128Ptr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualComplex128Ptr(nil, Complex128Ptr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualComplex128Ptr(Complex128Ptr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualComplex128Ptr(Complex128Ptr(1), Complex128Ptr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualComplex128Ptr(Complex128Ptr(1), Complex128Ptr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualRunePtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualRunePtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualRunePtr(nil, RunePtr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualRunePtr(RunePtr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualRunePtr(RunePtr(1), RunePtr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualRunePtr(RunePtr(1), RunePtr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

	t.Run("IsEqualBytePtr", func(t *testing.T) {
		cases := []struct {
			name     string
			actual   bool
			expected bool
		}{
			{
				name:     "both nil",
				actual:   IsEqualBytePtr(nil, nil),
				expected: true,
			},
			{
				name:     "first nil",
				actual:   IsEqualBytePtr(nil, BytePtr(1)),
				expected: false,
			},
			{
				name:     "second nil",
				actual:   IsEqualBytePtr(BytePtr(1), nil),
				expected: false,
			},
			{
				name:     "both non-nil, equal",
				actual:   IsEqualBytePtr(BytePtr(1), BytePtr(1)),
				expected: true,
			},
			{
				name:     "both non-nil, not equal",
				actual:   IsEqualBytePtr(BytePtr(1), BytePtr(2)),
				expected: false,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				assert.Equal(t, c.expected, c.actual)
			})
		}
	})

}
