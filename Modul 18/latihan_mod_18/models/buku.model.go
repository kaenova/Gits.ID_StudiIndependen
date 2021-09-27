package models

import (
	"errors"
	"fmt"
	"latihan_mod_18/db"
)

type Buku struct {
	ID   int    `json: "id"`
	Nama string `json: "nama"`
}

func GetAllBuku() (interface{}, error) {
	var (
		data   interface{}
		arrObj []Buku
		obj    Buku
	)

	db := db.CreateConnection()

	rows, err := db.Query(`SELECT * FROM "Buku".namabuku`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while fetching in database")
	}
	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Nama)
		if err != nil {
			return nil, errors.New("parsing database response")
		}
		arrObj = append(arrObj, obj)
	}

	data = arrObj
	return data, nil
}

func CreateBuku(nama string) (interface{}, error) {
	db := db.CreateConnection()
	result, err := db.Exec(fmt.Sprintf(`INSERT INTO "Buku".namabuku (nama) VALUES ('%s')`, nama))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while inserting")
	}
	if banyak, err := result.RowsAffected(); banyak < 1 || err != nil {
		return nil, errors.New("object is not inserted")
	}

	data := "book inserted"
	return data, nil
}
