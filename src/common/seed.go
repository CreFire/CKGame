package model

type Char uint64

func (c Char) CheckRight(index int32) bool {
	if c&(1<<index) != 0 {
		return true
	}
	return false
}
