package date_parser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParserLangDate(t *testing.T) {
	tval, _ := time.Parse("2006-01-02", "2018-09-13")

	eval, _ := ParserLangDate("en", "on September 13 2018", "on January 02 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("uk", "on 13 September 2018", "on 02 January 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("au", "on 13 September 2018", "on 02 January 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("ca", "on 13 September 2018", "on 02 January 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("in", "on 13 September 2018", "on 02 January 2006")
	assert.Equal(t, tval, eval)

	tde, _ := time.Parse("2006-01-02", "2018-12-26")
	eval, _ = ParserLangDate("de", "am 26. Dezember 2018", "am 02. January 2006")
	assert.Equal(t, tde, eval)

	eval, _ = ParserLangDate("jp", "on 2018年09月13日", "on 2006年January02日")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("fr", "le 13 septembre 2018", "le 02 January 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("it", "il 13 settembre 2018", "il 02 January 2006")
	assert.Equal(t, tval, eval)

	eval, _ = ParserLangDate("es", "el 13 septiembre 2018", "el 02 January 2006")
	assert.Equal(t, tval, eval)

	tnl, _ := time.Parse("2006-01-02", "2018-03-13")
	eval, _ = ParserLangDate("nl", "op 13 maart 2018", "op 02 January 2006")
	assert.Equal(t, tnl, eval)

	tbr, _ := time.Parse("2006-01-02", "2018-08-09")
	eval, _ = ParserLangDate("br", "em 9 de agosto de 2018", "em 2 de January de 2006")
	assert.Equal(t, tbr, eval)

	tmx, _ := time.Parse("2006-01-02", "2018-12-05")
	eval, _ = ParserLangDate("mx", "el 5 de diciembre de 2018", "el 2 de January de 2006")
	assert.Equal(t, tmx, eval)

	// ----------- amazon 还未开启 ---------

	eval, _ = ParserLangDate("ko", "2018년9월13일", "2006년January2일")
	assert.Equal(t, tval, eval)
}

func TestSigleParse(t *testing.T) {
	tbr, _ := time.Parse("2006-01-02", "2018-08-09")
	eval, _ := ParserLangDate("br", "em 9 de agosto de 2018", "em 2 de January de 2006")
	assert.Equal(t, tbr, eval)
}

