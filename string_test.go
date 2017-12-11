
package rstring

import (
    "testing"
)

func TestMatch(t *testing.T) {
    rs := String("http://github.com/akitanoyo/rstring")
    r := "//\\w+\\.\\w+/"
    if !rs.Match(r) {
        t.Errorf("got String.Match %v %v",
            rs.String(), r)
    }
}

func TestMatchList(t *testing.T) {
    rs := String("http://github.com/akitanoyo/rstring")
    r := "//(\\w+\\.\\w+)/"
    ms, ok := rs.MatchList(r)
    if !ok {
        t.Errorf("got String.MatchList %v %v",
            rs.String(), r)
    }

    if len(ms) != 1 || ms[0] != "github.com" {
        t.Errorf("got String.MatchList %v != %v",
            ms[0], "github.com")
    }
}

func TestReplaceAll(t *testing.T) {
    rs := String("http://github.com/akitanoyo/rstring")
    r := "/\\w+"
    rs.ReplaceAll(r, "X")
    if rs.String() != "http:/X.comXX"  {
        t.Errorf("got String.ReplaceAll %v != %v",
            rs.String(), "http:/X.comXX")
    }
}

func TestReplaceAllList(t *testing.T) {
    rs := String("http://github.com/akitanoyo/rstring")
    r := "/(\\w+)"
    _, res := rs.ReplaceAllList(r, "X")
    if rs.String() != "http:/X.comXX"  {
        t.Errorf("got String.ReplaceAllList %v != %v",
            rs.String(), "http:/X.comXX")
    }
    if len(res) != 3 {
        t.Errorf("got String.ReplaceAllList no 3match %v != %v",
            len(res), "3")
    }
    if res[2] != "rstring" {
        t.Errorf("got String.ReplaceAllList third match %v != %v",
            res[2], "rstring")
    }
   
}


