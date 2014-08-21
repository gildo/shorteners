package main

import (
    "github.com/go-martini/martini"
	"menteslibres.net/gosexy/redis"
    "encoding/base64"
    "crypto/rand"
    "net/http"
)

var client *redis.Client

func generateURL(url string) string {
    rb := make([]byte,3)
    _, err := rand.Read(rb)
    if err != nil {
        return "err"
    }
    rs := base64.URLEncoding.EncodeToString(rb)
    client.Set(rs, url)
    return url + " " + rs
}

func main() {
    m := martini.Classic()
    client = redis.New()
    client.Connect("127.0.0.1", uint(6379))

    m.Post("/:id", func(params martini.Params) string {
        return generateURL(params["id"])
    })

    m.Get("/:id", func(params martini.Params, res http.ResponseWriter, req *http.Request) {
        s, _ := client.Get(params["id"])

        http.Redirect(res, req, "http://"+s, 301)
    })

  m.Run()
}
