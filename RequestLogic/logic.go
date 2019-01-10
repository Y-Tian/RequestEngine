package RequestLogic

func MainLogic(args []string) {
	for _, s := range args {
		byteData := MakeRequest(s)
		Check(WriteToFile(s, byteData))
	}
	return
}
