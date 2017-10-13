package main

import "github.com/lflxp/ip2region/utils"

func main() {
	//utils.Reader("./data/GeoLite2-City-Locations-en.csv")
	utils.ReadRegion("./data/GeoLite2-City-Locations-en.csv")
}
