package bands

import (
  "net/http"
  "io"
  "io/ioutil"
  "encoding/json"
  "errors"
  "fmt"
)

func getError(responseBody io.ReadCloser) string {
  errorResponse := map[string][]string{}

  readBody(responseBody, &errorResponse)

  if len(errorResponse["errors"]) > 0 {
    return errorResponse["errors"][0]
  }

  return ""
}

func get(url string, obj interface{}) error {
  resp, err  := http.Get(url)

  if err != nil {
    trace("error GET %s", err)
    return err
  }

  if resp.StatusCode != 200 {
    trace("error status code %d", resp.StatusCode)
    
    errMessage := getError(resp.Body)
    if len(errMessage) > 0 {
      return errors.New(errMessage)
    }
    return errors.New(fmt.Sprintf("status code %d", resp.StatusCode))
  }

  defer resp.Body.Close()

  readBody(resp.Body, &obj)

  return nil
}

func readBody(responseBody io.ReadCloser, obj interface{}) error {
  body, err := ioutil.ReadAll(responseBody);
  if err != nil {
    return err
  }

  if err := json.Unmarshal(body, &obj); err != nil {
    trace("error %s", err)
    return err
  }

  return nil
}
