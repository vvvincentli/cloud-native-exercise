

# 1. 为 HTTPServer 添加 0-2 秒的随机延时

```go

// Healthz ,
func Healthz(w http.ResponseWriter, r *http.Request) {
	for k, _ := range r.Header {
		w.Header().Set(k, r.Header.Get(k))
	}
	time.Sleep(time.Duration(rand.Intn(2) * int(time.Second)))
	w.Header().Set("VERSION", os.Getenv("VERSION"))
	w.Write([]byte(fmt.Sprintf("i'm alive. [%s]", time.Now().Format(time.RFC3339))))
}
```



# 2.为 HTTPServer 项目添加延时 Metric

## 2.1.http-server中暴露Prometheus格式的metrics

## 2.2.Pod申明上报指标端口和地址




# 3.将 HTTPServer 部署至测试集群，并完成 Prometheus 配置

# 4.从 Promethus 界面中查询延时指标数据

# 5.（可选）创建一个 Grafana Dashboard 展现延时分配情况