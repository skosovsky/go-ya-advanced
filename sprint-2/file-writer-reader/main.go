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

type Producer struct {
	file   *os.File
	writer *bufio.Writer
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return &Producer{
		file:   file,
		writer: bufio.NewWriter(file),
	}, nil
}

func (p *Producer) WriteEvent(event *Event) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return fmt.Errorf("error marshalling event: %w", err)
	}

	if _, err = p.writer.Write(data); err != nil {
		return fmt.Errorf("error writing to buf: %w", err)
	}

	if err = p.writer.WriteByte('\n'); err != nil {
		return fmt.Errorf("error writing bytes to buf: %w", err)
	}

	if err = p.writer.Flush(); err != nil {
		return fmt.Errorf("error flushing buf: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	err := p.file.Close()
	if err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}

	return nil
}

type Consumer struct {
	file   *os.File
	reader *bufio.Reader
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return &Consumer{
		file:   file,
		reader: bufio.NewReader(file),
	}, nil
}

func (c *Consumer) ReadEvent() (*Event, error) {
	data, err := c.reader.ReadBytes('\n')
	if err != nil {
		return nil, fmt.Errorf("error reading from buf: %w", err)
	}

	var event Event
	if err = json.Unmarshal(data, &event); err != nil {
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
