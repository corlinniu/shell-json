### 介绍
在Shell中分析Json、Csv、Excel等等

### 编译安装
```shell
1. 安装golang: https://golang.google.cn/dl/
2. 配置golang环境变量
3. 执行编译安装命令: sh build.sh
```

### json
```shell
echo '{"name":1, "age":2}' | bin/json 'name = {name} age = {age}'
```

### excel
```shell
bin/excel data/excel.xlsx Sheet1 'Name = {Name} Age = {Age}'
```

### csv
```shell
bin/csv data/csv.csv 'Name = {Name} Age = {Age}'
```

### 其他好用的命令整理
#### join
```shell
# 类似于SQL中的JOIN

cat employee.txt
100 Jason Smith
200 John Doe
300 Sanjay Gupta
400 Ashok Sharma

cat bonus.txt
200 $500
300 $3,000
400 $1,250

# 左连接
join -a1 employee.txt bonus.txt
100 Jason Smith
200 John Doe $500
300 Sanjay Gupta $3,000
400 Ashok Sharma $1,250

# 右连接
join -a2 employee.txt bonus.txt
200 John Doe $500
300 Sanjay Gupta $3,000
400 Ashok Sharma $1,250
```