package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	My_redis "github.com/ev-go/myhttp3authjson/Cache"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type TMessage struct {
	FirstKey    string
	SecondKey   string
	Name        string
	PhoneNumber string
	ICQ         string
	LastKey     int64
}

type Gettokenanswerstruct struct {
	TokenRequestAt string
	User           string
	Login          string
	Password       string
	DataAnswer     string
	Token          string
}

// func (p *Gettokenanswerstruct) handler(w http.ResponseWriter, r *http.Request) {
// 	message := types.Message{}
// 	err := http_helper.HttpHelper{}.DecodePostRequest(r, &message)
// 	if err != nil {
// 		fmt.Println("can not decode post message", r.GetBody)
// 		return
// 	}
// }

// type httprequeststruct struct {
// 	Url string
// }

func main() {
	My_redis.Main()

	// // http_helper.HttpHelper("http://localhost:3000/get-token?login=root111&password=1111&data=21")

	// httprequest := httprequeststruct{"http://localhost:3000/get-token?login=root111&password=1111&data=21"}
	// message := Gettokenanswerstruct{}
	// http_helper.HttpHelper{}.DecodePostRequest(httprequest, &message)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	userid := "root25"
	currentusertoken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBVFRFTlRJT04hIjoi0J_RgNC40LLQtdGCLCDQnNCw0LrRgSA6KSIsIkRhdGEgYW5zd2VyIGlzIjoiMjExIiwiVG9rZW4gcmVxdWVzdCBhdCI6IjIwMjItMDUtMTJUMjI6MDI6MDMuNDIzNTc1NCswNTowMCIsImFkbWluIHBlcm1pc3Npb25zPyI6Im1heWJlIiwiZXhwIjoxNjUyMzc1NTIzLCJsb2dpbiI6InJvb3QifQ.9do8soXtimGxr9TDAd6EI2W0l-95U0SSJD_5GPz4kMA"

	node := rdb.Set(ctx, userid, currentusertoken, 0).Err()
	if node != nil {
		panic(node)
	}

	// err = rdb.Set(ctx, "key2", "74", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, node := rdb.Get(ctx, userid).Result()
	if node == redis.Nil {
		fmt.Println("key1 does not exist")
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println(userid, val)
	}
	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	val2, node := rdb.Get(ctx, "key2").Result()
	if node == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("key2", val2)
	}

	// Output: key value
	// key2 does not exist

	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("http://localhost:3000/get-token?login=root111&password=1111&data=21")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// bodyString := string(body)

	var Gettokenanswer = &Gettokenanswerstruct{}
	json.Unmarshal([]byte(body), Gettokenanswer)
	// fmt.Println(Gettokenanswer.Token)
	// fmt.Println(Gettokenanswerstruct)

	// }
	// fmt.Println(message)

	//var Gettokenanswer = Gettokenanswerstruct{bodyString.Token}

	// newToken := resp.Body
	// osStdout := os.Stdout

	// newTokenToRedis := rdb.Set(ctx, userid, newToken, 0).Err()
	// if newTokenToRedis != nil {
	// 	panic(newTokenToRedis)
	// }

	// tokenFromRedis, newTokenToRedis := rdb.Get(ctx, "userid").Result()
	// if newTokenToRedis == redis.Nil {
	// 	fmt.Println("userid does not exist")
	// } else if newTokenToRedis != nil {
	// 	panic(newTokenToRedis)
	// } else {
	// 	fmt.Println("userid", tokenFromRedis)
	// }

	// PostMessageEndpoint := "http://localhost:3000/get-token?login=root111&password=1111&data=21"

	

	func (p *PostMessageEndpoint) handler(w http.ResponseWriter, r *http.Request) {
		message := Gettokenanswerstruct{}
		err := http_helper.HttpHelper{}.DecodePostRequest(r, &message)
		if err != nil {
			fmt.Println("can not decode post message", r.GetBody)
			return
		}
	}

	// fmt.Println(PostMessage("http://localhost:3000/get-token?login=root111&password=1111&data=21"))

}
