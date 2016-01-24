package counters

import "io"

type countingWriter struct {
	count  *int64
	writer io.Writer
}

func (c countingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	if err == nil {
		*c.count += int64(n)
	}
	return n, err
}

// CountingWriter returns a new counting writer and a pointer to
//  the number of bytes written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cntr int64
	retval := countingWriter{&cntr, w}
	return retval, retval.count
}
