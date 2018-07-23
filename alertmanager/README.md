# Alertmanager sandbox

In this sandbox, a Prometheus [Alertmanager](https://github.com/prometheus/alertmanager) handles alerts thrown by a running Prometheus instance. The Alertmanager in this example is configured with one alert (specified in an [`alert.rules`](./prometheus/alert.rules) file) that Prometheus fires whenever the [`hello`](./hello/main.go) service---a simple web service---is down for more than 10 seconds. If Prometheus alerts the Alertmanager that the `hello` service is down, it will `POST` a webhook to the [`webhook`](./webhook/main.go) service, which will then log the alert sent by the Alertmanager to stdout.

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up five services:

Service | Description
:-------|:-----------
`prometheus` | A Prometheus instance that's [configured](./prometheus/alert.rules) to alert the Alertmanager whenever the `hello` web service has been down for more than 10 seconds
`alertmanager` | An Alertmanager instance that's configured to `POST` an alert message to the `webhook` service whenever Prometheus alerts Alertmanager that the `hello` service is down
`amtool` | [amtool](https://github.com/prometheus/alertmanager#amtool) is a CLI utility for interacting with the Alertmanager. This service is a utility container that enables you to run amtool against the Alertmanager instance included in the sandbox.
`hello` | A simple [web service](./hello/main.go) written in Go. The `hello` web service has just two endpoints: a `/hello` endpoint that returns a `{"hello":"world"}` JSON object and a `/metrics` endpoint for Prometheus instrumentation
`webhook` | A simple [web service](./webhook/main.go) written in Go. The `webhook` web service has just one `/alert` endpoint to which Alertmanager alerts are `POST`ed

## Creating an alert

When you first start up the containers, there's no alert-worthy behavior because the `hello` service is running as expected.

## amtool

The `amtool` service can be used as a proxy for the [amtool](https://github.com/prometheus/alertmanager#amtool) CLI utility. For ease of use, you can alias `amtool` to the container:

```bash
alias amtool='docker-compose run amtool amtool'
```

Now you can use amtool commands like this:

```bash
amtool alert
```



http://localhost:9090/graph?g0.range_input=1h&g0.expr=up&g0.tab=1

```bash
docker-compose stop hello
```

```bash
docker-compose start hello
```

```bash
docker-compose logs webhook
```

![Alertmanager dashboard](./alerts.png)