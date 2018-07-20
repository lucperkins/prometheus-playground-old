# The Prometheus Playground

This repo houses a variety of [Docker-Compose](https://docs.docker.com/compose/)-based "sandbox" projects showcasing the [Prometheus](https://prometheus.io) monitoring system.

Each sandbox project has a `README` with an explanation of the project and a `docker-compose.yml` configuration file for Docker Compose. To run a project, navigate to the appropriate directory and run `make run` (which is just an alias for `docker-compose up --build`). This will run the project in the *foreground*. To run the project in detached mode, use `make run-detached`.

## Prerequisites