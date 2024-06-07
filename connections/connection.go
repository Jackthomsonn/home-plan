package connections

type Connection interface {
	Connect()
	ControlTemperature(temperature uint32)
}
