package group

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	WebHookIsEmpty     = errors.New("webHook is empty")
	UnSupportedMsgType = errors.New("unSupported msg type")
)

var (
	TextContentFormat     = `{"content":"%s", "mentioned_list":%+v, "mentioned_mobile_list": %v}`
	MarkdownContentFormat = `{"content":"%s"}`
	ImageContentFormat    = `{"base64":"%s", "md5":"%s"}`
	NewsContentFormat     = `{"articles":%s}` // 图文消息，一个图文消息支持1到8条图文，超过九条则消息发送失败
	FileContentFormat     = `{"media_id":"%s"}`
)

var (
	contentType      = "application/json"
	msgFormatMapping = map[string]string{
		MsgTypeOfText:     `{"msgtype": "%s","text": %s}`,
		MsgTypeOfMarkdown: `{"msgtype": "%s","markdown": %s}`,
		MsgTypeOfImage:    `{"msgtype": "%s","image": %s}`,
		MsgTypeOfNews:     `{"msgtype": "%s","news": %s}`,
		MsgTypeOfFile:     `{"msgtype": "%s","file": %s}`,
	}
)

type NewsMsg struct {
	Title       string `json:"title"`       // 标题，不超过128个字节，超过会自动截断
	Description string `json:"description"` // 描述，不超过512个字节，超过会自动截断
	Url         string `json:"url"`         // 点击后跳转的链接。
	PicUrl      string `json:"pic_url"`     // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图 1068*455，小图150*150。
}

func SendEnterPriseWeChat(webHook, msgType, msg string) error {
	if webHook == "" {
		return WebHookIsEmpty
	}
	format, ok := msgFormatMapping[msgType]
	if !ok {
		return UnSupportedMsgType
	}
	fmt.Println(fmt.Sprintf(format, msgType, msg))
	_, err := http.Post(webHook, contentType, strings.NewReader(fmt.Sprintf(format, msgType, msg)))
	return err
}

// 组装消息需要通知的用户(支持)
func AssembleNoticeUserString(notices []string) string {
	if len(notices) == 0 {
		return ""
	}

	ret := ""
	for idx, n := range notices {
		if idx == 0 {
			ret += "["
		}
		if idx == len(notices)-1 {
			ret += "]"
		}
		ret += "\"" + n + "\""
	}

	return ret
}

// todo 图片base64的
// todo 获取文件的md5
