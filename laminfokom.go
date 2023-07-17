package hadbackend

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDataAkreditas(db string, dataakreditas Akreditas) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dataakreditas").InsertOne(context.TODO(), dataakreditas)
	if err != nil {
		fmt.Printf("InsertDataAkreditas: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataProgramStudi(db string, dataprogramstudi ProgramStudi) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("dataprogramstudi").InsertOne(context.TODO(), dataprogramstudi)
	if err != nil {
		fmt.Printf("InsertDataProgramStudi: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertProfile(db string, profile Profile) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection("profile").InsertOne(context.TODO(), profile)
	if err != nil {
		fmt.Printf("InsertProfile: %v\n", err)
	}
	return insertResult.InsertedID
}

//	func GetDataDataAkreditas(stats string) (data []DataAkreditas) {
//		user := MongoConnect("infokomdb").Collection("dataakreditas").
//		// filter := bson.M{"status": stats}
//		// cursor, err := user.Find(context.TODO(), filter)
//		fmt.Println("GetDataDataAkreditas :", err)
//		if err != nil {
//		}
//		err = cursor.All(context.TODO(), &data)
//		if err != nil {
//			fmt.Println(err)
//		}
//		return
//	}

func GetDataAkreditas(stats string) (data []Akreditas) {
	user := MongoConnect("infokomdb").Collection("dataAkreditas")
	filter := bson.M{"status": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataAkreditas :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// func GetDataProgramStudi(stats string) (data []DataProgramStudi) {
// 	user := MongoConnect("infokomdb").Collection("dataprogramstudi").
// 	filter := bson.M{"programstudi": stats}
// 	cursor, err := user.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetDataDataProgramStudi :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

func GetDataProgramStudi(stats string) (data []ProgramStudi) {
	user := MongoConnect("infokomdb").Collection("dataprogramstudi")
	filter := bson.M{"program": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataProgramStudi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// func GetDataProfile(stats string) (data []Profile) {
// 	user := MongoConnect("infokomdb").Collection("profile").
// 	filter := bson.M{"isi_satu": stats}
// 	cursor, err := user.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetDataProfile :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

func GetDataProfile(stats string) (data []Profile) {
	user := MongoConnect("infokomdb").Collection("Profile")
	filter := bson.M{"isi_satu": stats}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetDataProgramStudi :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
