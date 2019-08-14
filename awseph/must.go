package awseph

func MustExists(b bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return b
}

