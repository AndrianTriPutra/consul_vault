dev-up: 
	ENV=development HOST=consul-server:8500 FILE=api docker compose -f docker-compose.yml -p app_managenv up --force-recreate -d --build 
dev-down: 
	ENV=development HOST=consul-server:8500 FILE=api docker compose -f docker-compose.yml -p app_managenv up --force-recreate -d --build 
down:
	ENV=prod docker-compose -f docker/prod-docker-compose.yml down 
	

