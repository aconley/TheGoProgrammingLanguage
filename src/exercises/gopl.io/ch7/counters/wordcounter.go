package counters

import (
	"bufio"
	"bytes"
)

// WordCounter counts the number of words written to it before
//  discarding them
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, scanner.Err()
}
