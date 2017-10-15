package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lflxp/ip2region/utils"
	"os"
	"strings"
	"time"
	"net/http"
)

var Region *utils.Ip2Region
var Data *[]utils.Origin
var Locations *[]utils.CityLocations
var Asn *[]utils.AsnBlocks

func init() {
	//db := os.Args[1]
	db := "./data/ip2region.db"
	_, err := os.Stat(db)
	if os.IsNotExist(err) {
		panic("not found db " + db)
	}
	//全局变量的坑 不能 := 否则是创建一个新的指针对象
	Region, err = utils.New(db)
	//defer Region.Close()

	begin := time.Now()
	ip := utils.IpInfo{}

	ip, err = Region.BtreeSearch("8.8.8.8")

	if err != nil {
		fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("\x1b[0;32m%s  %s\x1b[0m", ip.String(), time.Since(begin).String()))
	}

	Data = utils.LoadCityBlocksIpv4("./data/GeoLite2-City-Blocks-IPv4.csv")
	Locations,_ = utils.GetCityLocations("./data/GeoLite2-City-Locations-zh-CN.csv")
	Asn,_ = utils.GetAsnBlocks("./data/GeoLite2-ASN-Blocks-IPv4.csv")
}

func mains() {
	db := os.Args[1]

	_, err := os.Stat(db)
	if os.IsNotExist(err) {
		panic("not found db " + db)
	}

	region, err := utils.New(db)
	defer region.Close()
	fmt.Println(`initializing
+-------------------------------------------------------+
| ip2region test script                                 |
| format 'ip type'                                      |
| type option 'b-tree','binary','memory' default b-tree |
| Type 'quit' to exit program                           |
+-------------------------------------------------------+`)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("ip2reginon >> ")
		data, _, _ := reader.ReadLine()
		begin := time.Now()
		commands := strings.Fields(string(data))
		ip := utils.IpInfo{}
		len := len(commands)
		if len == 0 {
			continue
		}

		if commands[0] == "quit" {
			break
		}

		if !(len > 1) {
			commands = append(commands, "b-tree")
		}
		switch commands[1] {
		case "b-tree":
			ip, err = region.BtreeSearch(commands[0])
		case "binary":
			ip, err = region.BinarySearch(commands[0])
		case "memory":
			ip, err = region.MemorySearch(commands[0])
		default:
			err = errors.New("parameter error")
		}

		if err != nil {

			fmt.Println(fmt.Sprintf("\x1b[0;31m%s\x1b[0m", err.Error()))
		} else {
			fmt.Println(fmt.Sprintf("\x1b[0;32m%s  %s\x1b[0m", ip.String(), time.Since(begin).String()))
		}
	}
}

func Test(this *gin.Context) {
	this.String(http.StatusOK, "ok")
}

//  "/checkip/type/ip"
func CheckIp(this *gin.Context) {
	var err error
	//db := "./data/ip2region.db"
	//_,err:= os.Stat(db)
	//if os.IsNotExist(err){
	//	panic("not found db " + db)
	//}
	//
	//Region, err := utils.New(db)
	//defer Region.Close()
	getip := this.Param("ip")   //data
	types := this.Param("type") //commands[1]

	//fmt.Println(getip,types)
	begin := time.Now()

	ip := utils.IpInfo{}
	len := len(getip)
	if len == 0 {
		this.String(http.StatusNotFound, "nothing")
	}

	if types != "b-tree" || types != "binary" || types != "memory" {
		types = "b-tree"
	}
	switch types {
	case "b-tree":
		ip, err = Region.BtreeSearch(getip)
	case "binary":
		ip, err = Region.BinarySearch(getip)
	case "memory":
		ip, err = Region.MemorySearch(getip)
	default:
		err = errors.New("parameter error")
	}

	if err != nil {
		//fmt.Println(err.Error())
		this.String(http.StatusOK, err.Error())
	} else {
		//fmt.Println( fmt.Sprintf("\x1b[0;32m%s  %s\x1b[0m",ip.String(),time.Since(begin).String()))
		this.String(http.StatusOK, ip.String()+" "+time.Since(begin).String())
	}
}

func Check(this *gin.Context) {
	getip := this.Param("ip")
	begin := time.Now()
	id := utils.BinarySearchCityBlocksIPv4(Data,getip)
	cityBlocks := (*Data)[id]

	var l utils.CityLocations
	var a utils.AsnBlocks
	for _,x := range *Locations {
		if x.GeonameId == cityBlocks.Geoname_id {
			l = x
			for _,y := range *Asn {
				if cityBlocks.Network == y.Network {
					a = y
				}
			}
		}
	}
	this.String(http.StatusOK,fmt.Sprintf("%d %s %s|%s|%s|%s|%s|%s",id,time.Since(begin).String(),l.ContinentName,l.CountryName,l.S1Name,l.S2Name,l.CityName,a.Autonomous_system_organization))
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})
	r.GET("/test", Test)
	r.GET("/checkip/:type/:ip", CheckIp)
	r.GET("/check/:ip", Check)
	r.Run(":8080")
}