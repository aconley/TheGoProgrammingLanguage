package movieposters

import (
  "io"
  "os"
  "fmt"
  "net/http"
)

func writeURLToJpg(imageURL, outname string) (int64, error) {
  out, err := os.Create(outname)
  defer out.Close()

  resp, err := http.Get(imageURL)
  defer resp.Body.Close()
  
  n, err := io.Copy(out, resp.Body)
  if err != nil {
    return n, fmt.Errorf("Error writing to %s from %s: %v",
      outname, imageURL, err)
  }
  
  return n, nil
}

// WritePosters downloads a set of urls and
//  writes them out.  The output filenames
//  are returned, and are based on the input title
func WritePosters(title string, urls []string) ([]string, error) {
  if len(urls) == 1 {
    outname := title + ".jpg"
    _, err := writeURLToJpg(urls[0], outname)
    if err != nil {
      return nil, err
    }
    return []string{outname}, nil
  }
  outputNames := make([]string, len(urls))
  for idx, url := range urls {
    outname := fmt.Sprintf("%s%d.jpg", title, idx)
    _, err := writeURLToJpg(url, outname)
    if err != nil {
      return nil, err
    }
    outputNames[idx] = outname
  }
  return outputNames, nil
}