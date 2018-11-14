package constants

// package requires own folder

import (
	"fmt"
)

const (
	kb = 1024
	mb = kb * 1024
	gb = mb * 1024
	tb = gb * 1024
	pb = tb * 1024
)

// ByteSize is the type
type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= pb:
		return "Very big"
	case b >= tb:
		return fmt.Sprintf("%.2ftb", b/tb)
	case b >= gb:
		return fmt.Sprintf("%.2fgb", b/gb)
	case b >= mb:
		return fmt.Sprintf("%.2fmb", b/mb)
	case b >= kb:
		return fmt.Sprintf("%.2fkb", b/kb)
	}
	return fmt.Sprintf("%.2dbytes", b)
}

// func main() {
// 	fmt.Println(ByteSize(2048))
// 	fmt.Println(ByteSize(87652873984987))
// }
