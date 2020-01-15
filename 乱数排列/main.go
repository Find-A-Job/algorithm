package main

import (
	"fmt"
	"math/rand"
	"time"
)

type sliceByte []byte

type fRule interface {
	New() fRule
	GetClone() fRule
	GetLenth() int
	AddToEnd(interface{})
	CutFromEnd() interface{}
	Swap(int, int)
	SwapAndDelete(int) interface{}
	Printf(int, bool) string
}

func (sb *sliceByte) New() fRule {
	r := sliceByte(make([]byte, 0, len([]byte(*sb))))
	return &r
}

func (sb *sliceByte) GetClone() fRule {

	cloneBody := make([]byte, len([]byte(*sb)))
	copy(cloneBody, []byte(*sb))
	r := sliceByte(cloneBody)

	return &r
}

func (sb *sliceByte) GetLenth() int {
	return len([]byte(*sb))
}

func (sb *sliceByte) AddToEnd(data interface{}) {
	slby := []byte(*sb)
	*sb = append(slby, data.(byte))

}

func (sb *sliceByte) CutFromEnd() interface{} {
	slby := []byte(*sb)
	retVal := slby[len(slby)-1]
	*sb = sliceByte(slby[:len(slby)-1])

	return retVal
}

func (sb *sliceByte) Swap(i, j int) {
	slby := []byte(*sb)
	slby[j], slby[len(slby)-1] = slby[len(slby)-1], slby[j]
}

func (sb *sliceByte) SwapAndDelete(index int) interface{} {
	slby := []byte(*sb)
	slby[index], slby[len(slby)-1] = slby[len(slby)-1], slby[index]

	return sb.CutFromEnd()
}

func (sb *sliceByte) Printf(col int, willPrint bool) string {
	slby := []byte(*sb)
	pString := "[ "
	for i, d := range slby {
		pString += fmt.Sprintf("%03d ", d)
		if (i+1)%col == 0 {
			pString += fmt.Sprintf("\n  ")
		}
	}
	pString += fmt.Sprintf("\n]\n")

	if willPrint == true {
		fmt.Printf("%v", pString)
	}

	return pString
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randomLen := rand.Intn(1000) + 100000
	orderSlice := make([]byte, randomLen)
	for i := 0; i < randomLen; i += 1 {
		orderSlice[i] = byte(i + 1)
	}
	// fmt.Printf("randomLen:%v, orderSlice:%v\n", randomLen, orderSlice)
	// fr := []byte{1, 2, 3, 4, 5}
	// r := sliceByte(fr)
	// result := messDataSort(&r)

	alienSlice := sliceByte(orderSlice)
	t1Begin := time.Now().UnixNano()
	// time.Sleep(1 * time.Second)
	res1 := messDataSort(&alienSlice)
	t1End := time.Now().UnixNano()
	time.Sleep(1 * time.Second)
	t2Begin := time.Now().UnixNano()
	res2 := otherAlgo(&alienSlice)
	t2End := time.Now().UnixNano()

	fmt.Printf("res1 spend time :%v, endTime:%v, startTime:%v\n", t1End, t1Begin, t1End-t1Begin)
	fmt.Printf("res2 spend time :%v, endTime:%v, startTime:%v\n", t2End, t2Begin, t2End-t2Begin)
	fmt.Printf("%v, %v, %v\n", t2Begin-t1End, len([]byte(*(res1.(*sliceByte)))), len([]byte(*(res2.(*sliceByte)))))
	res1.Printf(10, false)
	res2.Printf(10, false)
	fmt.Printf("zmx\n")
}

func messDataSort(fr fRule) fRule {
	frClone := fr.GetClone()
	frNew := fr.New()
	sliceLen := frClone.GetLenth()
	for sliceLen > 0 {
		//随机一个数
		index := rand.Intn(sliceLen)
		// fmt.Printf("randomNum:%v, ", index)
		//用数组中最后一个数和这个数进行交换，并从数组中删除这个数
		thisData := frClone.SwapAndDelete(int(index))
		// fmt.Printf("left:%v, ", frClone)
		// fmt.Printf("delData%v, ", thisData.(byte))
		//填入新容器
		frNew.AddToEnd(thisData)
		// fmt.Printf("after:%v", frNew)
		//重新计算长度
		sliceLen = frClone.GetLenth()
		// fmt.Printf("\n")
	}

	return frNew
}
