NAME=mvserver
SRCS=$(shell find . -type f -name '*.go')

run: $(NAME)
	./$(NAME)

$(NAME): $(SRCS)
	go build -o $(NAME) .

install:
	cp -f $(NAME) /usr/local/bin/
	INAME=$(NAME) IUSER=volumio ./service.sh
	systemctl daemon-reload
	systemctl enable $(NAME).service
	systemctl start $(NAME).service

uninstall:
	systemctl stop $(NAME).service
	systemctl disable $(NAME).service
	rm -f /etc/systemd/system/$(NAME).service
	rm -f /usr/local/sbin/$(NAME)

.PHONY: run build

