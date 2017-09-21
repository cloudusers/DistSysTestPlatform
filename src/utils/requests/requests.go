// Package requests implements functions to manipulate requests.
package requests

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	log "github.com/cihub/seelog"
)

const (
	// Default quantity of retries
	Retries = 3
	// Default timeout is 30 seconds
	Timeout = 1800 * time.Second
)

// NewRequest returns an Request with exponential backoff as default.
func NewRequest(method, urlStr string, body io.Reader) (*Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	return &Request{
			req, Retries, Timeout, backoff.NewExponentialBackOff()},
		nil
}

// Request type.
type Request struct {
	*http.Request
	retry   int
	timeout time.Duration
	backoff *backoff.ExponentialBackOff // Default Type of backoff.
}

// Set the amount of retries
func (r *Request) Retries(times int) *Request {
	r.retry = times
	return r
}

// Timeout specifies a time limit for requests made by the Client.
// A Timeout of zero means no timeout.
func (r *Request) Timeout(t time.Duration) *Request {
	r.timeout = t
	return r
}

// New Client with timeout
func (r *Request) newClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &http.Client{Timeout: r.timeout, Transport: tr}
}

func PostRequest(host string,
	port string,
	handler string,
	body string) bool {

	host = fmt.Sprintf("http://%s:%s/%s", host, port, handler)
	//fmt.Println(host)

	var ret bool = false

	req, err := NewRequest("POST",
		host,
		strings.NewReader(body))
	if err != nil {
		log.Error(err.Error())
	} else {
		//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Charset", "utf-8")
		res, err := req.Do()
		if err != nil {
			log.Error(err.Error())
		} else {
			ret = true
			if nil != res {
				defer res.Body.Close()
				_, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Error(fmt.Sprintf("read response failed! err:", err.Error()))
				}
			}
		}
	}

	return ret
}
