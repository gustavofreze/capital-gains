package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	// maximumScannerTokenSizeBytes is the maximum allowed size for a single scanned token (line).
	// This prevents bufio.Scanner from failing when a JSON line is very large (e.g., minified JSON).
	maximumScannerTokenSizeBytes = 10_485_760

	// defaultScannerBufferCapacityBytes is the initial buffer capacity used by bufio.Scanner.
	// It should be large enough for typical JSON inputs without causing reallocations.
	defaultScannerBufferCapacityBytes = 65_536
)

type DefaultConsole struct {
	inputScanner *bufio.Scanner
	outputWriter *bufio.Writer
}

func NewDefaultConsole() *DefaultConsole {
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Buffer(make([]byte, 0, defaultScannerBufferCapacityBytes), maximumScannerTokenSizeBytes)

	return &DefaultConsole{
		inputScanner: inputScanner,
		outputWriter: bufio.NewWriter(os.Stdout),
	}
}

func (adapter *DefaultConsole) ReadLines() []string {
	lines := make([]string, 0)

	for adapter.inputScanner.Scan() {
		line := adapter.inputScanner.Text()

		if strings.TrimSpace(line) == "" {
			break
		}

		lines = append(lines, line)
	}

	if scanError := adapter.inputScanner.Err(); scanError != nil {
		panic(scanError)
	}

	return lines
}

func (adapter *DefaultConsole) WriteLine(text string) {
	if _, writeError := fmt.Fprintln(adapter.outputWriter, text); writeError != nil {
		panic(writeError)
	}

	if flushError := adapter.outputWriter.Flush(); flushError != nil {
		panic(flushError)
	}
}
