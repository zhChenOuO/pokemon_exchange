package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

//CorsConfig ...
var CorsConfig = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
	},
	AllowHeaders: []string{
		"*",
		echo.HeaderAccept,
		echo.HeaderAcceptEncoding,
		echo.HeaderAuthorization,
		echo.HeaderContentDisposition,
		echo.HeaderContentEncoding,
		echo.HeaderContentLength,
		echo.HeaderContentType,
		echo.HeaderCookie,
		echo.HeaderSetCookie,
		echo.HeaderIfModifiedSince,
		echo.HeaderLastModified,
		echo.HeaderLocation,
		echo.HeaderUpgrade,
		echo.HeaderVary,
		echo.HeaderWWWAuthenticate,
		echo.HeaderXForwardedFor,
		echo.HeaderXForwardedProto,
		echo.HeaderXForwardedProtocol,
		echo.HeaderXForwardedSsl,
		echo.HeaderXUrlScheme,
		echo.HeaderXHTTPMethodOverride,
		echo.HeaderXRealIP,
		echo.HeaderXRequestID,
		echo.HeaderXRequestedWith,
		echo.HeaderServer,
		echo.HeaderOrigin,
	},
	ExposeHeaders: []string{
		"*",
		echo.HeaderAuthorization,
		echo.HeaderContentDisposition,
		echo.HeaderContentEncoding,
		echo.HeaderContentLength,
		echo.HeaderContentType,
		echo.HeaderCookie,
		echo.HeaderSetCookie,
		echo.HeaderIfModifiedSince,
		echo.HeaderLastModified,
		echo.HeaderLocation,
		echo.HeaderUpgrade,
		echo.HeaderVary,
		echo.HeaderWWWAuthenticate,
		echo.HeaderXUrlScheme,
		echo.HeaderXHTTPMethodOverride,
		echo.HeaderXRealIP,
		echo.HeaderXRequestID,
		echo.HeaderXRequestedWith,
		echo.HeaderServer,
		echo.HeaderOrigin,
	},
})
