package main

import (
	"fmt"
	"os"
	"protobuf/model"

	"google.golang.org/protobuf/encoding/protojson"
)

var user1 = &model.User{
	Id:       "01",
	Name:     "Jalaluddin",
	Password: "007",
	Gender:   model.UserGender_MALE,
}
var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  1,
		Longitude: 1,
	},
}
var garageList = &model.GarageList{
	List: []*model.Garage{
		garage,
	},
}
var garageListBuyer = &model.GarageListBuyer{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {
	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %s \n", user1.String())

	// =========== as JSON string
	jsonb, err1 := protojson.Marshal(garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As JSON\n       %s \n", string(jsonb))

	// json string to proto object
	protoObject := new(model.GarageList)
	err2 := protojson.Unmarshal(jsonb, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As json string to proto\n       %s \n", protoObject.String())

}
