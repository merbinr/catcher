set -xe


docker compose down
docker compose build
docker system prune -f
export $(cat .env | xargs)
docker-compose -f docker-compose-local.yml up
