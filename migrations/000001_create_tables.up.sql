CREATE TABLE teams (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100) NOT NULL
);

CREATE TABLE players (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         team_id INT NOT NULL,
                         FOREIGN KEY (team_id) REFERENCES teams (id)
);

CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       player_id INT NOT NULL,
                       points INT NOT NULL,
                       rebounds INT NOT NULL,
                       assists INT NOT NULL,
                       steals INT NOT NULL,
                       blocks INT NOT NULL,
                       fouls INT NOT NULL,
                       turnovers INT NOT NULL,
                       minutes_played FLOAT NOT NULL,
                       FOREIGN KEY (player_id) REFERENCES players (id)
);

CREATE INDEX idx_player_id ON games(player_id);