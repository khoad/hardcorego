package toy

//Toy ...
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// New creates a new instance of Toy
func New(name string, weight int) Toy {
	return Toy{
		Name:   name,
		Weight: weight,
	}
}

// OnHand returns the current number of this
// toy on hand.
func (t Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand updates onHand
func (t *Toy) UpdateOnHand(onHand int) int {
	t.onHand = onHand
	return t.onHand
}

// Sold returns the current number of this
// toy sold.
func (t Toy) Sold() int {
	return t.sold
}

// UpdateSold updates sold
func (t *Toy) UpdateSold(sold int) int {
	t.sold = sold
	return t.sold
}
