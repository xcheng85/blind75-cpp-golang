package unstable

import (
	"fmt"
)

// no class member in cpp struct
// implicit implement stable.Renderer interface
type SphereRenderer struct {
}

func (b *SphereRenderer) Render(s *Sphere) error {
	fmt.Printf("Render Sphere that radius is %f", s.r)
	return nil
}

// due to pointer receiver
func NewSphereRenderer() *SphereRenderer {
	return &SphereRenderer{}
}