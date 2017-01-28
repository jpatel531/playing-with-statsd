package main

import (
	"github.com/quipo/statsd"
	"log"
	"math/rand"
	"time"
)

const (
	prefix = "stickyd."
)

func main() {
	statsdclient := statsd.NewStatsdClient("127.0.0.1:18125", prefix)
	err := statsdclient.CreateSocket()

	if err != nil {
		log.Panicln(err)
	}

	interval := time.Second * 5
	stats := statsd.NewStatsdBuffer(interval, statsdclient)
	defer stats.Close()

	go countRunner(stats, time.NewTicker(time.Second*2))
	go gaugeRunner(stats, time.NewTicker(time.Second*2))
	// go setRunner(stats, time.NewTicker(time.Second*2))
	go timerRunner(stats, time.NewTicker(time.Second*2))

	select {}
}

func countRunner(stats *statsd.StatsdBuffer, ticker *time.Ticker) {
	for range ticker.C {
		log.Println("sending testcount")
		logErr(stats.Incr("testcount", int64(1)))
	}
}

func gaugeRunner(stats *statsd.StatsdBuffer, ticker *time.Ticker) {
	for range ticker.C {
		log.Println("sending testgauge")
		logErr(stats.Gauge("testgauge", rand.Int63n(100)))
	}
}

// func setRunner(stats *statsd.StatsdBuffer, ticker *time.Ticker) {
// 	for range ticker.C {
// 		log.Println("sending testset")
// 		logErr(stats.Gauge("testset", rand.Int63n(100)))
// 		// logErr(stats.)
// 	}
// }

func timerRunner(stats *statsd.StatsdBuffer, ticker *time.Ticker) {
	for range ticker.C {
		log.Println("sending testtimer")
		logErr(stats.Timing("testtimer", rand.Int63n(1000)))
	}
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
