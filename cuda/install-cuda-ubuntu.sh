#!/bin/bash

# Update Ubuntu Repositories
sudo apt update -y

# Install common tools
sudo apt install tmux tmate emacs-nox vim htop git gcc make jq curl linux-headers-$(uname -r) python -y

# Install CUDA Toolkit
wget https://developer.download.nvidia.com/compute/cuda/repos/ubuntu1804/x86_64/cuda-ubuntu1804.pin
sudo mv cuda-ubuntu1804.pin /etc/apt/preferences.d/cuda-repository-pin-600
wget http://developer.download.nvidia.com/compute/cuda/10.2/Prod/local_installers/cuda-repo-ubuntu1804-10-2-local-10.2.89-440.33.01_1.0-1_amd64.deb
sudo dpkg -i cuda-repo-ubuntu1804-10-2-local-10.2.89-440.33.01_1.0-1_amd64.deb
sudo apt-key add /var/cuda-repo-10-2-local-10.2.89-440.33.01/7fa2af80.pub
sudo apt update -y
sudo DEBIAN_FRONTEND=noninteractive apt -y install cuda
rm -f cuda-repo*

# Install Docker
sudo apt-get remove docker docker-engine docker.io -y
sudo apt install docker.io -y
sudo systemctl start docker

# Add user to docker group
sudo usermod -aG docker $(whoami)

# CUDA Samples tests
# Evironment Setup
echo "export PATH=/usr/local/cuda-10.2/bin/:\$PATH" >> $HOME/.profile
echo "export LD_LIBRARY_PATH=/usr/local/cuda-10.2/lib64\${LD_LIBRARY_PATH:+:\${LD_LIBRARY_PATH}}" >>  $HOME/.profile
source $HOME/.profile
cuda-install-samples-10.2.sh $HOME/cuda-samples
pushd $HOME/cuda-samples/NVIDIA_CUDA-10.2_Samples/0_Simple/simplePrintf
make
./simplePrintf
popd

# Echo
echo "That's all, if no errors you're all setup for CUDA work"
