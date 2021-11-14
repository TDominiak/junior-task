package domain

type DeviceService interface {
	Save(device *Device) error
	GetByID(id string) (*Device, error)
	GetAll() ([]Device, error)
	Delete(id string) error
	
}


type service struct {
	deviceRepo Repository
}

func NewDeviceService(deviceRepo Repository) DeviceService {
	return &service{deviceRepo: deviceRepo}
}

func (s *service) Save(device *Device) error {
	return s.deviceRepo.Save(device)
}

func (s *service) GetByID(id string) (*Device, error) {
	return s.deviceRepo.GetByID(id)
}

func (s *service) GetAll() ([]Device, error) {
	return s.deviceRepo.GetAll()
}

func (s *service) Delete(id string) error {
	return s.deviceRepo.Delete(id)
}
