#!/usr/bin/env bash

# Filesystem service
systemctl --user stop project-fs.service
systemctl --user disable project-fs.service

# Binaries & scripts
sudo rm -rf /opt/project-fs

# Dangling references
sudo rm -f /usr/local/bin/todo

echo -e "\n\nUninstall Complete\nTa"

