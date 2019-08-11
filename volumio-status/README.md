# volumio-status

* Volumioの状況をLEDでお知らせ
* シャットダウンボタン

# 準備

/boot/config.txtに追記してリブート

	# SPIを有効にする
	dtparam=spi=on
	
	# パワーオフ状態でなければ、GPIO22をHIGHにする
	dtoverlay=gpio-poweroff,gpiopin=22,active_low="y"

SPIの確認

	$ ls /dev/spidev0.*
	/dev/spidev0.0  /dev/spidev0.1

cgoを使うのでgccなどが必要

	$ sudo apt install build-essential

goのインストール

	$ cd /usr/local/src

Raspberry Pi arm7lの場合

	$ sudo wget https://dl.google.com/go/go1.12.7.linux-armv6l.tar.gz
	$ sudo tar zxvC /usr/local -f go1.12.7.linux-armv6l.tar.gz

goの設定

	$ cd
	$ mkdir ~/.go
	$ echo "export GOPATH=$HOME/.go" >> ~/.bashrc
	$ echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
	$ echo 'export "PATH=/usr/local/go/bin:$PATH"' >> ~/.bashrc
	$ exec $SHELL
	$ go env
	$ go version

# volumio-status

	$ cd volumio-status

## 実行

	$ make	

## インストール

	$ make build
	$ sudo make install

## アンインストール

	$ sudo make uninstall

