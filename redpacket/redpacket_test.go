// Revision History:
//     Initial: 2019-04-12 19:22    Jon Snow
package redpacket

import (
	"testing"
)

func TestRedPacket(t *testing.T) {
	testCount := 10000
	userNum := 10
	users := make([]float64, userNum)
	for i := 0; i < testCount; i++ {
		result := make([]float64, userNum)
		p := NewPacket(100, userNum)
		for uid := range users {
			if amount, status := p.Gain(uid); status == OK {
				users[uid] += amount
				result[uid] = amount
			}
		}
		// log.Println(result)
		// time.Sleep(time.Second)
	}

	total := .0
	for _, totalGain := range users {
		total += totalGain
	}
	t.Logf("total gain: %f", total)

	for uid, gain := range users {
		t.Logf("user %d gain: %f ex: %.4f", uid, gain, gain/total)
	}
}
