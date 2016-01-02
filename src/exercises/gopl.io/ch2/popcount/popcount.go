package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i / 2] + byte(i & 1)
    }
}

// Shift and Lookup based popcount
func PopCountByLookup(x uint64) int {
    return int(pc[byte(x)] +
               pc[byte(x >> 8)] +
               pc[byte(x >> 16)] +
               pc[byte(x >> 24)] +
               pc[byte(x >> 32)] +
               pc[byte(x >> 40)] +
               pc[byte(x >> 48)] +
               pc[byte(x >> 56)])
}

func PopCountByLookup2(x uint64) int {
    sum := pc[byte(x)]
    for i := uint(1); i < 8; i++ {
        sum += pc[byte(x >> i)]
    }
    return int(sum)
}

func PopCountByShifting(x uint64) int {
    n := 0
    shifted := x
    for i := uint(0); i < 64; i++ {
        if ((shifted & 1) != 0) {
            n++
        }
        shifted >>= 1
    }
    return n
}

func PopCountByClearing(x uint64) int {
    n := 0
    for x != 0 {
        x &= (x - 1) // Clear rightmost bit
        n++
    }
    return n
}

// This is the Hacker's Delight version
func PopCountHacker(x uint64) int {
        x = x - ((x >> 1) & 0x5555555555555555)
        x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
        x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
        x = x + (x >> 8)
        x = x + (x >> 16)
        x = x + (x >> 32)
        return int(x & 0x7f)
}
