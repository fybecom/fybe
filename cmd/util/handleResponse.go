package util

import (
	"encoding/json"

	cliCmd "fybe.com/cli/fybe/cmd"
	"fybe.com/cli/fybe/outputFormatter"
)

func HandleResponse(
	responseJson []byte,
	configFormatter outputFormatter.FormatterConfig,
) {
	var responseData []interface{}
	json.Unmarshal(responseJson, &responseData)

	lines := outputFormatter.Formatter(cliCmd.OutputFormat).
		Format(responseData, configFormatter)

	configPrinter := outputFormatter.PrinterConfig{
		Delimiter: cliCmd.OutputFormatDetails}

	outputFormatter.Printer(cliCmd.OutputFormat).Print(lines, configPrinter)
}
