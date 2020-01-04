# Harmony * Switch
Harmony * Switch运行在安卓上的Switch比价软件的后端支撑系统。

## API
### GameInfo
#### Paging
##### Request
GET /game?limit=24&offset=0
#### Response
```json
{
    "List": [
        {
            "ID": 1,
            "NameCN": "你裁我剪！斯尼帕",
            "NameEN": "Snipperclips: Cut It Out, Together",
            "ImgUrl": "http://media.eshop-switch.com/image_s/568AB7D675964B1540E265A5E2A2EABF.jpg",
            "Region": "日本",
            "PriceCNY": 117.28,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": true,
            "IsExclusive": true
        },
        {
            "ID": 2,
            "NameCN": "塞尔达传说：旷野之息",
            "NameEN": "The Legend of Zelda: Breath of the Wild",
            "ImgUrl": "http://media.eshop-switch.com/image_s/5766ECBAD6A26F5B0DC69FDB5E57C85E.jpg",
            "Region": "美国",
            "PriceCNY": 417.57,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": true,
            "HasDemo": false,
            "IsExclusive": true
        },
        {
            "ID": 3,
            "NameCN": "超级炸弹人R",
            "NameEN": "SUPER BOMBERMAN R",
            "ImgUrl": "http://media.eshop-switch.com/image_s/36A1CA95016E98EE395B8FBD183358D0.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 109.54,
            "SaleRate": 50,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": true,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 4,
            "NameCN": "1 2 Switch",
            "NameEN": "1-2-Switch",
            "ImgUrl": "http://media.eshop-switch.com/image_s/BD8D4B8483B8CA7716C46159F97F4E2E.jpg",
            "Region": "香港",
            "PriceCNY": 320.84,
            "SaleRate": 100,
            "LanguageTag": null,
            "HasSolidEdition": true,
            "HasDemo": false,
            "IsExclusive": true
        },
        {
            "ID": 5,
            "NameCN": "合金弹头3",
            "NameEN": "ACA NEOGEO METAL SLUG3",
            "ImgUrl": "http://media.eshop-switch.com/image_s/0B5103272BB1040ED7787F5E4EC9BB96.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 6,
            "NameCN": "黑白棋",
            "NameEN": "Othello",
            "ImgUrl": "http://media.eshop-switch.com/image_s/A9AD78A392FE839E8E0A5E1323C08A4A.jpg",
            "Region": "日本",
            "PriceCNY": 15.99,
            "SaleRate": 51,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 7,
            "NameCN": "拳皇98",
            "NameEN": "ACA NEOGEO THE KING OF FIGHTERS 98",
            "ImgUrl": "http://media.eshop-switch.com/image_s/5C618726064524F292ECF8265D8E3FA3.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 8,
            "NameCN": "兰空",
            "NameEN": "VOEZ",
            "ImgUrl": "http://media.eshop-switch.com/image_s/78D4585B0C7D870A80FA1F9BD2EEFBE1.jpg",
            "Region": "挪威",
            "PriceCNY": 126.82,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": true,
            "IsExclusive": false
        },
        {
            "ID": 9,
            "NameCN": "新大开拓时代：街道制造",
            "NameEN": "New Frontier Days: Founding Pioneers",
            "ImgUrl": "http://media.eshop-switch.com/image_s/807CF801BE1EEE1E146FEB345CC08701.jpg",
            "Region": "日本",
            "PriceCNY": 38.37,
            "SaleRate": 51,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 10,
            "NameCN": "噗哟噗哟俄罗斯方块S",
            "NameEN": "Puyo Puyo Tetris",
            "ImgUrl": "http://media.eshop-switch.com/image_s/5C249A604D97CB39B7110D9EB4F9BF19.jpg",
            "Region": "美国",
            "PriceCNY": 208.75,
            "SaleRate": 100,
            "LanguageTag": [
                1,
                5
            ],
            "HasSolidEdition": true,
            "HasDemo": true,
            "IsExclusive": false
        },
        {
            "ID": 11,
            "NameCN": "超惑星战记Zero",
            "NameEN": "BLASTER MASTER ZERO",
            "ImgUrl": "http://media.eshop-switch.com/image_s/6B5B842BE59B361F7369F067C92C957F.jpg",
            "Region": "日本",
            "PriceCNY": 31.34,
            "SaleRate": 50,
            "LanguageTag": null,
            "HasSolidEdition": false,
            "HasDemo": true,
            "IsExclusive": false
        },
        {
            "ID": 12,
            "NameCN": "魔界战记5：完整版",
            "NameEN": "Disgaea 5 Complete",
            "ImgUrl": "http://media.eshop-switch.com/image_s/02074D793EBBE9043E1C505E0F82672F.jpg",
            "Region": "美国",
            "PriceCNY": 139.14,
            "SaleRate": 50,
            "LanguageTag": [
                5
            ],
            "HasSolidEdition": true,
            "HasDemo": true,
            "IsExclusive": false
        },
        {
            "ID": 13,
            "NameCN": "勇者斗恶龙英雄I+II",
            "NameEN": "Dragon Quest Heroes I+II",
            "ImgUrl": "http://media.eshop-switch.com/image_s/C43AD1440847387A2682D1CDB52B6640.jpg",
            "Region": "日本",
            "PriceCNY": 371.42,
            "SaleRate": 40,
            "LanguageTag": null,
            "HasSolidEdition": true,
            "HasDemo": true,
            "IsExclusive": false
        },
        {
            "ID": 14,
            "NameCN": "世界英雄 完美版",
            "NameEN": "ACA NEOGEO WORLD HEROS PERFECT",
            "ImgUrl": "http://media.eshop-switch.com/image_s/3F8E31279CE3D1C17B48D788DB6827BE.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 15,
            "NameCN": "火热火热7",
            "NameEN": "ACA NEOGEO WAKU WAKU 7",
            "ImgUrl": "http://media.eshop-switch.com/image_s/9FD048FEFC7F51E857A0EE0751B3D9B4.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 16,
            "NameCN": "突击奇兵",
            "NameEN": "ACA NEOGEO SHOCK TROOPERS",
            "ImgUrl": "http://media.eshop-switch.com/image_s/869DB7DA89DBE89C528F5C21A53F815D.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 17,
            "NameCN": "信长之野望：创造with威力加强版",
            "NameEN": "Nobunaga no Yabou - Souzou with Power-Up Kit",
            "ImgUrl": "http://media.eshop-switch.com/image_s/40D732DF591103E1A61FE560EC4C9498.jpg",
            "Region": "日本",
            "PriceCNY": 689.38,
            "SaleRate": 100,
            "LanguageTag": null,
            "HasSolidEdition": true,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 20,
            "NameCN": "拳皇94",
            "NameEN": "ACA NEOGEO THE KING OF FIGHTERS 94",
            "ImgUrl": "http://media.eshop-switch.com/image_s/2D074BC638FDF278D00E52423888BAEB.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 21,
            "NameCN": "NAM-1975",
            "NameEN": "ACA NEOGEO NAM-1975",
            "ImgUrl": "http://media.eshop-switch.com/image_s/1B7C26E10D5B831725E40E35D36D5DF7.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 22,
            "NameCN": "合金弹头",
            "NameEN": "ACA NEOGEO METAL SLUG",
            "ImgUrl": "http://media.eshop-switch.com/image_s/62C20439133FD26E0BAD781AEEAFF64A.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 23,
            "NameCN": "三国志13威力加强版",
            "NameEN": "San Goku Shi 13 with Power-Up Kit",
            "ImgUrl": "http://media.eshop-switch.com/image_s/FE8097AC59BE3BFC1238F8171680BE01.jpg",
            "Region": "香港",
            "PriceCNY": 382.5,
            "SaleRate": 100,
            "LanguageTag": [
                1
            ],
            "HasSolidEdition": true,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 24,
            "NameCN": "高尔夫大赛",
            "NameEN": "ACA NEOGEO NEO TURF MASTERS",
            "ImgUrl": "http://media.eshop-switch.com/image_s/5422112C4AC562FDB1960996BAD0F758.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 25,
            "NameCN": "阿尔法任务",
            "NameEN": "ACA NEOGEO ALPHA MISSION II",
            "ImgUrl": "http://media.eshop-switch.com/image_s/344FE46373799CF5924AD249AD4FC53C.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 44.3,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": false
        },
        {
            "ID": 26,
            "NameCN": "神巫女",
            "NameEN": "KAMIKO",
            "ImgUrl": "http://media.eshop-switch.com/image_s/EF2A161F53868E83DF4F099A8F2F43E0.jpg",
            "Region": "澳大利亚",
            "PriceCNY": 31.64,
            "SaleRate": 100,
            "LanguageTag": [
                0
            ],
            "HasSolidEdition": false,
            "HasDemo": false,
            "IsExclusive": true
        }
    ]
}
```
#### Summary
##### Request
GET /game/:id/summary
##### Response
```json
{
    "ID": 1,
    "NameCN": "你裁我剪！斯尼帕",
    "NameEN": "Snipperclips: Cut It Out, Together",
    "ImgUrl": "http://media.eshop-switch.com/image_s/568AB7D675964B1540E265A5E2A2EABF.jpg",
    "Region": "日本",
    "PriceCNY": 117.28,
    "SaleRate": 100,
    "LanguageTag": null,
    "HasSolidEdition": false,
    "HasDemo": true,
    "IsExclusive": true
}
```
#### Detail
##### Request
GET /game/:id/detail
##### Response
```json
{
    "ID": 0,
    "GameNameCN": "塞尔达传说：旷野之息",
    "GameNameEN": "The Legend of Zelda: Breath of the Wild",
    "GameSize": "13.8GB",
    "ImgUrl": [
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192359.jpg",
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192403.jpg",
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192411.jpg",
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192424.jpg",
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192439.jpg",
        "http://image.eshop-switch.com/l/%E5%A1%9E%E5%B0%94%E8%BE%BE%E4%BC%A0%E8%AF%B4%EF%BC%9A%E8%8D%92%E9%87%8E%E4%B9%8B%E6%81%AF/20171002192444.jpg"
    ],
    "Price": [
        {
            "Region": "美国",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 417.57,
            "PriceLocal": 59.99,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "免税州邮编99613",
            "HistoryLowestPrice": 292.28
        },
        {
            "Region": "澳大利亚",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 437.9,
            "PriceLocal": 89.95,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 306.46
        },
        {
            "Region": "香港",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 445.96,
            "PriceLocal": 499,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "购买网址store.nintendo.com.hk",
            "HistoryLowestPrice": 0
        },
        {
            "Region": "加拿大",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 450.37,
            "PriceLocal": 83.99,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "魁北克省外Alberta，5%税率最低",
            "HistoryLowestPrice": 315.24
        },
        {
            "Region": "日本",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 491.01,
            "PriceLocal": 7678,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 0
        },
        {
            "Region": "瑞典",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 512.34,
            "PriceLocal": 689,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 311.57
        },
        {
            "Region": "挪威",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 538.18,
            "PriceLocal": 679,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 326.55
        },
        {
            "Region": "南非",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 538.21,
            "PriceLocal": 1079,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "不支持PayPal",
            "HistoryLowestPrice": 296.29
        },
        {
            "Region": "欧盟",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 545.64,
            "PriceLocal": 69.99,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "所有欧元使用国，这里以德国为例",
            "HistoryLowestPrice": 381.93
        },
        {
            "Region": "英国",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 551.53,
            "PriceLocal": 59.99,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 386.04
        },
        {
            "Region": "墨西哥",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 589.71,
            "PriceLocal": 1599,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 412.8
        },
        {
            "Region": "俄罗斯",
            "HasChinese": true,
            "IsLowestPrice": false,
            "IsOnSale": false,
            "SaleRate": 0,
            "EndDate": 0,
            "PriceCNY": 593.14,
            "PriceLocal": 5249,
            "GameLanguage": [
                " 中文",
                "日语",
                "英语",
                "西班牙语",
                "法语",
                "德语",
                "意大利语",
                "荷兰语",
                "俄语"
            ],
            "ReleaseDate": "2017-03-03",
            "RegionComment": "",
            "HistoryLowestPrice": 387.48
        }
    ],
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
    "LanguageTag": [
        0
    ],
    "HasSolidEdition": true,
    "HasDemo": false,
    "IsExclusive": true,
    "GameTypeTag": [
        "冒险",
        "动作",
        "角色扮演"
    ],
    "HasChinese": true,
    "GamePlayers": 1,
    "Description": "游戏变成了完全的开放地图，沙盒式的玩法拓展出了更多的内容，林克可以爬山、游泳，而且游戏对应外观变化，玩家在游戏过程中会找到不止一件的服装以及武器。而在大地图中也会找到更多可以拾取的道具。服装除了外观外也有着属性的设定，例如林克在寒冷的地方穿薄衣服就会哆嗦，需要换上厚一点的服装才会变正常。在地图上的互动要素也不少，玩家手上拿着火把就能把草堆点燃，拿着斧头就能砍树等。而且不同的武器也会反馈出不用的攻击动作，某些武器还能当作投掷道具扔出去。游戏的战斗中还有类似“子弹时间”的设定，当敌人漏出过大的破绽时就可以采取更有效的攻击，但具体发动条件未知。林克还拥有一个魔法炸弹一样的道具可以使用，而且可以改变其用法形态，通过炸弹可以炸掉矿点来获得素材道具，不过魔法炸弹每次使用之后都要经过冷却时间。林克还可以从高空滑翔至低处，不过滑翔也和攀爬一样需要消耗耐力。或者可以选择把盾牌当做滑板从高处滑下。",
    "GameScore": [
        "IGN:10",
        "MC:97"
    ]
}
```
