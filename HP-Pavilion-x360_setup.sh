#!/usr/bin/env sh
# This script will need internet to work
# Login to snap first too

set -x

sudo apt update
sudo apt install -y git

### Wifi drivers ###
#sudo apt update
#sudo apt install -y dkms git
#git clone https://github.com/tomaspinho/rtl8821ce
#cd rtl8821ce
#sudo ./dkms-install.sh
#sudo modprobe 8821ce
#cd ..
#rm -rf rtl8821ce

### Firefox Touch Scrolling ###
echo export MOZ_USE_XINPUT2=1 | sudo tee /etc/profile.d/use-xinput2.sh


### Apps ###
# Htop
snap install htop
snap connect htop:mount-observe 
snap connect htop:network-control 
snap connect htop:process-control 
snap connect htop:system-observe 

# VS Code
snap install code --classic

# Multipass
snap install multipass --classic --edge

# Snapcraft
snap install snapcraft --classic

# IoTop
sudo apt install iotop

# Utilities
sudo apt install curl
sudo apt install mosh
sudo apt install gddrescue


### Gnome Settings ###
gsettings set org.gnome.desktop.interface cursor-size 64
gsettings set org.gnome.shell favorite-apps "['firefox.desktop', 'thunderbird.desktop', 'org.gnome.Nautilus.desktop', 'org.gnome.Terminal.desktop', 'code_code.desktop', 'postman_postman.desktop']
"


