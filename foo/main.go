package main

import "fmt"

import "go get gopkg.in/yaml.v2"

func main() {
    m := make(map[interface{}]interface{})
    err = yaml.Unmarshal([]byte(`{"foo": "bar"}`), &m)
    if err != nil {
            log.Fatalf("error: %v", err)
    }

    fmt.Printf("Parsed: %v", m)
}
