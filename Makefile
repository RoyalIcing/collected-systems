dev:
	go run main.go

docker-build:
	docker build -t collected-systems .

docker-run:
	docker run -it collected-systems -p 3838:3838 --rm --name collected-systems
