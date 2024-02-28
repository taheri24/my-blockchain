package msg

// a VisitFunc is helper for constructor as a constructor parameter , it used closure
// benefits: versatile,customize-freindly,easy-abstraction,easy-maintable and all benefits of clousures
type (
	VisitFunc func(msg *Msg)
)
