# サブコントローラ

* BCK信号があるときLEDを点灯させる
* Piが起動していないとき、ボタンでRaspberry Piリセット(起動)信号を送る

* ATMEL(microchip) ATtiny13A
* AVRISPmk2

avr開発環境

	$ sudo apt update
	$ sudo apt install -y gcc-avr avrdude avr-libc build-essential

ヒューズ設定

	$ sudo make fuse

ビルド

	$ make

書込

	$ make flash


