package utils

import (
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(s, 6))
}

func TestLoadCityBlocksIpv4(t *testing.T) {
	LoadCityBlocksIpv4(Data)
}

func TestBinarySearchCityBlocksIPv4(t *testing.T) {
	ip := "1.0.99.33"
	data := LoadCityBlocksIpv4("../data/GeoLite2-City-Blocks-IPv4.csv")

	tmp := BinarySearchCityBlocksIPv4(data,ip)

	if tmp == -1 {
		t.Error("nothing found")
	} else {
		t.Log("Success ",tmp)
	}
}

func BenchmarkBinarySearchCityBlocksIPv4(b *testing.B) {
	data := LoadCityBlocksIpv4("../data/GeoLite2-City-Blocks-IPv4.csv")

	for i:=0;i<b.N;i++ {
		tmp := BinarySearchCityBlocksIPv4(data,fmt.Sprintf("%d.2.1.3",i+1))
		b.Log(fmt.Sprintf("%d.2.1.3",i),tmp)
	}

}