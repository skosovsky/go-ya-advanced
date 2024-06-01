package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	MyType struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}

	TypeString struct {
		Val string `json:"value"`
	}

	TypeFloat struct {
		Val float64 `json:"value"`
	}
)

func (t *MyType) UnmarshalJSON(data []byte) error {
	type Alias MyType

	aliasValue := &struct {
		*Alias
		Value json.RawMessage `json:"value"`
	}{
		Alias: (*Alias)(t),
		Value: nil,
	}

	log.Println(t.Type)

	if err := json.Unmarshal(data, aliasValue); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	log.Println(t.Type)

	switch t.Type {
	case "string":
		t.Value = &TypeString{
			Val: "",
		}
	case "float":
		t.Value = &TypeFloat{
			Val: 0.0,
		}
	}

	if t.Value != nil {
		if err := json.Unmarshal(aliasValue.Value, t.Value); err != nil {
			return fmt.Errorf("unmarshal %T: %w", t.Value, err)
		}
	}

	return nil
}

func main() {
	orig := MyType{Type: "string", Value: TypeString{Val: "some string"}}

	origJSON, err := json.MarshalIndent(orig, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("orig:", string(origJSON))

	dupl := MyType{Type: "", Value: nil}  // dupl.value = interface | nil
	err = json.Unmarshal(origJSON, &dupl) // dupl.value = interface | map[string]interface
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("dupl: %T, %v", dupl, dupl)
	log.Printf("dupl value: %T, %v", dupl.Value, dupl.Value)
}
