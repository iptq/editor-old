package ui

type SeekerConfig struct {
	Width int
}

type Seeker struct {
	config SeekerConfig
}

func NewSeeker(config SeekerConfig) *Seeker {
	return &Seeker{
		config: config,
	}
}

func (s Seeker) Draw(ctx Context) {
}
