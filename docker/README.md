# Docker stack

This folder contains local Docker support files for TZone:

- `docker-compose.yml` starts the API, frontend, PostgreSQL, and MongoDB services.
- `docker/postgres/init/001-init.sql` enables the `pgcrypto` extension required by GORM UUID defaults.
- `docker/mongo-seed/` seeds the MongoDB `Cluster0.brands` collection from `phone.json`.

