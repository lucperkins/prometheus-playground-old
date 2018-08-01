# Instrumented Go application

In this sandbox, a Prometheus intance scrapes metrics from a simple [Go](https://golang.org) web application.

## Usage

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up two services:

Service | Description
:-------|:-----------
`prometheus` | A Prometheus instance that's [configured](./prometheus/prometheus.yml) to scrape metrics from the Go application running on port 2112
`myapp` | A simple Go web application that pretty much only exports a simple metric (a `myapp_processed_ops_total` counter that is [incremented](./myapp/main.go#L27-L34) every 2 seconds)