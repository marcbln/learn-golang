

```shell
docker run --rm -p9090:9090  -v $(readlink -f ./prometheus.yml):/etc/prometheus/prometheus.yml --name prometheus prom/prometheus      main
```

then in browser:
http://localhost:9090
