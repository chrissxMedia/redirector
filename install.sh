curl -Lo /usr/bin/redirector https://github.com/chrissxYT/chrissx.de-80/releases/download/1.0/redirector-arm

echo "[Unit]
Description=A simple static HTML

[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/redirector

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/redirector.service

systemctl enable --now redirector.service
