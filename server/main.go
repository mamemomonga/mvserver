package main

import(
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/mamemomonga/rpi-volumio-status-led/server/hardwares"
	"github.com/mamemomonga/rpi-volumio-status-led/server/runners"
)

func main() {
	log.Println("*** START ***")

	hw := hardwares.New()
	if err := hw.Init(); err != nil {
		log.Fatal(err)
	}
	defer hw.Close()

	run := runners.New(hw)
	run.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT)

	select {
	case <-quit:
		log.Println("*** INTERRUPT ***")
	case <-run.RunnersDone:
		log.Println("*** RUNNERS DONE ***")
	}
	log.Println("*** STOP ***")
}
