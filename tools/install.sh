#!/usr/bin/env bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $DIR/..

# Project scripts
sudo mkdir -p /opt/project-fs
sudo cp -R todo /opt/project-fs/todo
sudo ln -s /opt/project-fs/todo/bin/todo /usr/local/bin

# FUSE binary
sudo cp target/projectfs /opt/project-fs/projectfs

# Filesystem service
mkdir -p ~/.config/systemd/user/
cp systemd/project-fs.service ~/.config/systemd/user/
systemctl --user enable ~/.config/systemd/user/project-fs.service
systemctl --user daemon-reload
systemctl --user start project-fs.service

echo -e "\n\nInstallation Complete\nPlease run 'todo config'"

