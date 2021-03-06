package triangolatte

// Element type is a wrapper for Point used in cyclic list.
type Element struct {
	Prev, Next *Element
	Point      Point
}

// Insert function takes Point and Element (optionally nil) and adds Point
// wrapped in a new Element after given Element, if present.
func Insert(p Point, e *Element) *Element {
	new := Element{Point: p}

	if e != nil {
		new.Next = e.Next
		new.Prev = e
		e.Next.Prev = &new
		e.Next = &new
	} else {
		new.Prev = &new
		new.Next = &new
	}
	return &new
}

// Remove detaches Element from the list (preserving its connections for reference).
// Be aware of potential memory leaks.
func (e *Element) Remove() {
	e.Next.Prev = e.Prev
	e.Prev.Next = e.Next
}
