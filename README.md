# Back

## Systemd

[Unit] Description=Go back-end

[Service] User=ubuntu WorkingDirectory=/home/ubuntu/code/back ExecStart=/usr/local/go/bin/go run main.go Restart=always Environment=GOMODCACHE=/home/ubuntu/go/pkg/mod Environment=GOPATH=/home/ubuntu/go Environment=GOCACHE=/home/ubuntu/go/pkg/cache StandardOutput=file:/var/log/go-back-service.log StandardError=file:/var/log/go-back-service.log

[Install] WantedBy=multi-user.target