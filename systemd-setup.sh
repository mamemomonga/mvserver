#!/bin/bash
set -eu

cat > /etc/systemd/system/$NAME.service << EOS
[Unit]
Description=$NAME
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/tmp
ExecStart=/usr/local/bin/$NAME
ExecReload=/bin/kill -SIGUSR1 \$MAINPID
TimeoutSec=15
Restart=always

[Install]
WantedBy=multi-user.target
EOS

