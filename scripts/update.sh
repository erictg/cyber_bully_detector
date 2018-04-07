set -e
sudo docker-compose -f docker-compose.yml down
git pull
sudo docker-compose -f docker-compose.yml build
sudo docker-compose -f docker-compose.yml up -d