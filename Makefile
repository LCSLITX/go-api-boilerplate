SHELL = /bin/zsh
# Uncomment the previous line and configurate it if you
# want make command to run using another terminal, as it
# implies on `source .env` command.

server-run:
	clear && source .env && go run main.go
