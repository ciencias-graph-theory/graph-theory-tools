package writer

import (
	"os"
)

func Write(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(err)
	}
}
