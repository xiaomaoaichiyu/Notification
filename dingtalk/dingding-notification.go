package dingtalk

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendDingMsg(msg string) {
	//请求地址模板
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=0d5f33a18069b75c8853025106a81d4f1b226db98969a1490a9403c74989b1ff`
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		// handle error
	}

	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	//发送请求
	resp, err := client.Do(req)

	p, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(p))

	//关闭请求
	defer resp.Body.Close()

	if err != nil {
		// handle error
		fmt.Println(err)
	}
}
