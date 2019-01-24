
package main


import ( "fmt"

_"net/http/pprof" "sync"

// "github.com/hpcloud/tail" "github.com/prometheus/client_golang/prometheus" "github.com/prometheus/common/log" "github.com/prometheus/common/version"

// //"net" "net/http"

// "net/url"

// "errors" "flag"

"os" "os/exec" "io/ioutil"

// "regexp"

// "strconv"

// "strings" "time"

)


const (

namespace = "my_exporter"

)


type Exporter struct { cmd string

mutex sync.RWMutex up,freeRam,TotalRam prometheus.Gauge Debug bool

upstreamResponseTimes *prometheus.HistogramVec


}


func (e *Exporter) scrape() {


log.Infoln("--------") e.up.Set(1)

}


func (e *Exporter) Describe(ch chan<- *prometheus.Desc) { e.upstreamResponseTimes.Describe(ch)

}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

e.m utex.Lock() // To protect metrics from concurrent collects defer e.mutex.Unlock()


e.scrape() e.upstreamResponseTimes.Collect(ch) e.up.Collect(ch)

}


func NewExporter(cmd string, timeout time.Duration) (*Exporter, error) {

return &Exporter{ cmd: cmd,

up: prometheus.NewGauge(prometheus.GaugeOpts{ Namespace: namespace,

Name: "up",

Help: "Was the last scrape of nginx exporter successful.",

}),

upstreamResponseTimes: prometheus.NewHistogramVec( prometheus.HistogramOpts{

Namespace: namespace,

Name: "http_upstream_request_times",

Help: "Response times from backends/upstreams",

},

[]string{"method", "endpoint", "response_code", "client_type"},

),

}, nil

}


func run() {


cmd := exec.Command("/bin/sh", "-c", `free -m `) stdout , err := cmd.StdoutPipe()

if err != nil { panic(err.Error())

}


if err != nil {

fmt.Println("Stdoutpipe error", err.Error())

}


stderr , err := cmd.StderrPipe()


if err != nil {

fmt.Println("Stderrpipe error", err.Error())

}


if err := cmd.Start(); err != nil { fmt.Println("error start", err )


}


bytesErr, err := ioutil.ReadAll(stderr)

if err != nil {

fmt.Println("ReadAll stderr: ", err.Error()) return

}


if len(bytesErr) != 0 {

fmt.Printf("stderr is not nil: %s", bytesErr) return

}


bytes, err := ioutil.ReadAll(stdout) if err != nil {

fmt.Println("ReadAll stdout: ", err.Error()) return

}


if err := cmd.Wait(); err != nil { fmt.Println("Wait: ", err.Error()) return

}


fmt.Printf("stdout: %s", string(bytes))


}


func main() { var (

listenAddress = flag.String("web.listen-address", ":9101", "Address to listen on for web interface and telemetry.")

metricsPath = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")

cmd = flag.String("free", "", "command to run")

nginxTimeout = flag.Duration("nginx.timeout", 5*time.Second, "Timeout for trying to get stats from HA.")

showVersion = flag.Bool("version", false, "Print version information.")

// debug = flag.Bool("debug", false, "Print debug information.")

)

flag.Parse()


if *showVersion {

fmt.Println(os.Stdout, version.Print("shannon_exporter")) os.Exit(0)

}


log.Infoln("Starting shannon_exporter", version.Info()) log.Infoln("Build context", version.BuildContext())


run()


exporter, _ := NewExporter(*cmd, *nginxTimeout)


prometheus.MustRegister(exporter) prometheus.MustRegister(version.NewCollector("nginx_exporte r"))

http.Handle(*metricsPath, prometheus.Handler()) http.HandleFunc("/", func(w http.ResponseWriter, r

*http.Request) { w.Write([]byte(`<html>

<head><title>Haproxy Exporter</title></head>

<body>

<h1>Haproxy Exporter</h1>

<p><a href='` + *metricsPath + `'>Metrics</a></p>

</body>

</html>`))

})


log.Fatal(http.ListenAndServe(*listenAddress, nil))


}