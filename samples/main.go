package main

import "fmt"

func main() {
	functionList := make(map[string]func())
	functionList["hello"] = hello
	functionList["print"] = printSample
	functionList["scan"] = scanSample
	functionList["type"] = typesSample
	functionList["if"] = ifSample
	functionList["switch"] = switchSample
	functionList["for"] = forSample
	functionList["function"] = funcSample
	functionList["pointer"] = pointerSample
	functionList["array"] = arraySample
	functionList["slice"] = sliceSample
	functionList["map"] = mapSample
	functionList["struct"] = structSample
	functionList["method"] = methodSample
	functionList["interface"] = interfaceSample
	functionList["goRoutine"] = goRoutineSample
	functionList["channel"] = channelSample
	functionList["chanBuffer"] = chanBufferSample
	functionList["chanSync"] = chanSyncSample
	functionList["chanDirection"] = chanDirectionSample
	input := ""
	fmt.Scanln(&input)
	functionList[input]()
}
