package awseph

func MustCreate(f func() error, err error) func() error {
	if err != nil {
		panic(err)
	}
	return f
}

func MustExists(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}

