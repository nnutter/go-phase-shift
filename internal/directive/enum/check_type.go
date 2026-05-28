package enum

//sumtype:decl
type CheckType interface {
	isCheckType()
}

func ParseCheckType(s string) CheckType {
	switch s {
	case NonmutatingName:
		return Nonmutating{}
	default:
		return nil
	}
}

const NonmutatingName = "nonmutating"

type Nonmutating struct{}

func (Nonmutating) isCheckType() {}
