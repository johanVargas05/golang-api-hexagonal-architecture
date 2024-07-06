#!/bin/bash

# Genera el perfil de cobertura incluyendo todos los paquetes
go test ./... -coverprofile=coverage.tmp

grep -v -E "(mocks|doc|src/domain/ports|src/domain/dtos|src/domain/entities|src/domain/errors|src/domain/validate_objects|src/infrastructure/primary/api/constants|src/infrastructure/primary/api/routers|src/infrastructure/secondary/pkg|src/main.go)" coverage.tmp > coverage.out

rm coverage.tmp

go tool cover -html=coverage.out
