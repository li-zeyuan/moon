package sequence

import (
	"errors"
	"net"
	"sync"
	"time"
)

const (
	bitLenTime    uint8 = 39 // 时间戳位数
	bitLenMachine uint8 = 16 // 机器码位数
	bitLenSeq     uint8 = 8  // 序列号位数
	machineShift        = bitLenSeq
	timeShift           = bitLenMachine + bitLenSeq
	maxTime             = -1 ^ (-1 << bitLenTime)
	maxSeq              = -1 ^ (-1 << bitLenSeq)
	maxMachine          = -1 ^ (-1 << bitLenMachine)

	maxSeqSecond = maxSeq * 1000 >> 3 // 单位为8ms，所以一秒最大可生成maxSeq * 125 = 32000, 可以用139年
	startTime    = 1476979200000      // 2016-10-21 00:00:00
)

type Sequence struct {
	mu        sync.Mutex
	monotonic time.Time
	MachineID uint16
	Sequence  uint16
	SeqSecond uint32
	LastMilli int64
}

type Setting struct {
	MachineID      func() (uint16, error)
	CheckMachineID func(uint16) bool
}

func NewSequence(setting Setting) *Sequence {
	seq := new(Sequence)
	seq.Sequence = 0

	var err error
	if setting.MachineID != nil {
		seq.MachineID, err = setting.MachineID()
	} else {
		seq.MachineID, err = getMachine()
	}

	if err != nil || (setting.CheckMachineID != nil && setting.CheckMachineID(seq.MachineID)) {
		panic("create machine id error")
	}

	// 使用单调时钟
	now := time.Now()
	seq.monotonic = now.Add(time.Unix(startTime/1000, (startTime%1000)*1000000).Sub(now))
	return seq
}

func getMachine() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3])&maxMachine, nil
}

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipNet, ok := a.(*net.IPNet)
		if !ok || ipNet.IP.IsLoopback() {
			continue
		}

		ip := ipNet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func (seq *Sequence) NextID() uint64 {
	seq.mu.Lock()
	sec, mSec := formatTime(seq.monotonic)
	lastSec := getLastSec(seq.LastMilli)
	if lastSec < sec {
		seq.Sequence = 0
		seq.SeqSecond = 0
		seq.LastMilli = mSec
	} else {
		seq.Sequence = (seq.Sequence + 1) & maxSeq
		seq.SeqSecond = (seq.SeqSecond + 1) % maxSeqSecond
		// 当前8ms内的数量已经用完,但1s内的数量还没用完，直接用下8ms的id
		if seq.SeqSecond != 0 && seq.Sequence == 0 {
			seq.LastMilli++
		}
		// 已经超过1s的数量，等待下一秒到来
		lastSec = getLastSec(seq.LastMilli)
		if seq.SeqSecond == 0 || lastSec > sec {
			time.Sleep(seq.getSleepTime(sec))
			seq.Sequence = 0
			_, seq.LastMilli = formatTime(seq.monotonic)
		}
	}

	ID := (uint64((seq.LastMilli)&maxTime) << timeShift) |
		(uint64(seq.MachineID) << machineShift) |
		(uint64(seq.Sequence))
	seq.mu.Unlock()
	return ID
}

func formatTime(start time.Time) (int64, int64) {
	now := time.Since(start).Nanoseconds()
	sec := now / 1e9
	mSec := (now / 1e6) >> 3
	return sec, mSec
}

func getLastSec(mSec int64) int64 {
	return (mSec << 3) / 1e3
}

func (seq *Sequence) getSleepTime(second int64) time.Duration { // ms
	return time.Duration(second+1)*1000*time.Millisecond -
		time.Millisecond*time.Duration(time.Since(seq.monotonic).Milliseconds())
}
