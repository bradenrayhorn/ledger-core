package server

type server struct{}

func CreateServer() *server {

	return &server{}
}

func (s server) Run() error {

	return nil
}
