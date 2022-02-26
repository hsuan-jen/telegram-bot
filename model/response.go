package model

type Command struct {
	ChatID int64                  `json:"chat_id"`
	CMD    string                 `json:"cmd"`    //指令
	Params map[string]interface{} `json:"params"` //参数
}
