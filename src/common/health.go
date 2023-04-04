package common

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

var HealthResponseBytes = []byte(`{"ok": true}`)

type HealthResponse struct {
	Ok bool `json:"ok"`
}

func EchoHealth(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		HealthResponse{Ok: true},
	)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	w.Write(HealthResponseBytes)
}

// minimalist server for health

func ServeHealth(port int) {
	r := http.NewServeMux()
	r.HandleFunc("/health", Health)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func TestServer(port int) {
	e := echo.New()
	e.GET("/", EchoHealth)
	e.GET("/health", EchoHealth)
	e.GET("/hello", EchoHealth)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
