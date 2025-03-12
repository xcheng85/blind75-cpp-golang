package stable

// golang does not have void, pattern is int, error
type Model interface {
	Render() error
}

