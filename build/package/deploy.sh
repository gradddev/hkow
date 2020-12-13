#!/bin/bash

file_path=$1

display_usage() { 
    echo "usage: deploy.sh /path/to/hkow.ipk" 
}

if [[ -z "$file_path" || ! -e $file_path || ${file_path: -4} != ".ipk" ]]; then
    display_usage
    exit 1
fi

scp "$file_path" root@192.168.1.1:~/hkow.ipk
ssh root@192.168.1.1 << 'EOF'
  opkg remove hkow;
  opkg install ~/hkow.ipk;
  rm ~/hkow.ipk;
  /etc/init.d/hkow enable;
  /etc/init.d/hkow start;
EOF
