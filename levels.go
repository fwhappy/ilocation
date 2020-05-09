package ilocation

var levels = map[string]string{
	"上海市":       "一线城市",
	"北京市":       "一线城市",
	"广州市":       "一线城市",
	"深圳市":       "一线城市",
	"天津市":       "一线城市",
	"重庆市":       "一线城市",
	"香港特区":      "一线城市",
	"澳门特区":      "一线城市",
	"台湾省":       "一线城市",
	"石家庄市":      "二线城市",
	"太原市":       "二线城市",
	"沈阳市":       "二线城市",
	"大连市":       "二线城市",
	"长春市":       "二线城市",
	"哈尔滨市":      "二线城市",
	"南京市":       "二线城市",
	"杭州市":       "二线城市",
	"宁波市":       "二线城市",
	"合肥市":       "二线城市",
	"福州市":       "二线城市",
	"厦门市":       "二线城市",
	"南昌市":       "二线城市",
	"济南市":       "二线城市",
	"青岛市":       "二线城市",
	"郑州市":       "二线城市",
	"武汉市":       "二线城市",
	"长沙市":       "二线城市",
	"南宁市":       "二线城市",
	"成都市":       "二线城市",
	"贵阳市":       "二线城市",
	"昆明市":       "二线城市",
	"西安市":       "二线城市",
	"乌鲁木齐市":     "二线城市",
	"唐山市":       "三线城市",
	"廊坊市":       "三线城市",
	"大同市":       "三线城市",
	"呼和浩特市":     "三线城市",
	"包头市":       "三线城市",
	"鞍山市":       "三线城市",
	"抚顺市":       "三线城市",
	"本溪市":       "三线城市",
	"锦州市":       "三线城市",
	"齐齐哈尔市":     "三线城市",
	"大庆市":       "三线城市",
	"无锡市":       "三线城市",
	"徐州市":       "三线城市",
	"常州市":       "三线城市",
	"苏州市":       "三线城市",
	"南通市":       "三线城市",
	"连云港市":      "三线城市",
	"扬州市":       "三线城市",
	"镇江市":       "三线城市",
	"泰州市":       "三线城市",
	"宿迁市":       "三线城市",
	"温州市":       "三线城市",
	"嘉兴市":       "三线城市",
	"湖州市":       "三线城市",
	"绍兴市":       "三线城市",
	"金华市":       "三线城市",
	"衢州市":       "三线城市",
	"芜湖市":       "三线城市",
	"蚌埠市":       "三线城市",
	"安庆市":       "三线城市",
	"泉州市":       "三线城市",
	"南平市":       "三线城市",
	"九江市":       "三线城市",
	"淄博市":       "三线城市",
	"烟台市":       "三线城市",
	"潍坊市":       "三线城市",
	"泰安市":       "三线城市",
	"威海市":       "三线城市",
	"黄石市":       "三线城市",
	"荆州市":       "三线城市",
	"咸宁市":       "三线城市",
	"株洲市":       "三线城市",
	"岳阳市":       "三线城市",
	"郴州市":       "三线城市",
	"珠海市":       "三线城市",
	"汕头市":       "三线城市",
	"佛山市":       "三线城市",
	"惠州市":       "三线城市",
	"东莞市":       "三线城市",
	"中山市":       "三线城市",
	"柳州市":       "三线城市",
	"桂林市":       "三线城市",
	"北海市":       "三线城市",
	"海口市":       "三线城市",
	"绵阳市":       "三线城市",
	"攀枝花市":      "三线城市",
	"拉萨市":       "三线城市",
	"宝鸡市":       "三线城市",
	"咸阳市":       "三线城市",
	"铜川市":       "三线城市",
	"兰州市":       "三线城市",
	"西宁市":       "三线城市",
	"银川市":       "三线城市",
	"秦皇岛市":      "四线城市",
	"张家口市":      "四线城市",
	"承德市":       "四线城市",
	"呼伦贝尔市":     "四线城市",
	"黄山市":       "四线城市",
	"景德镇市":      "四线城市",
	"洛阳市":       "四线城市",
	"张家界市":      "四线城市",
	"百色市":       "四线城市",
	"三亚市":       "四线城市",
	"广元市":       "四线城市",
	"乐山市":       "四线城市",
	"遵义市":       "四线城市",
	"丽江市":       "四线城市",
	"西双版纳傣族自治州": "四线城市",
	"延安市":       "四线城市",
	"嘉峪关市":      "四线城市",
	"吐鲁番市":      "四线城市",
	"哈密市":       "四线城市",
}
