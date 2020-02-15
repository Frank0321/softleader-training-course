# Fluent Bit

OpenShift 預設的 [Container Logs Aggregator](https://docs.openshift.com/container-platform/3.11/install_config/aggregate_logging.html#aggregated-fluentd) 是 EFK:

- **Elasticsearch** - Store, Search, and Analyze
- **Fluentd** - Log Collector, Processor, and Aggregator
- **Kibana** - Explore, Visualize, and Share

[Fluent Bit](https://fluentbit.io/) 跟 [Fluentd](https://www.fluentd.org/) 都是 [Treasure Data](https://www.treasuredata.com/opensource/) 的產品, 而 Fluent Bit 像是一個輕量級的 Fluentd, 專門處理 Log Collector 及 Processor, 而有較少的 Log Aggregator 功能

跟 Fluentd 相比效能更好, 不論 CPU 跟記憶體的使用率 Fluent Bit 都有更好的表現, 因此在公司轉換 OpenShfit 之前, 我們將使用 Fluent Bit 來取代 Logstach 來收集公司微服務中的 Log

> 延伸閱讀: [Fluent-bit rocks](https://gist.github.com/StevenACoffman/4e267f0f60c8e7fcb3f77b9e504f3bd7)

## Getting Started

Fluent Bit 在公司的配置中有制定一些規則, 只有符合規則的 App log 才會被收集, 你可以在這邊 [公司 fluent-bit 的配置](https://github.com/softleader/log-aggregation) 查看是怎麼配置的

此篇接下來會介紹及建議在一個 Spring Boot App 中要怎麼配置才能符合規則!

### pom.xml

我們建議使用 [logstash/logstash-logback-encoder](https://github.com/logstash/logstash-logback-encoder) 來做輸出 JSON 的 encoder, 因此在 `pom.xml` 中請加上:

```xml
<dependency>
  <groupId>net.logstash.logback</groupId>
  <artifactId>logstash-logback-encoder</artifactId>
  <version>${logstash-logback-encoder.version}</version>
</dependency>
```
> 請從 [Logstash Logback Encoder - Maven Repository](https://mvnrepository.com/artifact/net.logstash.logback/logstash-logback-encoder) 選擇版本

### logback.xml

在 `logback.xml` 中, 我們先加上以下 JSON appender, 並且設定 `app_name` 來自於 properties 中的 `spring.application.name`

```xml
<springProperty scope="context" name="appName" source="spring.application.name"/>

<appender class="ch.qos.logback.core.ConsoleAppender" name="JSON">
  <encoder class="net.logstash.logback.encoder.LoggingEventCompositeJsonEncoder">
    <providers>
      <mdc/>
      <timestamp>
        <timeZone>Asia/Taipei</timeZone>
      </timestamp>
      <logLevel/>
      <threadName/>
      <loggerName>
        <fieldName>class_name</fieldName>
      </loggerName>
      <message/>
      <pattern>
        <pattern>
        {
        "app_name": "${appName:-}",
        "pid": "${PID:-}"
        }
        </pattern>
      </pattern>
      <stackTrace>
        <throwableConverter class="net.logstash.logback.stacktrace.ShortenedThrowableConverter">
          <maxDepthPerThrowable>30</maxDepthPerThrowable>
          <!-- 最多帶 10k 的 stack trace, 超過應該也沒用了 -->
          <maxLength>10240</maxLength>
          <shortenedClassNameLength>20</shortenedClassNameLength>
          <!-- 將 root cause 擺最前面, 避免超過 10k 被 trim 掉 -->
          <rootCauseFirst>true</rootCauseFirst>
        </throwableConverter>
      </stackTrace>
    </providers>
  </encoder>
</appender>
```

> 延伸閱讀: [Providers 的配置說明](https://github.com/logstash/logstash-logback-encoder#providers-for-loggingevents)

最後在 root 的地方, 建議透過指定 Spring Profile 才啟動開 JSON appender, 以下範例是當啟動 spring profile `efk`  時, 就停用所有其他的 appender 只開啟 json appender

```xml
<root level="...">
  <springProfile name="!efk">
    <appender-ref ref="STDOUT" />
    <appender-ref ref="FILE" />
    ...
  </springProfile>
  <springProfile name="efk">
    <appender-ref ref="JSON" />
  </springProfile>
  ...
</root>
```
