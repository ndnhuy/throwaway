docker build --tag math .
docker run -d --name math -p 8989:8989 math
docker run -d --name math -p 8989:8989 -m 6m --cpus="0.01" math:1.0