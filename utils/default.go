package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func VarDump(data ...interface{}) {
	for _, val := range data {
		fmt.Printf("(%v) ", reflect.TypeOf(val))
		switch reflect.ValueOf(val).Kind() {
		case reflect.Slice:
			fallthrough
		case reflect.Map:
			fallthrough
		case reflect.Struct:
			bytes, _ := json.Marshal(val)
			fmt.Println(string(bytes))
		default:
			fmt.Println(val)
		}
	}
}

func RandIntArr(min, max float64, num int) []int {
	max = max + 1
	rand.Seed(time.Now().UnixMicro())
	ret := []int{}
	for i := 0; i <= num; i++ {
		tmp := rand.Float64()*(max-min) + min
		ret = append(ret, int(tmp))
	}
	return ret
}

func Check(arr []int) error {
	n := len(arr)
	for i := 0; i <= n-2; i++ {
		if arr[i] > arr[i+1] {
			return fmt.Errorf("第%d位的%d比第%d位的%d 大", i, arr[i], i+1, arr[i+1])
		}
	}
	return nil
}
