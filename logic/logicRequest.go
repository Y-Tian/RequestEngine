package logic

import (
	download "RequestEngine/tool"
)

func MainLogic(args []string) {
	for _, s := range args {
		byteData := download.MakeRequest(s)
		download.Check(download.WriteToFile(s, byteData))
	}
	return
}
