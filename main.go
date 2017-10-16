package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lflxp/ip2region/utils"
	"flag"
	"net/http"
)

var Data *[]utils.Origin
var Locations *map[string]utils.CityLocations
var Asn *[]utils.AsnBlocks
var path = flag.String("path", "./data", "GeoIP2 文件目录")

func init() {
	flag.Parse()
	Data, Locations, Asn = utils.NewOrigin(*path)
}

func Test(this *gin.Context) {
	this.String(http.StatusOK, "ok")
}

func Check(this *gin.Context) {
	getip := this.Param("ip")
	json := utils.ParseIp(Data, Locations, Asn, getip)
	this.JSON(http.StatusOK, gin.H{
		"time": json.Time,
		"ip":json.Ip,
		"GeoIP":gin.H{
			"Locations":gin.H{
				"Geoname_id":json.Locations.GeonameId,
				"LocaleCode":json.Locations.LocaleCode,
				"ContinentCode":json.Locations.ContinentCode,
				"ContinentName":json.Locations.ContinentName,
				"CountryIsoCode":json.Locations.CountryIsoCode,
				"CountryName":json.Locations.CountryName,
				"S1IsoCode":json.Locations.S1IsoCode,
				"S1Name":json.Locations.S1Name,
				"S2IsoCode":json.Locations.S2IsoCode,
				"S2Name":json.Locations.S2Name,
				"CityName":json.Locations.CityName,
				"MetroCode":json.Locations.MetroCode,
				"TimeZone":json.Locations.TimeZone,
			},
			"Blocks":gin.H{
				"Start":json.Blocks.Start,
				"End":json.Blocks.End,
				"FirstIp":json.Blocks.FirstIp,
				"EndIp":json.Blocks.EndIp,
				"Network":json.Blocks.Network,
				"Geoname_id":json.Blocks.Geoname_id,
				"Registered_country_geoname_id":json.Blocks.Registered_country_geoname_id,
				"Represented_country_geoname_id":json.Blocks.Represented_country_geoname_id,
				"Is_anonymous_proxy":json.Blocks.Is_anonymous_proxy,
				"Is_satellite_provider":json.Blocks.Is_satellite_provider,
				"Postal_code":json.Blocks.Postal_code,
				"Latitude":json.Blocks.Latitude,
				"Longitude":json.Blocks.Longitude,
				"Accuracy_radius":json.Blocks.Accuracy_radius,
			},
			"Asn":gin.H{
				"Network":json.Asn.Network,
				"Autonomous_system_number":json.Asn.Autonomous_system_number,
				"Autonomous_system_organization":json.Asn.Autonomous_system_organization,
			},
		},
		"status":json.Status,
	})
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
	r.GET("/check/:ip", Check)
	r.Run(":8080")
}