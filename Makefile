.PHONY: cadvisor nginx-proxy node-exporter

cadvisor:
	(cd cadvisor && make run)

nginx-proxy:
	(cd nginx-proxy && make run)

node-exporter:
	(cd node-exporter && make run)
