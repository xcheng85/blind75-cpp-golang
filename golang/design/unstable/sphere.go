package unstable

import (
	"github.com/xcheng85/blind75-cpp-golang/golang/design/stable"
	_ "github.com/xcheng85/blind75-cpp-golang/golang/design/stable"
)

// composition through struct embedding
type Sphere struct {
	r float32
	// unstable SphereRenderer which knows domain Sphere
	renderer *SphereRenderer
}

func (b *Sphere) Render() error {
	b.renderer.Render(b)
	return nil
}

// go style constructor
func NewSphereModel(r float32, renderer *SphereRenderer) stable.Model {
	// the & is needed due to pointer receiver impl
	return &Sphere{
		r,
		renderer}
}
