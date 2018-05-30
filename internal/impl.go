package sdt


type app struct {
}

func (a *app) Ready() bool {
	return true
}

func (a *app) Ping() error {
	return nil
}
