package controllers

import (
	"ben/haute/common/logz"
	"ben/haute/db"
	"ben/haute/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateOrUpdateBooking(c echo.Context) error {
	u := &models.Booking{}
	if err := c.Bind(u); err != nil {
		return err
	}

	if u.ID.ID() == 0 {
		u.ID = uuid.New()
		//u.CreatedAt = time.Now()
		db.DB.Create(&u)
	} else {
		//u.UpdatedAt = time.Now()
		result := db.DB.Model(&u).Where("id = ?", u.ID).Updates(&u)
		if result.Error != nil {
			return c.JSON(http.StatusNotFound, result.Error)
		}
	}
	return c.JSON(http.StatusOK, u)
}

func CreateBooking(c echo.Context) error {
	u := &models.Booking{}
	if err := c.Bind(u); err != nil {
		return err
	}
	result := db.DB.Create(&u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func GetBooking(c echo.Context) error {
	id := c.Param("id")
	u := &models.Booking{}
	result := db.DB.First(u, "id = ? and deleted_at is null", id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)

}

func UpdateBooking(c echo.Context) error {
	u := &models.Booking{}
	if err := c.Bind(u); err != nil {
		return err
	}
	result := db.DB.Updates(&u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func DeleteBooking(c echo.Context) error {
	ids := c.Param("id")
	id, err := uuid.Parse(ids)
	logz.LOGZ.Infof("DeleteBooking %v", ids)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Incorrect ID")
	}
	u := &models.Booking{ID: id}
	result := db.DB.Model(&u).Update(models.DeletedAt, time.Now())

	//result := db.DB.Delete(u)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, u)
}

func GetAllBookings(c echo.Context) error {
	var bookings []models.Booking
	result := db.DB.Find(&bookings, "deleted_at is null")
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, bookings)
}

func GetBookingPaymentStatus(c echo.Context) error {
	var bookingPaymentStatus []models.BookingPaymentStatus
	result := db.DB.Find(&bookingPaymentStatus)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, bookingPaymentStatus)
}

func GetBookingTransportStatus(c echo.Context) error {
	var bookingTransportStatus []models.BookingTransportStatus
	result := db.DB.Find(&bookingTransportStatus)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}
	return c.JSON(http.StatusOK, bookingTransportStatus)
}

type locationResp struct {
	Locations []models.Location
	Groups    []models.LocationGroup
}

func GetLocations(c echo.Context) error {
	var locations []models.Location
	result := db.DB.Find(&locations)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	var locationGroups []models.LocationGroup
	result = db.DB.Find(&locationGroups)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	res := locationResp{Locations: locations, Groups: locationGroups}
	//r, _ := json.Marshal(res)
	//println(string(r))

	return c.JSON(http.StatusOK, res)

}

/*
func GetLocations(c echo.Context) error {
	var locations []models.Location
	result := db.DB.Find(&locations)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	return c.JSON(http.StatusOK, locations)

}

func GetGroups(c echo.Context) error {
	var locationGroups []models.LocationGroup
	result := db.DB.Find(&locationGroups)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, result.Error)
	}

	return c.JSON(http.StatusOK, locationGroups)

}
*/
