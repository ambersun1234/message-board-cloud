all: hook

hook: ./scripts
	@bash ./scripts/install-git-hook.sh