package main

import "fmt"

// snapshot represents only indexes changed in current epoch
type Snapshot map[int]int

type SnapshotArray struct {
	snapshot  Snapshot    // Current (post Snap() changed values)
	history   [][]int     // reflog for index (of snap_id's)
	snapshots []*Snapshot // Snapshots storage
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapshot:  make(Snapshot),
		history:   make([][]int, length),
		snapshots: []*Snapshot{},
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.snapshot[index] = val
}

func (this *SnapshotArray) Snap() int {
	if len(this.snapshot) == 0 {
		this.snapshots = append(this.snapshots, nil)
		return len(this.snapshots) - 1
	}

	// ref of the new snapshot added to storage
	var snapshot = make(Snapshot)
	this.snapshots = append(this.snapshots, &snapshot)

	// range over with
	// 1. Copy value to snapshot in storage
	// 2. Adding reference snap_id to values reflog
	// 3. Delete value from curent snapsot to get it clean
	for k, v := range this.snapshot {
		snapshot[k] = v
		this.history[k] = append(this.history[k], len(this.snapshots)-1)
		delete(this.snapshot, k)
	}

	return len(this.snapshots) - 1
}

// ListSnapshots is debug fucntion, i used to check maps values.
func (this *SnapshotArray) ListSpanshots() {
	for k, i := range this.snapshots {
		fmt.Println(k, *i)
	}
}

func (this *SnapshotArray) Get(index int, snap_id int) int {

	for i := len(this.history[index]) - 1; i >= 0; i-- {
		// iterating history for index to find snapshots
		// with changed value.
		if snap_id < this.history[index][i] || *(this.snapshots[this.history[index][i]]) == nil {
			continue
		}
		// value found.
		snapshot := *(this.snapshots[this.history[index][i]])
		return snapshot[index]
	}

	return 0
}

func main() {
	sa := Constructor(3)
	sa.Set(0, 10)
	sa.Set(1, 1)
	sa.Set(2, 2)

	var snap_id int
	snap_id = sa.Snap()
	fmt.Println("Current Snapshot", snap_id)

	sa.Set(0, 3)
	sa.Set(1, 3)
	snap_id = sa.Snap()

	sa.Set(0, 8)
	sa.Set(1, 2)
	snap_id = sa.Snap()

	fmt.Println("Current Snapshot", snap_id)

	for _, v := range [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}} {
		fmt.Printf("Snapshot %[2]d, Values at index %[1]d, is %[3]d\n", v[0], v[1], sa.Get(v[0], v[1]))
	}

	sa.ListSpanshots()
}
