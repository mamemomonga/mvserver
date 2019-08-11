#include <avr/io.h>
#include <util/delay.h>

//      ATTINY13A
// RST | 1    8 | VCC
// PB3 | 2    7 | PB2
// PB4 | 3    6 | PB1
// GND | 4    5 | PB0

// PB3
#define PRST     ( 1<<PB3 )
#define PRST_INIT DDRB  |= PRST
#define PRST_H    PORTB |= PRST
#define PRST_L    PORTB &=~PRST

// PB4
#define BCK      ( 1<<PB4 )
#define BCK_INIT { DDRB &=~ BCK; PORTB |= BCK; }
#define BCK_IS_H ( PINB & BCK )

// PB2
#define LED1      ( 1<<PB2 )
#define LED1_INIT DDRB  |=  LED1
#define LED1_H    PORTB |=  LED1
#define LED1_L    PORTB &=~ LED1
#define LED1_I    PORTB ^=  LED1

// PB1
#define PON      ( 1<<PB1 )
#define PON_INIT { DDRB &=~ PON; PORTB &=~ PON; }
#define PON_IS_H ( PINB & PON )

// PB0
#define BTN      ( 1<<PB0 )
#define BTN_INIT { DDRB &=~ BTN; PORTB |= BTN; }
#define BTN_IS_H ( PINB & BTN )

int main(void) {
	BCK_INIT;
	PON_INIT;
	BTN_INIT;
	PRST_INIT;
	PRST_H;
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
		if(!PON_IS_H) {
			if(!BTN_IS_H) {
				PRST_L;
				_delay_ms(500);
				PRST_H;
			}
		}
    }
    return 0;
}
