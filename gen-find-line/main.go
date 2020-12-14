package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// 最终吐出去的结果
type Result struct {
	FindName string
	FindDir  []*FileDirFormat //
}

type FileDirFormat struct {
	Dir   string
	Files []*File
}

type File struct {
	Name  string
	Lines []int64
}

var (
	filePath string
	usage    = `
input
	Method
		getTagSearchList
	Found usages  (8 usages found)
	wechat-cx  (8 usages found)
			app/Http/App/V1_0/Controllers  (1 usage found)
				CategoryController.php  (1 usage found)
					220 $outData = $this->categoryLogic->getTagSearchList($params);

output
	find *** in the
		app/Http/BaiDuCx/V1_0/Controllers/CategoryController.php:195
		...
		...
	use`
)

func flagInit() {
	flag.StringVar(&filePath, "f", "", "please enter report.txt file path to export the IDE")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stdout, usage)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	flagInit()
	if filePath == "" {
		panic("input filePath is empty")
	}
	result, err := src(filePath)
	if err != nil {
		panic("read src is failed" + err.Error())
	}
	out(result, os.Stdout)
}

func src(filePath string) (*Result, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	ret := &Result{}
	// 读取文件
	findFlagIdx := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		findFlagIdx++
		line := scanner.Text()
		if findFlagIdx == 2 {
			ret.FindName = strings.TrimSpace(line)
		}
		// 正则匹配 /， 说明是文件夹
		match, _ := regexp.Match("\\/", []byte(line))
		if match {
			// 根据空格进行截取
			lines := strings.Fields(line)
			ret.FindDir = append(ret.FindDir, &FileDirFormat{
				Dir:   lines[0],
				Files: make([]*File, 0),
			})
			continue
		}
		// 匹配文件
		suffixMatch, _ := regexp.Match(".php", []byte(line))
		if suffixMatch {
			if len(ret.FindDir) == 0 {
				panic("file format is wrong")
			}
			// 获取对应的文件夹
			dir := ret.FindDir[len(ret.FindDir)-1]
			lines := strings.Fields(line)
			dir.Files = append(dir.Files, &File{
				Name:  lines[0],
				Lines: []int64{},
			})
			continue
		}
		// 最后获取查找到的行数
		lines := strings.Fields(line)
		if len(lines) < 1 {
			continue
		}
		// 正则匹配数字
		if ok, _ := regexp.Match("^[0-9]", []byte(lines[0])); ok {
			// 查找归属
			dir := ret.FindDir[len(ret.FindDir)-1]
			file := dir.Files[len(dir.Files)-1]
			file.Lines = append(file.Lines, ToInt64(lines[0]))
		}
	}

	return ret, nil
}

func out(result *Result, writer io.Writer) {
	if result == nil {
		panic("result is nil")
	}
	_, _ = writer.Write([]byte(fmt.Sprintf("find %s in the \n", result.FindName)))
	for _, r := range result.FindDir {
		for _, f := range r.Files {
			for _, l := range f.Lines {
				_, _ = writer.Write([]byte(fmt.Sprintf("\t %s/%s:%d\n", r.Dir, f.Name, l)))
			}
		}
	}
	_, _ = writer.Write([]byte("use"))
}

func ToInt64(i interface{}) (v int64) {
	switch t := i.(type) {
	case string:
		v, _ = strconv.ParseInt(t, 10, 64)
	case int:
		v = int64(t)
	case int8:
		v = int64(t)
	case int16:
		v = int64(t)
	case int32:
		v = int64(t)
	case int64:
		v = t
	case uint:
		v = int64(t)
	case uint8:
		v = int64(t)
	case uint16:
		v = int64(t)
	case uint32:
		v = int64(t)
	case uint64:
		v = int64(t)
	case float64:
		v = int64(t)
	case []uint8:
		v, _ = strconv.ParseInt(Ui8ToA(i), 10, 64)
	}

	return v
}

// []uint8 转字符串
func Ui8ToA(i interface{}) string {
	if v, ok := i.(string); ok {
		return v
	}

	return string(Ui8ToB(i))
}

// []uint8 转字符串字节
func Ui8ToB(i interface{}) (b []byte) {
	if v, ok := i.([]uint8); ok {
		for _, val := range v {
			b = append(b, val)
		}
	}

	return b
}
