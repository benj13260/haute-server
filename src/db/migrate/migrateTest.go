package main

import (
	"ben/haute/db"
	"ben/haute/models"
	"encoding/json"
	"fmt"
)

func testFk() {
	db.DB.Migrator().DropTable(&models.U{})
	db.DB.Migrator().DropTable(&models.E{})

	db.DB.AutoMigrate(&models.U{}, &models.E{})
	db.DB.Create(&models.E{ID: 1, Address: "Chamonix"})
	i := &models.U{ID: 2, Name: "Me", Eid: 1}
	r, _ := json.Marshal(i)
	println(string(r))
	fmt.Println("? Migration complete")

	db.DB.Create(i)
}

func migrateTest() {
	//testFk()

	//result := db.DB.Find(&bookingPaymentStatus)

}
