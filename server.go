package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// window下乱码问题 https://blog.csdn.net/qq_37493556/article/details/107541084
func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func cmdRun(name string, args []string) string {
	cmd := exec.Command(name, args...)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return ConvertByte2String(out.Bytes(), GB18030)
}

func resp(ctx *gin.Context, str string) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, str)
}

func main()  {

	// 配置日志
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	r.Any("/*id", func(ctx *gin.Context) {

		uri := ctx.Request.RequestURI
		re := regexp.MustCompile(`.*\.js|.*\.css|.*\.png|.*\.jpg`)

		if re.Match([]byte(uri)) {
			return
		}

		// 去除图标文件获取
		if uri == "/favicon.ico" {
		  	resp(ctx, "")
		  	return
		}

		output := cmdRun(
			"F:\\Python\\Kronos\\venv\\Scripts\\python.exe",
			[]string{
				"F:\\Python\\Kronos\\entry.py",
				"--params=proxy_url=http://localhost:8080" + ctx.Request.RequestURI,
			})

		print(output)
		resp(ctx, output)
	})

	r.Run("0.0.0.0:8081")
}
