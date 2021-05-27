// function json (input) {
//     jsonoutput.value = formatJSON(input);
// }

function json(input) {
    var result = formatJSON(input)
    if ((result != null) && ('error' in result)) {
        console.log("Go return value", result)
        jsonoutput.value = ""
        alert(result.error)
    }
}
