# 移动客户端数据库

public_mobile_data

## 取校验值（检查更新）

```
GET /API/v1_0/Products/public_mobile_data('{{language}}')?
  api_key={{api_key}}&
  format={{format}}&
  select={{select}}&
  {{timestamp}}
```

| 参数      | 描述            |
|-----------|-----------------|
| language  | 语言 `zh-cn`    |
| api_key   | hex             |
| format    | 固定 `json`     |
| select    | 固定 `checksum` |
| timestamp | unix时间戳 毫秒 |

注意最后一个参数不是键值对，直接是时间戳的值

示例

```
GET /API/v1_0/Products/public_mobile_data('zh-cn')?api_key=00112233445566778899AABBCCDDEEFF&$format=json&$select=checksum&1589130887710 HTTP/1.1
```

返回

```
{
	"d": {
		"__metadata": {
			"id": "https://odata.intel.com/API/v1_0/Products/public_mobile_data('zh-cn')",
			"uri": "https://odata.intel.com/API/v1_0/Products/public_mobile_data('zh-cn')",
			"type": "Intel.OData.v1_0.public_mobile_data"
		},
		"checksum": "683c87bbdf6db1e1e4aa51d892cf493acf7edfe8488394299dd0e5c0d61fffe6"
	}
}
```

校验和怎么计算，目前还不清楚

## 获取移动客户端数据库（sqlite）

```
GET /API/v1_0/Products/public_mobile_data('{{language}}')?
  api_key={{api_key}}&
  format={{format}}&
  {{timestamp}}
```

| 参数      | 描述            |
|-----------|-----------------|
| language  | 语言 `zh-cn`    |
| api_key   | hex             |
| format    | 固定 `json`     |
| timestamp | unix时间戳 毫秒 |

示例

```
GET /API/v1_0/Products/public_mobile_data('zh-cn')?api_key=00112233445566778899AABBCCDDEEFF&$format=json&1589131922417
```

返回

```
{
	"d": {
		"__metadata": {
			"id": "https://odata.intel.com/API/v1_0/Products/public_mobile_data('zh-cn')",
			"uri": "https://odata.intel.com/API/v1_0/Products/public_mobile_data('zh-cn')",
			"type": "Intel.OData.v1_0.public_mobile_data"
		},
		"language": "zh-cn",
		"checksum": "683c87bbdf6db1e1e4aa51d892cf493acf7edfe8488394299dd0e5c0d61fffe6",
		"data_bin": ""
	}
}
```

`data_bin`中为base64编码的字符串，解码后是sqlite数据库文件

## 目前已知可用的语言

```
https://odata.intel.com/API/v1_0/Products/public_mobile_data?$select=checksum,language
```

* de
* en
* es
* fr
* it
* ja
* ko
* pl
* pt-br
* ru
* zh-cn
* zh-tw
