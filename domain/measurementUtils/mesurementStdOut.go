package measurementUtils

import (
	"log"

	"github.com/influxdata/influxdb-client-go/api"
)

type MeasurementStdoutExchanger struct {
	Writer       api.WriteAPIBlocking
	measurements chan Measurement
}

func NewMeasurementStdoutExchanger(measurements chan Measurement) (MeasurementExchanger, error) {
	return &MeasurementStdoutExchanger{Writer: NewClientInflux(), measurements: measurements}, nil
}

func (m *MeasurementStdoutExchanger) Publish(id string, value float64) error {
	log.Printf("Measurements send from device %v", id)
	m.measurements <- Measurement{id, value}
	return nil
}

func (ms *MeasurementStdoutExchanger) StartConsuming() error {

	go func() {
		for m := range ms.measurements {
			savePoint(m, ms.Writer)
			log.Printf("Measurements recaived from device %v", m.Id)
		}
	}()

	return nil
}
