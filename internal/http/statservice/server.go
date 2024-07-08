//go:generate oapi-codegen --config=api.codegen.yaml ../../../api/stats_service_openapi.yaml

package statservice

import (
	"context"
	"encoding/json"
	"nba-task-main/internal/entity"
	"nba-task-main/internal/nats"
	"net/http"
	"time"

	na "github.com/nats-io/nats.go"
)

var _ StrictServerInterface = (*Server)(nil)

type Server struct {
	NatsClient *na.Conn
}

func NewServer(natsClien *na.Conn) *Server {
	return &Server{NatsClient: natsClien}
}

func (s Server) PostAddStat(ctx context.Context, request PostAddStatRequestObject) (PostAddStatResponseObject, error) {

	if request.Body == nil {
		return PostAddStat400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	if !validateStat(*request.Body) {

		return PostAddStat400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	r := entity.Stat{
		PlayerID:      request.Body.PlayerId,
		Points:        request.Body.Points,
		Rebounds:      request.Body.Rebounds,
		Assists:       request.Body.Assists,
		Steals:        request.Body.Steals,
		Blocks:        request.Body.Blocks,
		Fouls:         request.Body.Fouls,
		Turnovers:     request.Body.Turnovers,
		MinutesPlayed: request.Body.MinutesPlayed,
	}
	req, err := json.Marshal(r)
	if err != nil {

		return PostAddStat400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	re, err := s.NatsClient.Request(entity.AddStatNATSHandler, req, 10*time.Second)
	if err != nil {

		return PostAddStat400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	resp := nats.Responce{}
	err = json.Unmarshal(re.Data, &resp)
	if err != nil {

		return PostAddStat500JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}

	if resp.Code != http.StatusOK {

		return PostAddStat500JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}

	return PostAddStat201Response{}, nil
}

func (s Server) PostAddTeam(ctx context.Context, request PostAddTeamRequestObject) (PostAddTeamResponseObject, error) {
	r := entity.Team{
		Name: request.Body.Name,
	}
	req, err := json.Marshal(r)
	if err != nil {

		return PostAddTeam400JSONResponse(NewErrorResponse(http.StatusBadRequest, "team invitation failed")), nil
	}
	re, err := s.NatsClient.Request(entity.AddTeamNATSHandler, req, 10*time.Second)
	if err != nil {

		return PostAddTeam400JSONResponse(NewErrorResponse(http.StatusBadRequest, "team invitation failed")), nil
	}
	resp := nats.Responce{}
	err = json.Unmarshal(re.Data, &resp)
	if err != nil {

		return PostAddTeam500JSONResponse(NewErrorResponse(http.StatusBadRequest, "team invitation failed")), nil
	}

	if resp.Code != http.StatusOK {

		return PostAddTeam500JSONResponse(NewErrorResponse(http.StatusBadRequest, "team invitation failed")), nil
	}
	return PostAddTeam201Response{}, nil
}

func (s Server) PostAddPlayer(ctx context.Context, request PostAddPlayerRequestObject) (PostAddPlayerResponseObject, error) {
	r := entity.Player{
		Name:   request.Body.Name,
		TeamId: request.Body.TeamId,
	}
	req, err := json.Marshal(r)
	if err != nil {

		return PostAddPlayer400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	re, err := s.NatsClient.Request(entity.AddPlayerNATSHandler, req, 10*time.Second)
	if err != nil {

		return PostAddPlayer400JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}
	resp := nats.Responce{}
	err = json.Unmarshal(re.Data, &resp)
	if err != nil {

		return PostAddPlayer500JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}

	if resp.Code != http.StatusOK {

		return PostAddPlayer500JSONResponse(NewErrorResponse(http.StatusBadRequest, "invitation failed")), nil
	}

	return PostAddPlayer201Response{}, nil
}

func NewErrorResponse(code int, msg string) ErrorResponse {
	return ErrorResponse{
		Message:    &msg,
		StatusCode: &code,
	}
}

func validateStat(stat Stat) bool {
	if stat.Points < 0 || stat.Rebounds < 0 || stat.Assists < 0 || stat.Steals < 0 || stat.Blocks < 0 || stat.Turnovers < 0 {
		return false
	}
	if stat.Fouls < 0 || stat.Fouls > 6 {
		return false
	}
	if stat.MinutesPlayed < 0 || stat.MinutesPlayed > 48.0 {
		return false
	}
	return true
}
