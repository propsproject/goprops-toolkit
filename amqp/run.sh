docker stop rabbit
docker rm rabbit
docker build -t rabbit https://github.com/PROPSProject/dockerfiles.git#master:rabbitmq
docker run -d -p 5672:5672 -p 15672:15672 rabbit
sleep 5;
go test -tags integration -v