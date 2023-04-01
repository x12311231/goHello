package goHello

import (
    "testing"
)

func TestHello(t *testing.T) {
    msg, err := Hello("bob")
    if nil != err || msg != "hello, bob" {
        t.Fatalf("some err:%v or response should be hello, bob,but %q", err, msg)
    }
}

func TestHelloNameEmpty(t *testing.T) {
    msg, err := Hello("")
    if nil == err {
        t.Fatalf(`Hello("") should current a err,but it's not, msg is %s, err is %v`, msg, err)
    }
}
