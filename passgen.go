package main

import (
    "flag"
    "fmt"
    "math/rand"
    "time"
)

var (
    length    int
    lowercase bool
    uppercase bool
    numbers   bool
    symbols   bool
)

const (
    lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
    uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numberChars    = "0123456789"
    symbolChars    = "!@#$%^&*()-=_+[]{}|;':\"<>,.?/~`"
)

func init() {
    rand.Seed(time.Now().UnixNano())

    flag.IntVar(&length, "l", 12, "password length")
    flag.BoolVar(&lowercase, "lowercase", true, "include lowercase characters")
    flag.BoolVar(&uppercase, "uppercase", true, "include uppercase characters")
    flag.BoolVar(&numbers, "numbers", true, "include numbers")
    flag.BoolVar(&symbols, "symbols", true, "include symbols")
    flag.Parse()
}

func main() {
    var chars string

    if lowercase {
        chars += lowercaseChars
    }
    if uppercase {
        chars += uppercaseChars
    }
    if numbers {
        chars += numberChars
    }
    if symbols {
        chars += symbolChars
    }

    password := make([]byte, length)
    for i := 0; i < length; i++ {
        password[i] = chars[rand.Intn(len(chars))]
    }

    fmt.Println(string(password))
}

