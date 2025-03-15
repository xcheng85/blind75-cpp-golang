package visitor

// domain object
// class Volume
type Meta struct {
	Name       string
	Annotation string
}

// composition instead of inheritance
type Volume struct {
	Meta Meta
}

// class HeightField
type HeightField struct {
	Meta Meta
}

// generic
func NewDomainObj[T HeightField | Volume] (name string) *T {
	return &T{
		Meta: Meta{
			 Name: name,
		},
	}
}

// Annotation
type IObjectVisitor interface {
	// no function overload in go
	VisitVolume(*Volume)
	VisitHeightField(*HeightField)
}

type IVisitableObject interface {
	Accept(IObjectVisitor)
}

func (vol *Volume) Accept(v IObjectVisitor) {
	v.VisitVolume(vol)
}

func (hf *HeightField) Accept(v IObjectVisitor) {
	v.VisitHeightField(hf)
}

// class Annotation
type AnnotationVisitor struct {
}

func (a *AnnotationVisitor) VisitVolume(v *Volume) {
	v.Meta.Annotation = "Volume Annotation"
}

func (a *AnnotationVisitor) VisitHeightField(hf *HeightField) {
	hf.Meta.Annotation = "HeightField Annotation"
}

// naming convention, post_fix v2 for example
type NamingV2SpecVisitor struct {
}

const NameV2Postfix string = "_v2"

func (a *NamingV2SpecVisitor) VisitVolume(v *Volume) {
	v.Meta.Name += NameV2Postfix
}

func (a *NamingV2SpecVisitor) VisitHeightField(hf *HeightField) {
	hf.Meta.Name += NameV2Postfix
}

// ctor for AnnotationClass
func NewGenericVisitor[V NamingV2SpecVisitor | AnnotationVisitor]() *V {
	return &V{}
}
