docker build --tag math .
docker run -d --name math -p 8989:8989 math