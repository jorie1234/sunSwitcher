package main

import (
	"fmt"
	"github.com/kelvins/sunrisesunset"
	"log"
	"time"
)

func main() {

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Println(err)
	}
	t := time.Now().In(loc)
	zone, offset := t.Zone()
	fmt.Println(zone, offset)

	mqtt := NewMqtt("tcp://192.168.178.93:1883", "sunSwitch")
	for {
		t := time.Now().In(loc)
		// You can use the Parameters structure to set the parameters
		p := sunrisesunset.Parameters{
			Latitude:  52.297754,
			Longitude: 9.940568,
			UtcOffset: float64(offset / 3600),
			Date:      t,
		}
		// Calculate the sunrise and sunset times
		sunrise, sunset, err := p.GetSunriseSunset()
		sunrise = time.Date(t.Year(), t.Month(), t.Day(), sunrise.Hour(), sunrise.Minute(), sunrise.Second(), 0, t.Location())
		sunset = time.Date(t.Year(), t.Month(), t.Day(), sunset.Hour(), sunset.Minute(), sunset.Second(), 0, t.Location())
		// If no error has occurred,
		if err == nil {
			log.Println("Sunrise:", sunrise.Format("15:04:05")) // Sunrise: 06:11:44
			log.Println("Sunset:", sunset.Format("15:04:05"))   // Sunset: 18:14:27
		} else {
			log.Println(err)
		}

		if t.Before(sunrise) {
			//Wait until sunrise....
			timeToSunrise := sunrise.Sub(t)
			log.Printf("waiting for sunrise in %s", timeToSunrise)
			time.Sleep(timeToSunrise)
			log.Print("Lampe ausschalten")
			SetPowerState(mqtt, "cmnd/sonoff_6269/POWER", "0")
			SetPowerState(mqtt, "cmnd/sonoff_8125/POWER", "0")
		}

		t = time.Now().In(loc)
		if t.Before(sunset) {
			//Wait until sunset....
			timeToSunset := sunset.Sub(t)
			log.Printf("waiting for sunset in %s", timeToSunset)
			time.Sleep(timeToSunset)
			log.Print("Lampe einschalten")
			SetPowerState(mqtt, "cmnd/sonoff_6269/POWER", "1")
			SetPowerState(mqtt, "cmnd/sonoff_8125/POWER", "1")
		}
		//Wait for next day.....
		t = time.Now().In(loc)
		timeToTomorrow := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).Add(26 * time.Hour).Sub(t)
		log.Printf("waiting for tomorrow in %s", timeToTomorrow)
		time.Sleep(timeToTomorrow)

	}
}

func SetPowerState(mqtt *mqttClient, device, state string) error {
	err := mqtt.Publish(device, state)
	if err != nil {
		log.Printf("mqtt publish %s state %s error %s", device, state, err)
		err = mqtt.Publish(device, state)
		if err != nil {
			log.Printf("2. try: mqtt publish %s state %s error %s", device, state, err)
			return err
		}
	}
	return nil
}
