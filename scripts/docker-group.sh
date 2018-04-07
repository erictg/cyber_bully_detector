
#!/usr/bin/env bash


# Exit as failure
set -e

echo install docker group
sudo groupadd docker || true
sudo usermod -aG docker $USER
newgrp docker