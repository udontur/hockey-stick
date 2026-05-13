#!/bin/bash

# GitHub raw linked redirected on install.hadrian.cc/hockey-stick by cloudflare page rule

arch=$(uname -i)
# TODO: jq isn't preinstalled on some system
version=$(curl https://api.github.com/repos/udontur/hockey-stick/releases/latest -s | jq .name -r)
link="https://github.com/udontur/hockey-stick/releases/download/$version/hockey-stick_${version}_Linux"

dpkg --help &> /dev/null
if [ $? -eq 0 ]; then
    echo "Debian-based operating system detected. "
    if [ $arch = "x86_64" ]; then
        curl -LsSf "${link}_x86_64.deb" -o hockey-stick.deb
        sudo apt install ./hockey-stick.deb
    elif [ $arch = "aarch64" ]; then
        curl -LsSf "${link}_aarch64.deb" -o hockey-stick.deb
        sudo apt install ./hockey-stick.deb
    else
        echo "$arch architecture is not supported."
        exit 1
    fi
    exit 0
fi

rpm --help &> /dev/null
if [ $? -eq 0 ]; then
    echo "RPM-based operating system detected. "
    if [ $arch = "x86_64" ]; then
        curl -LsSf "${link}_x86_64.rpm" -o hockey-stick.rpm
        sudo dnf install ./hockey-stick.rpm
    elif [ $arch = "aarch64" ]; then
        curl -LsSf "${link}_aarch64.rpm" -o hockey-stick.rpm
        sudo dnf install ./hockey-stick.rpm
    else
        echo "$arch architecture is not supported."
        exit 1
    fi
    exit 0
fi

echo "Other Linux operating system detected. "
if [ $arch = "x86_64" ]; then
    curl -LsSf "${link}_x86_64.tar.gz" -o hockey-stick.tar.gz
    mkdir -p hockey-stick
    tar -xvzf hockey-stick.tar.gz -C hockey-stick
    cp ./hockey-stick/hockey-stick ~/.local/bin/hockey-stick
elif [ $arch = "aarch64" ]; then
    curl -LsSf "${link}_aarch64.tar.gz" -o hockey-stick.tar.gz
    mkdir -p hockey-stick
    tar -xvzf hockey-stick.tar.gz -C hockey-stick
    cp ./hockey-stick/hockey-stick ~/.local/bin/hockey-stick
else
    echo "$arch architecture is not supported."
    exit 1
fi
