run_agnft_db:
	@echo "Starting agnft db service..."
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml stop agnft-auth_pg_db
	docker-compose -f ./autogear_nft_deployment/docker-compose.yml up --build -d agnft-auth_pg_db
	@echo "agnft services started..."	