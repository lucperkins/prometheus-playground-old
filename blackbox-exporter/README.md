# Blackbox prober exporter sandbox

In this sandbox, the [Blackbox prober exporter](https://github.com/prometheus/blackbox_exporter) probes a [simple web service](./web/main.go), while Prometheus scrapes `probe_*` metrics from the Blackbox exporter.

## Usage

To start the sandbox:

```bash
# In the foreground
make run # docker-compose up --build

# In detached mode
make run-detached # docker-compose up --build --detach
```

This will start up a `blackbox`, `prometheus`, and `web` container.

The `web` container runs a web service with two endpoints:

* `/hello` returns a `{"hello": "world"}` JSON object
* `/health` returns an HTTP `200 OK` (indicating that the service is running)
]
http://localhost:9115/
http://localhost:9115/probe?target=web:2112/health&module=http_2xx
http://localhost:9090/graph?g0.range_input=1h&g0.expr=probe_success%7Binstance%3D%22web%3A2112%2Fhealth%22%7D&g0.tab=1

```bash
docker-compose stop web
```

```bash
docker-compose start web
```

## Assets

Folder | Assets
:------|:------
[`prometheus`](./prometheus) | A [`prometheus.yml`](./prometheus/prometheus.yml) configuration file for Prometheus
[`blackbox`](./blackbox) | A [`blackbox.yml`](./blackbox/blackbox.yml) configuration file for the Blackbox exporter
[`web`](./web) | Source files for the simple web server (written in Go)