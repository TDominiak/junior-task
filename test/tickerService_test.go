package domain

import (
	"testing"

	"github.com/TDominiak/junior-task/domain"
	"github.com/TDominiak/junior-task/domain/measurementUtils"
	"github.com/TDominiak/junior-task/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestTickerServiceStart(t *testing.T) {

	repo := repository.NewInMemortRepository()
	id := primitive.NewObjectID()
	d := &domain.Device{
		ID:       id,
		Name:     "test",
		Interval: 1,
		Value:    1,
	}
	repo.Save(d)
	serivce := domain.NewDeviceService(repo)
	chanMeasurement := make(chan measurementUtils.Measurement)
	publisher, _ := measurementUtils.GetPublisher("stdout", chanMeasurement)
	tickerService := domain.NewTickerService(serivce, publisher)
	tickerService.Start()

	result := <-chanMeasurement
	assert.Equal(t, measurementUtils.Measurement{Id: id.Hex(), Value: 1}, result)

}
