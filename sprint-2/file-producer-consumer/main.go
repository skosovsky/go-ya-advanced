package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Event struct {
	ID       uint    `json:"id"`
	CarModel string  `json:"carModel"`
	Price    float64 `json:"price"`
}

type Producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	encoder := json.NewEncoder(file)

	return &Producer{
		file:    file,
		encoder: encoder,
	}, nil
}

func (p *Producer) WriteEvent(event *Event) error {
	err := p.encoder.Encode(event)
	if err != nil {
		return fmt.Errorf("error encoding event: %w", err)
	}

	return nil
}

func (p *Producer) Close() error {
	if err := p.file.Close(); err != nil {
		return fmt.Errorf("error closing producer: %w", err)
	}

	return nil
}

type Consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(filename string) (*Consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	decoder := json.NewDecoder(file)

	return &Consumer{
		file:    file,
		decoder: decoder,
	}, nil
}

func (c *Consumer) ReadEvent() (*Event, error) {
	var event Event
	err := c.decoder.Decode(&event)
	if err != nil {
		return nil, fmt.Errorf("error decoding event: %w", err)
	}

	return &event, nil
}

func (c *Consumer) Close() error {
	if err := c.file.Close(); err != nil {
		return fmt.Errorf("error closing consumer: %w", err)
	}

	return nil
}

func main() {
	var events = []*Event{
		{ID: 1, CarModel: "Lada", Price: 400000},       //nolint:mnd // example
		{ID: 2, CarModel: "Mitsubishi", Price: 650000}, //nolint:mnd // example
		{ID: 3, CarModel: "Toyota", Price: 800000},     //nolint:mnd // example
		{ID: 4, CarModel: "BMW", Price: 875000},        //nolint:mnd // example
		{ID: 5, CarModel: "Mercedes", Price: 999999},   //nolint:mnd // example
	}

	fileName := "events.log"

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Printf("error removing file: %v", err)
		}
	}(fileName)

	producer, err := NewProducer(fileName)
	if err != nil {
		log.Panicf("error creating producer: %v", err)
	}

	defer func(producer *Producer) {
		err = producer.Close()
		if err != nil {
			log.Panicf("error closing producer: %v", err)
		}
	}(producer)

	consumer, err := NewConsumer(fileName)
	if err != nil {
		log.Printf("error creating consumer: %v", err)
	}

	defer func(consumer *Consumer) {
		err = consumer.Close()
		if err != nil {
			log.Printf("error closing consumer: %v", err)
		}
	}(consumer)

	for _, event := range events {
		if err = producer.WriteEvent(event); err != nil {
			log.Printf("error writing event: %v", err)
		}

		var readEvent *Event
		readEvent, err = consumer.ReadEvent()
		if err != nil {
			log.Printf("error reading event: %v", err)
		}

		log.Printf("event: %v", readEvent)
	}
}
