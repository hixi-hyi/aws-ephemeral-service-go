package awsephemeral

type Chain func() error

type Session struct {
	Defers []Chain
}

func NewSession() *Session {
	return &Session{}
}

func (m *Session) Teardown() func() {
	return func() {
		var errs []error
		for _, f := range m.Defers {
			errs = append(errs, f())
		}
		if len(errs) != 0 {
			panic(errs)
		}
	}
}

func (m *Session) Add(f Chain, err error) error {
	if err != nil {
		return err
	}
	m.Defers = append(m.Defers, f)
	return nil
}
