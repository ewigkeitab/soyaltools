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
