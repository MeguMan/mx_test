package xlsxDecoder

import (
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/tealeg/xlsx"
	"strconv"
)

func ParseFile(path string) []model.Offer{
	wb, err := xlsx.OpenFile("xlsx_files/samplefile.xlsx")
	if err != nil {
		panic(err)
	}
	sheetName := "SampleList"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return nil
	}

	var oo []model.Offer

	for i := 1; i < sh.MaxRow; i++ {
		o := model.Offer{}
		for i, v := range sh.Rows[i].Cells {
			switch i {
			case 0:
				o.OfferId, _ = strconv.Atoi(v.Value)
			case 1:
				o.Name = v.Value
			case 2:
				o.Price, _ = strconv.Atoi(v.Value)
			case 3:
				o.Quantity, _ = strconv.Atoi(v.Value)
			case 4:
				o.Available = false
				if v.Value == "true" {
					o.Available = true
				}
			}
		}
		oo = append(oo, o)
	}

	return oo
}
