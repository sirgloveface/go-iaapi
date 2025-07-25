package repository

import (
	"context"
	"errors"

	"time"

	"github.com/sirgloveface/go-iaapi/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(client *mongo.Client) *MongoTaskRepository {
	collection := client.Database("goapi").Collection("tasks")
	return &MongoTaskRepository{collection: collection}
}

func (r *MongoTaskRepository) Create(task model.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, task)
	return err
}

func (r *MongoTaskRepository) FindAll() ([]model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []model.Task
	for cursor.Next(ctx) {
		var task model.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *MongoTaskRepository) FindByID(id string) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task model.Task
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}
