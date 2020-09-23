package group

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestSendEnterPriseWeChat(t *testing.T) {
	type args struct {
		webHook string
		msgType string
		msg     string
	}
	picBase64, picMd5 := getImageMessage()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "standard-text",
			args: args{
				webHook: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ec479b3f-6173-4a18-8187-161b0473f240",
				msgType: MsgTypeOfText,
				msg:     fmt.Sprintf(TextContentFormat, "这是一个文本消息", []string{"\"@all\""}, "[]"),
			},
			wantErr: false,
		},
		{
			name: "standard-markdown",
			args: args{
				webHook: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ec479b3f-6173-4a18-8187-161b0473f240",
				msgType: MsgTypeOfMarkdown,
				msg: fmt.Sprintf(MarkdownContentFormat, "<font color='info'>绿色</font> \n "+
					"<font color='comment'>灰色</font> \n"+
					"<font color='warning'>橙红色</font> \n"+
					"[请点击这里,跳转](https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B) \n"+
					"## 这是一个标题 \n"+
					"**这是一个加粗的markdown消息** \n"+
					"> 引用 \n"+
					""),
			},
			wantErr: false,
		},
		{
			name: "standard-image",
			args: args{
				webHook: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ec479b3f-6173-4a18-8187-161b0473f240",
				msgType: MsgTypeOfImage,
				msg:     fmt.Sprintf(ImageContentFormat, picBase64, picMd5),
			},
			wantErr: false,
		},
		{
			name: "standard-news",
			args: args{
				webHook: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ec479b3f-6173-4a18-8187-161b0473f240",
				msgType: MsgTypeOfNews,
				msg:     fmt.Sprintf(NewsContentFormat, getNewsMessage()),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendEnterPriseWeChat(tt.args.webHook, tt.args.msgType, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("SendEnterPriseWeChat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getImageMessage() (string, string) {
	// 图片的
	f, _ := ioutil.ReadFile("/Users/zmm/Desktop/other/235000-1584114600db79.jpg")
	picBase64 := base64.StdEncoding.EncodeToString(f)
	picMd5 := fmt.Sprintf("%x", md5.Sum(f))
	return picBase64, picMd5
}

func getNewsMessage() string {
	news := []NewsMsg{
		{
			Title:       "这是一条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是二条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是三条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是四条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是五条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是六条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是七条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		{
			Title:       "这是八条图文消息",
			Description: "洗澡啦，洗澡啦",
			Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
			PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
		// 超过九条则消息发送失败
		// {
		// 	Title:       "这是九条图文消息",
		// 	Description: "洗澡啦，洗澡啦",
		// 	Url:         "https://work.weixin.qq.com/api/doc/90000/90136/91770#%E6%96%87%E6%9C%AC%E7%B1%BB%E5%9E%8B",
		// 	PicUrl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		// },
	}

	data, _ := json.Marshal(news)
	return string(data)
}
