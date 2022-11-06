### 介绍
在Shell中分析Json、Csv、Excel等等

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