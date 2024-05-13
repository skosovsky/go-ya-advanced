package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullName(t *testing.T) {
	t.Parallel()

	type fields struct {
		FirstName string
		LastName  string
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			fields: fields{
				FirstName: "",
				LastName:  "",
			},
			want: " ",
		},
		{
			name: "only first name",
			fields: fields{
				FirstName: "John",
				LastName:  "",
			},
			want: "John ",
		},
		{
			name: "only last name",
			fields: fields{
				FirstName: "",
				LastName:  "Smith",
			},
			want: " Smith",
		},
		{
			name: "both names",
			fields: fields{
				FirstName: "John",
				LastName:  "Smith",
			},
			want: "John Smith",
		},
		{
			name: "long names",
			fields: fields{
				FirstName: "John de Patrik",
				LastName:  "Smith Milano",
			},
			want: "John de Patrik Smith Milano",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			user := User{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}
			if got := user.FullName(); got != tt.want {
				t.Errorf("FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFullNameTestify(t *testing.T) {
	t.Parallel()

	type fields struct {
		FirstName string
		LastName  string
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty",
			fields: fields{
				FirstName: "",
				LastName:  "",
			},
			want: " ",
		},
		{
			name: "only first name",
			fields: fields{
				FirstName: "John",
				LastName:  "",
			},
			want: "John ",
		},
		{
			name: "only last name",
			fields: fields{
				FirstName: "",
				LastName:  "Smith",
			},
			want: " Smith",
		},
		{
			name: "both names",
			fields: fields{
				FirstName: "John",
				LastName:  "Smith",
			},
			want: "John Smith",
		},
		{
			name: "long names",
			fields: fields{
				FirstName: "John de Patrik",
				LastName:  "Smith Milano",
			},
			want: "John de Patrik Smith Milano",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			user := User{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}

			assert.Equal(t, tt.want, user.FullName())
		})
	}
}
