package measurementUtils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

type Measurement struct {
	Id    string
	Value float64
}

func savePoint(m Measurement, mw api.WriteAPIBlocking) {

	point := influxdb2.NewPointWithMeasurement("deviceValues").
		AddTag("deviceId", m.Id).
		AddField("value", m.Value).
		SetTime(time.Now().Round(time.Second))

	if err := mw.WritePoint(context.Background(), point); err != nil {
		log.Print(err)
	}

}

func NewClientInflux() api.WriteAPIBlocking {

	client := influxdb2.NewClient(os.Getenv("INFLUX_URL"),
		fmt.Sprintf("%s:%s", os.Getenv("INFLUXDB_INIT_USERNAME"), os.Getenv("INFLUXDB_INIT_PASSWORD")))
	if _, err := client.Health(context.Background()); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	writeAPI := client.WriteAPIBlocking(os.Getenv("INFLUXDB_INIT_ORG"), os.Getenv("INFLUXDB_INIT_BUCKET"))

	return writeAPI
}
