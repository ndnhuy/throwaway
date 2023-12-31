# latest RabbitMQ 3.12
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 --hostname rabbitmq -v $(pwd)/.rabbitmq:/var/lib/rabbitmq rabbitmq:3.12-management