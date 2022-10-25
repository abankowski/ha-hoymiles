package hoymiles

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"log"
	"strconv"
)

//go:embed status_codes
var srcStatusCodes []byte

var StatusCodes = make(map[int]string)

func init() {

	records, err := csv.NewReader(bytes.NewBuffer(srcStatusCodes)).ReadAll()

	if err != nil {
		return
	}

	for _, r := range records {
		if len(r) > 1 {
			code, err := strconv.Atoi(r[1])
			if err == nil {
				StatusCodes[code] = r[0]
			}
		} else {
			log.Fatal("cannot use row as error code", r)
		}
	}
}
