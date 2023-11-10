docker build -t math:1.1 -f math/Dockerfile .
docker build -t lb:latest -f load_balance/Dockerfile .

docker run -d --name math -p 8989:8989 math
docker run -d --name math -p 8989:8989 -m 6m --cpus="0.01" math:1.0
go test -bench=BenchmarkAdd -benchmem -count 10

#K6
docker run --rm --network host -i grafana/k6 run - <script.js
k6 run script.js

#Run k6 export to prometheus+grafana
K6_PROMETHEUS_RW_SERVER_URL=http://localhost:9090/api/v1/write \
./k6 run -o experimental-prometheus-rw script.js


#NOTE
Exp1: 1 service, 6M RAM, 0.01 CPU
- send 80~100 req/s, RPS 20~30