package utils

func AssertErr(err error) {
	if err != nil {
		panic(err)
	}
}
