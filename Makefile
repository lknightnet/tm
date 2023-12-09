run:
	docker compose up --build

build:
	CGO_ENABLED=0 GOOS=linux go build -o task-manager

stop:
	docker stop task-manager-server-1
	docker stop task-manager-postgres-1

rm:
	docker rm task-manager-server-1
	docker rm task-manager-postgres-1