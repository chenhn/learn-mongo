package model

type Student struct {
	Name      string        `bson:"name"`
	StudentId string        `bson:"student_id"`
	Sex       bool          `bson:"sex"`
	Address   string        `bson:"address"`
	Grades    []interface{} `bson:"grades"`
}

type Restaurant struct {
	Name         string
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}
