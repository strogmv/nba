package postgres

import (
	"nba-task-main/internal/entity"
)

func (r *Repository) AddTeam(team entity.Team) error {
	query := `INSERT INTO teams (name) VALUES ($1) RETURNING id`
	var teamID int
	err := r.db.QueryRow(query, team.Name).Scan(&teamID)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) AddPlayer(player entity.Player) error {
	_, err := r.db.Exec("INSERT INTO players (name, team_id) VALUES ($1, $2)", player.Name, player.TeamId)
	if err != nil {

		return err
	}

	return nil
}

func (r *Repository) AddStat(stat entity.Stat) error {
	_, err := r.db.Exec("INSERT INTO games (player_id, points, rebounds, assists, steals, blocks, fouls, turnovers, minutes_played) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		stat.PlayerID, stat.Points, stat.Rebounds, stat.Assists, stat.Steals, stat.Blocks, stat.Fouls, stat.Turnovers, stat.MinutesPlayed)
	if err != nil {

		return err
	}

	return nil
}

func (r *Repository) PlayerAverage(playerID int) (entity.PlayerAverage, error) {
	var avg entity.PlayerAverage
	err := r.db.QueryRow(`
        SELECT player_id, 
               AVG(points), 
               AVG(rebounds), 
               AVG(assists), 
               AVG(steals), 
               AVG(blocks), 
               AVG(fouls), 
               AVG(turnovers),
               AVG(minutes_played)
        FROM games
        WHERE player_id = $1
        GROUP BY player_id
    `, playerID).Scan(&avg.PlayerID, &avg.Points, &avg.Rebounds, &avg.Assists, &avg.Steals, &avg.Blocks, &avg.Fouls, &avg.Turnovers, &avg.MinutesPlayed)
	if err != nil {
		return avg, nil
	}
	return avg, nil
}

func (r *Repository) TeamAverage(teamID int) (entity.TeamAverage, error) {
	var avg entity.TeamAverage
	err := r.db.QueryRow(`
        SELECT p.team_id, 
               AVG(g.points), 
               AVG(g.rebounds), 
               AVG(g.assists), 
               AVG(g.steals), 
               AVG(g.blocks), 
               AVG(g.fouls), 
               AVG(g.turnovers),
               AVG(g.minutes_played)
        FROM games g
        JOIN players p ON g.player_id = p.id
        WHERE p.team_id = $1
        GROUP BY p.team_id
    `, teamID).Scan(&avg.TeamID, &avg.Points, &avg.Rebounds, &avg.Assists, &avg.Steals, &avg.Blocks, &avg.Fouls, &avg.Turnovers, &avg.MinutesPlayed)
	if err != nil {

		return avg, nil
	}

	return avg, nil
}
