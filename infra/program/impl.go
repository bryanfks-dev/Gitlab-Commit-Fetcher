package program

func (p *program) Start() error {
	_, err := p.Prog.Run()
	return err
}
