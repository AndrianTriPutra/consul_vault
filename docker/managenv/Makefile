network:
	@if ! docker network inspect managenv > /dev/null 2>&1; then \
		echo "Network managenv not found, creating..."; \
		docker network create --driver=bridge managenv; \
	else \
		echo "Network managenv already exists."; \
	fi
up: 
	docker compose -f docker-compose.yml -p managenv up --force-recreate -d --build 
down:
	# docker compose -f docker-compose.yml down  
	docker stop vault
	docker stop consul
	docker rm vault 
	docker rm consul
	
  


