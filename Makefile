dev:
	$(GOPATH)/bin/gin -p 3839 -a 3838 run main.go

docker-build:
	docker build -t collected-systems .

docker-run:
	docker run -it -p 3838:3838 --rm --name collected-systems collected-systems

docker-debug:
	docker run -it -p 3838:3838 --rm --name collected-systems collected-systems --entrypoint /bin/ls -c /go/src/app

now-prod:
	now --docker -e S3_ACCESS_KEY=$(PRODUCTION_S3_ACCESS_KEY) -e S3_SECRET_ACCESS_KEY=$(PRODUCTION_S3_SECRET_ACCESS_KEY) \
	&& now alias

up-env:
	up env add S3_ACCESS_KEY=$(PRODUCTION_S3_ACCESS_KEY) S3_SECRET_ACCESS_KEY=$(PRODUCTION_S3_SECRET_ACCESS_KEY)

up-staging: up-env
	up deploy

up-prod: up-env
	up deploy production

prod: now-prod up-prod
