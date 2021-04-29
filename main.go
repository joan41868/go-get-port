package go_get_port

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
)

const (
	MaxBoundary = 65536
)

func getRandomBetween() int {
	rnd, err := rand.Int(rand.Reader, big.NewInt(MaxBoundary))
	if err != nil {
		panic(err)
	}
	return int(rnd.Int64())
}

func attemptGetPort() int {
	randPort := getRandomBetween()
	srv := http.Server{
		Addr: fmt.Sprintf(":%d", randPort),
	}
	flg := false
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			flg = true
			return
		}
	}()

	_ = srv.Shutdown(context.TODO())
	if flg {
		return 0
	}
	return randPort
}

func GetPort() int {
	locked := make(map[int]bool)
	var i int
	for {
		i = attemptGetPort()
		if i > 0 {
			break
		} else {
			locked[i] = true
		}
	}
	return i
}
