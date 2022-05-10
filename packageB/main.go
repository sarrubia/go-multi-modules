package packageB

type Pack struct{}

func (p *Pack) Name() string    { return "PackageB" }
func (p *Pack) Version() string { return "1.4.1-B" }
