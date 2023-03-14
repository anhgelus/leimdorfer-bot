package utils

func HandlePanic(err error) {
	if err != nil {
		panic(err)
	}
}
