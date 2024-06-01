package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

type Slice []byte

func (s Slice) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(hex.EncodeToString(s))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal slice to JSON: %w", err)
	}

	return data, nil
}

func (s *Slice) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("json decoding failed: %w", err)
	}

	slice, err := hex.DecodeString(tmp)
	if err != nil {
		return fmt.Errorf("hex decoding failed: %w", err)
	}

	*s = slice

	return nil
}

type MySlice struct {
	ID    int
	Slice Slice
}

func main() {
	ret, err := json.Marshal(MySlice{ID: 7, Slice: []byte{1, 2, 3, 10, 11, 255}}) //nolint:mnd,musttag // example
	if err != nil {
		panic(err)
	}

	log.Println(string(ret))

	var result MySlice
	if err = json.Unmarshal(ret, &result); err != nil { //nolint:musttag // example
		panic(err)
	}

	log.Println(result)
}
