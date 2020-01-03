# Harmony * Switch
Harmony * Switch运行在安卓上的Switch比价软件的后端支撑系统。

## API
### GameInfo
#### Detail
##### Request
Get /info/detail/:id
##### Response
```json
{
	"ID": 0,
	"GameNameCN": "塞尔达传说：旷野之息",
	"GameNameEN": "The Legend of Zelda: Breath of the Wild",
	"GameSize": "13.8GB",
	"ImgUrl": ["http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192359.jpg", "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192403.jpg", "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192411.jpg", "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192424.jpg", "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192439.jpg", "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192444.jpg"],
	"Price": [{
		"Region": "美国",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 417.57,
		"PriceLocal": 59.99,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "免税州邮编99613",
		"HistoryLowestPrice": 292.28
	}, {
		"Region": "澳大利亚",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 437.9,
		"PriceLocal": 89.95,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 306.46
	}, {
		"Region": "香港",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 445.96,
		"PriceLocal": 499,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "购买网址store.nintendo.com.hk",
		"HistoryLowestPrice": 0
	}, {
		"Region": "加拿大",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 450.37,
		"PriceLocal": 83.99,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "魁北克省外Alberta，5%税率最低",
		"HistoryLowestPrice": 315.24
	}, {
		"Region": "日本",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 491.01,
		"PriceLocal": 7678,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 0
	}, {
		"Region": "瑞典",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 512.34,
		"PriceLocal": 689,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 311.57
	}, {
		"Region": "挪威",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 538.18,
		"PriceLocal": 679,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 326.55
	}, {
		"Region": "南非",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 538.21,
		"PriceLocal": 1079,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "不支持PayPal",
		"HistoryLowestPrice": 296.29
	}, {
		"Region": "欧盟",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 545.64,
		"PriceLocal": 69.99,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "所有欧元使用国，这里以德国为例",
		"HistoryLowestPrice": 381.93
	}, {
		"Region": "英国",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 551.53,
		"PriceLocal": 59.99,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 386.04
	}, {
		"Region": "墨西哥",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 589.71,
		"PriceLocal": 1599,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 412.8
	}, {
		"Region": "俄罗斯",
		"HasChinese": true,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 593.14,
		"PriceLocal": 5249,
		"GameLanguage": ["中文", "日语", "英语", "西班牙语", "法语", "德语", "意大利语", "荷兰语", "俄语"],
		"ReleaseDate": "2017-03-03",
		"RegionComment": "",
		"HistoryLowestPrice": 387.48
	}],
	"LowestPrice": {
		"Region": "",
		"HasChinese": false,
		"IsLowestPrice": false,
		"IsOnSale": false,
		"SaleRate": 0,
		"EndDate": 0,
		"PriceCNY": 0,
		"PriceLocal": 0,
		"GameLanguage": null,
		"ReleaseDate": "",
		"RegionComment": "",
		"HistoryLowestPrice": 0
	},
	"LanguageTag": [0],
	"HasSolidEdition": true,
	"HasDemo": false,
	"IsExclusive": true,
	"GameTypeTag": ["冒险", "动作", "角色扮演"],
	"HasChinese": true,
	"GamePlayers": 1,
	"Description": "游戏变成了完全的开放地图，沙盒式的玩法拓展出了更多的内容，林克可以爬山、游泳，而且游戏对应外观变化，玩家在游戏过程中会找到不止一件的服装以及武器。而在大地图中也会找到更多可以拾取的道具。服装除了外观外也有着属性的设定，例如林克在寒冷的地方穿薄衣服就会哆嗦，需要换上厚一点的服装才会变正常。在地图上的互动要素也不少，玩家手上拿着火把就能把草堆点燃，拿着斧头就能砍树等。而且不同的武器也会反馈出不用的攻击动作，某些武器还能当作投掷道具扔出去。游戏的战斗中还有类似“子弹时间”的设定，当敌人漏出过大的破绽时就可以采取更有效的攻击，但具体发动条件未知。林克还拥有一个魔法炸弹一样的道具可以使用，而且可以改变其用法形态，通过炸弹可以炸掉矿点来获得素材道具，不过魔法炸弹每次使用之后都要经过冷却时间。林克还可以从高空滑翔至低处，不过滑翔也和攀爬一样需要消耗耐力。或者可以选择把盾牌当做滑板从高处滑下。",
	"GameScore": ["IGN:10", "MC:97"]
}
```

#### Summary
##### Request
GET /info/summary/:id
##### Response
```json

```