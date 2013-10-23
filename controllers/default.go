package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/starchou/webcha/utils"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

const (
	TOKEN    = "starchouforwecha"
	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

type msgBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

type Request struct {
	XMLName                xml.Name `xml:"xml"`
	msgBase                         // base struct
	Location_X, Location_Y float32
	Scale                  int
	Label                  string
	PicUrl                 string
	MsgId                  int
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	ArticleCount int     `xml:",omitempty"`
	Articles     []*item `xml:"Articles>item,omitempty"`
	FuncFlag     int
}

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	signature := this.Input().Get("signature")
	beego.Info(signature)
	timestamp := this.Input().Get("timestamp")
	beego.Info(timestamp)
	nonce := this.Input().Get("nonce")
	beego.Info(nonce)
	echostr := this.Input().Get("echostr")
	beego.Info(echostr)
	beego.Info(Signature(timestamp, nonce))
	if Signature(timestamp, nonce) == signature {
		this.Ctx.WriteString(echostr)
	} else {
		this.Ctx.WriteString("")
	}
}

func (this *MainController) Post() {
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	beego.Info(string(body))
	var wreq *Request
	if wreq, err = DecodeRequest(body); err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	beego.Info(wreq.Content)
	wresp, err := dealwith(wreq)
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	data, err := wresp.Encode()
	if err != nil {
		beego.Error(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
		return
	}
	this.Ctx.WriteString(string(data))
	return
}

func dealwith(req *Request) (resp *Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Text
	beego.Info(req.MsgType)
	beego.Info(req.Content)
	if req.MsgType == Text {
		if strings.Trim(strings.ToLower(req.Content), " ") == "help" || req.Content == "Hello2BizUser" || req.Content == "subscribe" {
			resp.Content = "功能列表如下：查询手机归属地请输入：‘手机号码:13838385438’"
			return resp, nil
		}
		strs := strings.Split(req.Content, ":")
		//var resurl string
		//var a item
		beego.Info(req.Content)
		beego.Info(strs[0])
		switch strs[0] {
		case "手机号码":
			data, err := utils.GetDataString(strs[1], "")
			if err == nil {
				beego.Info(data)
				resp.Content = data
			} else {
				resp.Content = "查询出错!"
			}

		case "天气":
			resp.Content = "该功能正在开发中"
		}

		resp.MsgType = News
		resp.ArticleCount = 1
		resp.Articles = resp.Articles
		resp.FuncFlag = 1
	} else {
		resp.Content = "暂时还不支持其他的类型"
	}
	return resp, nil
}

func Signature(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DecodeRequest(data []byte) (req *Request, err error) {
	req = &Request{}
	if err = xml.Unmarshal(data, req); err != nil {
		return
	}
	req.CreateTime *= time.Second
	return
}

func NewResponse() (resp *Response) {
	resp = &Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}

func (resp Response) Encode() (data []byte, err error) {
	resp.CreateTime = time.Second
	data, err = xml.Marshal(resp)
	return
}
