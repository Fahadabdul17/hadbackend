package hadbackend

import (
	"fmt"
	"testing"
)

// func TestInsertDataAkreditas(t *testing.T) {
// 	dbname := "infokomdb"
// 	dataakreditas := DataAkreditas{
// 		ID:               primitive.NewObjectID(),
// 		Perguruan_Tinggi: "ULBI",
// 		Program_studi:    "Teknik Informatika",
// 		Peringkat:        "B",
// 		Status:           "Masih Berlaku",
// 	}
// 	insertedID := InsertDataAkreditas(dbname, dataakreditas)
// 	if insertedID == nil {
// 		t.Error("Failed to insert user")
// 	}
// }

// func TestInsertDataProgramStudi(t *testing.T) {
// 	dbname := "infokomdb"
// 	dataprogramstudi := DataProgramStudi{
// 		ID:            primitive.NewObjectID(),
// 		Program_studi: "Teknik Informatika",
// 		Program:       "Sarjana Terapan",
// 	}
// 	insertedID := InsertDataProgramStudi(dbname, dataprogramstudi)
// 	if insertedID == nil {
// 		t.Error("Failed to insert surat")
// 	}
// }

// func TestInsertProfile(t *testing.T) {
// 	dbname := "infokomdb"
// 	profile := Profile{
// 		ID:       primitive.NewObjectID(),
// 		Isi_satu: "kampus merdeka",
// 	}
// 	insertedID := InsertProfile(dbname, profile)
// 	if insertedID == nil {
// 		t.Error("Failed to insert lokasi")
// 	}
// }

func TestGetDataAkreditas(t *testing.T) {
	stats := "status"
	data := GetDataAkreditas(stats)
	fmt.Println(data)
}
func TestGetDataProgramStudi(t *testing.T) {
	stats := "programstudi"
	data := GetDataProgramStudi(stats)
	fmt.Println(data)
}

func TestGetDataProfile(t *testing.T) {
	stats := "isi_satu"
	data := GetDataProfile(stats)
	fmt.Println(data)
}
