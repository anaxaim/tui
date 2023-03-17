CLIENT_IMG=tui-client

ui: ## run ui locally
	npm i && npm run dev

build:
	docker build -t $(CLIENT_IMG) .

run:
	docker rm -f $(CLIENT_IMG) &>/dev/null && echo 'Removed old container' ; \
	docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8080:80 --name $(CLIENT_IMG) $(CLIENT_IMG)
