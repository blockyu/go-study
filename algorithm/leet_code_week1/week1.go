package leet_code_week1

func SingleNumber(nums []int) int {
	hashmap := make(map[int]int)
	for _, k := range nums {
		hashmap[k]++
	}
	for _, k := range nums {
		v, _ := hashmap[k]
		if v==1 {
			return k
		}
	}
	return 0
}
func IsHappy(n int) bool {
	hashmap := make(map[int]int)
	for {
		_, ok := hashmap[n]
		if ok {
			return false
		}
		hashmap[n] = n
		n=checkHappy(n)
		if n==1 {
			return true
		}
	}
}

func checkHappy(n int) int {
	var cur, sum int
	sum=0
	for {
		cur = n % 10
		sum += cur*cur
		n /= 10
		if  n == 0 {
			break
		}
	}
	return sum
}
