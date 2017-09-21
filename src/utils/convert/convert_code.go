package convert

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"

	log "github.com/cihub/seelog"
)

func DecodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		log.Error(err.Error())
		return text, err
	}

	return string(dst[:nDst]), nil
}

func GbkToUtf8(str string) (string, error) {
	s := []byte(str)
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Error(e.Error())
		return "", e
	}
	return string(d), nil
}

func Utf8ToGbk(str string) (string, error) {
	s := []byte(str)
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		log.Error(e.Error())
		return "", e
	}
	return string(d), nil
}
