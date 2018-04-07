#!/usr/bin/env bash

#!/usr/bin/env bash
# This script is as much a memory aid as a complete
# fire and forget recipe

# Exit as failure
set -e

# increase map size for elasticsearch
sudo sysctl -w vm.max_map_count=262144

# install docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
sudo apt-get update || true

#apt-cache policy docker-ce
sudo apt-get install -y docker-ce

sudo apt-get install -y unzip

# execute without sudo
sudo usermod -aG docker ${USER}

echo install docker-compose
#sudo curl -o /usr/local/bin/docker-compose -L "https://github.com/docker/compose/releases/download/1.15.0/docker-compose-$(uname -s)-$(uname -m)"
#sudo chmod +x /usr/local/bin/docker-compose