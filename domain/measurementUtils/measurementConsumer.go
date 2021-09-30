package measurementUtils

import "log"

type MeasurementConsumer struct {
	Measurements chan Measurement
}

func (m *MeasurementConsumer) Start() error {
	go func() {
		for m := range m.Measurements {
			log.Printf("Measurements recaived from device %v", m.Id)

		}
	}()

	return nil
}
