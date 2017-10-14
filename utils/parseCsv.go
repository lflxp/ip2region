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
	"github.com/lflxp/cidr"
	"time"
	"strconv"
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

type CityBlocks struct {
	Network 			string
	Geoname_id 			string
	Registered_country_geoname_id 	string
	Represented_country_geoname_id 	string
	Is_anonymous_proxy		string
	Is_satellite_provider 		string
	Postal_code 			string
	Latitude			string
	Longitude			string
	Accuracy_radius			string
}

type AsnBlocks struct {
	Network 			string
	Autonomous_system_number 	string
	Autonomous_system_organization 	string

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

func GetAsnBlocks(path string) (*[]AsnBlocks,*map[string]int) {
	fmt.Println("开始读取AsnBlocks文件"+path)
	index := map[string]int{}
	num := 0
	data := []AsnBlocks{}

	//locations := map[string]string{}

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
			return nil,nil
		}
		if record[0] != "network" {
			//fmt.Println(record)
			tmp := AsnBlocks{
				Network:record[0],
				Autonomous_system_number:record[1],
				Autonomous_system_organization:record[2],
			}
			if record[1] != "" {
				//locations[record[5]] = record[7]
				index[record[0]] = num
				data = append(data,tmp)
				num += 1
			}
		}
	}
	fmt.Println("读取AsnBlocks文件完毕"+path)
	return &data,&index
}

func GetCityLocations(path string) (*[]CityLocations,*map[string]int) {
	fmt.Println("开始读取CityLocations文件"+path)
	index := map[string]int{}
	num := 0
	data := []CityLocations{}

	//locations := map[string]string{}

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
			return nil,nil
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
				//locations[record[5]] = record[7]
				index[record[0]] = num
				data = append(data,tmp)
				num += 1
			}
		}
	}
	fmt.Println("读取CityLocations完毕")
	return &data,&index
}

func GetCityBlocks(path string) *[]CityBlocks {
	data := []CityBlocks{}
	fmt.Println("开始读取CityBlocks文件"+path)
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//读取数据到内存
	//踢重 国家 城市 省份
	num :=0
	for {
		record,err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:",err)
			return nil
		}
		if record[0] != "network" {
			//fmt.Println(record)
			tmp := CityBlocks{
				Network:record[0],
				Geoname_id:record[1],
				Registered_country_geoname_id:record[2],
				Represented_country_geoname_id:record[3],
				Is_anonymous_proxy:record[4],
				Is_satellite_provider:record[5],
				Postal_code:record[6],
				Latitude:record[7],
				Longitude:record[6],
				Accuracy_radius:record[9],
			}
			if record[5] != "" {
				data = append(data,tmp)
			}
		}
		num += 1
	}
	fmt.Println(fmt.Sprintf("CityBlocks 读取完毕"))
	return &data
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

