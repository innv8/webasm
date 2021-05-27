package wrappers

import (
	"fmt"
	"syscall/js"

	"github.com/innv8/webasm/cmd/wasm/logic"
)

func JsonWrapper() js.Func {

	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return errResult("invalid no of arguments passed")
		}

		// access the dom from Go
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return errResult("Unable to get document object")
		}
		// the element
		jsonOutputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOutputTextArea.Truthy() {
			return errResult("unable to get output text area")
		}

		inputJSON := args[0].String()
		fmt.Printf("input : %s\n", inputJSON)
		pretty, err := logic.PrettyJSON(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error : %s\n", err.Error())
			jsonOutputTextArea.Set("value", errStr)
			return errResult(errStr)
		}
		jsonOutputTextArea.Set("value", pretty)
		return nil
	})
	return jsonFunc
}

func errResult(err string) map[string]interface{} {
	return map[string]interface{}{
		"error": err,
	}
}
