package packageA

type Pack struct{}

func (p *Pack) Name() string    { return "PackageA" }
func (p *Pack) Version() string { return "1.5.0-A" }
