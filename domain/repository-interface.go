package domain


type Repository interface {
	Save(device *Device) error
	GetByID(id string) (*Device, error)
	GetAll() ([]Device, error)
	Delete(id string) error
}

