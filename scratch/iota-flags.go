// Explores iota as a binary flag enum
package main
type Flags uint

const (
	// remember, iota starts at 0
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast            // supports broadcast access capability
	FlagLoopback             // is a loopback interface
	FlagPointToPoint         // belongs to a point-to-point link
	FlagMulticast            // supports multicast access capability
)

func main() {

}

// functions to test, set and clear various flags

func isUp(v Flags) bool {
	return v & FlagUp == FlagUp
}
func setUp(v *Flags) {
	*v |= FlagUp
}
func clearUp(v *Flags) {
	*v &^= FlagUp // the difference operator
}
func isBroadcast(v Flags) bool {
	return v & FlagBroadcast == FlagBroadcast
}
func setBroadcast(v *Flags) {
	*v |= FlagBroadcast
}
func clearBroadcast(v *Flags) {
	*v &^= FlagBroadcast // the difference operator
}
// ..