package domain

import (
	"errors"
	"time"

	"github.com/TDominiak/junior-task/domain/measurementUtils"
)

type TickerService interface {
	Start() error
	Stop() error
}
type tickerService struct {
	deviceService DeviceService
	publisher     measurementUtils.MeasurementPublisher
	done          chan bool
}

func NewTickerService(deviceService DeviceService, publisher measurementUtils.MeasurementPublisher) TickerService {
	return &tickerService{deviceService: deviceService, publisher: publisher, done: make(chan bool)}
}

func (s *tickerService) Start() error {
	devices, err := s.deviceService.GetAll()
	if err != nil {
		return errors.New("failed to start measurements sending")
	}

	for _, device := range devices {
		go func(d Device) {
			s.tick(d)
		}(device)

	}

	return nil
}

func (s *tickerService) Stop() error {
	close(s.done)
	return nil
}

func (s *tickerService) tick(device Device) {
	ticker := time.NewTicker(time.Second * time.Duration(device.Interval))

	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			s.publisher.Publish(device.ID.Hex(), device.Value)
		}
	}
}
