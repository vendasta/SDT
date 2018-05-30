package sdt

type Application interface {
	// Ready will return true when the application is ready to serve traffic
	Ready() bool

	// Ping will return an error when the application is not healthy
	Ping() error
}

func NewApplication() Application {
	return &app{}
}
