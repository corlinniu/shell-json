# 介绍
在Shell中分析Json、Csv、Excel等等

# 命令介绍：json
## Usage
```
cat xxxx.json | json "字符串分隔符" "字符串" ["字符串" ["JsonPath"]]
JsonPath: 可以百度搜一下，一般写法是 $.xxx 表示取json文件中xxx字段的值
说明：上面的命令只会翻译JsonPath进行替换，其他字符串原样输出
```
## 示例
```
> cat raw.txt | json "" "update " $.table "set name = '" $.name "' where id = " $.id
update qx_userset name = 'xxxx' where id = 1;
update qx_userset name = 'xxxx' where id = 1;
```
# 命令介绍：json2
## Usage
```
cat xxxx.json | json2 "update user set name = '{name}' where id = {id}"
```
## 示例
```
> cat raw.txt | json2 "update {table} set name = '{name}' where id = {id};"
update qx_user set name = 'xxxx' where id = 1;
update qx_user set name = 'xxxx' where id = 1;
```

# 命令介绍：csv
## Usage
```azure
csv data/data.csv "ID = {ID} Name = {姓名} Age = {年龄}"
```
## 示例
```azure
> bin/csv data/data.csv "ID = {ID} Name = {姓名} Age = {年龄}"
rst = ID = 1 Name = xxxx Age = 32
rst = ID = 2 Name = yyyy Age = 45
```