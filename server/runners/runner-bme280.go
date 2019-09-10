package runners

import (
	"fmt"
	"log"
	"time"
)

func (t *runner) bme280() {
	ticker := time.NewTicker(1 * time.Second)
	bme := t.p.hw.Env
	for {
		err := bme.Sense()
		if err != nil {
			log.Println(err)

			return
		}
		thi := bme.THI()
		fmt.Printf("[BME280]\n")
		fmt.Printf("  気温: 摂氏 %2.2f 度\n", bme.Temperature())
		fmt.Printf("  湿度: %3.1f パーセント\n", bme.Humidity())
		fmt.Printf("  気圧: %4.2f ヘクトパスカル (%2.4f 気圧)\n", bme.Pressure(), bme.Atm())
		fmt.Printf("  不快指数: %2.2f %s(%d: %s)\n", thi.Value, thi.FeelJa, thi.Number, thi.FeelEn)

		<-ticker.C

		select {
		case <-t.p.ctx.Done():
			return
		default:
		}
	}
}
