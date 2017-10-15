package utils

import (
	"github.com/lflxp/cidr"
	"fmt"
	"os"
	"encoding/csv"
	"io"
)

const Data = `1.0.0.0/24,2151718,2077456,,0,0,3095,-37.7333,145.1500,1000
1.0.1.0/24,1810821,1814991,,0,0,,26.0614,119.3061,50
1.0.2.0/23,1810821,1814991,,0,0,,26.0614,119.3061,50
1.0.4.0/22,2077456,2077456,,0,0,,-33.4940,143.2104,1000
1.0.8.0/21,1809858,1814991,,0,0,,23.1167,113.2500,50
1.0.16.0/20,1850147,1861060,,0,0,190-0031,35.6850,139.7514,500
1.0.32.0/19,1809858,1814991,,0,0,,23.1167,113.2500,50
1.0.64.0/20,1858311,1861060,,0,0,710-0816,34.5833,133.7667,20
1.0.80.0/21,1858311,1861060,,0,0,710-0816,34.5833,133.7667,20
1.0.88.0/22,1854383,1861060,,0,0,700-0825,34.6617,133.9350,20
1.0.92.0/22,1861060,1861060,,0,0,190-0031,35.6854,139.7531,20
1.0.96.0/24,1849892,1861060,,0,0,682-0821,35.5036,134.2383,100
1.0.97.0/24,1850742,1861060,,0,0,706-0013,34.4833,133.9500,100
1.0.98.0/24,1854774,1861060,,0,0,694-0061,35.1833,132.5000,20
1.0.99.0/24,1861084,1861060,,0,0,693-0002,35.3667,132.7667,20
1.0.100.0/22,1861084,1861060,,0,0,693-0002,35.3667,132.7667,20
1.0.104.0/21,1862415,1861060,,0,0,732-0822,34.3963,132.4594,100
1.0.112.0/23,1849892,1861060,,0,0,682-0821,35.5036,134.2383,100
1.0.114.0/23,1858311,1861060,,0,0,710-0816,34.5833,133.7667,1
1.0.116.0/22,1852225,1861060,,0,0,750-0003,33.9500,130.9500,200
1.0.120.0/22,1862302,1861060,,0,0,747-0804,34.0500,131.5667,50
1.0.124.0/24,1861060,1861060,,0,0,190-0031,35.6854,139.7531,100
1.0.125.0/24,1849892,1861060,,0,0,682-0821,35.5036,134.2383,100
1.0.126.0/23,1849892,1861060,,0,0,682-0821,35.5036,134.2383,100
1.0.128.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.136.0/22,1609350,1605651,,0,0,10600,13.7333,100.4833,5
1.0.140.0/22,1608900,1605651,,0,0,,16.1848,103.3007,200
1.0.144.0/20,1153557,1605651,,0,0,,10.5000,99.1667,500
1.0.160.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.168.0/21,1151464,1605651,,0,0,,8.4509,98.5298,20
1.0.176.0/22,1150626,1605651,,0,0,64120,17.1640,99.8622,500
1.0.180.0/22,8355831,1605651,,0,0,,11.6586,102.5420,500
1.0.184.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.192.0/23,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.194.0/24,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.195.0/24,1605651,1605651,,0,0,,13.7500,100.4667,1000
1.0.196.0/22,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.200.0/22,1609350,1605651,,0,0,10250,13.8583,100.4688,20
1.0.204.0/22,1607737,1605651,,0,0,67000,16.4433,101.1475,200
1.0.208.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.216.0/22,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.220.0/22,7284248,1605651,,0,0,,9.4936,99.9492,1
1.0.224.0/21,1151254,1605651,,0,0,,7.8333,98.3833,1
1.0.232.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.240.0/21,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.248.0/23,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.250.0/24,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.251.0/24,1151254,1605651,,0,0,,7.8333,98.3833,20
1.0.252.0/23,1605651,1605651,,0,0,,13.7500,100.4667,500
1.0.254.0/24,1151254,1605651,,0,0,,7.8333,98.3833,1000
1.0.255.0/24,1605651,1605651,,0,0,,13.7500,100.4667,500
1.1.0.0/24,1810821,1814991,,0,0,,26.0614,119.3061,50`

type Origin struct {
	Start 				int64
	End 				int64
	FirstIp				string
	EndIp				string
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

//二分法查找
//切片s是升序的
//k为待查找的整数
//如果查到有就返回对应角标,
//没有就返回-1
func BinarySearch(s []int, k int) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

//解析数据
func LoadCityBlocksIpv4(path string) *[]Origin {
	result := []Origin{}
	fmt.Println("开始读取CityBlocks文件"+path)
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//读取数据到内存
	//踢重 国家 城市 省份
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if record[0] != "network" {
			tmp := Origin{}
			count := cidr.NewCidr(record[0]).GetCidrIpRange()
			tmp.Start,_ = ip2long(count.Min)
			tmp.End,_ = ip2long(count.Max)
			tmp.FirstIp = count.Min
			tmp.EndIp = count.Max
			tmp.Network = record[0]
			tmp.Geoname_id = record[1]
			tmp.Registered_country_geoname_id = record[2]
			tmp.Represented_country_geoname_id = record[3]
			tmp.Is_anonymous_proxy = record[4]
			tmp.Is_satellite_provider = record[5]
			tmp.Postal_code = record[6]
			tmp.Latitude = record[7]
			tmp.Longitude = record[8]
			tmp.Accuracy_radius = record[9]
			//fmt.Println("start ",record[0],count.Min,count.Max,tmp.Start)
			result = append(result,tmp)
		}
	}
	fmt.Println("读取完毕CityBlocks文件")
	return &result
}

//通过二分法解析ip对应的ip段
func BinarySearchCityBlocksIPv4(data *[]Origin,ip string) int {
	lo,hi := 0,len(*data)-1
	k,err := ip2long(ip)
	//fmt.Println("ip ",ip,k)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	for lo <= hi {
		m := (lo+hi)>>1
		//fmt.Println(m,k,(*data)[m].Start,(*data)[m].End,(*data)[m].FirstIp,(*data)[m].EndIp)
		if (*data)[m].Start < k {
			if (*data)[m].End < k {
				lo = m + 1
			} else if (*data)[m].End > k {
				return m
			}
		} else if (*data)[m].Start > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}