package limitreader

import "io"

type limitReader struct {
  nRem int
  reader io.Reader
}

func (l *limitReader) Read(p []byte) (int, error) {
  if l.nRem <= 0 {
    return 0, io.EOF
  }
  if len(p) > l.nRem {
    p = p[0:l.nRem]
  }
  n, err := l.reader.Read(p)
  l.nRem -= n
  return n, err
}

// LimitReader returns a Reader that reads from r
// but reports an EOF condition after n bytes.
func LimitReader(r io.Reader, n int) io.Reader {
  return &limitReader{n, r}
}