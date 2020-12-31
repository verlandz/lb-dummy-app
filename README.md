# lb-dummy-app
Dummy app for testing load balancer with `Golang`

### Docker
You may pull from [docker hub](https://hub.docker.com/repository/docker/verlandz/lb-dummy-app)
and simply run this cmd
```
docker run -d --name test_lb -p 8080:8080 verlandz/lb-dummy-app:1.0
```
You can also change the port, ex to 8081
```
docker run -d --name test_lb -p 8081:8080 verlandz/lb-dummy-app:1.0
```