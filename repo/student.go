package repo

import (
	"context"
	"fmt"
	"learn_mongo/config"
	"learn_mongo/model"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	studentRepo StudentRepo
)

func StudentRepoInit() {
	config.GetMongoClient().Database(viper.GetString("mongo.database"))
}

type StudentRepo struct {
	coll *mongo.Collection
}

func (s StudentRepo) Create() error {
	student := model.Student{Name: "chn"}
	x, err := s.coll.InsertOne(context.TODO(), student)
	if err != nil {
		return err
	}
	fmt.Println(x)
	return nil
}
