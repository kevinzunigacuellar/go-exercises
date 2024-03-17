package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Name string
	Id   int
	err  error
}

func (u User) String() string {
	return fmt.Sprintf("%d. %s", u.Id, u.Name)
}

func getUser(s string) (User, error) {
	sp := strings.Split(s, ":")
	if len(sp) != 2 {
		return User{}, fmt.Errorf("record(%s) was not in the correct format", s)
	}

	id, err := strconv.Atoi(sp[1])

	if err != nil {
		return User{}, fmt.Errorf("record(%s) id was not an integer", s)
	}

	return User{Name: strings.TrimSpace(sp[0]), Id: id}, nil
}

func decodeUsers(ctx context.Context, r io.Reader) chan User {
	ch := make(chan User, 1)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if ctx.Err() != nil {
				ch <- User{err: ctx.Err()}
				return
			}
			u, err := getUser(scanner.Text())
			if err != nil {
				u.err = err
				ch <- u
				return
			}
			ch <- u
		}
	}()
	return ch
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	f, err := os.Open("files/users.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ch := decodeUsers(ctx, f)

	for u := range ch {
		if u.err != nil {
			panic(u.err)
		}
		fmt.Println(u)
	}
}
