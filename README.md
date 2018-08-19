# Go-Spider


> Golang 爬虫上手指南



### 1. 基本思路

- 清晰需要的内容
- 分析网页
- 获取网页信息
- 解析网页信息

### 2. 分析网页

- Chrome 浏览器审查元素，查看网页源代码

### 3. 网页响应值的类型

- json: 一般是调用的API，比较好分析，解析json 数据即可
- xml: 不常见
- html: 常见，使用正则表达式、CSS 选择器、XPATH 获取需要的内容

### 4. 请求的类型

- Get : 常见，直接请求即可
- Post : 需要分析请求的参数，构造请求，向对方服务器端发送请求，再解析响应值

### 5. 请求头部信息

- Uer-Agent 头部信息

### 6. 存储

- 本地: text、json、csv
- 数据库: 关系型（postgres、MySQL）, 非关系型（mongodb）, 搜索引擎（elasticsearch） 

### 7. 图片处理

- 请求
- 存储

### 8. 其他

- 代理: ip 池
- User-Agent: 模拟浏览器
- APP:  APP 数据需要使用抓包工具：Mac(Charles)、Windows(Fiddler)(分析出Api)


### 9. 难点

- 分布式
- 大规模抓取


>  本教程讲述前8节
>  演示使用 Mac, IDE 选择的是 Goland




## 各章节目录

- [中国票房](domain/cbooo/cbo.md)
- [中影指数](domain/chinafilm/chinafilm.md)
- [懂球帝](domain/dongqiudi/dongqiudi.md)
- [GithubTrending](domain/githubtrending/githubtrending.md)
- [古诗文](domain/gushiwen/gushiwen.md)
- [猫眼票房](domain/maoyan/maoyan.md)
- [糯米票房](domain/nuomi/nuomi.md)
- [Pexels图片社区](domain/pexels/pexels.md)
- [全球票房排行榜](domain/piaofang/piaofang.md)

## 几大要点

**如何获取网页源代码**


- 原生 net/http
- gorequest (基于原生的net/http 封装)

**Web客户端请求方法**

- Get 绝大多少数
- Post


**Web服务端响应**

- json
- html

**Web服务端响应的处理方式**

- json: 使用原生的json 序列化，或者使用 gjson （第三方）
- html: 正则表达式、 Css 选择器、Xpath

**存储数据方式**

- Text
- Json
- Csv
- db

前三种，涉及文件读写；最后者涉及数据库操作

**展示方式**

- 图表

**最终目的**

- 数据分析
