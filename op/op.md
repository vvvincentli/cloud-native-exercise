

# 1. 为 HTTPServer 添加 0-2 秒的随机延时

> http-server\pkg\controller\healthz.go 
```go

// delay time 
func intn(min, max int) int {
	rand.Seed(time.Now().Unix())//reset seed
	return min + rand.Intn(max-min)
}


// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	tm := metrics.NewTimer() //新建指标
	defer tm.ObserveTotal()//
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	time.Sleep(time.Duration(intn(50, 2000) * int(time.Microsecond)))
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Write([]byte(fmt.Sprintf("i'm alive. [%s]", time.Now().Format(time.RFC3339))))
}

```



# 2.为 HTTPServer 项目添加延时 Metric


## 2.1.http-server中暴露Prometheus格式的metrics
> 新建指标
``` go
package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	NameSpace = "http-server"
)

var (
	functinLatency = CreateExecutionTimeMetric(NameSpace, "time spent.")
)

type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functinLatency)
}

func NewExecutionTimer(h *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: h,
		start: now,
		last:  now,
	}
}

func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}

```


> 暴露接口，http-server\middleware\router.go

```go
// route binding
func RouterBinding() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", withLogging(controller.Healthz))
	mux.HandleFunc("/error", withLogging(controller.Error))
	mux.Handle("/metrics", promhttp.Handler)
	return mux
}


```
> main函数中注册采集指标，

```go

func main() {
	defer shutdown()

	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := ""
	flag.StringVar(&configFile, "app-config", "", "application config file path.")
	flag.Parse()
	if configFile == "" {
		configFile = os.Getenv("app-config")
	}

	log.Println(fmt.Sprintf("LoadConfig %s", configFile))
	if err := configs.LoadConfig(configFile); err != nil {
		log.Println(fmt.Sprintf("loadConfig %s failed, error:%s", configFile, err))
		os.Exit(-1)
	}
	c := configs.GetConfig()

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err := metrics.Register()  //注册采集指标
	if err != nil {
		errs <- err
	}

	hostServer(c.App.Host, c.App.Port, errs)
	fmt.Println(fmt.Sprintf("exit: %v", <-errs))
}
```



## 2.2.Pod申明上报指标端口和地址

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    keel.sh/pollSchedule: '@every 30m'
    prometheus.io/port: http-metrics
    prometheus.io/scrap: "true"
  labels:
    app: http-server
    keel.sh/approvals: "1"
    keel.sh/policy: patch
    keel.sh/trigger: poll
  name: http-server
  namespace: demo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: http-server
  template:
    metadata:
      annotations:
        prometheus.io/scrape: "true"
        prometheus.io/port: "8080"
      labels:
        app: http-server

#其他配置不变
```


# 3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置
> 重新打镜像，部署


# 4.从 Promethus 界面中查询延时指标数据

# 5.（可选）创建一个 Grafana Dashboard 展现延时分配情况