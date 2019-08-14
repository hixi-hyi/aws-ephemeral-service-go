package awsephemeral

type Session struct {
	Defers []Chain
}

func NewSession() *Session {
	return &Session{}
}

func New() (*Session, func()) {
	m := &Session{}
	return m, m.Defer()
}

func (m *Session) Defer() func() {
	return func() {
		m.Teardown()
	}
}

func (m *Session) Teardown() {
	var errs []error
	for _, f := range m.Defers {
		err := f()
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		panic(errs)
	}
}

func (m *Session) Add(f Chain) {
	m.Defers = append(m.Defers, f)
}
