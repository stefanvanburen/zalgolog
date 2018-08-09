package main

import (
	"github.com/apex/log"
	"github.com/svanburen/zalgolog"
)

func main() {
	log.SetHandler(zalgolog.Default)
	for {
		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   "10",
		}).Info("To invoke the hive-mind representing chaos.")

		log.WithFields(log.Fields{
			"omg":    true,
			"number": 122,
		}).Warn("Invoking the feeling of chaos.")

		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   "10",
		}).Info("With out order.")

		log.WithFields(log.Fields{
			"animal": "walrus",
			"size":   "9",
		}).Error("The Nezperdian hive-mind of chaos. Zalgo.")

		log.WithFields(log.Fields{
			"omg":    true,
			"number": 100,
		}).Warn("He who Waits Behind The Wall.")

		log.Fatal("ZALGO !")
	}
}