func ReadBlocks(path string) []string {
	data := []string{}
	//获取城市ip端信息
	cityBlocks := GetCityBlocks(path+"/GeoLite2-City-Blocks-IPv4.csv")
	//获取城市地域信息和索引
	cityLocations,index := GetCityLocations(path+"/GeoLite2-City-Locations-zh-CN.csv")

	asnBlocks,bindex := GetAsnBlocks(path+"/GeoLite2-ASN-Blocks-IPv4.csv")

	fmt.Println("开始解析文件致ip2region")
	for _,x := range *cityBlocks {
		tmp := cidr.NewCidr(x.Network).GetCidrIpRange()
		//获取geoname_id 查询城市地域信息
		ind := (*index)[x.Geoname_id]
		loc := (*cityLocations)[ind]
		//获取asn
		aind := (*bindex)[x.Network]
		asn := (*asnBlocks)[aind]
		if loc.S1Name == "" {
			loc.S1Name = "0"
		} else if loc.CityName == "" {
			loc.CityName = "0"
		} else if asn.Autonomous_system_organization == "" {
			asn.Autonomous_system_organization = "0"
		}
		data = append(data,fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s,(%s,%s),%s",tmp.Min,tmp.Max,loc.ContinentName,loc.CountryName,loc.S1Name,loc.CityName,asn.Autonomous_system_organization,x.Latitude,x.Longitude,x.Accuracy_radius))
	}
	fmt.Println("分析完毕")
	return data
}

func WriteFile(path string,info []string) {
	fmt.Println("开始写入文件...")
	file,err := os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _,x := range info {
		file.WriteString(x+"\n")
	}
	file.Close()
	fmt.Println("写入完毕")
}

/**
 * 数组去重 去空
 */
func removeDuplicatesAndEmpty(a []string) *[]string {
    	tmp := map[string]string{}
	for _,x := range a {
		if x != "" {
			tmp[x] = x
		}
	}
	rs := []string{}
	for k,_ := range tmp {
		rs = append(rs,k)
	}
	sort.Strings(rs)
	return &rs
}

//字段排序
func mapPx(data map[string][]string) map[string]string {
	result := []string{}
	for k,_ := range data {
		result = append(result,k)
	}
	sort.Strings(result)

	tmp := map[string]string{}
	for _,x := range result {
		//w := md5.New()
		//io.WriteString(w, x)   //将str写入到w中
		//tmp[x] = fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
		tmp[x] = strconv.Itoa(time.Now().Nanosecond())
	}
	return tmp
}

//字段排序
func mapPx2(data map[string]string) map[string]string {
	result := []string{}
	for k,_ := range data {
		result = append(result,k)
	}
	sort.Strings(result)

	tmp := map[string]string{}
	for _,x := range result {
		//w := md5.New()
		//io.WriteString(w, x)   //将str写入到w中
		//tmp[x] = fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
		tmp[x] = strconv.Itoa(time.Now().Nanosecond())
	}
	return tmp
}


func cleanmap(info *map[string][]string) {
	for k,x := range *info {
		(*info)[k] = *removeDuplicatesAndEmpty(x)
	}
}

//from GeoLite2-City-Locations-zh-CN
func ReadRegion(path string,writePath string) {
	//获取城市地域信息和索引
	cityLocations,_ := GetCityLocations(path+"/GeoLite2-City-Locations-zh-CN.csv")
	//大陆板块
	D1 := map[string][]string{}
	//国家
	D2 := map[string][]string{}
	//省
	D3 := map[string][]string{}
	//城市
	D4 := map[string]string{}

	//过滤数据
	for _,x := range *cityLocations {
		D1[x.ContinentName] = append(D1[x.ContinentName],x.CountryName)

		D2[x.CountryName] = append(D2[x.CountryName],x.S1Name)

		D3[x.S1Name] = append(D3[x.S1Name],x.CityName)
		D4[x.CityName] = x.CityName
	}

	/**
	清洗数据
	 */
	cleanmap(&D1)
	cleanmap(&D2)
	cleanmap(&D3)

	//最终乱序字典
	//lresult := map[string]string{}

	fmt.Println("开始写入文件...")
	fd,_:=os.OpenFile(writePath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	defer fd.Close()

	DD1 := mapPx(D1)
	for k,x := range DD1 {
		//k 大洲名称
		//x 大洲编号
		//fmt.Println("大洲 ",x,0,k,1,0)
		//lresult[x] = fmt.Sprintf("%s,%d,%s,%d,%d",x,0,k,1,0)
		fd.WriteString(fmt.Sprintf("%s,%s,%s,%d,%d\n",x,0,k,1,0))
		//获取国家编号
		DD2 := mapPx(D2)

		//遍历国家名
		for _,guojia := range D1[k] {
			//fmt.Println("国家 ",DD2[guojia],x,guojia,2,0)
			//lresult[DD2[guojia]] = fmt.Sprintf("%s,%d,%s,%d,%d",DD2[guojia],x,guojia,2,0)
			fd.WriteString(fmt.Sprintf("%s,%s,%s,%d,%d\n",DD2[guojia],x,guojia,2,0))
			//获取省编号
			DD3 := mapPx(D3)

			//遍历省
			for _,sheng := range D2[guojia] {
				//fmt.Println("省 ",DD3[sheng],DD2[guojia],sheng,3,0)
				//lresult[DD3[sheng]] = fmt.Sprintf("%s,%d,%s,%d,%d",DD3[sheng],DD2[guojia],sheng,3,0)
				fd.WriteString(fmt.Sprintf("%s,%s,%s,%d,%d\n",DD3[sheng],DD2[guojia],sheng,3,0))
				//获取城市编号
				DD4 := mapPx2(D4)
				//遍历城市
				for _,city := range D3[sheng] {
					//fmt.Println("城市 ",DD4[city],DD3[sheng],city,4,0)
					//lresult[DD4[city]] = fmt.Sprintf("%s,%d,%s,%d,%d",DD4[city],DD3[sheng],city,4,0)
					fd.WriteString(fmt.Sprintf("%s,%s,%s,%d,%d\n",DD4[city],DD3[sheng],city,4,0))
				}
			}
		}
	}

	fmt.Println("写入完毕")
}
