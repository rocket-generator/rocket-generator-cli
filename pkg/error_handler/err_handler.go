package error_handler

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
