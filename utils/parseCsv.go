package utils

/**
将geoip的数据解析成ip2region数据
 */
import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	//"strconv"
)

type CityLocations struct {
	GeonameId	string
	LocaleCode	string
	ContinentCode	string
	ContinentName	string
	CountryIsoCode	string
	CountryName	string
	S1IsoCode	string
	S1Name		string
	S2IsoCode	string
	S2Name		string
	CityName	string
	MetroCode	string
	TimeZone 	string
}

func Reader(path string) {
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record,err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:",err)
			return
		}
		fmt.Println(record)
	}
}

//func DeleteMore(a *[]CityLocations,b []string) []string {
//	//第一层
//	//country[x] = strconv.Itoa(n)
//	fmt.Println(n+1,0,x,1,0)
//	S1 := []string{}
//	//去重
//	tmp1 := map[string]string{}
//	for _,xx := range data {
//		if x == xx.CountryName {
//			tmp1[xx.S1Name] = xx.S1Name
//		}
//	}
//	//排序
//	for k,_ := range tmp1 {
//		S1 = append(S1,k)
//	}
//	//第二层
//	sort.Strings(S1)
//	for _,xxx := range S1 {
//		fmt.Println(len(Countrys) + n + 1, n + 1, xxx, 2, 0)
//	}
//}

//from GeoLite2-City-Locations-zh-CN
func ReadRegion(path string) {
	data := []CityLocations{}

	country := map[string]string{}

	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//读取数据到内存
	//踢重 国家 城市 省份
	for {

		record,err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:",err)
			return
		}
		if record[0] != "geoname_id" {
			//fmt.Println(record)
			tmp := CityLocations{
				GeonameId:record[0],
				LocaleCode:record[1],
				ContinentCode:record[2],
				ContinentName:record[3],
				CountryIsoCode:record[4],
				CountryName:record[5],
				S1IsoCode:record[6],
				S1Name:record[7],
				S2IsoCode:record[6],
				S2Name:record[9],
				CityName:record[10],
				MetroCode:record[11],
				TimeZone:record[12],
			}
			if record[5] != "" {
				country[record[5]] = record[7]
				data = append(data,tmp)
			}
		}
	}

	//定位国家 按字母进行排序 确定顶层顺序
	Countrys := []string{}
	for k,_ := range country {
		Countrys = append(Countrys,k)
	}
	sort.Strings(Countrys)
	for n,x := range Countrys {
		//第一层
		//country[x] = strconv.Itoa(n)
		fmt.Println(n+1,0,x,1,0)
		S1 := []string{}
		//去重
		tmp1 := map[string]string{}
		for _,xx := range data {
			if x == xx.CountryName {
				tmp1[xx.S1Name] = xx.S1Name
			}
		}
		//排序
		for k,_ := range tmp1 {
			S1 = append(S1,k)
		}
		//第二层
		sort.Strings(S1)
		for _,xxx := range S1 {
			fmt.Println(len(Countrys)+n+1,n+1,xxx,2,0)
			//第三层
			S2 := []string{}
			//去重
			tmp2 := map[string]string{}
			for _,xx := range data {
				if xxx == xx.S1Name {
					if xx.S2Name != "" {
						tmp2[xx.S2Name] = xx.S2Name
					}
				}
			}
			//排序
			for k,_ := range tmp2 {
				S2 = append(S2,k)
			}
			sort.Strings(S2)
			for _,xxxx := range S2 {
				fmt.Println(len(Countrys)+len(S1)+n+1,n+2,xxxx,3,0)
				//第四层
				City := []string{}
				//去重
				tmp3 := map[string]string{}
				for _,xx := range data {
					if xxx == xx.S2Name {
						if xx.CityName != "" {
							tmp2[xx.CityName] = xx.CityName
						}
					}
				}
				//排序
				for k,_ := range tmp3 {
					City = append(City,k)
				}
				sort.Strings(City)
				for _,xxxx := range City {
					fmt.Println(len(Countrys)+len(S2)+n+1,n+2,xxxx,4,0)
				}
			}
		}
	}
}
