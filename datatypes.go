package soyaltools

// NodeCardlist card list data
type NodeCardlist struct {
	Nid  int
	Tags []NodeCard
}

// NodeCard card data
type NodeCard struct {
	Nid      int
	UserAddr int
	Tag      string
	PIN      string
	Expire   string
	UserName string
}
type NodeLoglist struct {
	Nid  int
	Tags []NodeEvenlog
}
type NodeEvenlog struct {
	Nid      int
	Sec      byte
	Min      byte
	Hour     byte
	Weekday  byte
	Date     byte
	Month    byte
	Year     byte
	Port     byte
	UserAddr int
	Tag      string
}
