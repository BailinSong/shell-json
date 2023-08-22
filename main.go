package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		printHelp()
		return
	}

	render := bufio.NewReader(os.Stdin)

LOOP:
	input, err := render.ReadString('\n')
	if err != nil {
		return
	}

	jsonStr := gojsonq.New().FromString(input)
	for index, v := range os.Args {
		if index > 0 {
			value := jsonStr.Find(v)

			switch value.(type) {
			case string:
				fmt.Print(value)
			default:
				nJson, _ := json.Marshal(value)
				fmt.Print(string(nJson))
			}
			jsonStr.Reset()
			fmt.Print("\t")
		}
	}
	fmt.Print("\n")
	goto LOOP
}

func printHelp() {
	fmt.Println(
		`
Usage: Provide JSON input and command-line arguments.
    Example: echo '{"key": "value","key1": {"key2": "value2","key3": "value3"},"key4":["value4","value5","value6","value7"]}' | json <keyPath> [... <keyPath>]
    <keyPath>    json key path
        Example:
            key output value
            key1.key2 output value2
            key4.[0] output value4
            key1 output {"key2": "value2","key3": "value3"}
            key key1.key2 output value	value2`)

}
