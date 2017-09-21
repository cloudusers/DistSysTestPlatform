package convert

import (
	"fmt"
	"strconv"
)

func Int322String(i int32) string {
	ss := strconv.Itoa(int(i))
	return ss
}
func Int642String(i int64) string {
	ss := strconv.FormatInt(i, 10)
	return ss
}

func String2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

//X=32,int32
func String2Int64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	return i
}

func String2Int32(s string) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	return int32(i)
}

func String2Float32(s string) float32 {
	i, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Println(err)
	}

	return float32(i)
}
func String2Float64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(err)
	}

	return float64(i)
}
