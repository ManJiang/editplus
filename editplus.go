package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var user string
	fmt.Printf("This is a Editplus editor Keygen application.\nAuthor: MorePHP TEAM\n")
	fmt.Print("Please input Name: ")
	fmt.Scanf("%s\n", &user)
	fmt.Println("Your Key:", keygen(user))
}

func keygen(username string) string {
	list := []int{
		0, 49345, 49537, 320, 49921, 960, 640, 49729, 50689, 1728,
		1920, 51009, 1280, 50625, 50305, 1088, 52225, 3264, 3456, 52545, 3840,
		53185, 52865, 3648, 2560, 51905, 52097, 2880, 51457, 2496, 2176, 51265,
		55297, 6336, 6528, 55617, 6912, 56257, 55937, 6720, 7680, 57025, 57217,
		8000, 56577, 7616, 7296, 56385, 5120, 54465, 54657, 5440, 55041, 6080,
		5760, 54849, 53761, 4800, 4992, 54081, 4352, 53697, 53377, 4160, 61441,
		12480, 12672, 61761, 13056, 62401, 62081, 12864, 13824, 63169, 63361,
		14144, 62721, 13760, 13440, 62529, 15360, 64705, 64897, 15680, 65281,
		16320, 16000, 65089, 64001, 15040, 15232, 64321, 14592, 63937, 63617,
		14400, 10240, 59585, 59777, 10560, 60161, 11200, 10880, 59969, 60929,
		11968, 12160, 61249, 11520, 60865, 60545, 11328, 58369, 9408, 9600,
		58689, 9984, 59329, 59009, 9792, 8704, 58049, 58241, 9024, 57601, 8640,
		8320, 57409, 40961, 24768, 24960, 41281, 25344, 41921, 41601, 25152,
		26112, 42689, 42881, 26432, 42241, 26048, 25728, 42049, 27648, 44225,
		44417, 27968, 44801, 28608, 28288, 44609, 43521, 27328, 27520, 43841,
		26880, 43457, 43137, 26688, 30720, 47297, 47489, 31040, 47873, 31680,
		31360, 47681, 48641, 32448, 32640, 48961, 32000, 48577, 48257, 31808,
		46081, 29888, 30080, 46401, 30464, 47041, 46721, 30272, 29184, 45761,
		45953, 29504, 45313, 29120, 28800, 45121, 20480, 37057, 37249, 20800,
		37633, 21440, 21120, 37441, 38401, 22208, 22400, 38721, 21760, 38337,
		38017, 21568, 39937, 23744, 23936, 40257, 24320, 40897, 40577, 24128,
		23040, 39617, 39809, 23360, 39169, 22976, 22656, 38977, 34817, 18624,
		18816, 35137, 19200, 35777, 35457, 19008, 19968, 36545, 36737, 20288,
		36097, 19904, 19584, 35905, 17408, 33985, 34177, 17728, 34561, 18368,
		18048, 34369, 33281, 17088, 17280, 33601, 16640, 33217, 32897, 16448,
	}

	hexchars := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
	}
	regcode := [29]byte{}
	// regcode := [29]byte{'2', 'B', '6', '5', '0', '-', '1', '3', '5', 'E', '0', '-', '6', '0', '4', '2', 'D', '-', 'B', '9', '3', 'B', 'C', '-', '0', '5', '7', '5', '0'}
	// username := "MorePHP TEAM"

	for i := 0; i < 29; i++ {
		rand.Seed(time.Now().UnixNano())
		if i == 5 || i == 11 || i == 17 || i == 23 {
			regcode[i] = '-'
			continue

		}
		regcode[i] = hexchars[rand.Intn(16)]
	}

	total := 1
	for _, v := range []byte(username) {
		s := fmt.Sprintf("%d", v)
		k, _ := strconv.Atoi(s)
		total += k
	}

	temp := ((total+23)/6 + 3) * 7 % 16
	regcode[6] = hexchars[temp&0xF]
	// fmt.Println(string(regcode[6]))

	temp = (3*total + 39) / 8 % 16
	regcode[9] = hexchars[temp&0xf]

	// fmt.Println(string(regcode[9]))

	temp = (3*total + 19) / 9 % 16
	regcode[7] = hexchars[temp&0xf]
	// fmt.Println(string(regcode[7]))

	temp = (total + 10) / 3 * 8 % 16
	regcode[10] = hexchars[temp&0xf]
	// fmt.Println(string(regcode[10]))

	temp = ((9*total+10)/3 + 36) % 16
	regcode[4] = hexchars[temp&0xf]
	// fmt.Println(string(regcode[4]))

	temp = ((5*total + 11) / 5) % 16
	regcode[8] = hexchars[temp&0xf]
	// fmt.Println(string(regcode[8]))

	result := 0
	for _, v := range []byte(username) {
		s := fmt.Sprintf("%d", v)
		key, _ := strconv.Atoi(s)
		result = ((result >> 8) & 0xFF) ^ list[key^(result&0xFF)]
	}
	// fmt.Println(result)
	str := strings.ToUpper(fmt.Sprintf("%x", result))

	for k, v := range []byte(str) {
		if k == 0 {
			regcode[2] = v
		}
		if k == 1 {
			regcode[3] = v
		}
	}

	// fmt.Println(string(regcode[2]))
	// fmt.Println(string(regcode[3]))

	result = 0
	for i := 2; i < len(regcode); i++ {
		s := fmt.Sprintf("%d", regcode[i])
		key, _ := strconv.Atoi(s)
		result = ((result >> 8) & 0xFF) ^ list[key^(result&0xFF)]
	}
	str = strings.ToUpper(fmt.Sprintf("%x", result))
	for k, v := range []byte(str) {
		if k == 0 {
			regcode[0] = v
		}
		if k == 1 {
			regcode[1] = v
		}
	}
	// fmt.Println(string(regcode[0]))
	// fmt.Println(string(regcode[1]))
	return (strings.ToUpper(string(regcode[:29])))
}
