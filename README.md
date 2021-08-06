[![Build Status](https://travis-ci.org/softleader/softleader-training-course.svg?branch=master)](https://travis-ci.org/softleader/softleader-training-course)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/softleader/softleader-training-course/graphs/commit-activity)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

# 松凌科技教育訓練專區

[http://bit.ly/softleader-training-course](http://bit.ly/softleader-training-course)

<img src="./qr-code.svg" width="500">

![](./training.png)

## Directory Structure

請講師將上課教材，依照下述目錄結構擺放, **課程名稱不可以有空白!**

```sh
.
└── <年度>
    └── <季>
        ├── <課程A>
        ├── <課程B>
        └── ...
```

如:

```sh
.
└── 2016
    ├── Q1
    │   ├── spring-transaction
    │   └── stream-and-lambda
    └── Q2
        ├── design-patterns
        └── ...
```

在你的課程目錄中, 如果有放一個 *README.md*, 會自動擷取該檔案中第一個 header (`#`, `##` ... `######`) 作為課程標題

```sh
.
└── <年度>
    └── <季>
        └── <課程A>
             └── README.md
```

## Courses

> Auto-generated by *toc*

#### 2021
- Q1 - [base-jre-image](/2021/Q1/base-jre-image) - JRE Base Image
- Q1 - [springboot24-config](/2021/Q1/springboot24-config) - Spring Boot 2.4 Externalized Configuration
- Q2 - [java-io](/2021/Q2/java-io) - Java-IO
- Q2 - [react-crud](/2021/Q2/react-crud) - React CURD
- Q2 - [spring-boot-demo](/2021/Q2/spring-boot-demo)
- Q2 - [vue-v2](/2021/Q2/vue-v2) - vue2-training-course
- Q3 - [build-docker-image](/2021/Q3/build-docker-image)
- Q3 - [spring-boot(jakarta99)](/2021/Q3/spring-boot(jakarta99)) - Spring-boot 學習手冊(by jakarta99)
- Q3 - [spring-web-jpa](/2021/Q3/spring-web-jpa) - Spring-MVC

#### 2020
- Q1 - [fluent-bit](/2020/Q1/fluent-bit) - Fluent Bit

#### 2019
- Q1 - [microservice](/2019/Q1/microservice) - Microservice workshop
- Q1 - [react-redux](/2019/Q1/react-redux) - Available Scripts
- Q2 - [a-beginners-guide-to-kubernetes](/2019/Q2/a-beginners-guide-to-kubernetes) - A Beginner’s Guide to Kubernetes
- Q4 - [mapstruct](/2019/Q4/mapstruct) - MapStruct - Java bean mappings [WIP]

#### 2018
- Q1 - [ECMAScript6](/2018/Q1/ECMAScript6) - ECMAScript 6
- Q1 - [logging](/2018/Q1/logging)
- Q1 - [softleader-ldap](/2018/Q1/softleader-ldap) - LDAP 基本介紹
- Q1 - [stream-adv](/2018/Q1/stream-adv) - 2018Q1 Stream
- Q2 - [a-h](/2018/Q2/a-h) - A&H
- Q2 - [docker](/2018/Q2/docker) - 手把手帶你學 Docker
- Q2 - [exception-handling](/2018/Q2/exception-handling) - Best practices to handle exceptions
- Q2 - [microservices](/2018/Q2/microservices)
- Q2 - [tensorflow](/2018/Q2/tensorflow)
- Q3 - [Implements-SSO-with-Spring-Session](/2018/Q3/Implements-SSO-with-Spring-Session)
- Q3 - [building-docker-images](/2018/Q3/building-docker-images) - Building Docker Images
- Q3 - [jmeter](/2018/Q3/jmeter) - Jmeter功能介紹
- Q3 - [microservice-with-springcloud](/2018/Q3/microservice-with-springcloud)
- Q3 - [tensorflow](/2018/Q3/tensorflow) - Getting Started with TensorFlow

#### 2017
- Q1 - [redux](/2017/Q1/redux)
- Q1 - [regex](/2017/Q1/regex) - 一上來就是最終Boss關的Regex
- Q1 - [web-security-protection](/2017/Q1/web-security-protection)
- Q1 - [webpack](/2017/Q1/webpack) - Webpack-Demo
- Q4 - [softleader-day](/2017/Q4/softleader-day)

#### 2016
- Q1 - [spring-transaction](/2016/Q1/spring-transaction) - AOP基本概念
- Q1 - [stream-and-lambda](/2016/Q1/stream-and-lambda)
- Q2 - [design-patterns](/2016/Q2/design-patterns)
- Q2 - [ireport](/2016/Q2/ireport)
- Q2 - [react](/2016/Q2/react)
- Q2 - [zookeeper](/2016/Q2/zookeeper)
- Q3 - [APS](/2016/Q3/APS)
- Q3 - [JPA](/2016/Q3/JPA)
- Q3 - [RabbitMQ](/2016/Q3/RabbitMQ)
- Q3 - [react-workshop](/2016/Q3/react-workshop)
- Q4 - [softleader-day](/2016/Q4/softleader-day)

## GitHub Action

本 Repo 設有 [GitHub Action](./.github/workflows) 在 Push 或 PR 時自動的觸發 *toc*
