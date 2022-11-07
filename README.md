### 介绍
在Shell中分析Json、Csv、Excel等等

### 编译安装
```shell
1. 需要安装golang
2. 执行编译安装命令: sh build.sh
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