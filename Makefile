all: hook

hook: ./scripts/install-git-hook.sh
	@bash $<

format:
	@gofmt -w .

install: install-go install-protoc

install-protoc: ./tools/install-protoc.sh
	@bash $<

docker-login:
	@docker login -u ambersun1234