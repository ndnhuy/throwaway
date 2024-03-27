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


#####
thread pool size is 200 -> 1581.066829/s
thread pool size is 10 -> 1896.851584/s
thread pool size is 5 -> 1920.889062/s
thread pool size is 1 -> 1954.460304/s
=> the task is heavily CPU-bound so the less number of threads, the better it performs