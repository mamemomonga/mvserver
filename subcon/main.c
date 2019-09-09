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

#define FALSE 0
#define TRUE  1

void blink_fast(uint8_t count);
void fadein();
void fadeout();

uint8_t led_state;

int main(void) {
	BCK_INIT;
	PON_INIT;
	BTN_INIT;
	PRST_INIT;
	PRST_H;
	LED1_INIT;
	LED1_L;

	led_state=FALSE;

	blink_fast(5);
	_delay_ms(500);

    for(;;){
		uint8_t count=0;
		for(uint8_t i=0; i< 255; i++) {
			if(!BCK_IS_H) {
				count++;
			}
		}
		if(count>0) {
			fadein();
		} else {
			fadeout();
		}
		if(!BTN_IS_H) {
			if(!PON_IS_H) {
				PRST_L;
				_delay_ms(100);
				PRST_H;
			}
		}
    }
    return 0;
}

void blink_fast(uint8_t count) {
	for(uint8_t i=0; i<count; i++) {
		LED1_H;
		_delay_ms(50);
		LED1_L;
		_delay_ms(50);
	}
}

void fadein() {
	if(led_state) {
		return;
	}
	led_state=TRUE;
	for(uint8_t i=0;i<255;i++) {
		for(uint8_t j=0;i<255;i++) {
			if(i>=j) {
				LED1_H;
			} else {
				LED1_L;
			}
		}
	}
	LED1_H;
}
void fadeout() {
	if(!led_state) {
		return;
	}
	LED1_L;
	led_state=FALSE;
}

