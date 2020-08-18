package member

var (
	joinCallbacks  = make([]func(*Member), 0)
	leaveCallbacks = make([]func(*Member), 0)
)

func RegisterJoinCallback(callback func(*Member)) {
	joinCallbacks = append(joinCallbacks, callback)
}
func RegisterLeaveCallbacks(callback func(*Member)) {
	leaveCallbacks = append(leaveCallbacks, callback)
}
