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
	///////////
	client := http.Client{
		Timeout: time.Duration(6) * time.Second,
	}
	resp, err := client.Get("http://localhost:3000/get-token?login=root2&password=2&data=21")
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

	fmt.Printf("Body blog.logrocket.com : %s\t", body)

	var Gettokenanswer = &Gettokenanswerstruct{}
	json.Unmarshal([]byte(body), Gettokenanswer)
	fmt.Println("token from struct:", Gettokenanswer.Token)
	//////////
	// respo, erro := client.Get("http://localhost:3000/products")
	// if erro != nil {
	// 	fmt.Println(erro)
	// 	return
	// }
	// defer respo.Body.Close()

	// bodyo, err := ioutil.ReadAll(respo.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// // bodyString := string(body)

	// fmt.Printf("Body blog.logrocket.com : %s", bodyo)
	bearer := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBVFRFTlRJT04hIjoi0J_RgNC40LLQtdGCLCDQnNCw0LrRgSA6KSIsIkRhdGEgYW5zd2VyIGlzIjoiMjExIiwiVG9rZW4gcmVxdWVzdCBhdCI6IjIwMjItMDUtMjRUMjI6MzM6NTkuMDY5ODcwNyswNTowMCIsImFkbWluIHBlcm1pc3Npb25zPyI6Im1heWJlIiwiZXhwIjoxNjUzNDc4NDM5LCJsb2dpbiI6InJvb3QyIn0.IZV4284A8Ss9bnlsRe_WTZr2l7XEtioXbG_m9pLQtQY"
	cli := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", "http://localhost:3000/products", nil)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", `application/json`)
	if err != nil {
		panic(err)
	}
	// defer request.Body.Close() /where to add this?

	response, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("response?:", string(responseData))
	// if err != nil {
	// 	fmt.Println("err")
	// 	fmt.Println(err)
	// 	return
	// }
	//defer request.Body.Close()
	// io.Copy(os.Stdout, req.Body)
	// bodyprod, erra := ioutil.ReadAll(request.Body)
	// if erra != nil {
	// 	fmt.Println("bodyprod error")
	// 	panic(erra)
	// }
	// bodyprodString := string(bodyprod)
	// fmt.Println(bodyprodString)

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

	// func (p *PostMessageEndpoint) handler(w http.ResponseWriter, r *http.Request) {
	// 	message := Gettokenanswerstruct{}
	// 	err := http_helper.HttpHelper{}.DecodePostRequest(r, &message)
	// 	if err != nil {
	// 		fmt.Println("can not decode post message", r.GetBody)
	// 		return
	// 	}
	// }

	// fmt.Println(PostMessage("http://localhost:3000/get-token?login=root111&password=1111&data=21"))

}
