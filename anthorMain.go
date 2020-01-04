package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

type Item struct {
	Weight int
	Value  int
	//cmp    interface{}
}

type funcCmp func(i, j int) bool

type valueItem []Item
type WeightItem []Item

func main() {
	rand.Seed(time.Now().UnixNano())

	strCompare()
	return

	CharIssue()
	return

	BackpackIssue()
	return

}

func strCompare() {
	// strA := "ababababca"
	// strB := "abababca"
	// times := KMPAlgo(strA, strB)
	// fmt.Printf("匹配次数:%v\n", times)

	sliceStrA, sliceStrB := randomString(10, 30000, 10000, 9, 5, int('a'), int('d'))

	for i := 0; i < len(sliceStrA); i += 1 {
		start := time.Now().UnixNano()
		ikm := KMPAlgo(sliceStrA[i], sliceStrB[i])
		// icou := strings.Count(sliceStrA[i], sliceStrB[i])
		elapsed := time.Now().UnixNano() - start
		start2 := time.Now().UnixNano()
		icou := strings.Count(sliceStrA[i], sliceStrB[i])
		// ikm := KMPAlgo(sliceStrA[i], sliceStrB[i])
		elapsed2 := time.Now().UnixNano() - start2
		// fmt.Printf("%v\t\t\t%v\t\t%v:%v\t\t%v:%v\n",
		// 	sliceStrA[i], sliceStrB[i], KMPAlgo(sliceStrA[i], sliceStrB[i]), elapsed,
		// 	strings.Count(sliceStrA[i], sliceStrB[i]), elapsed2)
		fmt.Printf("%v:%v\t\t%v:%v\n", ikm, elapsed, icou, elapsed2)
	}
}

//样本个数maxData， 母串范围AMax, AMin，子串范围BMax, BMin， 随机字符范围rcA, rcB(rcA < rcB)
func randomString(maxData, AMax, AMin, BMax, BMin, rcA, rcB int) ([]string, []string) {
	strA := []string{}
	strB := []string{}
	for i := 0; i < maxData; i += 1 {
		am := rand.Intn(AMax-AMin) + AMin
		bm := rand.Intn(BMax-BMin) + BMin
		s1 := ""
		s2 := ""
		for j1 := 0; j1 < am; j1 += 1 {
			// s1 += fmt.Sprintf("%c", rand.Intn(26)+int('a'))
			s1 += fmt.Sprintf("%c", rand.Intn(rcB-rcA)+rcA)
		}
		strA = append(strA, s1)
		for j2 := 0; j2 < bm; j2 += 1 {
			// s2 += fmt.Sprintf("%c", rand.Intn(26)+int('a'))
			s2 += fmt.Sprintf("%c", rand.Intn(rcB-rcA)+rcA)
		}
		strB = append(strB, s2)
	}

	return strA, strB
}

func KMPAlgo(strA string, strB string) int {
	//strA是长串，strB是pattern匹配串
	next := PMTAlgo(strB)
	fmt.Printf("next:%v\n", next)

	totalMatchTime := 0 //匹配次数
	matchNum := 0       //匹配字符数
	for i, j := 0, 0; i < len([]rune(strA)); {
		if string(([]rune(strA))[i]) == string(([]rune(strB))[j]) {
			matchNum += 1
			j += 1
			i += 1

			if matchNum == len([]rune(strB)) {
				//成功匹配字符串
				totalMatchTime += 1
				j = 0
				matchNum = 0
				// fmt.Printf("1\n")
				continue
			}
			if matchNum > len([]rune(strB)) {
				panic(nil)
			}
		} else {
			if matchNum == 0 {
				i += 1
				j = 0
				// fmt.Printf("2\n")
				continue
			} else {
				j = next[matchNum-1]
				matchNum = j
				// fmt.Printf("3\n")
				continue
			}
		}
	}

	return totalMatchTime
}

func PMTAlgo(str string) []int {
	next := []int{}
	for i := 1; i < len([]rune(str)); i += 1 { //匹配成功i个
		ind := 0
		for j := i; j > 1; j -= 1 {
			// fmt.Printf("%v, %v, i:%v, j:%v\n", string(([]rune(str))[:j-1]), string(([]rune(str))[i-j+1:i]), i, j)
			if string(([]rune(str))[:j-1]) == string(([]rune(str))[i-j+1:i]) {
				ind = j - 1
				break
				// next = append(next, j-1)
			}
		}
		next = append(next, ind)
	}

	return next
}

