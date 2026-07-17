package internal

type GreetServiceImpl struct{}

func NewGreetServiceImpl() *GreetServiceImpl {
	return &GreetServiceImpl{}
}

func (g *GreetServiceImpl) Greet(name string) string {
	return "Hello " + name + "!"
}
