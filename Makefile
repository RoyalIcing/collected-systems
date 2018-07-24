dev:
	go run main.go

docker-build:
	docker build -t collected-systems .

docker-run:
	docker run -it -p 3838:3838 --rm --name collected-systems collected-systems

docker-debug:
	docker run -it -p 3838:3838 --rm --name collected-systems collected-systems --entrypoint /bin/ls -c /go/src/app

