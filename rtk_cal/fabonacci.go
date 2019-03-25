package rtk_cal

var map_fibcache = make(map[int]uint64)

func Getfibonacci(x int, ch chan uint64) {
	if isGeted(x) {
		ch <- map_fibcache[x]
		return
	} else {
		res := fibN(x)
		map_fibcache[x] = res
		ch <- res
		return
	}
}

func fibN(x int) uint64 {
	if x < 3 {
		return 1
	}
	var a, b uint64 = 1, 1
	for i := 3; i <= x; i++ {
		b, a = a+b, b
	}
	return b
}

func isGeted(key int) bool {
	if _, ok := map_fibcache[key]; ok {
		return true
	}
	return false
}