func CharIssue() {
	allStr := []string{
		"abcde",
		"abcd",
		"abc",
		"ab",
		"a",
		"野原新之助",
		"蜡笔小新",
		"新之助",
		"野原",
		"新",
		"野原广志a",
		"野原广a志",
		"野原a广志",
		"野原广志ab",
		"野原广ab志",
		"野原ab广志",
		"广志abc",
		"广abc志",
		"abc广志",
		"广志abcd",
		"广abcd志",
		"abcd广志",
		"新a",
		"a新",
		"广志a",
		"广a志",
		"a广志",
		"志ab",
		"ab志",
		"志abc",
		"abc志",
	}
	for _, s := range allStr {
		fmt.Printf("before:%v,\t\t after:%v\n", s, CoverNickName(s))
	}
	fmt.Printf("\n")

	return

	ch1 := "this is a example for test 1"
	ch2 := "这是测试2的例子"
	ch3 := "这是中英混合example"

	fmt.Printf("ch1:%v, ch2:%v, ch3:%v\n", len(ch1), len(ch2), len(ch3))
	fmt.Printf("ch1:%v, ch2:%v, ch3:%v\n", utf8.RuneCountInString(ch1), utf8.RuneCountInString(ch2), utf8.RuneCountInString(ch3))

	for _, c := range ch1 {
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
	for _, c := range ch2 {
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
	for _, c := range ch3 {
		fmt.Printf("%c", c)
	}
	fmt.Printf("\n")
	fmt.Printf("------------------------------\n")
	for i := 0; i < utf8.RuneCountInString(ch1); i += 1 {
		fmt.Printf("%c", ch1[i])
	}
	fmt.Printf("\n")
	for i := 0; i < utf8.RuneCountInString(ch2); i += 1 {
		fmt.Printf("%c", ch2[i])
	}
	fmt.Printf("\n")
	for i := 0; i < utf8.RuneCountInString(ch3); i += 1 {
		fmt.Printf("%c", ch3[i])
	}
	fmt.Printf("\n")
	fmt.Printf("------------------------------\n")
	//fmt.Printf("%v\n", string(([]rune(ch2))[1:3]))
	fmt.Printf("%v\n", CoverNickName(ch2))
}

func CoverNickName(nick string) string {
	//字符串分析,记录每个字换算成字符的个数
	strLen := []int{}
	for _, r := range nick {
		if unicode.Is(unicode.Han, r) {
			strLen = append(strLen, 2)
		} else {
			strLen = append(strLen, 1)
		}
	}

	//空字符串
	if len(strLen) < 1 {
		panic(nil)
	}

	//至少留一个
	if len(strLen) == 1 {
		return "****" + nick
	}

	cutNum := 0 //应该截取的字个数
	tLen := 0   //截取x个字后，字符的个数，x个字转为y字符数
	//从后往前截，直到剩余一个字或达到最大截取字符数
	for i := len(strLen) - 1; i > 0; i -= 1 {
		tLen += strLen[i]

		if i < 1 || tLen > 4 {
			break
		} else {
			cutNum += 1
		}
	}

	return "****" + string(([]rune(nick))[len(strLen)-cutNum:])
}

func BackpackIssue() {
	//有编号分别为a,b,c,d,e的五件物品，
	//它们的重量分别是2,2,6,5,4，
	//它们的价值分别是6,3,5,4,6，
	//每件物品数量只有一个，现在给你个承重为10的背包，
	//如何让背包里装入的物品具有最大的价值总和

	items := []Item{
		Item{
			Weight: 2,
			Value:  6,
		},
		Item{
			Weight: 2,
			Value:  3,
		},
		Item{
			Weight: 6,
			Value:  5,
		},
		Item{
			Weight: 5,
			Value:  4,
		},
		Item{
			Weight: 4,
			Value:  6,
		},
	}

	items = []Item{}
	for i := 0; i < 50; i += 1 {
		it := Item{
			Weight: rand.Intn(8) + 3,
			Value:  rand.Intn(8) + 3,
		}
		items = append(items, it)
	}

	items2 := make([]Item, len(items))
	copy(items2, items)
	sort.Sort(valueItem(items2))
	valueItem(items2).PrintfMsg()

	capacity := 50
	value := 0
	pick := []Item{}
	for i := 0; i < len(items2); i += 1 {
		if capacity-items2[i].Weight >= 0 {
			pick = append(pick, items2[i])
			value += items2[i].Value
			capacity -= items2[i].Weight
		}
	}

	fmt.Printf("new:, left:%v, value:%v\n", capacity, value)
	valueItem(pick).PrintfMsg()

	fmt.Printf("-------------------------------\n")
	sort.Sort(WeightItem(items2))
	valueItem(items2).PrintfMsg()

	capacity2 := 50
	value2 := 0
	pick2 := []Item{}
	for i := 0; i < len(items2); i += 1 {
		if capacity2-items2[i].Weight >= 0 {
			pick2 = append(pick2, items2[i])
			value2 += items2[i].Value
			capacity2 -= items2[i].Weight
		}
	}

	fmt.Printf("new:, left:%v, value:%v\n", capacity2, value2)
	valueItem(pick2).PrintfMsg()

	p3 := Backpack01(items, 50)
	fmt.Printf("--------new line---------\n\n")
	valueItem(p3).PrintfMsg()

	fmt.Printf("\n\n")
	items22 := []Item{
		Item{
			Weight: 5,
			Value:  1,
		},
		Item{
			Weight: 4,
			Value:  4,
		},
		Item{
			Weight: 1,
			Value:  1,
		},
		Item{
			Weight: 3,
			Value:  5,
		},
		Item{
			Weight: 2,
			Value:  3,
		},
	}
	Backpack01(items22, 10)
	fmt.Printf("\n")
	// p4 := Backpack01(items22, 10)
	// fmt.Printf("\n--------new line---------\n\n")
	// valueItem(p4).PrintfMsg()

	items33 := []Item{
		Item{
			Weight: 5,
			Value:  1,
		},
		Item{
			Weight: 1,
			Value:  1,
		},
		Item{
			Weight: 2,
			Value:  3,
		},
		Item{
			Weight: 4,
			Value:  4,
		},
		Item{
			Weight: 3,
			Value:  5,
		},
	}
	Backpack01(items33, 10)
	fmt.Printf("\n")
	// p5 := Backpack01(items33, 10)
	// fmt.Printf("\n--------new line---------\n\n")
	// valueItem(p5).PrintfMsg()

	items44 := []Item{
		Item{
			Weight: 7,
			Value:  13,
		},
		Item{
			Weight: 5,
			Value:  8,
		},
		Item{
			Weight: 4,
			Value:  20,
		},
		Item{
			Weight: 8,
			Value:  11,
		},
		Item{
			Weight: 12,
			Value:  6,
		},
	}
	Backpack01(items44, 20)
	fmt.Printf("\n")
}

func Backpack01(items []Item, capacity int) []Item {
	//思路:
	//    建立一张表来记录每种情况下的最优值
	//    类似于穷举，但比穷举效率高，
	//    原因是其每一次穷举的情况会被记录在表中，后续的穷举会使用到前一次的记录
	//
	//    同一份数据，以不同的顺序排列，其构建的表也会不一样，但其二维表的最后一行总是相同，也就是说顺序不会影响其结果但会影响表中部分内容

	valueItem(items).PrintfMsg()
	bpTable := make([]([]int), len(items)+1)
	for i := 0; i < len(bpTable); i += 1 {
		bpTable[i] = make([]int, capacity+1)
	}

	//填表
	for i := 0; i < len(items)+1; i += 1 {
		for j := 0; j < capacity+1; j += 1 {
			if i == 0 || j == 0 {
				(bpTable[i])[j] = 0
				continue
			} else {
				//假设先装入第i个物品，在容量为j的情况下，
				oldVal := (bpTable[i-1])[j]
				newVal := 0
				if j-items[i-1].Weight >= 0 {
					newVal = items[i-1].Value + (bpTable[i-1])[j-items[i-1].Weight]
				}

				if newVal > oldVal {
					(bpTable[i])[j] = newVal
				} else {
					(bpTable[i])[j] = oldVal
				}
			}

		}
	}

	Print2Slice(bpTable)

	retVal := []Item{}
	//fmt.Printf("%v, %v\n", capacity, len(bpTable))
	lc := capacity
	for i := len(bpTable) - 1; i > 0; i -= 1 {
		if lc <= 0 {
			break
		}
		if bpTable[i][lc] == bpTable[i-1][lc] {
			continue
		} else {
			retVal = append(retVal, items[i-1])
			lc -= items[i-1].Weight
			continue
		}
	}

	return retVal
}

//
func Print2Slice(s []([]int)) {
	fmt.Printf("-----------填表-----------\n")
	for _, v1 := range s {
		for _, v2 := range v1 {
			fmt.Printf("%03d ", v2)
		}
		fmt.Printf("\n")
	}

}

//
func (vi valueItem) PrintfMsg() {
	for _, v := range vi {
		fmt.Printf("value:%v, weight:%v\n", v.Value, v.Weight)
	}
}

//
func (vi valueItem) Len() int {
	return len(vi)
}
func (vi valueItem) Less(i, j int) bool {
	if vi[i].Value == vi[j].Value {
		return vi[i].Weight < vi[j].Weight

	} else {
		return vi[i].Value > vi[j].Value
	}

	//return (vi.(funcCmp))(i, j)
}
func (vi valueItem) Swap(i, j int) {
	vi[i].Value, vi[j].Value = vi[j].Value, vi[i].Value
	vi[i].Weight, vi[j].Weight = vi[j].Weight, vi[i].Weight
}

//
func (wi WeightItem) Len() int {
	return len(wi)
}
func (wi WeightItem) Less(i, j int) bool {
	if (wi[i].Value / wi[i].Weight) == (wi[j].Value / wi[j].Weight) {
		return wi[i].Weight < wi[j].Weight

	} else {
		return (wi[i].Value / wi[i].Weight) > (wi[j].Value / wi[j].Weight)
	}

	//return (vi.(funcCmp))(i, j)
}
func (wi WeightItem) Swap(i, j int) {
	wi[i].Value, wi[j].Value = wi[j].Value, wi[i].Value
	wi[i].Weight, wi[j].Weight = wi[j].Weight, wi[i].Weight
}
