package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

const MockXMLDocument = `
<?xml version="1.0" encoding="UTF-8"?>
<storage-report version="1.0" exported-on="2021-01-31" location-id="001">
<item barcode="000000000001">
  <quantity>100</quantity>
</item>
<item barcode="000000000002">
  <quantity>500</quantity>
</item>
</storage-report>
`

type Item struct {
	XMLName  xml.Name `xml:"item"`
	Barcode  string   `xml:"barcode,attr"`
	Quantity int64    `xml:"quantity"`
}

func StorageReportStream(stream io.Reader) error {
	decoder := xml.NewDecoder(stream)

	for {
		xmlToken, err := decoder.Token()
		if err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("xml parse error: %w", err)
		}

		if errors.Is(err, io.EOF) {
			break
		}

		xmlElement, ok := xmlToken.(xml.StartElement)
		if !ok || xmlElement.Name.Local != "item" {
			continue
		}

		var item Item
		if err = decoder.DecodeElement(&item, &xmlElement); err != nil {
			return fmt.Errorf("xml decode error: %w", err)
		}

		HandleReportItem(item)
	}

	return nil
}

func HandleReportItem(item Item) {
	log.Println("Handle position:", item.Barcode)
	log.Println("Quantity:", item.Quantity)
}

func main() {
	reader := strings.NewReader(MockXMLDocument)
	if err := StorageReportStream(reader); err != nil {
		log.Fatal(err)
	}
}
