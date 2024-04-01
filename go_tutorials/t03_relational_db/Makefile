.PHONY: start-db
start-db:
	docker run --name t03-mysql -e MYSQL_ROOT_PASSWORD=admin -e MYSQL_DATABASE=go_tutorial -p 3306:3306 --expose 3306 -d mysql:latest
	docker cp ./mysql.sql t03-mysql:/mysql.sql

stop-db:
	docker container stop t03-mysql
	docker rm t03-mysql

ssh-db:
	docker exec -it t03-mysql /bin/bash

check:
	docker inspect t03-mysql | grep "IPAddress"
