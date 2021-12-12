package node

func (n *Node) Lock() {
	n.locked = true
}

func (n *Node) Unlock() {
	n.locked = false
}

func (n *Node) IsLocked() bool {
	return n.locked
}

func (n *Node) Select() {
	n.selected = true
}

func (n *Node) Unselect() {
	n.selected = false
}

func (n *Node) IsSelected() bool {
	return n.selected
}
