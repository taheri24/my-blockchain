package msg

type (
	OptionType int
	Option     struct {
		optionType OptionType
	}
)

const (
	optMsgKind OptionType = iota
)

func MsgKind(kind int) Option {
	return Option{optionType: optMsgKind}
}

func applyOptions(m *Msg, opts []Option) {
	for _, opt := range opts {
		switch opt.optionType {
		case optMsgKind:
			m.ContentsKind = int8(opt.optionType)
		}
	}
}
