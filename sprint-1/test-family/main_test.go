package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFamily_AddNew(t *testing.T) { //nolint:paralleltest // consistent test
	type fields struct {
		r Relationship
		p Person
	}

	testCases := []struct {
		name string
		args fields
		want error
	}{
		{
			name: "First father",
			args: fields{
				r: "father",
				p: Person{
					FirstName: "John",
					LastName:  "Smith",
					Age:       30,
				},
			},
			want: nil,
		},
		{
			name: "Second father",
			args: fields{
				r: "father",
				p: Person{
					FirstName: "John",
					LastName:  "Dra",
					Age:       35,
				},
			},
			want: ErrRelationshipAlreadyExists,
		},
	}

	family := Family{
		Members: nil,
	}

	for _, tt := range testCases { //nolint:paralleltest // consistent test
		t.Run(tt.name, func(t *testing.T) {
			if got := family.AddNew(tt.args.r, tt.args.p); !errors.Is(got, tt.want) {
				t.Errorf("AddNew() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFamily_AddNewTestify(t *testing.T) { //nolint:paralleltest // consistent test
	type fields struct {
		r Relationship
		p Person
	}

	testCases := []struct {
		name string
		args fields
		want error
	}{
		{
			name: "First father",
			args: fields{
				r: "father",
				p: Person{
					FirstName: "John",
					LastName:  "Smith",
					Age:       30,
				},
			},
			want: nil,
		},
		{
			name: "Second father",
			args: fields{
				r: "father",
				p: Person{
					FirstName: "John",
					LastName:  "Dra",
					Age:       35,
				},
			},
			want: ErrRelationshipAlreadyExists,
		},
	}

	family := Family{
		Members: nil,
	}

	for _, tt := range testCases { //nolint:paralleltest // consistent test
		t.Run(tt.name, func(t *testing.T) {
			assert.ErrorIs(t, tt.want, family.AddNew(tt.args.r, tt.args.p))
		})
	}
}

func TestFamily_AddNewFull(t *testing.T) {
	t.Parallel()

	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers Family
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: Family{
				Members: map[Relationship]Person{ //nolint:exhaustive // false positive
					Mother: {
						FirstName: "Maria",
						LastName:  "Popova",
						Age:       36,
					},
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: Family{
				Members: map[Relationship]Person{ //nolint:exhaustive // false positive
					Father: {
						FirstName: "Misha",
						LastName:  "Popov",
						Age:       42,
					},
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			family := &tt.existedMembers
			err := family.AddNew(tt.newPerson.r, tt.newPerson.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFamily_AddNewFullTestify(t *testing.T) {
	t.Parallel()

	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers Family
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: Family{
				Members: map[Relationship]Person{ //nolint:exhaustive // false positive
					Mother: {
						FirstName: "Maria",
						LastName:  "Popova",
						Age:       36,
					},
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: Family{
				Members: map[Relationship]Person{ //nolint:exhaustive // false positive
					Father: {
						FirstName: "Misha",
						LastName:  "Popov",
						Age:       42,
					},
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			family := &tt.existedMembers
			err := family.AddNew(tt.newPerson.r, tt.newPerson.p)
			if !tt.wantErr {
				require.NoError(t, err)
				assert.Contains(t, family.Members, tt.newPerson.r)

				return
			}

			assert.Error(t, err)
		})
	}
}
