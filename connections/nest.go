package connections

type NestConnection struct {
}

func (n *NestConnection) Connect() {}

func (n *NestConnection) ControlTemperature(temperature float64) {}

func NewNestConnection() *NestConnection {
	return &NestConnection{}
}

var _ Connection = &NestConnection{}
