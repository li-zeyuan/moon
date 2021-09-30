package sequence

var seq *Sequence

func init() {
	setting := Setting{
		MachineID:      nil,
		CheckMachineID: nil,
	}
	seq = NewSequence(setting)
	//time.Sleep(time.Second)   // 因为会向前接1s内的id，启动时需要确保
}

func NewID() int64 {
	return int64(seq.NextID())
}

func Decompose(id int64) map[string]int64 {
	var result = map[string]int64{
		"mask":    id >> 63,
		"shift":   ((id >> timeShift & maxTime) << 3) / 1000,
		"machine": id >> machineShift & maxMachine,
		"seq":     id & maxSeq,
	}
	result["time"] = result["shift"] + startTime/1000
	return result
}
