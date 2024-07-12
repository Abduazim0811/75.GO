package mongodb

import (
	"context"
	"time"

	"75.GO/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StudentMongoDb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewStudent(uri, dbname, collectionName string) (*StudentMongoDb, error) {
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbname).Collection(collectionName)
	return &StudentMongoDb{client: client, collection: collection}, nil
}

func (s *StudentMongoDb) StoreNewStudents(student *models.Student) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.collection.InsertOne(ctx, student)
}

func (s *StudentMongoDb) StoreGetbyIdStudent(id primitive.ObjectID) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var student models.Student
	err := s.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
func (s *StudentMongoDb) StoreUpdateStudent(id primitive.ObjectID, updatedStudent *models.Student) (*mongo.UpdateResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": updatedStudent,
	}

	return  s.collection.UpdateOne(ctx, bson.M{"_id":id}, update)
}

func (s *StudentMongoDb) StoreDeleteStudent(id primitive.ObjectID) (*mongo.DeleteResult, error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.collection.DeleteOne(ctx, bson.M{"_id":id})
}