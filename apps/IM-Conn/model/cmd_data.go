package model

import "encoding/json"

type CmdData struct {
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`  // obj
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func NewCmdData() *CmdData {
	return new(CmdData)
}

func (c *CmdData) Byte() []byte {
	if data, err := json.Marshal(c); err == nil {
		return data
	}else{
		return []byte(err.Error())
	}
}

func (c *CmdData) SendMsg(str string, code int) []byte {
	c.Msg = str
	return c.Byte()
}
func (c *CmdData) SendCmd(cmd string, data interface{}) []byte {
	c.Cmd = cmd
	c.Data = data
	return c.Byte()
}