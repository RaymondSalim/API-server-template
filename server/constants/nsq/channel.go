package constants

const service = "web_server_template"

const (
	Increment = iota
	Decrement
)

var Channels = map[int]string{
	Increment: service + ".increment",
	Decrement: service + ".decrement",
}
