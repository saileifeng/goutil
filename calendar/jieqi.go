package calendar

import (
	"fmt"
	"strconv"
)

var JieQiTable = []string{
	"小寒", "大寒", "立春", "雨水", "惊蛰", "春分",
	"清明", "谷雨", "立夏", "小满", "芒种", "夏至",
	"小暑", "大暑", "立秋", "处暑", "白露", "秋分",
	"寒露", "霜降", "立冬", "小雪", "大雪", "冬至",
}

var JieQiCTable =[2][24]float64{
	{6.11, 20.84, 4.6295, 19.4599, 6.3826, 21.4155, 5.59, 20.888, 6.318, 21.86,
	6.5, 22.2, 7.928, 23.65, 8.35, 23.95, 8.44, 23.822, 9.098, 24.218, 8.218, 23.08, 7.9, 22.6},
	{5.4055, 20.12, 3.87, 18.73, 5.63, 20.646, 4.81, 20.1, 5.52, 21.04, 5.678,
	21.37, 7.108, 22.83, 7.5, 23.13, 7.646, 23.042, 8.318, 23.438, 7.438, 22.36, 7.18, 21.94},
}

var JieQiTableBase = []int{4, 19, 3, 18, 4, 19, 4, 19, 4, 20, 4, 20, 6, 22, 6, 22, 6, 22, 7, 22, 6, 21, 6, 21}
var JieQiTableIdx = "0123415341536789:;<9:=<>:=1>?012@015@015@015AB78CDE8CD=1FD01GH01GH01IH01IJ0KLMN;LMBEOPDQRST0RUH0RVH0RWH0RWM0XYMNZ[MB\\]PT^_ST`_WH`_WH`_WM`_WM`aYMbc[Mde]Sfe]gfh_gih_Wih_WjhaWjka[jkl[jmn]ope]qph_qrh_sth_W"
var JieQiTableOffset = "211122112122112121222211221122122222212222222221222122222232222222222222222233223232223232222222322222112122112121222211222122222222222222222222322222112122112121222111211122122222212221222221221122122222222222222222222223222232222232222222222222112122112121122111211122122122212221222221221122122222222222222221211122112122212221222211222122222232222232222222222222112122112121111111222222112121112121111111222222111121112121111111211122112122112121122111222212111121111121111111111122112122112121122111211122112122212221222221222211111121111121111111222111111121111111111111111122112121112121111111222111111111111111111111111122111121112121111111221122122222212221222221222111011111111111111111111122111121111121111111211122112122112121122211221111011111101111111111111112111121111121111111211122112122112221222211221111011111101111111110111111111121111111111111111122112121112121122111111011111121111111111111111011111111112111111111111011111111111111111111221111011111101110111110111011011111111111111111221111011011101110111110111011011111101111111111211111001011101110111110110011011111101111111111211111001011001010111110110011011111101111111110211111001011001010111100110011011011101110111110211111001011001010011100110011001011101110111110211111001010001010011000100011001011001010111110111111001010001010011000111111111111111111111111100011001011001010111100111111001010001010000000111111000010000010000000100011001011001010011100110011001011001110111110100011001010001010011000110011001011001010111110111100000010000000000000000011001010001010011000111100000000000000000000000011001010001010000000111000000000000000000000000011001010000010000000"


func JieQiDay(year,month int)(firstDayName string,firstDay int,secondDayName string,secondDay int){
	century := year / 100 + 1
	centuryYear := year % 100
	d := 0.2422
	firstDayName = JieQiTable[month*2-2]
	firstDay = int((float64(centuryYear) * d + JieQiCTable[century%10][month*2-2]) - float64((centuryYear-1)/4))
	secondDayName = JieQiTable[month*2-1]
	secondDay = int((float64(centuryYear) * d + JieQiCTable[century%10][month*2-1]) - float64((centuryYear-1)/4))
	if centuryYear % 4 == 0 && month > 2{
		firstDay -= 1
		secondDay -= 1
	}
	//day := (float64(centuryYear) * d + JieQiCTable[century%10][month-1]) - ((centuryYear-1)/4)
	//fmt.Println(century,"世纪",centuryYear,"年代",month,"月 [",firstDayName,"是",firstDay,"号 ] [",secondDayName,"是",secondDay,"号 ]")
	return
}

// y年的第n个节气为几日
// n range [1,24]
func JieQi(year, n int) int {
	n -= 1
	charcodeAt := int(JieQiTableIdx[year-MinYear])
	offset, err := strconv.Atoi(string(JieQiTableOffset[(charcodeAt-48)*24+n]))
	if err != nil {
		fmt.Println("strconv.Atoi error")
	}
	//return JieQiTableBase[n] + JieQiTableOffset.charAt((JieQiTableIdx.charCodeAt(year-MinYear)-48)*24+n)
	return JieQiTableBase[n] + offset
}

// month range [1,12]
func JieQisOfMonth(year, m int) (first, second int) {
	return JieQi(year, (m-1)*2), JieQi(year, m*2-1)
}

func JieQisOfYear(year int) (list [12][2]int) {
	for i := 1; i < 13; i++ {
		list[i-1][0], list[i-1][1] = JieQisOfMonth(year, i)
	}
	return
}
