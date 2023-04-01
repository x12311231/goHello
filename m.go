package goHello

import (
    "fmt"
    "errors"
)

func Hello(name string) (string, error) {
    if "" == name {
        return "", errors.New("name cannot be empty")
    }
    return fmt.Sprintf("hello, %s", name), nil
}
