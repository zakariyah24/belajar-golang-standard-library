package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Zakariyah", "Firmansyah"})
	_ = writer.Write([]string{"Danang", "Rafirda"})
	_ = writer.Write([]string{"Dendy", "Virganata"})
	writer.Flush()
}
