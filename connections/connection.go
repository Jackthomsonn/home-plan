package connections

type Connection interface {
	Connect()
	ControlTemperature(temperature float64)
}
