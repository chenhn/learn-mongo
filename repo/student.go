package repo

import (
	"context"
	"fmt"
	"learn_mongo/config"
	"learn_mongo/model"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	StRepo StudentRepo
)

func StudentRepoInit() {
	StRepo = StudentRepo{
		coll: config.GetMongoClient().Database(viper.GetString("mongo.database")).Collection("Student"),
	}
}

type StudentRepo struct {
	coll *mongo.Collection
}

func (repo StudentRepo) Create(s model.Student) error {
	x, err := repo.coll.InsertOne(context.TODO(), s)
	if err != nil {
		return err
	}
	fmt.Println(*x)
	return nil
}

func (repo StudentRepo) Search(s bson.D) ([]model.Student, error) {
	cursor, err := repo.coll.Find(context.TODO(), s)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var sts []model.Student
	err = cursor.All(context.TODO(), &sts)
	if err != nil {
		return nil, err
	}
	return sts, nil
}
