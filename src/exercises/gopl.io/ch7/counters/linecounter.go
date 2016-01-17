package counters

import (
	"bufio"
	"bytes"
)

// LineCounter counts the number of lines written to it before
//  discarding them
type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, scanner.Err()
}
