package ilocation

import (
	"errors"
	"fmt"
	_ "github.com/golang/glog"
)

func getPart(ip string) (low int, high int, n int, err error) {
	var ip1, ip2, ip3, ip4 int
	fmt.Sscanf(ip, "%d.%d.%d.%d", &ip1, &ip2, &ip3, &ip4)

	// if ip1 < 0 || ip1 > 223 {
	//  return 0, 0, 0, errors.New("Invalid ip")
	// }
	if v, ok := partIndex[ip1]; ok {
		low = v[0]
		high = v[1]
	} else {
		return 0, 0, 0, errors.New("Invalid ip")
	}
	// low = partIndex{ip1}{0}
	// high = partIndex{ip1}{1}
	n = ip1<<24 + ip2<<16 + ip3<<8 + ip4
	return low, high, n, nil
}

func GetLocation(ip string) (country, province, city string) {
	l, h, n, err := getPart(ip)

	if err != nil {
		return "", "", ""
	}

	i, j, k := innerGetLocation(n, l, h)
	//fmt.Printf("ip %s/%d, inner get (%d, %d, %d)\n",ip, n, i, j, k)

	country = locations[k]
	province = locations[j]
	city = locations[i]

	//glog.V(5).Infof("get location %s, %s,%s,%s", ip, country, province, city)

	if country == "N/A" {
		country = ""
	}
	if province == "N/A" {
		province = ""
	}
	if city == "N/A" {
		city = ""
	}

	return country, province, city
}

//非法的IP会返回一个不准确的值
func innerGetLocation(ip int, l int, h int) (i, j, k int) {
	//fmt.Printf("to (%d, %d)\n", l, h)
	if l == h {
		i = ipData[l][0] // country
		j = ipData[l][1] // provice
		k = ipData[l][2] // city
		return i, j, k
	}

	m := l + (h-l)/2

	if ipData[m][3] > ip {
		return innerGetLocation(ip, l, m)
	} else if ipData[m][4] < ip {
		return innerGetLocation(ip, m+1, h)
	} else {
		return innerGetLocation(ip, m, m)
	}
}

// GetCityLevel 获取城市属于几线城市
func GetCityLevel(city string) string {
	return levels[city]
}
