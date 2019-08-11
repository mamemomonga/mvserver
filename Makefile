NAME=rpi-volumio-status-led

run: build
	sudo $(NAME)

build:
	go build -o $(NAME) .

install:
	cp -f $(NAME) /usr/local/bin/
	NAME=$(NAME) ./systemd-setup.sh
	systemctl daemon-reload
	systemctl enable $(NAME).service
	systemctl start $(NAME).service

uninstall:
	systemctl stop $(NAME).service
	systemctl disable $(NAME).service
	rm -f /etc/systemd/system/$(NAME).service
	rm -f /usr/local/sbin/$(NAME)

.PHONY: run build
