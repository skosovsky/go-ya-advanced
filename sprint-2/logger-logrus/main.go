package main

import (
	"os"

	rlog "github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		rlog.Fatal(err)
	}

	defer file.Close()

	rlog.SetOutput(file)
	rlog.SetFormatter(&rlog.JSONFormatter{}) //nolint:exhaustruct // example
	rlog.SetLevel(rlog.WarnLevel)

	rlog.WithFields(rlog.Fields{
		"genre": "metal",
		"name":  "Rammstein",
	}).Info("Немецкая метал-группа, образованная в январе 1994 года в Берлине.")

	rlog.WithFields(rlog.Fields{
		"omg":  true,
		"name": "Garbage",
	}).Warn("В 2021 году вышел новый альбом No Gods No Masters.")

	rlog.WithFields(rlog.Fields{
		"omg":  true,
		"name": "Linkin Park",
	}).Fatal("Группа Linkin Park взяла паузу после смерти вокалиста Честера Беннингтона 20 июля 2017 года.")
}
