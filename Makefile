.PHONY: hello
hello: ## echo
	echo Hello

###############
# nginxまわり
###############
.PHONY: nginx-copy-conf
nginx-copy-conf: ## nginx.confをコピーして上書き
	sudo cp ./nginx/nginx.conf /etc/nginx/nginx.conf

.PHONY: nginx-restart
nginx-restart: ## nginxの再起動
	sudo systemctl restart nginx

.PHONY: nginx
nginx: ## nginxのセットアップ
	make nginx-copy-conf
	make nginx-restart

###############
# デプロイまわり
###############
.PHONY: kill-app-process
kill-app-process: ## ローカルのアプリプロセスを殺す
	kill $(shell lsof -i :8080 -t)

.PHONY: build-app
build-app: ## アプリのビルド
	go build ./main.go

.PHONY: run-app-with-background
run-app-with-background: ## アプリを起動
	./main &

.PHONY: deploy
deploy: ## アプリをデプロイ
	make kill-app-process
	make build-app
	make run-app-with-background
	curl localhost
