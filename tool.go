package gotool

import (
	"github.com/jiaoningbo/gotool/array"
	"github.com/jiaoningbo/gotool/bcrypt"
	"github.com/jiaoningbo/gotool/captcha"
	"github.com/jiaoningbo/gotool/compression"
	"github.com/jiaoningbo/gotool/conversion"
	"github.com/jiaoningbo/gotool/convert"
	"github.com/jiaoningbo/gotool/date"
	"github.com/jiaoningbo/gotool/desensitized"
	"github.com/jiaoningbo/gotool/id_utils"
	"github.com/jiaoningbo/gotool/logs"
	"github.com/jiaoningbo/gotool/openfile"
	"github.com/jiaoningbo/gotool/page"
	"github.com/jiaoningbo/gotool/pretty"
	"github.com/jiaoningbo/gotool/request"
	"github.com/jiaoningbo/gotool/str"
	"github.com/jiaoningbo/gotool/tree"
)

// @Title tool
// @Description A simple, semantic and developer-friendly Golang tool development set
// @Page github.com/jiaoningbo/gotool
// @Version v0.0.1
// @Author druidcaesa
// @Email hbsjzfynxm@gmail.com

var (
	StrArrayUtils     array.StrArray            //String 数据工具声明
	Logs              logs.Logs                 //log日志声明
	BcryptUtils       bcrypt.BcryptUtil         //加密工具声明
	DateUtil          date.Date                 //时间操作
	StrUtils          str.StrUtils              //字符串操作
	HttpUtils         request.Request           //http工具
	ConvertUtils      convert.Convert           //公历转农历相关操作
	PageUtils         page.Page                 //分页插件
	IdUtils           id_utils.IdUtils          //id生成工具
	ZipUtils          compression.ZipUtils      //压缩和解压缩工具
	FileUtils         openfile.FileUtils        //IO操作工具
	CaptchaUtils      captcha.Captcha           //验证码工具
	DesensitizedUtils desensitized.Desensitized //敏感数据脱敏
	TreeUtils         tree.Tree                 //菜单树结构化工具
	PrettyUtils       pretty.PrettyUtils        //JSON打印格式化工具
	TypeConversion    conversion.Conversion     //类型转换，用于string，int，int64，float等数据转换，免去err的接收，和设置默认值
)
