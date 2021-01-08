package xlsxDecoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MeguMan/mx_test/internal/app/model"
	"github.com/tealeg/xlsx"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Resp struct {
	Href string
}

func ParseFile(path string, rs *model.RowsStats) ([]model.Offer, error){
	url, err := getURLForDownloading(path)
	if err != nil {
		return nil, err
	}
	err = downloadFile(url)
	if err != nil {
		return nil, err
	}
	wb, err := xlsx.OpenFile("downloadedFile.xlsx")
	if err != nil {
		return nil, err
	}
	sheetName := "SampleList"
	sh, ok := wb.Sheet[sheetName]
	if !ok {
		err = errors.New("sheet does not exist")
		return nil, err
	}

	var oo []model.Offer

	for i := 1; i < sh.MaxRow; i++ {
		o := model.Offer{}
		for i, c := range sh.Rows[i].Cells {
			switch i {
			case 0:
				v, err := strconv.Atoi(c.Value)
				if err != nil {
					rs.ErrorRows += 1
				}
				o.OfferId = v
			case 1:
				if c.Value == "" {
					rs.ErrorRows += 1
				}
				o.Name = c.Value
			case 2:
				v, err := strconv.Atoi(c.Value)
				if err != nil {
					rs.ErrorRows += 1
				}
				o.Price = v
			case 3:
				v, err := strconv.Atoi(c.Value)
				if err != nil {
					rs.ErrorRows += 1
				}
				o.Quantity = v
			case 4:
				if c.Value == "true" {
					o.Available = true
				} else if c.Value == "false" {
					o.Available = false
				} else {
					rs.ErrorRows += 1
				}
			}
		}
		oo = append(oo, o)
	}

	return oo, nil
}

func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create("downloadedFile.xlsx")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func getURLForDownloading(path string) (string, error){
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources/download?path=%s", path), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "OAuth AgAAAAA1z4O5AADLW7ibSa25TUIVocRFVAYdP1Q")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	r := Resp{}
	err =json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	return r.Href, nil
}
