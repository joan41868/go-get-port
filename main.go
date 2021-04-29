package go_get_port

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	MIN_BOUNDARY = 1
	MAX_BOUNDARY = 65536
)

func getRandomBetween() int {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(MAX_BOUNDARY-MIN_BOUNDARY+1) + MIN_BOUNDARY
	return rnd
}

func attemptGetPort() int {
	randPort := getRandomBetween()
	srv := http.Server{
		Addr: fmt.Sprintf(":%d", randPort),
	}

	go srv.ListenAndServe()

	srv.Shutdown(context.TODO())
	return randPort
}

func GetPort() int {
	var i int
	for {
		i = attemptGetPort()
		if i > 0 {
			break
		}
	}
	return i
}
