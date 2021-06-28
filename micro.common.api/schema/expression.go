package schema

type Expression string

func (e *Expression) Eq() string {
	return " = "
}
