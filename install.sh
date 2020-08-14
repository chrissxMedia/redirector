curl -Lo /usr/bin/redirector https://github.com/chrissxYT/chrissx.de-80/releases/download/3.1/redirector-arm
chmod +x /usr/bin/redirector

echo "[Unit]
Description=A simple HTTPS upgrador

[Service]
Type=simple
Restart=always
ExecStart=/usr/bin/redirector

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/redirector.service

systemctl enable --now redirector.service
