package address

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type Province struct {
	Code string
	Name string
}

// City 市级区划
type City struct {
	Code         string
	Name         string
	ProvinceCode string
}

// District 区级区划
type District struct {
	Code         string
	Name         string
	CityCode     string
	ProvinceCode string
}

var (
	provinces  = make(map[string]Province)
	cities     = make(map[string]City)
	districts  = make(map[string]District)
	dataLoaded sync.Once
)

func ReadCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// LoadData 加载数据到内存
func LoadData() {
	dataLoaded.Do(func() {
		provinceRecords, err := ReadCSV("data/provinces.csv")
		if err != nil {
			fmt.Println("failed to read provinces.csv:", err)
			return
		}
		for _, record := range provinceRecords {
			province := Province{
				Code: record[0],
				Name: record[1],
			}
			provinces[province.Code] = province
		}

		cityRecords, err := ReadCSV("data/cities.csv")
		if err != nil {
			fmt.Println("failed to read cities.csv:", err)
			return
		}
		for _, record := range cityRecords {
			city := City{
				Code:         record[0],
				Name:         record[1],
				ProvinceCode: record[2],
			}
			cities[city.Code] = city
		}

		districtRecords, err := ReadCSV("data/districts.csv")
		if err != nil {
			fmt.Println("failed to read districts.csv:", err)
			return
		}
		for _, record := range districtRecords {
			district := District{
				Code:         record[0],
				Name:         record[1],
				CityCode:     record[2],
				ProvinceCode: record[3],
			}
			districts[district.Code] = district
		}
	})
}
