# lb-dummy-app
Dummy app for testing load balancer with `Golang`

## How to run the app

### Golang
- go to this filepath
- `go run main.go`, it'll run in port `:8080`
- `APP_PORT=8081 go run main.go`, it'll run in port `:8081`

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

## How to check the app
`curl localhost:8080` or hit `http://localhost:8080/` in browser.
