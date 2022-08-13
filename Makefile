
build-image:
	docker image prune --force --filter='label=crescent-lang'
	docker build -t crescent-lang:latest --label crescent-lang .