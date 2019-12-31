// Revision history:
//     Init: 2019/12/27    Jon Snow

package redpacket

import (
	"math/rand"
	"time"
)

var (
	float64Pool = make(chan float64, 1<<7)
	r           = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func init() {
	go func() {
		for {
			float64Pool <- r.Float64()
		}
	}()
}

func randomFloat64() float64 {
	return <-float64Pool
}
