package measurementUtils

import (
	"errors"
	"fmt"
	"log"
)

type MeasurementPublisher interface {
	Publish(id string, value float64) error
}

func GetPublisher(publisherType string, measurements chan Measurement) (MeasurementPublisher, error) {
	if publisherType == "stdout" {
		return newMeasurementStdout(measurements), nil
	}
	if publisherType == "rabbit" {
		return nil, errors.New("not implemented")
	}
	return nil, fmt.Errorf("wrong publisher type passed")
}

type MeasurementStdout struct {
	measurements chan Measurement
}

func newMeasurementStdout(measurements chan Measurement) MeasurementPublisher {
	return &MeasurementStdout{measurements: measurements}
}

func (m *MeasurementStdout) Publish(id string, value float64) error {
	log.Printf("Measurements send from device %v", id)
	m.measurements <- Measurement{id, value}
	return nil
}
