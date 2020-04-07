package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type _GitHubUser struct {
	ID    uint64
	Login string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cab username1 username2 username3 ...\n" +
			"  Will print:\n  'Co-authored-by: username1 <100001+username1@users.noreply.github.com>'\n" +
			"  'Co-authored-by: username2 <100002+username2@users.noreply.github.com>'\n  'Co-authored-by:" +
			" username3 <100003+username3@users.noreply.github.com>'\n  ...")
		os.Exit(0)
	}

	for i := 1; i < len(os.Args); i++ {
		resp, err := http.Get("https://api.github.com/users/" + os.Args[i])
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}
		var u _GitHubUser
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Co-authored-by: %s <%d+%s@users.noreply.github.com>\n", u.Login, u.ID, u.Login)
	}
}
