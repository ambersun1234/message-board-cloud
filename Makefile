all: hook

hook: ./scripts/install-git-hook.sh
	@bash $<

install: install-go install-protoc install-go-module

install-go: ./tools/install-go.sh
	@bash $<

install-protoc: ./tools/install-protoc.sh
	@bash $<

install-go-module: ./tools/install-go-module.sh
	@bash $<

docker-login:
	@docker login -u ambersun1234