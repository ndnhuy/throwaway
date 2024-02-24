# docker build
```
docker build --platform linux/amd64 -t demo-orders .
```
# K6
docker run --rm --network host -i grafana/k6 run - <script.js

k6 run script.js

### Run k6 export to prometheus+grafana
K6_PROMETHEUS_RW_SERVER_URL=http://localhost:9090/api/v1/write \
./k6 run -o experimental-prometheus-rw script.js