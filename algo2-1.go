func BackpackIssue() {
	items = []Item{}
	for i := 0; i < 50; i += 1 {
		it := Item{
			Weight: rand.Intn(8) + 3,
			Value:  rand.Intn(8) + 3,
		}
		items = append(items, it)
	}


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

value:1, weight:5
value:4, weight:4
value:1, weight:1
value:5, weight:3
value:3, weight:2
-----------填表-----------
000 000 000 000 000 000 000 000 000 000 000 
000 000 000 000 000 001 001 001 001 001 001 
000 000 000 000 004 004 004 004 004 005 005 
000 001 001 001 004 005 005 005 005 005 006 
000 001 001 005 006 006 006 009 010 010 010 
000 001 003 005 006 008 009 009 010 012 013 

value:1, weight:5
value:1, weight:1
value:3, weight:2
value:4, weight:4
value:5, weight:3
-----------填表-----------
000 000 000 000 000 000 000 000 000 000 000 
000 000 000 000 000 001 001 001 001 001 001 
000 001 001 001 001 001 002 002 002 002 002 
000 001 003 004 004 004 004 004 005 005 005 
000 001 003 004 004 005 007 008 008 008 008 
000 001 003 005 006 008 009 009 010 012 013 
