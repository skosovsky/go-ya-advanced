package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "zero value",
			arg:  0,
			want: 0,
		},
		{
			name: "negative value",
			arg:  -1,
			want: 1,
		},
		{
			name: "positive value",
			arg:  1,
			want: 1,
		},
		{
			name: "small value",
			arg:  -0.00000000001,
			want: 0.00000000001,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Abs(tt.arg); got != tt.want {
				t.Errorf("Abs(%f) = %f, want %f", tt.arg, got, tt.want)
			}
		})
	}
}

func TestAbsTestify(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		arg  float64
		want float64
	}{
		{
			name: "zero value",
			arg:  0,
			want: 0,
		},
		{
			name: "negative value",
			arg:  -1,
			want: 1,
		},
		{
			name: "positive value",
			arg:  1,
			want: 1,
		},
		{
			name: "small value",
			arg:  -0.00000000001,
			want: 0.00000000001,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.InDelta(t, tt.want, Abs(tt.arg), 0.0001)
			assert.InEpsilon(t, tt.want, Abs(tt.arg), 0.02)
		})
	}
}
