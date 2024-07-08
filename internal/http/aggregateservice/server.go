//go:generate oapi-codegen --config=api.codegen.yaml ../../../api/aggregate_service_openapi.yaml

package aggregateservice

import (
	"context"
	"encoding/json"
	"fmt"
	"nba-task-main/internal/entity"
	"net/http"
	"time"

	na "github.com/nats-io/nats.go"
)

var _ StrictServerInterface = (*Server)(nil)

type Server struct {
	NatsClient *na.Conn
}

func (s Server) GetPlayerAverageId(ctx context.Context, request GetPlayerAverageIdRequestObject) (GetPlayerAverageIdResponseObject, error) {
	req, err := json.Marshal(request.Id)
	if err != nil {
		return GetPlayerAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	re, err := s.NatsClient.Request(entity.GetPlayerNATSHandler, req, 10*time.Second)
	if err != nil {

		return GetPlayerAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	resp := entity.PlayerAverage{}
	err = json.Unmarshal(re.Data, &resp)
	if err != nil {

		return GetPlayerAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	fmt.Println("entity.PlayerAverage", resp)
	return GetPlayerAverageId200JSONResponse{
		Assists:       resp.Assists,
		Blocks:        resp.Blocks,
		Fouls:         resp.Fouls,
		MinutesPlayed: resp.MinutesPlayed,
		PlayerId:      request.Id,
		Points:        resp.Points,
		Rebounds:      resp.Rebounds,
		Steals:        resp.Steals,
		Turnovers:     resp.Turnovers,
	}, nil
}

func (s Server) GetTeamAverageId(ctx context.Context, request GetTeamAverageIdRequestObject) (GetTeamAverageIdResponseObject, error) {
	req, err := json.Marshal(request.Id)
	if err != nil {

		return GetTeamAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	re, err := s.NatsClient.Request(entity.GetTeamNATSHandler, req, 10*time.Second)
	if err != nil {

		return GetTeamAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	var resp = entity.TeamAverage{}
	err = json.Unmarshal(re.Data, &resp)
	if err != nil {

		return GetTeamAverageId400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), err
	}
	return GetTeamAverageId200JSONResponse{
		Assists:       resp.Assists,
		Blocks:        resp.Blocks,
		Fouls:         resp.Fouls,
		MinutesPlayed: resp.MinutesPlayed,
		Points:        resp.Points,
		Rebounds:      resp.Rebounds,
		Steals:        resp.Steals,
		Turnovers:     resp.Turnovers}, nil
}

// TODO Implement it
func (s Server) GetPlayerAverage(ctx context.Context, request GetPlayerAverageRequestObject) (GetPlayerAverageResponseObject, error) {
	return nil, nil
}

// TODO Implement it
func (s Server) GetTeamAverage(ctx context.Context, request GetTeamAverageRequestObject) (GetTeamAverageResponseObject, error) {
	return nil, nil
}

func NewServer(natsClien *na.Conn) *Server {
	return &Server{NatsClient: natsClien}
}

func NewErrorResponse(code int, msg string) ErrorResponse {
	return ErrorResponse{
		Message:    &msg,
		StatusCode: &code,
	}
}
