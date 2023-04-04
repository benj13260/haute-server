package servers

import (
	"ben/haute/common"
	"net/http"
	"strings"

	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router(allowUpload bool) *echo.Echo {
	router := echo.New()
	router.HideBanner = true
	router.HidePort = true
	// Cors setup
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, // all /query
			http.MethodOptions, // CORS preflight
		},
		AllowHeaders: []string{
			"Authorization",
			"Content-Type",
			"Cookie",
			"Accept",
			"Set-Cookie",
			"Content-Length",
			"Accept-Encoding",
		},
		AllowCredentials: true,
		// AllowOrigins:     cfg.Server.AllowedOrigins, // Set this with an envar. split(',')
		MaxAge: int(12 * time.Hour), // Idem
	}))

	if allowUpload {

		router.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
			Skipper: func(ctx echo.Context) bool {
				var contentType = ctx.Request().Header.Get("Content-Type")
				return strings.HasPrefix(contentType, "multipart/form-data")
			},
			Limit: "100K"}))
		//router.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))
		router.Use(middleware.BodyLimit(common.MAX_FILE_SIZE_STR))
	} else {
		router.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{Limit: "100K"}))
	}

	// GRAPHQL Server Routes
	return router
}
