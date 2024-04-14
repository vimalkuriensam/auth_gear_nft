run_agnft_db:
	@echo "Starting agnft db service..."
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml stop agnft-auth_pg_db
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml up --build -d agnft-auth_pg_db
	@echo "agnft auth db services started..."

run_agnft_queue:
	@echo "Starting agnft queue service..."
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml stop agnft-queue
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml up --build -d agnft-queue
	@echo "agnft queue services started..."	

run_agnft_services:
	@echo "Starting agnft services..."
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml stop
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml up --build -d
	@echo "agnft services started..."		