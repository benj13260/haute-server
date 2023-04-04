package servers

import (
	"ben/haute/common"
	"ben/haute/controllers"
	"fmt"
)

func ServePost(port int, path string, allowUpload bool) error {
	//cfg := config.Get()
	router := Router(allowUpload)

	router.GET("/health", common.EchoHealth)

	router.GET("/usertypes", controllers.GetUserTypes)

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.GET("/bookings", controllers.GetAllBookings)
	router.POST("/bookings", controllers.CreateOrUpdateBooking)
	router.GET("/bookings/:id", controllers.GetBooking)
	//router.PUT("/bookings/:id", controllers.UpdateBooking)
	router.DELETE("/bookings/:id", controllers.DeleteBooking)

	router.GET("/bookings/booking_payment_status", controllers.GetBookingPaymentStatus)
	router.GET("/bookings/booking_transport_status", controllers.GetBookingTransportStatus)
	router.GET("/bookings/locations", controllers.GetLocations)

	//router.Use(authentication.CookieMiddleWare)
	//router.Any(path, handlers.Handler())
	return router.Start(fmt.Sprintf(":%d", port))
}
