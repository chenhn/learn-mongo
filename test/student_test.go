package test

import (
	"fmt"
	"learn_mongo/config"
	"learn_mongo/model"
	"learn_mongo/repo"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestCreate(t *testing.T) {
	err := repo.StRepo.Create(model.Student{Name: "fuck", Address: "chengdu", Sex: false, StudentId: "1245"})
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second * 3)
}

func TestFind(t *testing.T) {
	s, err := repo.StRepo.Search(bson.D{{"name", "fuck"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func TestMain(m *testing.M) {
	os.Setenv("test", "1")
	config.Init()
	config.InitMongoClient()
	repo.InitRepo()
	m.Run()
}
