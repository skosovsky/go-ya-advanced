package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type (
	RaceReport struct {
		XMLName      xml.Name      `xml:"report"`
		Competition  Competition   `xml:"competition"`
		RacerResults []RacerResult `xml:"racer"`
	}

	Competition struct {
		XMLName  xml.Name `xml:"competition"`
		Location string   `xml:"location"`
		Class    string   `xml:"class"`
	}

	RacerResult struct {
		XMLName   xml.Name `xml:"racer"`
		GlobalID  int      `xml:"global_id,attr,omitempty"`
		Nick      string   `xml:"nick"`
		BestLapMs int64    `xml:"best_lap_ms"`
		Laps      float32  `xml:"laps"`
		Comment   string   `xml:",comment"`
	}
)

func FilterXML(input string, laps float32) (string, error) {
	var report RaceReport

	if err := xml.Unmarshal([]byte(input), &report); err != nil {
		return "", fmt.Errorf("error unmarshalling xml: %w", err)
	}

	filter := make([]RacerResult, 0, len(report.RacerResults))
	for _, racer := range report.RacerResults {
		if racer.Laps > laps {
			filter = append(filter, racer)
		}
	}

	report.RacerResults = filter

	var data []byte
	data, err := xml.MarshalIndent(report, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling xml: %w", err)
	}

	return string(data), nil
}

func main() {
	xmlData := `<report>
  <competition>
    <location>РФ, Санкт-Петербург, Дворец творчества юных техников</location>
    <class>ТА-24</class>
  </competition>
  <racer global_id="100">
    <nick>RacerX</nick>
    <best_lap_ms>61012</best_lap_ms>
    <laps>52.3</laps>
  </racer>
  <racer global_id="127">
    <nick>Иван The Шумахер</nick>
    <best_lap_ms>61023</best_lap_ms>
    <laps>51</laps>
  </racer>
  <racer global_id="203">
    <nick>Петя Иванов</nick>
    <best_lap_ms>63000</best_lap_ms>
    <laps>49.9</laps>
    <!--Болид не соответствует техническому регламенту,
    результат не учитывается в общем рейтинге-->
  </racer>
  <racer>
    <nick>Гость 1</nick>
    <best_lap_ms>123001</best_lap_ms>
    <laps>25.8</laps>
  </racer>
</report>`

	answer, err := FilterXML(xmlData, 50) //nolint:mnd // example
	if err != nil {
		log.Fatalf("error filtering xml: %v", err)
	}

	log.Println(answer)
}
