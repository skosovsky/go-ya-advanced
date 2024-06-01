package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Event struct {
	ID int `json:"id"`
}

type Consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return &Consumer{
		file:    file,
		scanner: bufio.NewScanner(file),
	}, nil
}

func (c *Consumer) ReadEvent() (*Event, error) {
	if !c.scanner.Scan() {
		err := c.scanner.Err()

		return nil, fmt.Errorf("error scanning event: %w", err)
	}

	data := c.scanner.Bytes()

	var event Event
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, fmt.Errorf("error unmarshalling event: %w", err)
	}

	return &event, nil
}

func (c *Consumer) Close() error {
	err := c.file.Close()
	if err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}

	return nil
}
