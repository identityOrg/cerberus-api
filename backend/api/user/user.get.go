// Package user provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package user

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// Error defines model for Error.
type Error struct {

	// Detail description
	Description *string `json:"description,omitempty"`

	// Error code
	Error *string `json:"error,omitempty"`
}

// UserDetails defines model for UserDetails.
type UserDetails struct {

	// A username tontainer
	Name UserName `json:"name"`

	// Username for login
	Username string `json:"username"`
}

// UserName defines model for UserName.
type UserName struct {

	// Surname or family name
	FamilyName *string `json:"family_name,omitempty"`

	// Given name
	GivenName *string `json:"given_name,omitempty"`
}

// UserSummary defines model for UserSummary.
type UserSummary struct {

	// User is active or inactive
	Active *bool `json:"active,omitempty"`

	// User is blocked by invalid attempt
	Blocked *bool `json:"blocked,omitempty"`

	// Email address
	Email *string `json:"email,omitempty"`

	// Login username
	Username *string `json:"username,omitempty"`
}

// UserSummaryPage defines model for UserSummaryPage.
type UserSummaryPage struct {

	// Page number of the response chunk
	PageNumber *int32 `json:"page_number,omitempty"`

	// Total available pages
	PageTotal *int32 `json:"page_total,omitempty"`

	// A collection of users
	Users []UserSummary `json:"users"`
}

// UnAuthorized defines model for UnAuthorized.
type UnAuthorized Error

// ListUserParams defines parameters for ListUser.
type ListUserParams struct {

	// Number of items in each page.
	// Default value is 10.
	PageSize *int `json:"page_size,omitempty"`

	// Page number to return in response.
	// Starts with 1 and default value is 1 too.
	PageNumber *int `json:"page_number,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List All Users
	// (GET /v1/api/user)
	ListUser(ctx echo.Context, params ListUserParams) error
	// Fetch user details
	// (GET /v1/api/user/{id})
	UserDetail(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListUser converts echo context to params.
func (w *ServerInterfaceWrapper) ListUser(ctx echo.Context) error {
	var err error

	ctx.Set("OAuth.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListUserParams
	// ------------- Optional query parameter "page_size" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_size", ctx.QueryParams(), &params.PageSize)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page_size: %s", err))
	}

	// ------------- Optional query parameter "page_number" -------------

	err = runtime.BindQueryParameter("form", true, false, "page_number", ctx.QueryParams(), &params.PageNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page_number: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ListUser(ctx, params)
	return err
}

// UserDetail converts echo context to params.
func (w *ServerInterfaceWrapper) UserDetail(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set("OAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UserDetail(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v1/api/user", wrapper.ListUser)
	router.GET("/v1/api/user/:id", wrapper.UserDetail)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RWS2/bRhD+K4Ntj6weSdoDTzWSpjBQuEEcnxLDGJEjaWNyl9kdqlEM/vdidpeiHpTj",
	"AM5J4j5mv5n5vpl5UIWtG2vIsFf5g3LkG2s8hY8bc9Hy2jr9jUr5LqxhMix/sWkqXSBra6afvTWy5os1",
	"1Sj/fnW0VLn6ZToYn8ZdP/3LOetU13WZKskXTjdiROXqdaXJMGgPrQHcPQxsAYuCvAdeEzj60pJnJfeT",
	"SXkxWs0fjmxeGCDZgd4vsIvPVLDKVONsQ451dPXg2rGVN8SoK9hfzBRvG1K58uy0WakuUzQOISCDwpZ0",
	"eqnbrSRYXaZuPLn4oB/xB8qIpfXkki+gzdK6GhOuQ78M1vS9jMiLV3Kuy5TY7e8cPn2TdmBpHVR2pc2o",
	"P5If7YQvHwdjWcRxe8bdq9EHL6C/D2wNozbkVKboK9ZNFc6v9IbMXYSrVKZ86/qP7jgQS6x1tb0b9+06",
	"XgTrIJ6DBPskyftPHlv5W/bO3JRMaxbY6r21DB+2TYzkzv8zsblu6xrddiw8Pm6Bo8aRJ8OBAmCX++w4",
	"jBgWrDekcnYtZWpR2eJeUrXEylOmqEZdqTzkbf5naWvUZlLYWu3zIu6eBrg3PcYbEXXclxBrk87uXF5Y",
	"WxEaMbrDdM5OOgCLLWizwUqXgMxUNzxqL/l0okpZBixLR96PJfq8EP4R7sMet38k1X06H8/2O1yNCqLB",
	"Fe3y63em9hIsJ+5MWy/Iqfz3V1lcYMtYqfyPefTLq/zjWS7Ez54K3or86CsfcmBY7rJns3R7wqkDb47D",
	"IUGCuCkxic0hVfli3Zp7lalYGFWutOGXL4aoa8O0IhdYvBeg4yc+yDLgBnWFi4pC/P3TzKY4n+awsFVF",
	"xb5SxaJmqv1TCnVPoKF1oHO4Ha29fqTiStukonWat9diNgb6X2n0p2jDMvTnge09SdUPeILOCF0oyumR",
	"NXMTW7v0pFN7r8ktyLU+SAEu3l2qQStF2vtNkKtMbcj5eGs+mU1m4rBtyGCjVa5eTmYTiXuDvA4OTDfz",
	"KTZ6Gi6PRN2Rt60raCiW2qwAo5L2UoKmBGFgKKWSmBXxSAnQngGr2IY9/Kd5bVsGT+iKNRROMzmNKiCO",
	"pi7LdO0metegw5q41+Kh+asdqQMtQBsgLNaBfpNP5g0tsa0YNli1JCVxPpsIheTql5ZCSUjqCuT2+hv1",
	"WQvj2TFbRcPntcUWHHHrjODoFTb5ZK4ZHUfnYR4CV54AA7b2UWxJ3I+iu80Ox9IXs9mzTaPHRXdkLr1u",
	"4/T5PmEQJr56RghnB+LL1N4uTdPyJ3NFKwwtdMedGOiIZ37umV3opgfjfCgE/WwRCX1RVUGYXh3s3RxJ",
	"pJeSnNqX3fRBl92j2gtMSvIDFD6VeqPLFtM8uzsoZGrISZkd1AjWAK+Rz4nyLXGxTvOxB1wERWqzqiha",
	"X2zjry5PhDnM3Oonc62f7J/Msx/Oq1x69fPJGVhxZRne2tYc0ymmIkS77B3+Ts2L410phXmpyWVgLIOX",
	"oRz9/rAVKomU/aGQhIQOvS+OHY/XkwHq5RkOHnTKgDf1yI+33W33fwAAAP//JvABSjcPAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}