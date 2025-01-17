// Package aggregateservice provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package aggregateservice

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message    *string `json:"message,omitempty"`
	StatusCode *int    `json:"statusCode,omitempty"`
}

// PlayerAverage defines model for PlayerAverage.
type PlayerAverage struct {
	Assists       float32 `json:"assists,omitempty"`
	Blocks        float32 `json:"blocks,omitempty"`
	Fouls         float32 `json:"fouls,omitempty"`
	MinutesPlayed float32 `json:"minutes_played,omitempty"`
	PlayerId      int     `json:"player_id,omitempty"`
	Points        float32 `json:"points,omitempty"`
	Rebounds      float32 `json:"rebounds,omitempty"`
	Steals        float32 `json:"steals,omitempty"`
	Turnovers     float32 `json:"turnovers,omitempty"`
}

// TeamAverage defines model for TeamAverage.
type TeamAverage struct {
	Assists       float32 `json:"assists,omitempty"`
	Blocks        float32 `json:"blocks,omitempty"`
	Fouls         float32 `json:"fouls,omitempty"`
	MinutesPlayed float32 `json:"minutes_played,omitempty"`
	Points        float32 `json:"points,omitempty"`
	Rebounds      float32 `json:"rebounds,omitempty"`
	Steals        float32 `json:"steals,omitempty"`
	TeamId        int     `json:"team_id,omitempty"`
	Turnovers     float32 `json:"turnovers,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get average statistics of all players
	// (GET /player_average)
	GetPlayerAverage(ctx echo.Context) error
	// Get average statistics of a player by ID
	// (GET /player_average/{id})
	GetPlayerAverageId(ctx echo.Context, id int) error
	// Get average statistics of all teams
	// (GET /team_average)
	GetTeamAverage(ctx echo.Context) error
	// Get average statistics of a team by ID
	// (GET /team_average/{id})
	GetTeamAverageId(ctx echo.Context, id int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPlayerAverage converts echo context to params.
func (w *ServerInterfaceWrapper) GetPlayerAverage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPlayerAverage(ctx)
	return err
}

// GetPlayerAverageId converts echo context to params.
func (w *ServerInterfaceWrapper) GetPlayerAverageId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetPlayerAverageId(ctx, id)
	return err
}

// GetTeamAverage converts echo context to params.
func (w *ServerInterfaceWrapper) GetTeamAverage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTeamAverage(ctx)
	return err
}

// GetTeamAverageId converts echo context to params.
func (w *ServerInterfaceWrapper) GetTeamAverageId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTeamAverageId(ctx, id)
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
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/player_average", wrapper.GetPlayerAverage)
	router.GET(baseURL+"/player_average/:id", wrapper.GetPlayerAverageId)
	router.GET(baseURL+"/team_average", wrapper.GetTeamAverage)
	router.GET(baseURL+"/team_average/:id", wrapper.GetTeamAverageId)

}

type GetPlayerAverageRequestObject struct {
}

type GetPlayerAverageResponseObject interface {
	VisitGetPlayerAverageResponse(w http.ResponseWriter) error
}

type GetPlayerAverage200JSONResponse []PlayerAverage

func (response GetPlayerAverage200JSONResponse) VisitGetPlayerAverageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPlayerAverageIdRequestObject struct {
	Id int `json:"id"`
}

type GetPlayerAverageIdResponseObject interface {
	VisitGetPlayerAverageIdResponse(w http.ResponseWriter) error
}

type GetPlayerAverageId200JSONResponse PlayerAverage

func (response GetPlayerAverageId200JSONResponse) VisitGetPlayerAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetPlayerAverageId400JSONResponse ErrorResponse

func (response GetPlayerAverageId400JSONResponse) VisitGetPlayerAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetPlayerAverageId404JSONResponse ErrorResponse

func (response GetPlayerAverageId404JSONResponse) VisitGetPlayerAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

type GetTeamAverageRequestObject struct {
}

type GetTeamAverageResponseObject interface {
	VisitGetTeamAverageResponse(w http.ResponseWriter) error
}

type GetTeamAverage200JSONResponse []TeamAverage

func (response GetTeamAverage200JSONResponse) VisitGetTeamAverageResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTeamAverageIdRequestObject struct {
	Id int `json:"id"`
}

type GetTeamAverageIdResponseObject interface {
	VisitGetTeamAverageIdResponse(w http.ResponseWriter) error
}

type GetTeamAverageId200JSONResponse TeamAverage

func (response GetTeamAverageId200JSONResponse) VisitGetTeamAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTeamAverageId400JSONResponse ErrorResponse

func (response GetTeamAverageId400JSONResponse) VisitGetTeamAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type GetTeamAverageId404JSONResponse ErrorResponse

func (response GetTeamAverageId404JSONResponse) VisitGetTeamAverageIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get average statistics of all players
	// (GET /player_average)
	GetPlayerAverage(ctx context.Context, request GetPlayerAverageRequestObject) (GetPlayerAverageResponseObject, error)
	// Get average statistics of a player by ID
	// (GET /player_average/{id})
	GetPlayerAverageId(ctx context.Context, request GetPlayerAverageIdRequestObject) (GetPlayerAverageIdResponseObject, error)
	// Get average statistics of all teams
	// (GET /team_average)
	GetTeamAverage(ctx context.Context, request GetTeamAverageRequestObject) (GetTeamAverageResponseObject, error)
	// Get average statistics of a team by ID
	// (GET /team_average/{id})
	GetTeamAverageId(ctx context.Context, request GetTeamAverageIdRequestObject) (GetTeamAverageIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetPlayerAverage operation middleware
func (sh *strictHandler) GetPlayerAverage(ctx echo.Context) error {
	var request GetPlayerAverageRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetPlayerAverage(ctx.Request().Context(), request.(GetPlayerAverageRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPlayerAverage")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetPlayerAverageResponseObject); ok {
		return validResponse.VisitGetPlayerAverageResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetPlayerAverageId operation middleware
func (sh *strictHandler) GetPlayerAverageId(ctx echo.Context, id int) error {
	var request GetPlayerAverageIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetPlayerAverageId(ctx.Request().Context(), request.(GetPlayerAverageIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPlayerAverageId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetPlayerAverageIdResponseObject); ok {
		return validResponse.VisitGetPlayerAverageIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTeamAverage operation middleware
func (sh *strictHandler) GetTeamAverage(ctx echo.Context) error {
	var request GetTeamAverageRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTeamAverage(ctx.Request().Context(), request.(GetTeamAverageRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTeamAverage")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTeamAverageResponseObject); ok {
		return validResponse.VisitGetTeamAverageResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTeamAverageId operation middleware
func (sh *strictHandler) GetTeamAverageId(ctx echo.Context, id int) error {
	var request GetTeamAverageIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTeamAverageId(ctx.Request().Context(), request.(GetTeamAverageIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTeamAverageId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTeamAverageIdResponseObject); ok {
		return validResponse.VisitGetTeamAverageIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RWTW/jNhD9K8S0R0GWHaXr6uZ+oPClWLR7WwQLWhrZ3EokS46CGoH/ezGU7Fi2sMs4",
	"RYo9xSE5M2/mPT7qCUrTWqNRk4fiCXy5w1aGn786Z9wf6K3RHnnBOmPRkcKw3aL3chs28B/Z2gahgJUW",
	"yGHClGXnHFaQAO0tb3lySm/hkIAnSZ3/2VTj4DzLT4eVJtyig8PhtGQ2n7Ekjn/fyD261SO6of4YmPRe",
	"+b6bU+536SKB2rhWEhRQN0bSMzLdtRuulcCmMeVf48gsfRcTWZuuGQcu0ruYwFbpjtB/stxUNcpwl6f3",
	"MSlCqPukxtHz62EmYI3SF6NZ3Kc/xFRxuDGdrsbByzSPifWE8mI887jeqHPaPKIbB9+l868HT0nnA8r2",
	"RcJZ3MfROKGcPC7yWjnzZZxavyCdRZ5FKedaDvNsEcfppB7yPO66TAhiGcNpAoSyjZL6tHLmefrjLdLh",
	"JaVrw8kq9KVTlpTRUMCHnfJCeSG1WL1fCzJiiyR+/2kl+nsp2O6UJ1X6lCspCk7JB3ojE3+eDnAGSIBR",
	"98nnaZZm3I2xqKVVUMBdWErAStqF3mbD/ZfPut4iXQP9DUkMZ84wCVML2TQDWC9q4wTtUHiU3mgIlZ3k",
	"FOuqTzK2XxZC/0QEMIss4z+l0YQ6gJDWNqoMGWafOeXpleFfirANgd87rKGA72bP79FseIxm44rP5Ejn",
	"5L7nZtzqSjTKE7c2cHDddyDZd20r3T52OEyf3HooPg6eCw+c5YKA2ZOqDrewcAQ7pkBs9seN9S9f5WNd",
	"BWk42SIF7X98AsV1WS6QgJZtuCdVIO7vTvEbXZDrMDlj5Us36/DwSspfwPQEs5OjC9OyWKpaYTVMi3WS",
	"/4fQxp9DE9DW+lE2qjojKyDI3w7B4CfakKjZm+M1fgS92fcim5R5sN7XuQyniPGY83f6LRzmvN6L/IUb",
	"eq27hKGcDZ3/nxj57b4SQF67Slie9pSzeXzjjjJi9jY/CXz8X25y/N55ay/hud3kJEFVlz5ylDRnQfd4",
	"1FHnGihgR2SL2awxpWx2xlOxzJZzODwc/g0AAP//1YTU6p0OAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
