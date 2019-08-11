# rpi-volumio-status-led

VolumioにLEDをつけたりなどする

## PCM5102Aの接続

![image](images/pcm5102a.png)

## SN74HC595の接続

シフトレジスタ、QAからQHまでは、1KΩの抵抗を挟んでLEDと接続

![image](images/sn74hc595.png)
![image](images/595chart.png)

## Raspberry Pi J8 Header

![image](images/pi.png)

# 準備

/boot/config.txtに追記してリブート

	dtparam=spi=on
	dtoverlay=gpio-poweroff,gpiopin=12,active_low="y"
	dtoverlay=gpio-shutdown,gpio_pin=16

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

