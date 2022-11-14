# go-template

Template for GO projects

Steps to run a project locally

1. Create a file in the `config` folder called `.env` based on `.env.example`.

2. Run command `docker-compose -f docker-compose.yml up`. This will start `postgres` and `redis`

3. Use Dockerfile.DEV. This will start project and rebuild it when .go files changed.

Project struct

1. cmd - Main folder for run service.
2. config - Contains config structs, initialization, example.
3. internal
    - dto - Data transfer objects
    - handlers - Handlers
    - models - Models
    - repository - Repository
    - services - Services
4. pkg - packages.
