package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/TDominiak/junior-task/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func newMongoClient(mongoUrl string, timeout int) (*mongo.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoUrl, mongoDb string, timeout int) domain.Repository {

	mongoClient, err := newMongoClient(mongoUrl, timeout)
	if err != nil {
		log.Printf("Can not create Mongo client: %s", err)
		os.Exit(1)
	}

	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}

	return repo

}

func (i *mongoRepository) Save(device *domain.Device) error {
	ctx, cancel := context.WithTimeout(context.Background(), i.timeout)
	defer cancel()
	collection := i.client.Database(i.db).Collection("devices")
	_, err := collection.InsertOne(ctx, device)
	if err != nil {
		return err
	}
	return nil
}

func (i *mongoRepository) GetByID(id string) (*domain.Device, error) {
	var device domain.Device
	ctx, cancel := context.WithTimeout(context.Background(), i.timeout)
	defer cancel()
	collection := i.client.Database(i.db).Collection("devices")
	deviceId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": deviceId}
	err = collection.FindOne(ctx, filter).Decode(&device)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// implent const err
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, err
	}

	return &device, nil
}

func (i *mongoRepository) GetAll() ([]domain.Device, error) {

	collection := i.client.Database(i.db).Collection("devices")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var results []domain.Device
	if err = cur.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results, nil

}

func (i *mongoRepository) Delete(id string) error {

	return nil

}
