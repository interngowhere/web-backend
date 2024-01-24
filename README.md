# web-backend

### Project structure

### Local Development

To run docker containers locally, run

```
docker-compose -f docker-compose-dev.yml up -d
```

This will spin up docker containers for the backend go server, PostgreSQL database, and nginx reverse proxy respectively.

In dev environment, nginx is configured to listen on port 80. In production environment, nginx is configured to handle TLS/SSL encryption for HTTPS traffic and will listen on port 443.

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
