# Challenge Service

Go NBA Code Challenge

## How to run it?

To run the server, you can use the command below. It runs all dependencies (postgres, NATS)
and applies postgres migrations

```bash
docker compose up
```

## How to test it?

The codebase contains integration tests stored in `test` folder. It uses Ginkgo as a testing framework helping write
tests in the BDD style.

```bash
make run-it-tests
```
also  use this curl:
```
curl --location 'http://localhost:8081/add_team' \
--header 'Content-Type: application/json' \
--data '{"name": "Lakers3"}'

curl --location 'http://localhost:8081/add_player' \
--header 'Content-Type: application/json' \
--data '{"name": "LeBron James", "team_id": 1}'

curl --location 'http://localhost:8081/add_stat' \
--header 'Content-Type: application/json' \
--data '{
    "player_id": 6,
    "points": 30,
    "rebounds": 10,
    "assists": 5,
    "steals": 2,
    "blocks": 1,
    "fouls": 3,
    "turnovers": 2,
    "minutes_played": 35.5
}'

curl --location 'http://localhost:8080/team_average/1'
curl --location 'http://localhost:8080/player_average/1'

```
## How to contribute? 

### Project structure
The application's structure was inspired by [Clean Architecture](https://github.com/bxcodec/go-clean-arch) (without `usecase` layer) and 
uses [Standard Golang Project Layout](https://github.com/golang-standards/project-layout). 
```
-- api      - openapi specification
-- cmd      - main files
-- config   - config files
-- internal
    -- app      - app
    -- entity   - domain entities, repository interface
    -- http     - implementation of http controllers
        -- aggregateservice  - codegenerated server by oapi-codegen
        -- statservice  - codegenerated server by oapi-codegen      
    -- postgres - mongodb repository
    -- nats    - nats client
-- migrations   - mongodb migrations
-- test     - integration tests
-- third_party  - code related to third_party service
```

### Run locally
You can use the command below to run it using `go` locally

```bash
make run-env
CONFIG_PATH=config/aggregate.qds.stats.config.local.yaml go run cmd/aggregate/main.go
```

### Used tools
All tools can be installed by the command below
```
make install-go-tools
```

You can use the tools below to:
* [`oapi-codegen`](github.com/deepmap/oapi-codegen) - generate a code and server code from an OpenAPI spec

### Before commit
If you want to add new changes you should use the command below before these changes
```bash
make pre-commit-check
```