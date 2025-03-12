package stable

// golang does not have void, pattern is int, error
type Renderer interface {
	Render(Model) error
}
