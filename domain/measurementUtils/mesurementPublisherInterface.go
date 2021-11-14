package measurementUtils

type MeasurementExchanger interface {
	Publish(id string, value float64) error
	StartConsuming() error
}
