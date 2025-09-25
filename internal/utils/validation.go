package utils

import (
	"fmt"
	"os"
)

func Must(err error, context string) {
	if err != nil {
		Fatal("%s: %v", context, err)
	}
}

func Fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
