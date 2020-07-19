package main

import (
    "flag"
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/emersion/go-msgauth/dkim"
)

func main() {

    fname := flag.String("fname", "mailprueba.txt", "input file name containing email headers and body")
    flag.Parse()

    content, err := ioutil.ReadFile(*fname)
    if err != nil {
        fmt.Println("ERROR: " + err.Error())
        return
    }

    text := string(content)

    r := strings.NewReader(text)

    verifications, err := dkim.Verify(r)
    if err != nil {
        fmt.Println("ERROR: " + err.Error())
        return
    }

    fmt.Printf("verifications found: %d\n", len(verifications))

    for _, v := range verifications {
        if v.Err == nil {
            fmt.Println("Valid signature for:", v.Domain)
        } else {
            fmt.Println("Invalid signature for:", v.Domain, v.Err)
        }
    }
}
