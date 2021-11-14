package repository

import (
	"errors"
	"sync"
	"github.com/TDominiak/junior-task/domain"
)

var ErrNotFound = errors.New("not found")

type inMemoryRepository struct {
	mu      sync.Mutex
	devices map[string]domain.Device
}

func NewInMemortRepository() domain.Repository {
	inMemoryRepo := inMemoryRepository{devices: make(map[string]domain.Device)}
	return &inMemoryRepo

}

func (i *inMemoryRepository) Save(device *domain.Device) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.devices[device.ID.Hex()] = *device

	return nil
}

func (i *inMemoryRepository) GetByID(id string) (*domain.Device, error) {
	if device, found := i.devices[id]; found {
		return &device, nil
	} else {
		return nil, ErrNotFound
	}
}

func (i *inMemoryRepository) GetAll() ([]domain.Device, error) {
	all := make([]domain.Device, 0, len(i.devices))

	for _, value := range i.devices {
		all = append(all, value)
	}
	return all, nil
}

func (i *inMemoryRepository) Delete(id string) error {

	if _, ok := i.devices[id]; ok {
		delete(i.devices, id)
		return nil
	} else {
		return errors.New("no id")
	}
}
