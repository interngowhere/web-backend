# web-backend

### Project structure

### Local Development

**Environment variables**

Set the environmental variables as follow
| Key | Value | Remarks |
|--------------|--------------------------------------|-----------------------------------------------------------------------|
| DB_DIALECT | postgres | |
| DB_HOST | database | `database` is the name of the golang server service in docker-compsoe |
| DB_PORT | 5432 | |
| DB_USER | | e.g. user |
| DB_NAME | | e.g. postgres |
| DB_PASSWORD | | e.g. mysecretpassword |
| SERVER_PORT | 8000 | |
| SERVER_HOST | server | `server` is the name of the golang server service in docker-compsoe |
| JWT_SIGKEY | | Private key used to sign the JWT |
| JWT_VALIDITY | | Time in minutes for which the token is valid |
| DEV_ORIGIN | http://localhost:8000 | CORS allowed origin |
| PROD_ORIGIN | https://interngowhere.netlify.app | CORS allowed origin |

**Docker**

To run docker containers locally, run

```
docker-compose -f docker-compose-dev.yml up -d
```

This will spin up docker containers for the backend go server, PostgreSQL database, and nginx reverse proxy respectively.

In dev environment, nginx is configured to listen on port 80. In production environment, nginx is configured to handle TLS/SSL encryption for HTTPS traffic and will listen on port 443.

**Seeding the database**

As the database port 5432 is also exposed on localhost in dev environment, you can simply run the following to seed the database.

```
go run cmd/seed/main.go
```

Alternatively, you may run the command inside the docker container with the backend server.

This will get you shell access in the docker container:

```
docker exec -it server sh
```

You can then run the same command directly inside the container:

```
go run cmd/seed/main.go
```

### Commit messages

This project follows [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) guideline for commit messages. See the table below for the list of commit types.

| Type     | Description                                                                                            |
| -------- | ------------------------------------------------------------------------------------------------------ |
| feat     | A new feature                                                                                          |
| fix      | Bug fixes                                                                                              |
| test     | Adding missing tests or correcting existing tests                                                      |
| refactor | Changes to source code that neither add a feature nor fixes a bug                                      |
| build    | Changes to CI or build configuration files (Docker, github actions)                                    |
| chore    | Anything else that doesn't modify any `internal` or `test` files (linters, configs, etc.)              |
| revert   | Reverts a previous commit                                                                              |
| docs     | Documentation only changes                                                                             |
| perf     | A code change that improves performance                                                                |
| style    | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc) |
