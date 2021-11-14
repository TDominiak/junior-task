package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TDominiak/junior-task/domain"
	"github.com/TDominiak/junior-task/repository"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSaveOK(t *testing.T) {

	handlerTest := NewDeviceHandler(domain.NewDeviceService(repository.NewInMemortRepository()))
	values := map[string]interface{}{"name": "test", "interval": 1, "value": 2.0}
	jsonValue, _ := json.Marshal(values)
	req, err := http.NewRequest(http.MethodPost, "/device", bytes.NewBuffer(jsonValue))
	req = req.WithContext(context.WithValue(req.Context(), "id", "123"))

	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handlerTest.Save(w, req)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestGetByIDOK(t *testing.T) {

	id := primitive.NewObjectID()
	inMemoryTest := repository.NewInMemortRepository()
	d := &domain.Device{
		ID:       id,
		Name:     "test",
		Interval: 1,
		Value:    1.0,
	}

	inMemoryTest.Save(d)

	handlerTest := NewDeviceHandler(domain.NewDeviceService(inMemoryTest))
	req, err := http.NewRequest(http.MethodPost, "/device", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"id": id.Hex(),
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerTest.GetByID)
	handler.ServeHTTP(rr, req)

	assert.EqualValues(t, http.StatusOK, rr.Code)
	out, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	assert.EqualValues(t, string(out), rr.Body.String())

}
