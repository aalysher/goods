POSTGRES_USER := $(shell cat config.yaml | grep 'username:' | awk '{print $$2}')
POSTGRES_PASSWORD := $(shell cat config.yaml | grep 'password:' | awk '{print $$2}')
POSTGRES_HOST := $(shell cat config.yaml | grep 'host:' | awk '{print $$2}')
POSTGRES_PORT := $(shell cat config.yaml | grep 'port:' | awk '{print $$2}')
POSTGRES_DBNAME := $(shell cat config.yaml | grep 'dbname:' | awk '{print $$2}')

migrate:
	@echo "Running migrations..."
	@echo "Database connection string: postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DBNAME)"
	@migrate -source file: migration/ -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DBNAME) up"