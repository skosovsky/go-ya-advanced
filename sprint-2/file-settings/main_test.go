package main

import (
	"os"
	"testing"
)

func TestSettings(t *testing.T) {
	t.Parallel()

	filename := "settings.json"
	settings := Settings{
		Port: 8080,
		Host: "localhost",
	}

	if err := settings.Save(filename); err != nil {
		t.Error(err)
	}

	var result = new(Settings)
	if err := result.Load(filename); err != nil {
		t.Error(err)
	}

	if settings != *result {
		t.Errorf("expected settings to be equal: %+v != %+v", settings, result)
	}

	if err := os.Remove(filename); err != nil {
		t.Error(err)
	}
}
