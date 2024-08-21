package twitch

type Prediction struct {
	Title   string
	Options []string
}

func (p *Prediction) Create() error {
	return nil
}

func (p *Prediction) Resolve(winningOption string) error {
	return nil
}

func (p *Prediction) Cancel() error {
	return nil
}
