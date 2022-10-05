package constants

const (
	VisitCounter = iota
	ResetCounter
)

var Topics = map[int]string{
	VisitCounter: "visit_counter",
	ResetCounter: "reset_counter",
}
