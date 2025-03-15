package main

// unidirection dependencies, unstable ---> stable only
import (
	"github.com/xcheng85/blind75-cpp-golang/golang/design/unstable"
	"github.com/xcheng85/blind75-cpp-golang/golang/design/visitor"
	"fmt"
)

// int main() cpp
func main() {
	// strategy example
	renderer := unstable.NewSphereRenderer()
	sphere := unstable.NewSphereModel(10, renderer)
	sphere.Render()

	// adaptor example
	s := unstable.NewSliceStackAdaptor[int]()
	fmt.Println(s.Empty())
	fmt.Println(s.Size())
	s.Push(100)
	s.Push(200)
	fmt.Println(s.Size())
	fmt.Println(s.Pop())
	s.Push(300)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// visitor pattern 
	vol := visitor.NewDomainObj[visitor.Volume]("vol_0")
	hf := visitor.NewDomainObj[visitor.HeightField]("hf_0")
	
	annotationVisitor := visitor.NewGenericVisitor[visitor.AnnotationVisitor]()
	namev2Visitor := visitor.NewGenericVisitor[visitor.NamingV2SpecVisitor]()

	vol.Accept(annotationVisitor)
	vol.Accept(namev2Visitor)

	hf.Accept(namev2Visitor)
	fmt.Println(vol.Meta)
	fmt.Println(hf.Meta)
}