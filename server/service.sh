#!/bin/bash
set -eu

cat > /etc/systemd/system/$INAME.service << EOS
[Unit]
Description=$INAME
After=network.target

[Service]
Type=simple
User=$IUSER
WorkingDirectory=/tmp
ExecStart=/usr/local/bin/$INAME
ExecReload=/bin/kill -SIGUSR1 \$MAINPID
TimeoutSec=15
Restart=always

[Install]
WantedBy=multi-user.target
EOS

