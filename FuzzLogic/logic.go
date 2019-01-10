package FuzzLogic

func MainLogic(args []string) {
	for _, filename := range args {
		MakeFuzz(filename)
	}
}
