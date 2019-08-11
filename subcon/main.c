#include <avr/io.h>
#include <util/delay.h>

#define BCK      ( 1<<PB4 )
#define BCK_INIT { DDRB &=~ BCK; PORTB |= BCK; }
#define BCK_IS_H ( PINB & BCK )

#define LED1      ( 1<<PB2 )
#define LED1_INIT DDRB  |=  LED1
#define LED1_H    PORTB |=  LED1
#define LED1_L    PORTB &=~ LED1
#define LED1_I    PORTB ^=  LED1

int main(void) {
	BCK_INIT;
	LED1_INIT;
	LED1_L;

    for(;;){
		uint8_t count=0;
		uint8_t i;
		for(i=0; i< 255; i++) {
			if(!BCK_IS_H) {
				count++;
			}
		}
		if(count>0) {
			LED1_H;
		} else {
			LED1_L;
		}
    }
    return 0;
}
