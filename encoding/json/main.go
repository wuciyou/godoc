package main

import (
	"encoding/json"
	"fmt"
)

// type content struct {
// interface{}
// }
type value struct {
	Msg_id     int64
	From_uin   int64
	To_uin     int64
	Msg_id2    int64
	Reply_ip   int64
	Msg_type   int64
	Group_code int64
	Send_uin   int64
	Seq        int64
	Time       int64
	Info_seq   int64
	Content    []interface{}
}

type result struct {
	Poll_type string
	Value     value
}

type body struct {
	Result  []result
	P       string
	Errmsg  string
	Retcode int
}

var resultStr = `{
	"result": [
		{
			"poll_type": "group_message",
			"value": {
				"content": [
					[
						"font",
						{
							"color": "000000",
							"name": "宋体",
							"size": 10,
							"style": [
								0,
								0,
								0
							]
						}
					],
					"睡醒了 "
				],
				"from_uin": 1225509606,
				"group_code": 2003831713,
				"info_seq": 106211945,
				"msg_id": 9734,
				"msg_id2": 745861,
				"msg_type": 43,
				"reply_ip": 176756824,
				"send_uin": 640396092,
				"seq": 12256,
				"time": 1446367921,
				"to_uin": 554319697
			}
		}
	],
	"retcode": 0
}`

func main() {
	fmt.Println("json unmarshal example")
	unmarshal()
}

func unmarshal() {
	b := &body{}
	json.Unmarshal([]byte(resultStr), b)
	fmt.Println(b)
}
