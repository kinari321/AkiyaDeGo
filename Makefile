.PHONY: hello
hello: ## echo
	echo Hello

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
