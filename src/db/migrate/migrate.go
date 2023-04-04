package main

import (
	"ben/haute/common/logz"
	"ben/haute/db"
	"ben/haute/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var log = logz.LOGZ

var loc map[string]models.Location
var locG map[string][]uint

func locations() {

	byt, err := os.ReadFile("/home/benj/work/haute/haute-server/migration/locations.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(byt, &loc); err != nil {
		panic(err)
	}

	byt, err = os.ReadFile("/home/benj/work/haute/haute-server/migration/locationsGroup.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(byt, &locG); err != nil {
		panic(err)
	}

	//Inverse input locationGroupMap
	locT := make(map[uint]uint)
	for k, i := range locG {
		for j := range i {
			v, _ := strconv.Atoi(k)
			//println(strconv.Itoa(int(i[j])) + " " + k)
			locT[uint(i[j])] = uint(v)
		}
	}

	for _, i := range loc {

		n := &models.Location{
			ID:              i.ID,
			Name:            i.Name,
			Country:         i.Country,
			Code:            i.Code,
			IsTransitPoint:  i.IsTransitPoint,
			TransitMethodId: i.TransitMethodId,
			Parent:          i.Parent,
			Order:           i.Order,
			LocationGroupID: locT[i.ID],
		}

		//r, _ := json.Marshal(n)
		//println(string(r))

		db.DB.Create(n)

	}

}

func migrate() {
	db.DB.Migrator().DropTable(&models.User{}, &models.UserType{},
		&models.BookingPaymentStatus{}, &models.BookingTransportStatus{}, &models.Location{}, &models.LocationGroup{})

	db.DB.AutoMigrate(&models.User{}, &models.UserType{},
		&models.Booking{}, &models.BookingPaymentStatus{}, &models.BookingTransportStatus{},
		&models.LocationGroup{}, &models.Location{})

	// Create user type
	db.DB.Create(&models.UserType{ID: "0", Title: "Customer"})
	db.DB.Create(&models.UserType{ID: "1", Title: "Partner"})
	db.DB.Create(&models.UserType{ID: "2", Title: "Driver"})

	// Create booking type - transportation point of view
	db.DB.Create(&models.BookingTransportStatus{ID: "0", Title: "Not planned"})
	db.DB.Create(&models.BookingTransportStatus{ID: "1", Title: "Planned"})
	db.DB.Create(&models.BookingTransportStatus{ID: "2", Title: "Complete"})

	// Create booking payment status - transportation point of view
	db.DB.Create(&models.BookingPaymentStatus{ID: "0", Title: "Not paid"})
	db.DB.Create(&models.BookingPaymentStatus{ID: "1", Title: "Paid"})
	db.DB.Create(&models.BookingPaymentStatus{ID: "2", Title: "Issue"})

	// Create LocationGroup
	db.DB.Create(&models.LocationGroup{ID: 458, Name: "Chamonix"})
	db.DB.Create(&models.LocationGroup{ID: 29, Name: "Airports"})
	db.DB.Create(&models.LocationGroup{ID: 83, Name: "Cities"})
	db.DB.Create(&models.LocationGroup{ID: 30, Name: "Other Resorts"})

	db.DB.Create(&models.User{Name: "aaaa"})
	db.DB.Create(&models.User{Name: "aaab"})
	db.DB.Create(&models.User{Name: "aaac"})
	db.DB.Create(&models.User{Name: "accc"})
	db.DB.Create(&models.User{Name: "bccc"})
	db.DB.Create(&models.User{Name: "cccc"})

	fmt.Println("? Migration complete")
}

func main() {
	//migrateTest()

	migrate()
	locations()
}
