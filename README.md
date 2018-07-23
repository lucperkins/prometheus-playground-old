# The Prometheus Playground

This repo houses a variety of [Docker-Compose](https://docs.docker.com/compose)-based "sandbox" projects showcasing the [Prometheus](https://prometheus.io) monitoring system. All projects are "turnkey" and require just a single `docker-compose up` command to run.

Each sandbox project has a `README` with an explanation of the project, a `docker-compose.yml` configuration file for Docker Compose, and other necessary resources (config files, `Dockerfile`s, etc.). To run a project, navigate to the appropriate directory and run `make run` (which is just an alias for `docker-compose up --build`). This will run the project in the *foreground*. To run the project in detached mode, use `make run-detached`.

## Prerequisites

In order to run the sandbox projects you'll need to install [Docker](https://docker.com) and [Docker Compose](https://docs.docker.com/compose) and have a Docker daemon running locally.

## Projects

Directory | Scenario
:---------|:--------
[`alertmanager`](./alertmanager) | Prometheus monitors a basic web service and notifies [Alertmanager](https://prometheus.io/docs/alerting/alertmanager/) if the service is down; Alertmanager, in turns, notifies a web service via webhook
[`blackbox`](./blackbox) | A [BlackBox prober exporter](https://github.com/prometheus/blackbox_exporter) probes a simple web service and provides probe-based metrics to Prometheus
[`cadvisor`](./cadvisor) | Prometheus scrapes [cAdvisor](https://github.com/google/cadvisor)-gathered metrics for several containers
[`haproxy`](./haproxy) | Prometheus runs behind [HAProxy](https://haproxy.org/), which acts as a reverse proxy and provides basic auth and TLS encryption
[`nginx`](./nginx) | Prometheus runs behind [nginx](https://nginx.org), which acts as a reverse proxy and provides basic auth and TLS encryption
[`node-exporter`](./node-exporter) | Prometheus scrapes Linux host metrics from a [Node Exporter](https://prometheus.io/docs/guides/node-exporter/)