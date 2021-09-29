package domain

import "errors"


type TickerService interface {
	Start() error
	Stop() error
	
}
type tickerService struct {
	deviceService DeviceService
	publisher MeasurementPublisher
	done chan bool
}

func NewTickerService(deviceService DeviceService) TickerService {
	return &tickerService{deviceService: deviceService}
}


func (s *tickerService) Start() error {
	devices, err := s.deviceService.GetAll()
	if err != nil {
		return errors.New("failed to start measurements sending")
	}

	for _, device range devices {

	}

	return nil
}

func (s *tickerService) Stop() error {
	return nil
}

func (s *tickerService) tick(device Device) error {
	ticker := time.NewTicker(time.Second * time.Duration(device.Interval))

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			s.publisher.publish(device.ID.Hex(), device.Value)		
		}
	}
}

