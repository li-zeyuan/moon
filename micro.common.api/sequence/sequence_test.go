package sequence

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSequence_NextID(t *testing.T) {
	ID := NewID()
	log.Print(ID, Decompose(ID))
}

func TestSequence_NextIDOnce(t *testing.T) {
	ID := NewID()
	time.Sleep(50 * time.Millisecond)
	ID2 := NewID()
	map1 := Decompose(ID)
	map2 := Decompose(ID2)
	time1 := map1["time"]
	time2 := map2["time"]
	if time1 > time2 {
		panic("time")
	}
	log.Printf("%d %d %d", time1, time2, time2-time1)
	machineID1 := map1["machine"]
	machineID2 := map2["machine"]
	if machineID1 != machineID2 {
		panic("machine")
	}

	seq1 := map1["seq"]
	seq2 := map2["seq"]
	log.Printf("%d %d", seq1, seq2)
	mask1 := map1["mask"]
	mask2 := map2["mask"]
	if mask1 != mask2 {
		panic("mask")
	}
	log.Printf("%d %d", mask1, mask2)
}

func BenchmarkSequence_NextID(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = NewID()
	}
}

func BenchmarkSequence_NextIDDuplicate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	uniqueMap := make(map[int64]struct{})
	for n := 0; n < b.N; n++ {
		ID := NewID()
		if _, ok := uniqueMap[ID]; ok {
			panic("ID duplicate")
		}
		uniqueMap[ID] = struct{}{}
	}
}

func BenchmarkSequence_NextIDParallel(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = NewID()
		}
	})
}

func TestGetDbTbName(t *testing.T) {
	fmt.Println(100290 % 1000)
}