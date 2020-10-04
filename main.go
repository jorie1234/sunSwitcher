package main

import (
"fmt"
	"log"
	"time"
"github.com/kelvins/sunrisesunset"
)

func main() {

	t := time.Now()
	zone, offset := t.Zone()
	fmt.Println(zone, offset)

	// You can use the Parameters structure to set the parameters
	p := sunrisesunset.Parameters{
		Latitude:  52.297754,
		Longitude: 9.940568,
		UtcOffset: float64(offset/3600),
		Date:      t,
	}
	// Calculate the sunrise and sunset times
	sunrise, sunset, err := p.GetSunriseSunset()

	// If no error has occurred,
	if err == nil {
		log.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
		log.Println("Sunset:", sunset.Format("15:04:05")) // Sunset: 18:14:27
	} else {
		log.Println(err)
	}
 mqtt:=NewMqtt("ubuntu", "sunSwitch")
	for {
		// Calculate the sunrise and sunset times
		sunrise, sunset, err = p.GetSunriseSunset()

		t := time.Now()
		if t.Before(sunrise) {
			//Wait until sunrise....
			timeToSunrise:=sunrise.Sub(t)
			log.Printf("waiting for sunrise in %s", timeToSunrise)
			time.Sleep(timeToSunrise)
			log.Print("Lampe ausschalten")
			mqtt.Publish("cmnd/sonoff_6269/POWER", "0")
		}

		t = time.Now()
		if t.Before(sunset) {
			//Wait until sunset....
			timeToSunset:=sunset.Sub(t)
			log.Printf("waiting for sunset in %s", timeToSunset)
			time.Sleep(timeToSunset)
			log.Print("Lampe einschalten")
			mqtt.Publish("cmnd/sonoff_6269/POWER", "1")
		}
		//Wait for next day.....
		t = time.Now()
		timeToTomorrow:=time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Add(26*time.Hour).Sub(t)
		log.Printf("waiting for tomorrow in %s", timeToTomorrow)
		time.Sleep(timeToTomorrow)

	}
}