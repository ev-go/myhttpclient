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

// type httpRequestMessageStruct struct {
// 	requestUseLogin
// 	requestUsePassword
// 	requestUseData
// }

// type httpRequestStruct struct {
// 	requestUseUrl string
// 	requestUsePort string
// 	requestUseRout string
// 	httpRequestMessage httpRequestMessageStruct
// }

// requestUseUrl + requestUsePort + requestUseRout + "?" + requestUseLogin + "&" + requestUsePassword + "&" + requestUseData
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

	requestUseUrl := "http://localhost"
	requestUsePort := ":3000"
	requestUseRout := "/get-token"
	requestUseLogin := "login=root1"
	requestUsePassword := "password=1"
	requestUseData := "data=21"

	httpRequestString := requestUseUrl + requestUsePort + requestUseRout + "?" + requestUseLogin + "&" + requestUsePassword + "&" + requestUseData
	// Menu
	fmt.Println("\n********************************/ Menu /****************************************")
	fmt.Println("\nThis is client for sending http requests to server")
	fmt.Println("\nDefault URL: ", requestUseUrl, ";",
		"\nDefault Port: ", requestUsePort, ";",
		"\nDefault Rout: ", requestUseRout, ";",
		"\nDefault Login: ", requestUseLogin, ";",
		"\nDefault Password: ", requestUsePassword, ";",
		"\nDefault Data: ", requestUseData, ";",
		"\nDefault http request: ", httpRequestString, ";")

	// fmt.Println("Логин")
	// fmt.Scanf("%s\n", &Log)

	// fmt.Println("Пароль")
	// fmt.Scanf("%s\n", &Pass)
	var readFromTerminal string
	fmt.Println("\nDo you want to change defaults? (y/n)")
	fmt.Scanf("%s\n", &readFromTerminal)
	if readFromTerminal == "y" {
		fmt.Println("\nWhat part of http request need to change?",
			"Write in terminal: 'url' or 'port' or 'rout' or 'login' or 'password' or 'data'")
		fmt.Scanf("%s\n", &readFromTerminal)
		if readFromTerminal == "login" {
			fmt.Println("\nEnter new login")
			fmt.Scanf("%s\n", &readFromTerminal)
			requestUseLogin = "login=" + readFromTerminal
			fmt.Println("\nLogin changed for:", requestUseLogin)
		}
		if readFromTerminal == "password" {
			fmt.Println("\nEnter new password")
			fmt.Scanf("%s\n", &readFromTerminal)
			requestUsePassword = "password=" + readFromTerminal
			fmt.Println("\nPassword changed for:", requestUsePassword)
		}
		if readFromTerminal == "data" {
			fmt.Println("\nEnter new data")
			fmt.Scanf("%s\n", &readFromTerminal)
			requestUseData = "data=" + readFromTerminal
			fmt.Println("\nData changed for:", requestUseData)
		}
		if readFromTerminal == "port" {
			fmt.Println("\nEnter new port")
			fmt.Scanf("%s\n", &readFromTerminal)
			requestUsePort = ":" + readFromTerminal
			fmt.Println("\nPort changed for:", requestUsePort)
		}
	} else {
		fmt.Println("No changes")
	}
	httpRequestString = requestUseUrl + requestUsePort + requestUseRout + "?" + requestUseLogin + "&" + requestUsePassword + "&" + requestUseData

	fmt.Println("\n******************************/ Menu End /**************************************")
	//Menu end
	//	"http://localhost:3000/get-token?login=root111&password=1111&data=21"

	client := http.Client{
		Timeout: time.Duration(6) * time.Second,
	}
	resp, err := client.Get(httpRequestString)
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

	fmt.Printf("\nAnswer for request : %s\t", body)

	var Gettokenanswer = &Gettokenanswerstruct{}
	json.Unmarshal([]byte(body), Gettokenanswer)
	fmt.Println("\ntoken from struct:", Gettokenanswer.Token)
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
	bearer := "Bearer " + Gettokenanswer.Token
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

	fmt.Println("\nresponse?:", string(responseData))
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
