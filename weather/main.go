package weather

type Weather struct{}

func NewWeatherService() *Weather {
	return &Weather{}
}

func (svc *Weather) Collect() {}
