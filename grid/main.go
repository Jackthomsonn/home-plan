package grid

type Grid struct{}

func NewGridService() *Grid {
	return &Grid{}
}

func (svc *Grid) Collect() {
}
