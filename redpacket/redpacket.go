// Revision History:
//     Initial: 2019-04-12 19:11    Jon Snow
//     Description: 红包核心实现
package redpacket

import (
	"math"
	"time"
)

const minGain = 0.01

func round(amount float64) float64 {
	x := 1 / minGain
	return math.Round(amount*x) / x
}

type GainHistory struct {
	UserId int
	Amount float64
	GainAt time.Time
}

const (
	OK   = iota + 1 // 领取成功
	Gain            // 已领取
	None            // 已抢光
)

type Packet struct {
	Amount    float64
	Remaining float64

	Total   int // total of this red packet
	History []*GainHistory
}

func NewPacket(amount float64, total int) *Packet {
	if amount < float64(total)*minGain {
		return nil
	}

	return &Packet{
		Amount:    amount,
		Remaining: amount,
		Total:     total,
		History:   make([]*GainHistory, 0, total),
	}
}

func (p *Packet) AppendHistory(uid int, amount float64) {
	p.Remaining -= amount
	p.History = append(p.History, &GainHistory{
		UserId: uid,
		Amount: amount,
		GainAt: time.Now(),
	})
}

func (p *Packet) Gain(uid int) (amount float64, status int8) {
	for _, h := range p.History {
		if h.UserId == uid {
			return h.Amount, Gain
		}
	}

	l := len(p.History)

	if p.Total == l {
		return 0, None
	}

	if p.Total-l == 1 {
		p.Remaining = round(p.Remaining)
		amount = p.Remaining
		p.AppendHistory(uid, p.Remaining)
		return amount, OK
	}

	remaining := p.Total - l // 剩余红包个数
	maxAmount := (p.Remaining - (float64(remaining-1) * minGain)) / float64(remaining) * 2
	amount = round(randomFloat64() * maxAmount)
	p.AppendHistory(uid, amount)
	return amount, OK
}
