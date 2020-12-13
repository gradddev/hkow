#!/bin/bash

file_path=$1

display_usage() { 
    echo "usage: deploy.sh /path/to/hkow.ipk" 
}

if [[ -z "$file_path" || ! -e $file_path || ${file_path: -4} != ".ipk" ]]; then
    display_usage
    exit 1
fi

file_name=$(basename $file_path)
scp $file_path root@192.168.1.1:/tmp/
ssh root@192.168.1.1 "opkg remove hkow; opkg install /tmp/$file_name"
