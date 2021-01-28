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
	"path/filepath"
	"strconv"
)

type Resp struct {
	Href string
	Error string
}

func ParseFile(rs *model.RowsStats, uuid string, sellerId int) ([]model.Offer, error){
	wb, err := xlsx.OpenFile(fmt.Sprintf("xlsxFiles/%s.xlsx", uuid))
	if err != nil {
		return nil, err
	}
	sh, ok := wb.Sheet[wb.Sheets[0].Name]
	if !ok {
		err = errors.New("sheet does not exist")
		return nil, err
	}
	var oo []model.Offer

	for i := 1; i < sh.MaxRow; i++ {
		o := model.Offer{
			SellerId: sellerId,
		}
		errExist := false
		for i, c := range sh.Rows[i].Cells {
			switch i {
			case 0:
				v, err := strconv.Atoi(c.Value)
				if err != nil {
					rs.ErrorRows += 1
					errExist = true
				}
				o.OfferId = v
			case 1:
				if c.Value == "" {
					rs.ErrorRows += 1
					errExist = true
				}
				o.Name = c.Value
			case 2:
				v, err := strconv.Atoi(c.Value)
				if err != nil || v < 0 {
					rs.ErrorRows += 1
					errExist = true
				}
				o.Price = v
			case 3:
				v, err := strconv.Atoi(c.Value)
				if err != nil || v < 0 {
					rs.ErrorRows += 1
					errExist = true
				}
				o.Quantity = v
			case 4:
				if c.Value == "true" {
					o.Available = true
				} else if c.Value == "false" {
					o.Available = false
				} else {
					rs.ErrorRows += 1
					errExist = true
				}
			}
		}
		if errExist {
			continue
		}
		oo = append(oo, o)
	}

	return oo, nil
}

func DownloadFile(url string, name string) error {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(fmt.Sprintf("xlsxFiles/%s.xlsx", name))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetURLForDownloading(path string, token string) (string, error){
	req, err := http.NewRequest("GET", fmt.Sprintf("https://cloud-api.yandex.net/v1/disk/resources/download?path=%s", path), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", token)
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
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}
	if r.Error == "UnauthorizedError" {
		return "", errors.New("UnauthorizedError")
	}
	return r.Href, nil
}
