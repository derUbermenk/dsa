package data_structures

type LinkedList interface {
	find_item(item int) *LinkedList
	insert_item(item int) *LinkedList
	delete_item(item int)
}

func NewLinkedList(item int) *linkedList {
	return &linkedList{
		Item: item,
		Next: nil,
	}
}

type linkedList struct {
	Item int
	Next *linkedList
}

func (l *linkedList) find_item(item int) *linkedList {
	// base case: list is empty
	if l == nil {
		return nil
	}

	if l.Item == item {
		return l
	} else {
		return l.Next.find_item(item)
	}
}

func (l *linkedList) insert_item(item int) *linkedList {
	// let p the new head
	// create a linked_list with
	// 	p.Item = item
	//  p.Next = l
	//  make l point to p instead
	var p linkedList
	p.Item = item
	p.Next = l
	l = &p
	return l
}

// non struct method of insertion
func insert_item(l **linkedList, item int) {
	// let p the new head
	// create a linked_list with
	// 	p.Item = item
	//  p.Next = l
	//  make l point to p instead
	var p linkedList
	p.Item = item
	p.Next = *l
	*l = &p
}

// the predecessor is confirmed by the fact that the next node's
// item is equal to the item being described
func (l *linkedList) predecessor_list(item int) (predecessor *linkedList) {
	// Base Case:
	// the predecessor is proven to be not be present
	// 		as in the case when the item is already infront
	// if the given list is empty
	if l == nil || l.Next == nil {
		return nil
	}

	// the current node is the predecessor, i.e. it is confirmed
	// if the next node's item is equal to the item given.
	if l.Next.Item == item {
		return l
	} else {
		// else keep traversing
		return l.Next.predecessor_list(item)
	}
}

// method implementation
func (l *linkedList) delete_item(item int) *linkedList {
	var p *linkedList    // the node to be deleted
	var pred *linkedList // the predecessor to be deleted

	// find item in the list
	p = l.find_item(item)

	// if p exists
	if p != nil {
		// get its predecessor and point it to
		// if there is a predecessor, point it to the node
		// next to p
		// else if the predecessor does not exists
		// such is the case when the item is in front
		// return the next node instead
		pred = l.predecessor_list(item)
		if pred != nil {
			pred.Next = p.Next
		} else {
			l = p.Next
		}
	}

	// return l in the case that the item does not exist
	return l
}

// non method based delete approach
// manipulates the list itself
func delete_item(l **linkedList, item int) {
	var p *linkedList    // node to delete
	var pred *linkedList // predecessor of node to delete

	// find the node to delete
	p = (*l).find_item(item)
	if p != nil { // if the node is present
		// delete the node and manipulate the list
		// according to node position
		// positions is determined to pred value
		pred = (*l).predecessor_list(item)
		if pred != nil {
			// this means that node is along the middle
			// point the predecessors Next node to p's next Node
			pred.Next = p.Next
		} else if pred == nil {
			// this means node is in front
			*l = p.Next
		}
	}

	// if p is nil, such is the case when item does not exist
	// do nothing
}

// prints the list
// just for debugging
func (l *linkedList) get_items() []int {
	var items []int

	for current_node := l; current_node != nil; current_node = current_node.Next {
		items = append(items, current_node.Item)
	}

	return items
}
