package repository

import (
	"testing"

	"github.com/TDominiak/junior-task/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateDeviceOK(t *testing.T) {
	inMemoryTest := NewInMemortRepository()
	d := &domain.Device{
		ID:       primitive.NewObjectID(),
		Name:     "test",
		Interval: 1,
		Value:    1,
	}

	inMemoryTest.Save(d)

}

func TestGetByIDOK(t *testing.T) {
	inMemoryTest := NewInMemortRepository()
	id := primitive.NewObjectID()
	d := &domain.Device{
		ID:       id,
		Name:     "test",
		Interval: 1,
		Value:    1,
	}
	inMemoryTest.Save(d)
	device, err := inMemoryTest.GetByID(string((id.Hex())))
	assert.Nil(t, err)
	assert.Equal(t, device, d)
}

func TestGetByIDNoFound(t *testing.T) {
	inMemoryTest := NewInMemortRepository()
	id := primitive.NewObjectID()
	d := &domain.Device{
		ID:       id,
		Name:     "test",
		Interval: 1,
		Value:    1,
	}
	inMemoryTest.Save(d)
	device, _ := inMemoryTest.GetByID(string((primitive.NewObjectID().Hex())))
	assert.Nil(t, device)
}

func TestGetAll(t *testing.T) {
	inMemoryTest := NewInMemortRepository()
	d := &domain.Device{
		ID:       primitive.NewObjectID(),
		Name:     "test",
		Interval: 1,
		Value:    1,
	}
	inMemoryTest.Save(d)
	device, _ := inMemoryTest.GetByID(string((primitive.NewObjectID().Hex())))
	assert.Nil(t, device)
}
