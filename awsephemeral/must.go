package awsephemeral

type Chain func() error

func MustCreate(f Chain, err error) Chain {
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

