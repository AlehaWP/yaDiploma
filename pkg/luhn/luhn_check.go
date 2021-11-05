package luhn

import (
	"strconv"

	"github.com/AlehaWP/yaDiploma.git/pkg/logger"
)

func CheckInteger(i int) bool {
	str := strconv.Itoa(i)
	return CheckString(str)
}

func CheckString(s string) bool {
	var rArr []int
	for _, v := range s {
		// Десятичные значения
		res, err := strconv.Atoi(string(v))
		if err != nil {
			logger.Info("Luhn", "Ошибка конвертации строки", err)
		}
		rArr = append(rArr, res)
	}
	return Check(rArr)
}

func Check(arr []int) bool {
	l := len(arr)
	controlSum := arr[l-1]
	cArr := arr[:l-1]
	var cd int
	for i, v := range cArr {
		if i%2 == 0 {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}
		cd += v
	}

	return (cd+controlSum)%10 == 0
}
