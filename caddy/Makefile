build:
	docker build -t caddy .

tag-latest:
	sudo docker tag caddy eu.gcr.io/bluff-229901/caddy:latest

publish:
	build
	tag-latest
	sudo docker push caddy eu.gcr.io/bluff-229901/caddy:latest
