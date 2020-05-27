package main

import (
	"compress/gzip"
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

/**
 * Author:  WangDepeng
 * Date :   2020-05-25 
 * Time :   10:26
 * Package: 
 * Mail :   wangdepeng@cmss.chinamobile.com
 * Project: phonedata
 * Description: 

 */

const (
	fileData = "idcard.dat"
	ID_LENGTH = 17
	NUM_0 = "0"
)
var RATIO_ARR = []int32{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var CHECK_CODE_LIST = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

var wg sync.WaitGroup
func main() {
	fmt.Println("hello world !!!")
	f, err := os.Open(fileData)
	fmt.Println(err)
	r := csv.NewReader(f)
	d, _ := r.ReadAll()
	wg.Add(len(d))
	for _, line := range d{
		go getIdCard(line[0])
	}
	wg.Wait()
}


func getIdCard(addrCode string) {
	s := time.Now().AddDate(0,0,-1)
	e := time.Now().AddDate(0, 0, 0)
	fn, _ := os.Create(addrCode+".dat")
	gw := gzip.NewWriter(fn)
	defer func() {
		_ = gw.Close()
		_ = fn.Close()
		wg.Done()
	}()
	for s.Before(e){
		birthDay := s.Format("20060102")
		for i := 0; i < 1000; i++{
			seq := fmt.Sprintf("%03d", i)
			idNo := addrCode + birthDay + seq
			idNo = idNo + getVerifyCode(idNo)
			line := fmt.Sprintf("%s,%x\n", idNo, md5.Sum([]byte(idNo)))
			fmt.Println(line)
			_, _ = gw.Write([]byte(line))
		}
		s = s.AddDate(0,0,1)
	}
}

/**
生成身份证号码的检验位
 */
func getVerifyCode(idNo string) string{
	if len(idNo) != ID_LENGTH{
		return ""
	}
	var sum int32 = 0
	for i, item := range idNo{
		sum += (item-'0') * RATIO_ARR[i]
	}
	return CHECK_CODE_LIST[sum % 11]
}