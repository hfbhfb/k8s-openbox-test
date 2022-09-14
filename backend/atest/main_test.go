package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type EmuHttp struct {
	done chan struct{} // created lazily, closed by first cancel call
}

func (e *EmuHttp) Prepare() {
	if e.done == nil {
		e.done = make(chan struct{})
	}
	go func() {
		rand.Seed(time.Now().UnixNano())
		a := rand.Int()
		if a%2 == 0 {
			time.Sleep(time.Second * 9)
		}
		if a%2 == 1 {
			time.Sleep(time.Second * 1)
		}

		e.done <- struct{}{}
	}()
	return
}

func (e *EmuHttp) GetBytes() []byte {
	return []byte("emu http get byte")
}

func CallLogin() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	doneChanHttp := EmuHttp{}
	doneChanHttp.Prepare()
	select {
	case <-ctx.Done():
		fmt.Println("return err: no need call http call linksssssss")
		return nil, errors.New("time out get from http")
	case <-doneChanHttp.done:
		return doneChanHttp.GetBytes(), nil
	}
}

func TestDocumentationExamples(t *testing.T) {
	bs, err := CallLogin()
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Log(string(bs))
	t.Log("done")
}
