package limitreader

import (
  "io"
  "strings"
  "testing"
)

func TestLimitReader(t *testing.T) {
  t.Log("Testing reading from string reader")
  testString := "Now is the winter of our discontent"
  var strRdr io.Reader = strings.NewReader(testString)

  b := make([]byte, 10)

  n, err := strRdr.Read(b)
  if err != nil {
    t.Errorf("Got error %v reading from strReader", err)
  }
  if (n != 10) {
    t.Errorf("Got %d bytes, expected %d", n, 10)
  }

  limRdr := LimitReader(strRdr, 5)
  n, err = limRdr.Read(b)
  if err != nil {
    t.Errorf("Got error %v reading from limRdr", err)
  }
  if (n != 5) {
    t.Errorf("Got %d bytes, expected %d from limit reader", n, 5)
  }

  n, err = limRdr.Read(b)
  if err != io.EOF {
    t.Errorf("Should have gotten EOF error on second read")
  }

}