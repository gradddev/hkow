#!/bin/bash

docker build --tag openwrt-sdk .

docker run --rm \
           --tty \
           --volume "$(pwd)"/packages/:/home/build/openwrt/__packages__ \
           openwrt-sdk \
           /bin/bash -c "make package/hkow/{download,prepare,compile} V=s; mv bin/packages/mipsel_24kc/base/* __packages__"
