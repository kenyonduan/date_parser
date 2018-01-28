package date_parser

import (
	"regexp"
	"time"
	"strings"
	"fmt"
	"strconv"
	"github.com/wppurking/date_parser/i18n"
)

// 默认的日期布局
var (
	defaultLayouts = map[string]string{
		"en": "on January 2, 2006",      // 美国
		"uk": "on 2 January 2006",       // 英国
		"de": "am 2. January 2006",      // 德语
		"jp": "2006年January2日",          // 日语
		"fr": "le 2 January 2006",       // 法语
		"it": "il 2 January 2006",       // 意大利语
		"es": "el 2 de January de 2006", // 西班牙语
	}

	// 含有中文的字符
	asiaRegex = regexp.MustCompile(`((\d{1,2})月)`)
	// 韩国的月份
	koRegex = regexp.MustCompile(`((\d{1,2})월)`)

	// 默认的 en 的 longMonthName
	enLongMonthNames = i18n.LongMonthNames["en"]
)

// ParserLangDate 各国日期月份翻译,支持语言列表:
// ar 阿根廷
// de 德国
// en 美国
// es 西班牙
// fr 法国
// hu 匈牙利
// it 意大利
// jp 日本
// ko 韩国
// nl 荷兰
// pt 葡萄牙
// ro 罗马尼亚
// ru 俄罗斯
// sv 萨尔瓦多
// zh 中国
// 其他英语国家指定en同时需要自定义日期模板,比如英国,意大利,印度等
//--------------------------------------------------------------
// 传入语言标识,待处理日期,自定义日期布局
// 自定义日期模板需要将月份换成英语,如日本: 2016年11月13日 对应的模板 2006年January2日
// 自定义日期模板需要将月份换成英语,如法国: le 12 février 2016 对应的模板 le 2 January 2006
// 自定义日期模板需要将月份换成英语,如韩国: 2016년 11월 23일 对应的模板 2006년 January 2일
// 日本,中国,韩国,美国这些国家其实不用将月份替换成英语,但是为了统一好理解,粗暴的全部替换
//
// 解析思路为沿用 time.Parse 中的算法, 在上层对输入的字符串进行一些重写处理, 完后再进行 parse
// time.Parse 传入 layout 后, 核心会对 `月` 进行特别的解析, 其他的会跟着 Layout 进行匹配.
// 所以
// 1. 只需要对输入的 value 将月份调整到 en 格式即可
// 2. 同时对 layout 也映射处理, 将不同语言的 month 转换为 en 的 `月`
func ParserLangDate(lang, value string, layout ...string) (time.Time, error) {
	lastLayout := defaultLayouts["en"]
	switch lang {
	case "jp", "zh":
		v, err := regexReplaceVal(asiaRegex, value, enLongMonthNames)
		if err != nil {
			return time.Time{}, err
		}
		value = v
	case "ko":
		v, err := regexReplaceVal(koRegex, value, enLongMonthNames)
		if err != nil {
			return time.Time{}, err
		}
		value = v
	case "en", "us", "ca", "in", "au":
		// refs: https://github.com/theplant/cldr/blob/master/resources/locales/de/calendar.go
		// 所有 i18n 的语言格式化
		lang = "en"
	default:
		for k, v := range i18n.LongMonthNames[lang] {
			if strings.Contains(value, v) {
				month := enLongMonthNames[k]
				value = strings.Replace(value, v, month, -1)
				break
			}
		}
	}
	if _, ok := defaultLayouts[lang]; ok {
		lastLayout = defaultLayouts[lang]
	}

	if _, ok := defaultLayouts[lang]; ok {
		lastLayout = defaultLayouts[lang]
	} else if len(layout) == 0 {
		return time.Time{}, fmt.Errorf("没有找到默认语言的日期布局,请传入一个自定义布局")
	}
	// 只选取第一个, 为了默认值
	if len(layout) > 0 {
		lastLayout = layout[0]
	}
	return time.Parse(lastLayout, value)
}

func regexReplaceVal(reg *regexp.Regexp, val string, lm []string) (string, error) {
	mcs := reg.FindStringSubmatch(val)
	if mcs != nil && len(mcs) == 3 {
		// 获取索引
		idx, err := strconv.ParseInt(mcs[2], 10, 32)
		if err != nil {
			return "", err
		}
		return reg.ReplaceAllString(val, lm[idx]), nil
	}
	return "", nil
}
