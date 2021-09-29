package domain

import (
	"fmt"
	"log"
)

type MeasurementPublisher interface {
	publish(id string, value float64) error
}

type MeasurementStdout struct {}

func (m *MeasurementStdout) publish(id string, value float64) error {
	log.Printf("Measurements send from device %v, value %v", id, value)
	return nil
}


