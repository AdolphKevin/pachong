package main

var provinceMap map[string]string
var cityMap map[string]string

func init() {
	provinceMap = make(map[string]string)
	cityMap = make(map[string]string)
}

func doneProvinceMap() {
	provinceMap["Penza"] = "奔萨"
	provinceMap["Santa Catarina"] = "圣卡塔琳娜"
	provinceMap["Gaziantep"] = "加济安泰普"
	provinceMap["Izmir"] = "伊兹密尔省"
	provinceMap["La Coruna"] = "拉科鲁尼亚"
	provinceMap["Maykop"] = "迈科普"
	provinceMap["Mukdahan"] = "穆达汉(莫拉限)"
	provinceMap["Ushuaia"] = "乌斯怀亚"
	provinceMap["Delaware"] = "特拉华"
	provinceMap["Faro"] = "法罗"
	provinceMap["Jujuy"] = "胡胡伊"
	provinceMap["Posadas"] = "波萨达斯"
	provinceMap["Clare"] = "克莱尔"
	provinceMap["Norrbottens"] = "北博滕"
	provinceMap["California"] = "加利福尼亚"
	provinceMap["Catanzaro"] = "卡坦扎罗"
	provinceMap["Goa"] = "果阿"
	provinceMap["Prato"] = "普拉托"
	provinceMap["King"] = "金"
	provinceMap["Mayo"] = "梅奥"
	provinceMap["Rotterdam"] = "鹿特丹"
	provinceMap["Middelburg"] = "米德尔堡"
	provinceMap["San Luis Potosi"] = "圣路易斯波托西"
	provinceMap["San Luis"] = "圣路易斯"
	provinceMap["Tottori"] = "鸟取"
	provinceMap["Tottori-ken"] = "鸟取县"
	provinceMap["Yellowknife"] = "耶洛奈夫"
	provinceMap["Perm"] = "彼尔姆"
	provinceMap["Halifax"] = "哈利法克斯"
	provinceMap["Hermosillo"] = "埃莫西约"
	provinceMap["Kars"] = "卡尔斯"
	provinceMap["Chiba"] = "千叶"
	provinceMap["Chiba-ken"] = "千叶县"
	provinceMap["Chieti"] = "基耶蒂"
	provinceMap["Ulsan"] = "蔚山"
	provinceMap["Leitrim"] = "利特里姆"
	provinceMap["Maha Sarakham"] = "马哈沙拉堪"
	provinceMap["Rondonia"] = "朗多尼亚"
	provinceMap["Resistencia"] = "雷西斯滕西亚"
	provinceMap["Canakkale"] = "恰纳卡莱"
	provinceMap["Petropavlovsk-Kamchatskiy"] = "堪察加彼得罗巴甫洛夫斯克"
	provinceMap["Kalmar"] = "卡尔马"
	provinceMap["Puerto Vallarta"] = "巴亚尔塔港"
	provinceMap["Daejeon"] = "大田"
	provinceMap["Gerona"] = "赫罗纳"
	provinceMap["Quebec"] = "魁北克"
	provinceMap["Salerno"] = "萨莱诺"
	provinceMap["Orizaba"] = "奥里萨巴"
	provinceMap["Pest"] = "佩斯"
	provinceMap["Cuanza Sul"] = "南宽扎"
	provinceMap["Formosa"] = "福莫萨"
	provinceMap["Guipuzcoa"] = "吉普斯夸"
	provinceMap["Lisboa"] = "里斯本"
	provinceMap["Sumatera Utara"] = "北苏门答腊"
	provinceMap["Istanbul"] = "伊斯坦布尔省"
	provinceMap["Pahang"] = "彭亨"
	provinceMap["Rajkot"] = "拉杰果德"
	provinceMap["Mount Ayliff"] = "芒特艾利夫"
	provinceMap["Palermo"] = "巴勒莫"
	provinceMap["Port Shepstone"] = "谢普斯通港"
	provinceMap["Cochin"] = "科钦"
	provinceMap["Lodi"] = "洛迪"
	provinceMap["Los Angeles"] = "洛杉矶"
	provinceMap["Mexicali"] = "墨西加利"
	provinceMap["Gumushane"] = "居米什哈内"
	provinceMap["Huelva"] = "韦尔瓦"
	provinceMap["Vermont"] = "佛蒙特"
	provinceMap["Galway"] = "戈尔韦"
	provinceMap["Iwate-ken"] = "岩手县"
	provinceMap["Kemerovo"] = "克麦罗沃"
	provinceMap["Kanagawa-ken"] = "神奈川县"
	provinceMap["Kanagawa"] = "神奈川县"
	provinceMap["Kerry"] = "凯里"
	provinceMap["Schwyz"] = "施维茨"
	provinceMap["Tocantins"] = "托坎廷斯"
	provinceMap["Brasilia"] = "巴西利亚"
	provinceMap["Sudbury"] = "萨德伯里"
	provinceMap["Enschede"] = "恩斯赫德"
	provinceMap["Piacenza"] = "皮亚琴察"
	provinceMap["Vancouver"] = "温哥华"
	provinceMap["Zaire"] = "扎伊尓"
	provinceMap["Falcon"] = "法尔孔"
	provinceMap["Nebraska"] = "内布拉斯加"
	provinceMap["Tabuk"] = "泰布克"
	provinceMap["Lausanne"] = "洛桑"
	provinceMap["Murmansk"] = "摩尔曼斯克"
	provinceMap["Tyumen"] = "秋明"
	provinceMap["Vladivostok"] = "符拉迪沃斯托克(海参崴)"
	provinceMap["Kalimantan Selatan"] = "南加里曼丹"
	provinceMap["New York"] = "纽约"
	provinceMap["Portuguesa"] = "波图格萨"
	provinceMap["Trieste"] = "的里雅斯特"
	provinceMap["Hokkaido"] = "北海道"
	provinceMap["Negeri Sembilan"] = "森美兰州"
	provinceMap["Edmonton"] = "埃德蒙顿"
	provinceMap["Izhevsk"] = "伊热夫斯克"
	provinceMap["Salta"] = "萨尔塔"
	provinceMap["Pathum Thani"] = "巴吞他尼"
	provinceMap["Pistoia"] = "皮斯托亚"
	provinceMap["Randfontein"] = "兰德方丹"
	provinceMap["Louisiana"] = "路易斯安那"
	provinceMap["Ordu"] = "奥尔杜"
	provinceMap["Samut Sakhon"] = "沙没沙空(龙仔厝)"
	provinceMap["Sulawesi Tengah"] = "中苏拉维西"
	provinceMap["Lara"] = "拉腊"
	provinceMap["Meath"] = "米斯"
	provinceMap["Neuchatel"] = "纳沙泰尔"
	provinceMap["Zhytomyr"] = "日托米尔"
	provinceMap["Tijuana"] = "蒂华纳"
	provinceMap["Cabinda"] = "卡宾达"
	provinceMap["Davao"] = "达沃(纳卯)"
	provinceMap["Mysore"] = "迈索尔"
	provinceMap["Nagasaki"] = "长崎"
	provinceMap["Imphal"] = "因帕尔"
	provinceMap["Ladysmith"] = "莱迪史密斯"
	provinceMap["Songkhla"] = "宋卡"
	provinceMap["Victoria"] = "维多利亚州"
	provinceMap["Tokyo"] = "东京"
	provinceMap["Nonthaburi"] = "暖武里"
	provinceMap["Nusa Tenggara Timur"] = "东努沙登加拉"
	provinceMap["Richmond"] = "里士满"
	provinceMap["Ticino"] = "提契诺"
	provinceMap["Orebro"] = "厄勒布鲁"
	provinceMap["Vijayawada"] = "维杰亚瓦达"
	provinceMap["Busan"] = "釜山"
	provinceMap["Grosseto"] = "格罗塞托"
	provinceMap["Indiana"] = "印第安纳"
	provinceMap["Kalimantan Timur"] = "东加里曼丹"
	provinceMap["Lamphun"] = "南奔"
	provinceMap["Orense"] = "奥伦塞"
	provinceMap["Burgos"] = "布尔戈斯"
	provinceMap["Dordrecht"] = "多德雷赫特"
	provinceMap["Firenze"] = "佛罗伦萨"
	provinceMap["Krasnodar"] = "克拉斯诺达尔"
	provinceMap["Ravenna"] = "拉韦纳"
	provinceMap["Cebu"] = "宿务"
	provinceMap["Caen"] = "卡昂"
	provinceMap["Cologne"] = "科隆"
	provinceMap["Salem"] = "萨伦"
	provinceMap["Waterford"] = "沃特福德"
	provinceMap["Lorain"] = "洛雷恩"
	provinceMap["Minas Gerais"] = "米纳斯吉拉斯"
	provinceMap["Tunceli"] = "通杰利"
	provinceMap["Concordia"] = "康科迪亚"
	provinceMap["Trujillo"] = "特鲁希略"
	provinceMap["Cunene"] = "库内内"
	provinceMap["Gunma"] = "群马"
	provinceMap["Gunma-ken"] = "群马县"
	provinceMap["Lunda Norte"] = "北隆达"
	provinceMap["Muenster"] = "曼斯特"
	provinceMap["Wan Chai"] = "湾仔"
	provinceMap["Chieti"] = "基耶蒂"
	provinceMap["Dallas"] = "达拉斯"
	provinceMap["Seoul"] = "首尔(2005年10月前译为汉城)"
	provinceMap["Carmen"] = "卡门"
	provinceMap["Maricopa"] = "马里科帕"
	provinceMap["Roma"] = "罗马"
	provinceMap["Catanzaro"] = "卡坦扎罗"
	provinceMap["Obwalden"] = "上瓦尔登"
	provinceMap["Scotland"] = "苏格兰"
	provinceMap["Sligo"] = "斯莱戈"
	provinceMap["Sura"] = "苏拉"
	provinceMap["Cuenca"] = "昆卡"
	provinceMap["Southern"] = "南方"
	provinceMap["Wexford"] = "韦克斯福德"
	provinceMap["Lunda Sul"] = "南隆达"
	provinceMap["Saskatoon"] = "萨斯卡通"
	provinceMap["Sulawesi Selatan"] = "南苏拉维西"
	provinceMap["Toulouse"] = "图卢兹"
	provinceMap["Chihuahua"] = "奇瓦瓦"
	provinceMap["Kagawa-ken"] = "香川县"
	provinceMap["Krasnodar"] = "克拉斯诺达尔"
	provinceMap["Middelburg"] = "米德尔堡"
	provinceMap["North Shore"] = "北岸"
	provinceMap["Thohoyandou"] = "托霍扬多"
	provinceMap["Richards Bay"] = "理查兹贝"
	provinceMap["Guanajuato"] = "瓜纳华托"
	provinceMap["Gyeongsangbuk-do"] = "庆尚北道"
	provinceMap["Kentucky"] = "肯塔基"
	provinceMap["Kyoto-fu"] = "京都府"
	provinceMap["Lyon"] = "里昂"
	provinceMap["Nograd"] = "诺格拉德"
	provinceMap["Celaya"] = "塞拉亚"
	provinceMap["Geneve"] = "日内瓦"
	provinceMap["Konya"] = "科尼亚"
	provinceMap["Ludhiana"] = "卢迪亚纳"
	provinceMap["Virginia"] = "弗吉尼亚"
	provinceMap["Hyogo"] = "兵库县"
	provinceMap["Kanpur"] = "坎普尔"
	provinceMap["Magas"] = "马加斯"
	provinceMap["Pike"] = "派克"
	provinceMap["Mugla"] = "穆拉"
	provinceMap["Santa Catarina"] = "圣卡塔琳娜"
	provinceMap["Uruapan"] = "乌鲁阿潘"
	provinceMap["Chachoengsao"] = "差春骚(北柳)"
	provinceMap["Chalons-en-Champagne"] = "香槟地区沙隆"
	provinceMap["Cork"] = "科克"
	provinceMap["Ferrara"] = "费拉拉"
	provinceMap["Petropavlovsk-Kamchatskiy"] = "堪察加彼得罗巴甫洛夫斯克"
	provinceMap["Tomsk"] = "托木斯克"
	provinceMap["Brunswick"] = "不伦瑞克"
	provinceMap["Selangor"] = "雪兰莪"
	provinceMap["Hoofddorp"] = "霍夫多普"
	provinceMap["Maluku"] = "马卢库"
	provinceMap["Espirito Santo"] = "圣埃斯皮里图"
	provinceMap["Morelia"] = "莫雷利亚"
	provinceMap["Pesaro"] = "佩萨罗"
	provinceMap["Thunder Bay"] = "桑德贝"
	provinceMap["Osaka"] = "小坂"
	provinceMap["Springbok"] = "斯普林博克"
	provinceMap["Bokaro"] = "波卡罗"
	provinceMap["Idaho"] = "爱达荷"
	provinceMap["Wisconsin"] = "威斯康星"
	provinceMap["Ponta Delgada"] = "蓬塔德尔加达"
	provinceMap["Sondrio"] = "松德里奥"
	provinceMap["Phayao"] = "帕尧"
	provinceMap["Sakarya"] = "萨卡里亚"
	provinceMap["Sulawesi Utara"] = "北苏拉维西"
	provinceMap["Pattani"] = "北大年"
	provinceMap["Sabah"] = "沙巴"
	provinceMap["Canakkale"] = "恰纳卡莱"
	provinceMap["Chabarovsk"] = "哈巴罗夫斯克"
	provinceMap["Pahang"] = "彭亨"
	provinceMap["Dalarnas"] = "达拉纳"
	provinceMap["East London"] = "东伦敦"
	provinceMap["Gwalior"] = "瓜廖尔"
	provinceMap["Hannover"] = "汉诺威"
	provinceMap["Leeuwarden"] = "吕伐登"
	provinceMap["Norrbottens"] = "北博滕"
	provinceMap["Wyoming"] = "怀俄明"
	provinceMap["Caloocan"] = "卡洛奥坎"
	provinceMap["Tabuk"] = "泰布克"
	provinceMap["Tolna"] = "托尔瑙"
	provinceMap["Ehime-ken"] = "爱媛县"
	provinceMap["Ehime"] = "爱媛县"
	provinceMap["Georgia"] = "格鲁吉亚"
	provinceMap["Kamphaeng Phet"] = "甘烹碧"
	provinceMap["Krasnojarsk"] = "克拉斯诺亚尔斯克"
	provinceMap["Nusa Tenggara Timur"] = "东努沙登加拉"
	provinceMap["Porto"] = "波尔托"
	provinceMap["Lodi"] = "洛迪"
	provinceMap["Nijmegen"] = "奈梅亨"
	provinceMap["San Miguel de Tucuman"] = "图库曼省圣米格尔"
	provinceMap["Tambov"] = "坦波夫"
	provinceMap["Vizcaya"] = "比斯开"
	provinceMap["Heves"] = "赫维什"
	provinceMap["Lucknow"] = "勒克瑙"
	provinceMap["Maranhao"] = "马腊尼昂"
	provinceMap["New Brunswick"] = "新不伦瑞克"
	provinceMap["Gerona"] = "赫罗纳"
	provinceMap["Grand Forks"] = "大福克斯"
	provinceMap["Pordenone"] = "波代诺内"
	provinceMap["Terengganu"] = "丁加奴"
	provinceMap["Fukuoka-ken"] = "福冈县"
	provinceMap["Huila"] = "威拉"
	provinceMap["Orange"] = "奥朗日"
	provinceMap["Kalimantan Selatan"] = "南加里曼丹"
	provinceMap["Kazan"] = "喀山"
	provinceMap["Lausanne"] = "洛桑"
	provinceMap["Caceres"] = "卡塞雷斯"
	provinceMap["Hiroshima-ken"] = "广岛县"
	provinceMap["Jiddah"] = "吉达"
	provinceMap["Odessa"] = "敖德萨"
	provinceMap["Krasnojarsk"] = "克拉斯诺亚尔斯克"
	provinceMap["Mar del Plata"] = "马德普拉塔"
	provinceMap["Mexico City"] = "墨西哥城"
	provinceMap["Skane"] = "斯科耐"
	provinceMap["Yuen Long"] = "元朗"
	provinceMap["Cologne"] = "科隆"
	provinceMap["Dependencias Federales"] = "联邦属地"
	provinceMap["Gunma-ken"] = "群马县"
	provinceMap["Neuquen"] = "内乌肯"
	provinceMap["Toronto"] = "多伦多"
	provinceMap["Roi Et"] = "黎逸"
	provinceMap["Shimane-ken"] = "岛根县"
	provinceMap["Burgos"] = "布尔戈斯"
	provinceMap["Kagoshima-ken"] = "鹿儿岛县"
	provinceMap["Lampung"] = "楠榜"
	provinceMap["Surat Thani"] = "素叻他尼"
	provinceMap["Terengganu"] = "丁加奴"
	provinceMap["Vladimir"] = "弗拉基米尔"
	provinceMap["Bologna"] = "博洛尼亚(波伦亚)"
	provinceMap["Penza"] = "奔萨"
	provinceMap["Macau"] = "马科"
	provinceMap["Parma"] = "帕马"
	provinceMap["Santiago del Estero"] = "圣地亚哥-德尔埃斯特罗"
	provinceMap["Florida"] = "佛罗里达"
	provinceMap["Kudymkar"] = "库德姆卡尔"
	provinceMap["Maha Sarakham"] = "马哈沙拉堪"
	provinceMap["Blenheim"] = "布莱尼姆"
	provinceMap["Faridabad"] = "法里达巴德"
	provinceMap["Tochigi"] = "枥木"
	provinceMap["Tochigi-ken"] = "枥木县"
	provinceMap["Viedma"] = "别德马"
	provinceMap["Krasnodar"] = "克拉斯诺达尔"
	provinceMap["Mendoza"] = "门多萨"
	provinceMap["Huesca"] = "韦斯卡"
	provinceMap["Kentucky"] = "肯塔基"
	provinceMap["Kildare"] = "基尔代尔"
	provinceMap["Louisiana"] = "路易斯安那"
	provinceMap["Mardin"] = "马尔丁"
	provinceMap["Rovigo"] = "罗维戈"
	provinceMap["Mukdahan"] = "穆达汉(莫拉限)"
	provinceMap["Saga"] = "佐贺"
	provinceMap["Saga-ken"] = "佐贺县"
	provinceMap["Sakaka"] = "塞卡卡"
	provinceMap["Bochum"] = "波鸿"
	provinceMap["Phitsanulok"] = "彭世洛"
	provinceMap["Wisconsin"] = "威斯康星"
	provinceMap["Delta Amacuro"] = "阿马库罗三角洲"
	provinceMap["Groblersdal"] = "格罗布莱斯达尔"
	provinceMap["Pordenone"] = "波代诺内"
	provinceMap["Tabuk"] = "泰布克"
	provinceMap["Trier"] = "特里尔"
	provinceMap["Freiburg"] = "弗里堡"
	provinceMap["Lamphun"] = "南奔"
	provinceMap["Northern Territory"] = "北部地区"
	provinceMap["Sabah"] = "沙巴"
	provinceMap["Giyani"] = "基雅尼"
	provinceMap["Maranhao"] = "马腊尼昂"
	provinceMap["Nidwalden"] = "下瓦尔登"
	provinceMap["Sulawesi Selatan"] = "南苏拉维西"
	provinceMap["Edirne"] = "埃迪尔内"
	provinceMap["Louth"] = "劳斯"
	provinceMap["Leitrim"] = "利特里姆"
	provinceMap["Mugla"] = "穆拉"
	provinceMap["Sevilla"] = "塞维利亚"
	provinceMap["Valencia"] = "瓦伦西亚"
	provinceMap["Dresden"] = "德累斯顿"
	provinceMap["Eindhoven"] = "艾恩德霍芬"
	provinceMap["Voronezh"] = "沃罗涅日"
	provinceMap["West Virginia"] = "西维吉尼亚"
	provinceMap["Genova"] = "热那亚"
	provinceMap["Limerick"] = "利默里克"
	provinceMap["Tver"] = "特维尔"
	provinceMap["Varmlands"] = "韦姆兰"
	provinceMap["Ferrara"] = "费拉拉"
	provinceMap["Phetchaburi"] = "碧武里(佛丕)"
	provinceMap["Salerno"] = "萨莱诺"
	provinceMap["Indian River"] = "印第安里弗"
	provinceMap["Kirov"] = "基洛夫"
	provinceMap["Richards Bay"] = "理查兹贝"
	provinceMap["Namibe"] = "纳米贝"
	provinceMap["Sakarya"] = "萨卡里亚"
	provinceMap["Saskatoon"] = "萨斯卡通"
	provinceMap["Chiang Rai"] = "清莱"
	provinceMap["De Aar"] = "德阿尔"
	provinceMap["Vizcaya"] = "比斯开"
	provinceMap["Latina"] = "拉蒂纳"
	provinceMap["Diyarbakir"] = "迪亚巴克尔"
	provinceMap["Brescia"] = "布雷西亚"
	provinceMap["Ufa"] = "乌法"
	provinceMap["Sura"] = "苏拉"
	provinceMap["Bolu"] = "博卢"
	provinceMap["Maluku"] = "马卢库"
	provinceMap["King"] = "金"
	provinceMap["Lampung"] = "楠榜"
	provinceMap["Patna"] = "帕特纳"
	provinceMap["Rize"] = "里泽"
	provinceMap["Chemnitz"] = "开姆尼茨"
	provinceMap["La Spezia"] = "拉斯佩齐亚"
	provinceMap["Osaka"] = "小坂"
	provinceMap["Parana"] = "帕拉南"
	provinceMap["Port Shepstone"] = "谢普斯通港"
	provinceMap["Buraydah"] = "布赖代"
	provinceMap["Iowa"] = "艾奥瓦"
	provinceMap["North Shore"] = "北岸"
	provinceMap["Ordu"] = "奥尔杜"
	provinceMap["Mannheim"] = "曼海姆"
	provinceMap["Nalchik"] = "纳尔奇克"
	provinceMap["Niznij Novgorod"] = "下诺夫哥罗德"
	provinceMap["Obregon"] = "奥夫雷贡"
	provinceMap["Ticino"] = "提契诺"
	provinceMap["Wan Chai"] = "湾仔"
	provinceMap["Erzincan"] = "埃尔津詹"
	provinceMap["Kolkata"] = "加尔各答"
	provinceMap["Zaporizhzhya"] = "札波罗热"
	provinceMap["Kilkenny"] = "基尔肯尼"
	provinceMap["Burgenland"] = "布尔根兰"
	provinceMap["Jefferson"] = "杰斐逊"
	provinceMap["Hyderabad"] = "海得拉巴"
	provinceMap["Orizaba"] = "奥里萨巴"
	provinceMap["Sassari"] = "萨萨里"
	provinceMap["Ceara"] = "塞阿拉(福塔莱萨的旧称)"
	provinceMap["Freiburg"] = "弗里堡"
	provinceMap["Cuanza Norte"] = "北宽扎"
	provinceMap["Fermo"] = "费尔莫"
	provinceMap["Orel"] = "奥廖尔"
	provinceMap["Valence"] = "瓦朗斯"
	provinceMap["Funchal"] = "丰沙尔"
	provinceMap["Lampung"] = "楠榜"
	provinceMap["Niznij Novgorod"] = "下诺夫哥罗德"
	provinceMap["Northern Ireland"] = "北爱尔兰"
	provinceMap["Preah Vihear"] = "柏威夏"
	provinceMap["Kracheh"] = "桔井"
	provinceMap["Maryland"] = "马里兰"
	provinceMap["Ivanovo"] = "伊万诺沃"
	provinceMap["Kalimantan Selatan"] = "南加里曼丹"
	provinceMap["Kalmar"] = "卡尔马"
	provinceMap["Limerick"] = "利默里克"
	provinceMap["Longford"] = "朗福德"
	provinceMap["Sing Buri"] = "信武里"
	provinceMap["Cape Town"] = "开普敦"
	provinceMap["Groblersdal"] = "格罗布莱斯达尔"
	provinceMap["Sasolburg"] = "萨索尔堡"
	provinceMap["Ghaziabad"] = "加济阿巴德"
	provinceMap["Gwangju"] = "光州湖"
	provinceMap["Parma"] = "帕马"
	provinceMap["Tokushima-ken"] = "德岛县"
	provinceMap["Ufa"] = "乌法"
	provinceMap["Yozgat"] = "约兹加特"
	provinceMap["Chiba"] = "千叶"
	provinceMap["Chiba-ken"] = "千叶县"
	provinceMap["Maricopa"] = "马里科帕"
	provinceMap["Faro"] = "法罗"
	provinceMap["Giresun"] = "吉雷松"
	provinceMap["Incheon"] = "仁川"
	provinceMap["Santiago del Estero"] = "圣地亚哥-德尔埃斯特罗"
	provinceMap["Solothurn"] = "索洛图恩"
	provinceMap["Chita"] = "赤塔"
	provinceMap["Maykop"] = "迈科普"
	provinceMap["Sakarya"] = "萨卡里亚"
	provinceMap["Irkutsk"] = "伊尔库茨克"
	provinceMap["Loei"] = "黎"
	provinceMap["Ubon Ratchathani"] = "乌汶(乌汶叻差他尼)"
	provinceMap["Omsk"] = "鄂木斯克"
	provinceMap["Trompsburg"] = "特龙普斯堡"
	provinceMap["Male Atoll"] = "马累环礁"
	provinceMap["Margao"] = "马尔冈(默德冈的旧称)"
	provinceMap["Pesaro"] = "佩萨罗"
	provinceMap["Siena"] = "锡耶纳"
	provinceMap["Venezia"] = "威尼斯"
	provinceMap["Caserta"] = "卡塞塔"
	provinceMap["Daejeon"] = "大田"
	provinceMap["Portuguesa"] = "波图格萨"
	provinceMap["Vienna"] = "维也纳"
	provinceMap["Nidwalden"] = "下瓦尔登"
	provinceMap["Philadelphia"] = "费城"
	provinceMap["Thiruvananthapuram"] = "特里凡得琅"
	provinceMap["Vercelli"] = "韦尔切利"
	provinceMap["Vizcaya"] = "比斯开"
	provinceMap["Isparta"] = "厄斯帕尔塔"
	provinceMap["Kazan"] = "喀山"
	provinceMap["Lelystad"] = "莱利斯塔德"
	provinceMap["Dublin"] = "都柏林"
	provinceMap["Jiddah"] = "吉达"
	provinceMap["Lodi"] = "洛迪"
	provinceMap["Uttaradit"] = "程逸(乌达腊迪)"
	provinceMap["Hakkari"] = "哈卡里"
	provinceMap["Labuan"] = "拉布安"
	provinceMap["Oita"] = "大分"
	provinceMap["Tarragona"] = "塔拉戈纳"
	provinceMap["Terengganu"] = "丁加奴"
	provinceMap["Trat"] = "达叻(桐艾)"
	provinceMap["Tsuen Wan"] = "荃湾"
	provinceMap["Hail"] = "哈伊勒"
	provinceMap["Vicenza"] = "维琴察"
	provinceMap["Zonguldak"] = "宗古尔达克"
	provinceMap["Pachuca"] = "帕丘卡"
	provinceMap["Sura"] = "苏拉"
	provinceMap["Ranchi"] = "兰契"
	provinceMap["Springbok"] = "斯普林博克"
	provinceMap["Cuenca"] = "昆卡"
	provinceMap["Maluku"] = "马卢库"
	provinceMap["Lille"] = "里尔"
	provinceMap["Tachira"] = "塔奇拉"
	provinceMap["Veracruz"] = "韦拉克鲁斯"
	provinceMap["Calgary"] = "卡尔加里"
	provinceMap["Donegal"] = "多尼戈尔"
	provinceMap["Male"] = "马累"
	provinceMap["Saltillo"] = "索尔蒂洛"
	provinceMap["Salzburg"] = "萨尔茨堡"
	provinceMap["Dallas"] = "达拉斯"
	provinceMap["Kharkiv"] = "哈尔科夫"
	provinceMap["Kyoto-fu"] = "京都府"
	provinceMap["Pordenone"] = "波代诺内"
	provinceMap["Tula"] = "图拉"
	provinceMap["Hokkaido"] = "北海道"
	provinceMap["Karabuk"] = "卡拉比克"
	provinceMap["Pretoria"] = "比勒陀利亚"
	provinceMap["Haarlem"] = "哈勒姆"
	provinceMap["Nakhon Sawan"] = "那空沙旺;北榄坡"
	provinceMap["Lecce"] = "莱切"
	provinceMap["Bolu"] = "博卢"
	provinceMap["Como"] = "科莫"
	provinceMap["Perugia"] = "佩鲁贾"
	provinceMap["Upington"] = "阿平顿"
	provinceMap["Macerata"] = "马切拉塔"
	provinceMap["North Carolina"] = "北卡罗来纳"
	provinceMap["Kansas"] = "堪萨斯"
	provinceMap["Kildare"] = "基尔代尔"
	provinceMap["Oregon"] = "俄勒冈"
	provinceMap["Reynosa"] = "雷诺萨"
	provinceMap["Western Australia"] = "西澳大利亚"
	provinceMap["Hyderabad"] = "海得拉巴"
	provinceMap["Zug"] = "楚格"
	provinceMap["Pest"] = "佩斯"
	provinceMap["Sligo"] = "斯莱戈"
	provinceMap["Frosinone"] = "弗罗西诺内"
	provinceMap["Lamphun"] = "南奔"
	provinceMap["Shimane-ken"] = "岛根县"
	provinceMap["Skane"] = "斯科耐"
	provinceMap["Tambov"] = "坦波夫"
	provinceMap["Jodhpur"] = "焦特布尔"
	provinceMap["Salem"] = "萨伦"
	provinceMap["Stavropol"] = "斯塔夫罗波尔"
	provinceMap["Stratford"] = "斯特拉特福"
	provinceMap["West Virginia"] = "西维吉尼亚"
	provinceMap["Ceara"] = "塞阿拉(福塔莱萨的旧称)"
	provinceMap["Goa"] = "果阿"
	provinceMap["Bloemfontein"] = "布隆方丹"
	provinceMap["Wisconsin"] = "威斯康星"
	provinceMap["La Rioja"] = "拉里奥哈"
	provinceMap["Pisa"] = "比萨"
	provinceMap["Malatya"] = "马拉蒂亚"
	provinceMap["Prato"] = "普拉托"
	provinceMap["San Miguel de Tucuman"] = "图库曼省圣米格尔"
	provinceMap["Longford"] = "朗福德"
	provinceMap["Pachuca"] = "帕丘卡"
	provinceMap["Reynosa"] = "雷诺萨"
	provinceMap["Schaffhausen"] = "沙夫豪森"
	provinceMap["Ticino"] = "提契诺"
	provinceMap["Lunda Norte"] = "北隆达"
	provinceMap["Maluku"] = "马卢库"
	provinceMap["North Dakota"] = "北达科他"
	provinceMap["Whakatane"] = "瓦卡塔尼"
	provinceMap["Buenos Aires"] = "布宜诺斯艾利斯"
	provinceMap["Foggia"] = "福贾"
	provinceMap["Macau"] = "马科"
	provinceMap["Malaga"] = "马拉加"
	provinceMap["Ponta Delgada"] = "蓬塔德尔加达"
	provinceMap["Resistencia"] = "雷西斯滕西亚"
	provinceMap["Skane"] = "斯科耐"
	provinceMap["Campeche"] = "坎佩切"
	provinceMap["Dalarnas"] = "达拉纳"
	provinceMap["Fukushima-ken"] = "福岛县"
	provinceMap["Nelson"] = "纳尔逊"
	provinceMap["Phetchaburi"] = "碧武里(佛丕)"
	provinceMap["Pondicherry"] = "本地治里"
	provinceMap["Tachira"] = "塔奇拉"
	provinceMap["Wellington"] = "惠灵顿"
	provinceMap["Bochum"] = "波鸿"
	provinceMap["Cavado"] = "卡瓦多"
	provinceMap["Cojedes"] = "科赫德斯"
	provinceMap["Kirovohrad"] = "基洛沃赫拉德"
	provinceMap["Ranchi"] = "兰契"
	provinceMap["St.Gallen"] = "圣加仑"
	provinceMap["Trapani"] = "特拉帕尼"
	provinceMap["Uttaradit"] = "程逸(乌达腊迪)"
	provinceMap["Bolzano"] = "博尔扎诺"
	provinceMap["Cochin"] = "科钦"
	provinceMap["Malanje"] = "马兰热"
	provinceMap["Moorreesburg"] = "穆里斯堡"
	provinceMap["Zoetermeer"] = "佐特尔梅"
	provinceMap["Voronezh"] = "沃罗涅日"
	provinceMap["Mendoza"] = "门多萨"
	provinceMap["Ryazan"] = "梁赞"
	provinceMap["Stoeng Treng"] = "上丁"
	provinceMap["Kagoshima"] = "鹿儿岛"
	provinceMap["Kagoshima-ken"] = "鹿儿岛县"
	provinceMap["Tyumen"] = "秋明"
	provinceMap["Wicklow"] = "威克洛"
	provinceMap["Rieti"] = "列蒂"
	provinceMap["Shiga-ken"] = "滋贺县"
	provinceMap["Shiga"] = "滋贺〈滋贺县〉"
	provinceMap["Southern"] = "南方"
	provinceMap["Zacatecas"] = "萨卡特卡斯"
	provinceMap["Catamarca"] = "卡塔马卡"
	provinceMap["Catania"] = "卡塔尼亚"
	provinceMap["Hannover"] = "汉诺威"
	provinceMap["Modena"] = "摩德纳"
	provinceMap["Parma"] = "帕马"
	provinceMap["Sivas"] = "锡瓦斯"
	provinceMap["Tambov"] = "坦波夫"
	provinceMap["Carmen"] = "卡门"
	provinceMap["Kandal"] = "干丹"
	provinceMap["Victoria"] = "维多利亚州"
	provinceMap["Blekinge"] = "布莱金厄"
	provinceMap["Indian River"] = "印第安里弗"
	provinceMap["Manitoba"] = "马尼托巴省"
	provinceMap["Jamnagar"] = "贾姆讷格尔"
	provinceMap["Ottawa"] = "渥太华"
	provinceMap["Toyama-ken"] = "富山县"
	provinceMap["Ghaziabad"] = "加济阿巴德"
	provinceMap["Ludhiana"] = "卢迪亚纳"
	provinceMap["Vastra Gotalands"] = "西约特兰"
	provinceMap["Cabinda"] = "卡宾达"
	provinceMap["Pest"] = "佩斯"
	provinceMap["Puerto Vallarta"] = "巴亚尔塔港"
	provinceMap["Gandhinagar"] = "甘地讷格尔"
	provinceMap["Zug"] = "楚格"
	provinceMap["Elazig"] = "埃拉泽"
	provinceMap["Limerick"] = "利默里克"
	provinceMap["Magdeburg"] = "马格德堡"
	provinceMap["Trieste"] = "的里雅斯特"
	provinceMap["Breda"] = "布雷达"
	provinceMap["Budapest"] = "布达佩斯"
	provinceMap["Buraydah"] = "布赖代"
	provinceMap["Igdir"] = "厄德尔"
	provinceMap["Chachoengsao"] = "差春骚(北柳)"
	provinceMap["Jabalpur"] = "贾巴尔普尔"
	provinceMap["Koblenz"] = "科布伦茨"
	provinceMap["Rostov-na-Donu"] = "顿河畔罗斯托夫"
	provinceMap["Villahermosa"] = "比亚埃尔莫萨"
	provinceMap["Heves"] = "赫维什"
	provinceMap["Orebro"] = "厄勒布鲁"
	provinceMap["Rivne"] = "里夫内"
	provinceMap["Phrae"] = "帕府"
	provinceMap["Burgenland"] = "布尔根兰"
	provinceMap["Detmold"] = "代特莫尔德"
	provinceMap["Ibaraki-ken"] = "茨城县"
	provinceMap["Ibaraki"] = "茨木〈大阪府〉"
	provinceMap["Juarez"] = "华雷斯"
	provinceMap["Cosenza"] = "科森扎"
	provinceMap["Tijuana"] = "蒂华纳"
	provinceMap["Windsor"] = "温莎"
	provinceMap["Ushuaia"] = "乌斯怀亚"
	provinceMap["West Virginia"] = "西维吉尼亚"
	provinceMap["Siirt"] = "锡尔特"
	provinceMap["Ede"] = "埃德"
	provinceMap["Khon Kaen"] = "孔敬"
	provinceMap["Sevilla"] = "塞维利亚"
	provinceMap["Enschede"] = "恩斯赫德"
	provinceMap["Gyeonggi-do"] = "京畿道"
	provinceMap["Kalmar"] = "卡尔马"
	provinceMap["Leiden"] = "莱登"
	provinceMap["Gwangju"] = "光州湖"
	provinceMap["Waterford"] = "沃特福德"
	provinceMap["Irkutsk"] = "伊尔库茨克"
	provinceMap["Okinawa"] = "冲绳"
	provinceMap["Okinawa-ken"] = "冲绳县"
	provinceMap["Vinnytsya"] = "文尼察"
	provinceMap["Hamburg"] = "汉堡"
	provinceMap["Tak"] = "达府"
	provinceMap["Cuando Cubango"] = "宽多库邦戈"
	provinceMap["Donegal"] = "多尼戈尔"
	provinceMap["Geneve"] = "日内瓦"
	provinceMap["Tomsk"] = "托木斯克"
	provinceMap["Gunma-ken"] = "群马县"
	provinceMap["Irapuato"] = "伊拉普阿托"
	provinceMap["Chiang Rai"] = "清莱"
	provinceMap["Pulau Pinang"] = "槟榔屿"
	provinceMap["Northern Ireland"] = "北爱尔兰"
	provinceMap["Osaka"] = "小坂"
	provinceMap["Jamtlands"] = "耶姆特兰"
	provinceMap["Lille"] = "里尔"
	provinceMap["Kildare"] = "基尔代尔"
	provinceMap["Malanje"] = "马兰热"
	provinceMap["Maykop"] = "迈科普"
	provinceMap["Phetchabun"] = "碧差汶"
	provinceMap["Erzincan"] = "埃尔津詹"
	provinceMap["Faridabad"] = "法里达巴德"
	provinceMap["Gurgaon"] = "古尔冈"
	provinceMap["Brindisi"] = "布林迪西"
	provinceMap["Chumphon"] = "春蓬(尖喷)"
	provinceMap["Falcon"] = "法尔孔"
	provinceMap["Guadalajara"] = "瓜达拉哈拉"
	provinceMap["Poltava"] = "波尔塔瓦"
	provinceMap["Rawson"] = "罗森"
	provinceMap["Chaiyaphum"] = "猜也蓬"
	provinceMap["Huambo"] = "万博"
	provinceMap["Karaman"] = "卡拉曼"
	provinceMap["Thohoyandou"] = "托霍扬多"
	provinceMap["Usak"] = "乌沙克"
	provinceMap["Gunma-ken"] = "群马县"
	provinceMap["Leiden"] = "莱登"
	provinceMap["Palmerston North"] = "北帕默斯顿"
	provinceMap["Tijuana"] = "蒂华纳"
	provinceMap["Zaragoza"] = "萨拉戈萨"
	provinceMap["Jabalpur"] = "贾巴尔普尔"
	provinceMap["Tasmania"] = "塔斯马尼亚岛"
	provinceMap["Eskisehir"] = "埃斯基谢希尔"
	provinceMap["Madurai"] = "马杜赖"
	provinceMap["Suphan Buri"] = "素攀(素攀武里)"
	provinceMap["Cunene"] = "库内内"
	provinceMap["Lara"] = "拉腊"
	provinceMap["Nograd"] = "诺格拉德"
	provinceMap["Pest"] = "佩斯"
	provinceMap["Roma"] = "罗马"
	provinceMap["Vastmanlands"] = "西曼兰"
	provinceMap["Caloocan"] = "卡洛奥坎"
	provinceMap["Georgia"] = "格鲁吉亚"
	provinceMap["Goa"] = "果阿"
	provinceMap["Buraydah"] = "布赖代"
	provinceMap["Munich"] = "慕尼黑"
	provinceMap["Kazan"] = "喀山"
	provinceMap["George"] = "乔治"
	provinceMap["Limoges"] = "利摩日"
	provinceMap["Modena"] = "摩德纳"
	provinceMap["Palana"] = "帕拉纳"
	provinceMap["South Carolina"] = "南卡罗来纳"
	provinceMap["Tula"] = "图拉"
	provinceMap["Funchal"] = "丰沙尔"
	provinceMap["Vinnytsya"] = "文尼察"
	provinceMap["Rieti"] = "列蒂"
	provinceMap["Sumatera Utara"] = "北苏门答腊"
	provinceMap["Jamnagar"] = "贾姆讷格尔"
	provinceMap["Granada"] = "格拉纳达"
	provinceMap["Sudbury"] = "萨德伯里"
	provinceMap["Veracruz"] = "韦拉克鲁斯"
	provinceMap["Victoria"] = "维多利亚州"
	provinceMap["Burgenland"] = "布尔根兰"
	provinceMap["Huesca"] = "韦斯卡"
	provinceMap["Istanbul"] = "伊斯坦布尔省"
	provinceMap["Massachusetts"] = "马萨诸塞州"
	provinceMap["Wisconsin"] = "威斯康星"
	provinceMap["Eastern"] = "东方"
	provinceMap["Kharkiv"] = "哈尔科夫"
	provinceMap["Nusa Tenggara Timur"] = "东努沙登加拉"
	provinceMap["Skane"] = "斯科耐"
	provinceMap["Dubayy"] = "迪拜"
	provinceMap["Taumarunui"] = "陶马鲁努伊"
	provinceMap["Mantova"] = "曼托瓦"
	provinceMap["Miyagi-ken"] = "宫城县"
	provinceMap["Nagano"] = "长野"
	provinceMap["Salerno"] = "萨莱诺"
	provinceMap["Toluca"] = "托卢卡"
	provinceMap["Cuanza Sul"] = "南宽扎"
	provinceMap["Carabobo"] = "卡拉沃沃"
	provinceMap["Luzern"] = "卢塞恩"
	provinceMap["Nevsehir"] = "内夫谢希尔"
	provinceMap["Bologna"] = "博洛尼亚(波伦亚)"
	provinceMap["Ragusa"] = "拉古萨"
	provinceMap["Detmold"] = "代特莫尔德"
	provinceMap["Kandal"] = "干丹"
	provinceMap["Konya"] = "科尼亚"
	provinceMap["Ivanovo"] = "伊万诺沃"
	provinceMap["Magas"] = "马加斯"
	provinceMap["Rize"] = "里泽"
	provinceMap["Toyama-ken"] = "富山县"
	provinceMap["Cova da Beira"] = "贝拉洞"
	provinceMap["Giyani"] = "基雅尼"
	provinceMap["Orange"] = "奥朗日"
	provinceMap["Espirito Santo"] = "圣埃斯皮里图"
	provinceMap["Kyzyl"] = "克孜勒"
	provinceMap["Ohio"] = "俄亥俄"
	provinceMap["Coimbatore"] = "哥印拜陀"
	provinceMap["Philadelphia"] = "费城"
	provinceMap["South Australia"] = "南澳大利亚"
	provinceMap["Verona"] = "维罗纳"
	provinceMap["Irkutsk"] = "伊尔库茨克"
	provinceMap["Mangalore"] = "门格洛尔"
	provinceMap["Muenster"] = "曼斯特"
	provinceMap["Mus"] = "穆什"
	provinceMap["Omsk"] = "鄂木斯克"
	provinceMap["Ponta Delgada"] = "蓬塔德尔加达"
	provinceMap["Canakkale"] = "恰纳卡莱"
	provinceMap["Chalons-en-Champagne"] = "香槟地区沙隆"
	provinceMap["Paris"] = "巴黎"
	provinceMap["Potsdam"] = "波茨坦"
	provinceMap["Cagliari"] = "卡利亚里"
	provinceMap["Uruapan"] = "乌鲁阿潘"
	provinceMap["Entre Douro e Vouga"] = "恩特拉杜罗伏日"
	provinceMap["Gandhinagar"] = "甘地讷格尔"
	provinceMap["Pune"] = "浦那"
	provinceMap["Ryazan"] = "梁赞"
	provinceMap["Sakaka"] = "塞卡卡"
	provinceMap["Clark"] = "克拉克"
	provinceMap["Monaghan"] = "莫纳亨"
	provinceMap["Preah Vihear"] = "柏威夏"
	provinceMap["Wink"] = "温克"
	provinceMap["Monagas"] = "莫纳加斯"

}

func doneCityMap() {
	cityMap["Mount Vernon"] = "芒特弗农"
	cityMap["Gongju"] = "公州"
	cityMap["Beluran"] = "勿卢兰"
	cityMap["Franklin"] = "富兰克林"
	cityMap["Fort Franklin"] = "富兰克林堡"
	cityMap["Huntington"] = "亨廷顿"
	cityMap["Huntington Beach"] = "亨廷顿比奇"
	cityMap["Jiayuguan"] = "嘉峪关"
	cityMap["Pembroke Pines"] = "彭布罗克派恩斯"
	cityMap["Tokyo"] = "东京"
	cityMap["Birmingham"] = "伯明翰"
	cityMap["Ontario"] = "安大略"
	cityMap["Scranton"] = "斯克兰顿"
	cityMap["Chesapeake"] = "切萨皮克"
	cityMap["Guilin"] = "桂林"
	cityMap["Lloydminster"] = "劳埃德明斯特"
	cityMap["Santa Clara"] = "圣克拉拉"
	cityMap["Santa Clara County"] = "圣克拉拉县"
	cityMap["Beaufort"] = "保佛"
	cityMap["Cathedral City"] = "卡西德勒尔城"
	cityMap["Sepang"] = "雪邦"
	cityMap["Dezhou"] = "徳州"
	cityMap["Sacramento"] = "萨克拉门托"
	cityMap["Ulsan"] = "蔚山"
	cityMap["Launceston"] = "朗斯顿"
	cityMap["Liupanshui"] = "六盘水"
	cityMap["Gaithersburg"] = "盖瑟斯堡"
	cityMap["Gimhae"] = "金海"
	cityMap["Bucheon"] = "富川"
	cityMap["Dandong"] = "丹东"
	cityMap["Grand Forks"] = "大福克斯"
	cityMap["Saskatoon"] = "萨斯卡通"
	cityMap["Bradford"] = "布拉德福德"
	cityMap["El Monte"] = "埃尔蒙特"
	cityMap["Rockville"] = "罗克维尔"
	cityMap["Hengyang"] = "衡阳"
	cityMap["Zibo"] = "淄博"
	cityMap["Salford"] = "索尔福德"
	cityMap["Pine Valley"] = "派恩瓦利"
	cityMap["Port Augusta"] = "奥古斯塔港"
	cityMap["Shaoguan"] = "韶关"
	cityMap["Kent City"] = "肯特城"
	cityMap["Schenectady"] = "斯克内克塔迪"
	cityMap["Xianyang"] = "咸阳"
	cityMap["Conway"] = "康韦"
	cityMap["Oxnard"] = "奥克斯纳德"
	cityMap["Goose Bay"] = "古斯贝(哈皮瓦利-古斯贝的旧称)"
	cityMap["Whitehorse"] = "怀特霍斯"
	cityMap["Tuscaloosa"] = "塔斯卡卢萨县"
	cityMap["Hay River"] = "黑伊里弗"
	cityMap["Heihe"] = "黒河"
	cityMap["Saginaw"] = "萨吉诺"
	cityMap["Syracuse"] = "锡拉丘兹"
	cityMap["Paju"] = "坡州"
	cityMap["Suncheon"] = "顺天"
	cityMap["Dothan"] = "多森"
	cityMap["New York"] = "纽约"
	cityMap["Taebaek"] = "太白"
	cityMap["Tai Po"] = "大埔"
	cityMap["Auburn"] = "奥本"
	cityMap["Baoshan"] = "保山"
	cityMap["Edinburgh"] = "爱丁堡"
	cityMap["New Bedford"] = "纽柏德福特"
	cityMap["Whittier"] = "惠蒂尔"
	cityMap["Calgary"] = "卡尔加里"
	cityMap["Dongducheon"] = "东豆川"
	cityMap["Coventry"] = "考文垂"
	cityMap["Marysville"] = "马里斯维尔"
	cityMap["Eastern"] = "东方"
	cityMap["Worcester"] = "伍斯特"
	cityMap["Grand Prairie"] = "大草原城"
	cityMap["Hobart"] = "霍巴特"
	cityMap["Yorkton"] = "约克顿"
	cityMap["Osaka"] = "小坂"
	cityMap["Yichun"] = "伊春"
	cityMap["Tempe"] = "坦佩"
	cityMap["Kansas City"] = "堪萨斯城"
	cityMap["Palmdale"] = "帕姆代尔"
	cityMap["Cathedral City"] = "卡西德勒尔城"
	cityMap["Cheongju"] = "清州"
	cityMap["Lille"] = "里尔"
	cityMap["Rocky Mount"] = "落基山城"
	cityMap["Buena Park"] = "比埃纳帕克"
	cityMap["Collier County"] = "科利尔县"
	cityMap["Davis City"] = "戴维斯城"
	cityMap["London"] = "伦敦"
	cityMap["Fort Saskatchewan"] = "萨斯喀彻温堡"
	cityMap["Inverness"] = "因弗内斯"
	cityMap["St. Paul"] = "圣保罗"
	cityMap["Bortala"] = "博尔塔拉"
	cityMap["Columbus"] = "哥伦布"
	cityMap["Redding"] = "雷丁"
	cityMap["Vaughan"] = "沃恩"
	cityMap["Wabash"] = "沃巴什"
	cityMap["Richmond Hill"] = "里士满希尔"
	cityMap["Fredericton"] = "弗雷德里克顿"
	cityMap["Hesperia"] = "希斯皮里亚"
	cityMap["Waterloo"] = "滑铁卢"
	cityMap["Dongguan"] = "东莞"
	cityMap["Stamford"] = "斯坦福德"
	cityMap["Pasadena"] = "帕萨迪纳"
	cityMap["South San Francisco"] = "南圣弗朗西斯科"
	cityMap["San Francisco"] = "圣弗朗西斯科"
	cityMap["Gwangyang"] = "光阳"
	cityMap["Yangjiang"] = "阳江"
	cityMap["Delta"] = "德尔塔"
	cityMap["Las Vegas"] = "拉斯维加斯"
	cityMap["Avondale"] = "埃文代尔"
	cityMap["Dauphin"] = "多芬"
	cityMap["Heyuan"] = "河源"
	cityMap["Bethlehem"] = "伯利恒"
	cityMap["Penghu County"] = "彭湖县(台湾省下辖的县)"
	cityMap["Boise"] = "博伊西"
	cityMap["Eugene"] = "尤金"
	cityMap["Launceston"] = "朗斯顿"
	cityMap["Lacombe"] = "拉科姆"
	cityMap["Liverpool"] = "利物浦"
	cityMap["Xingtai"] = "邢台"
	cityMap["Eastmain"] = "伊斯特梅恩"
	cityMap["Holman"] = "霍尔曼"
	cityMap["San Diego"] = "圣迭戈"
	cityMap["Kuala Pilah"] = "瓜拉比拉"
	cityMap["Palmerston"] = "帕默斯顿"
	cityMap["Washington County"] = "华盛顿县"
	cityMap["Kudat"] = "古达"
	cityMap["Regina"] = "里贾纳"
	cityMap["Rochester"] = "罗切斯特"
	cityMap["San Antonio"] = "圣安东尼奥"
	cityMap["Troy"] = "特洛伊"
	cityMap["Hollywood"] = "好莱坞"
	cityMap["Geelong"] = "吉朗"
	cityMap["Beaverton"] = "比弗顿"
	cityMap["Baoji"] = "宝鸡"
	cityMap["Gimcheon"] = "金泉"
	cityMap["Burlington"] = "伯灵顿"
	cityMap["Jinan"] = "镇安"
	cityMap["Orlando"] = "奥兰多"
	cityMap["Salisbury"] = "索尔兹伯里"
	cityMap["Bowling Green"] = "鲍灵格林"
	cityMap["Rockhampton"] = "罗克汉普顿"
	cityMap["New Westminster"] = "新威斯敏斯特"
	cityMap["Plantation"] = "普兰泰申"
	cityMap["Po"] = "波村"
	cityMap["Chiayi City"] = "嘉义市(台湾省下辖的市)"
	cityMap["Zibo"] = "淄博"
	cityMap["Castle Rock"] = "罗克堡"
	cityMap["El Paso"] = "埃尔帕索"
	cityMap["Ganzi"] = "甘济"
	cityMap["Meadow Lake"] = "梅多莱克"
	cityMap["Sanya"] = "萨尼亚"
	cityMap["Altay"] = "阿尔泰"
	cityMap["Pendang"] = "彭当"
	cityMap["Tacoma"] = "塔科马"
	cityMap["Goyang"] = "高阳"
	cityMap["Kunak"] = "库纳克"
	cityMap["Tambunan"] = "淡汶安(淡巫南)"
	cityMap["Yongin"] = "龙仁"
	cityMap["Trenton"] = "托伦顿"
	cityMap["Old Crow"] = "旧克罗"
	cityMap["Palm Bay"] = "棕榈湾"
	cityMap["Columbia"] = "哥伦比亚"
	cityMap["Drummondville"] = "德拉蒙德维尔"
	cityMap["Houston"] = "休斯敦"
	cityMap["Gaspe"] = "加斯佩"
	cityMap["Xiangtan"] = "湘潭"
	cityMap["Brampton"] = "布兰普顿"
	cityMap["Bunbury"] = "班伯里"
	cityMap["Fort Lee"] = "利堡"
	cityMap["Zibo"] = "淄博"
	cityMap["Detroit"] = "底特律"
	cityMap["Liupanshui"] = "六盘水"
	cityMap["Ann Arbor"] = "安阿伯"
	cityMap["Kaohsiung City"] = "高雄市(台湾省下辖的市)"
	cityMap["Pingliang"] = "平凉"
	cityMap["Wan Chai"] = "湾仔"
	cityMap["Lille"] = "里尔"
	cityMap["Carson City"] = "卡森城"
	cityMap["Dalian"] = "大连"
	cityMap["Winchester"] = "温切斯特"
	cityMap["Liuzhou"] = "柳州"
	cityMap["Thorold"] = "索罗尔德"
	cityMap["Worcester"] = "伍斯特"
	cityMap["Yorkton"] = "约克顿"
	cityMap["Miaoli County"] = "苗栗县(台湾省下辖的县)"
	cityMap["Anqing"] = "安庆"
	cityMap["Longyan"] = "龙岩"
	cityMap["Pasir Mas"] = "巴西马士"
	cityMap["Arar"] = "阿尔阿尔"
	cityMap["Jasin"] = "野新"
	cityMap["Mayo"] = "梅奥"
	cityMap["Dauphin"] = "多芬"
	cityMap["Edinburgh"] = "爱丁堡"
	cityMap["Toledo"] = "托莱多"
	cityMap["Jecheon"] = "堤川"
	cityMap["Malden"] = "莫尔登"
	cityMap["Cold Lake"] = "科尔德莱克"
	cityMap["Salt Lake City"] = "盐湖城"
	cityMap["Montgomery"] = "蒙哥马利"
	cityMap["Montgomery City"] = "蒙哥马利城"
	cityMap["Wrigley"] = "里格利"
	cityMap["Beihai"] = "北海"
	cityMap["Gloucester"] = "格洛斯特"
	cityMap["Ningbo"] = "宁波"
	cityMap["Stamford"] = "斯坦福德"
	cityMap["Port Pirie"] = "皮里港"
	cityMap["Maoming"] = "茂名"
	cityMap["Clearwater"] = "克利尔沃特"
	cityMap["Greensboro"] = "格林斯伯勒"
	cityMap["Morden"] = "莫登"
	cityMap["Linfen"] = "临汾"
	cityMap["Gongju"] = "公州"
	cityMap["Winnipeg"] = "温尼伯"
	cityMap["Lumut"] = "卢穆特(红土坎)"
	cityMap["Port Colborne"] = "科尔本港"
	cityMap["Donghae"] = "东海"
	cityMap["Oklahoma City"] = "俄克拉何马城"
	cityMap["Sandy"] = "桑迪"
	cityMap["Shaoguan"] = "韶关"
	cityMap["Inverness"] = "因弗内斯"
	cityMap["Oxford"] = "牛津"
	cityMap["Beibei"] = "北碚"
	cityMap["Butterworth"] = "北海"
	cityMap["Boynton Beach"] = "博因顿比奇"
	cityMap["Chesapeake"] = "切萨皮克"
	cityMap["Santa Rosa"] = "圣罗莎"
	cityMap["Kuala Lumpur"] = "吉隆坡"
	cityMap["Paris"] = "巴黎"
	cityMap["Cheongju"] = "清州"
	cityMap["Vancouver"] = "温哥华"
	cityMap["Ontario"] = "安大略"
	cityMap["Caloundra"] = "日光海岸"
	cityMap["Kimberley"] = "金伯利"
	cityMap["Terrace"] = "特勒斯"
	cityMap["Hialeah"] = "海厄利亚"
	cityMap["Kitchener"] = "基奇纳"
	cityMap["Bengo"] = "本戈"
	cityMap["Karamay"] = "克拉玛依"
	cityMap["Kota Tinggi"] = "哥打丁宜"
	cityMap["Duncan"] = "邓肯"
	cityMap["Gaithersburg"] = "盖瑟斯堡"
	cityMap["Guilin"] = "桂林"
	cityMap["Mokpo"] = "木浦"
	cityMap["Flin Flon"] = "弗林弗伦"
	cityMap["Blackfoot"] = "布莱克富特"
	cityMap["Yiyang"] = "益阳"
	cityMap["Bangkok"] = "曼谷"
	cityMap["Denton"] = "登顿"
	cityMap["Hsinchu City"] = "新竹市(台湾省下辖的市)"
	cityMap["Santa Clara"] = "圣克拉拉"
	cityMap["Santa Clara County"] = "圣克拉拉县"
	cityMap["Brooklyn Park"] = "布鲁克林帕克"
	cityMap["Dearborn Heights"] = "迪尔伯恩海茨"
	cityMap["Lexington"] = "列克星敦"
	cityMap["Wuzhou"] = "梧州"
	cityMap["Baoding"] = "保定"
	cityMap["Heyuan"] = "河源"
	cityMap["Nanning"] = "南宁"
	cityMap["Ankang"] = "安康"
	cityMap["Ningde"] = "宁徳"
	cityMap["Sai Kung"] = "西贡"
	cityMap["Sejong"] = "世宗"
	cityMap["Des Moines"] = "得梅因"
	cityMap["Jining"] = "済宁"
	cityMap["Whitehorse"] = "怀特霍斯"
	cityMap["Clearwater"] = "克利尔沃特"
	cityMap["Tambunan"] = "淡汶安(淡巫南)"
	cityMap["Aba"] = "奥包"
	cityMap["Ocala"] = "奥卡拉"
	cityMap["Pawtucket"] = "波塔基特"
	cityMap["Philadelphia"] = "费城"
	cityMap["Edison"] = "爱迪生"
	cityMap["Halifax"] = "哈利法克斯"
	cityMap["Quincy"] = "昆西"
	cityMap["Roma"] = "罗马"
	cityMap["Saint Albans"] = "圣阿本斯"
	cityMap["Niagara Falls"] = "尼加拉瀑布城"
	cityMap["Sarikei"] = "泗里奎"
	cityMap["Union City"] = "尤宁城"
	cityMap["Offaly"] = "奥法利"
	cityMap["Overland Park"] = "欧弗兰帕克"
	cityMap["Rockford"] = "罗克福德"
	cityMap["Darwin"] = "达尔文"
	cityMap["Iowa City"] = "艾奥瓦城"
	cityMap["Los Angeles County"] = "洛杉矶县"
	cityMap["Townsville"] = "汤斯维尔"
	cityMap["Santa Monica"] = "圣莫尼卡"
	cityMap["Bowie"] = "鲍伊"
	cityMap["Labuan"] = "拉布安"
	cityMap["Taitung County"] = "台东县(台湾省下辖的县)"
	cityMap["Baldwin Park"] = "鲍德温帕克"
	cityMap["Manhattan"] = "曼哈顿"
	cityMap["Ogden"] = "奥格登"
	cityMap["Cambridge City"] = "剑桥城"
	cityMap["Franklin"] = "富兰克林"
	cityMap["Fort Franklin"] = "富兰克林堡"
	cityMap["Lacombe"] = "拉科姆"
	cityMap["Pocatello"] = "波卡特洛"
	cityMap["Eastvale"] = "伊斯特韦尔"
	cityMap["Norwich"] = "诺里季"
	cityMap["Kaohsiung City"] = "高雄市(台湾省下辖的市)"
	cityMap["Weifang"] = "潍坊"
	cityMap["Brant"] = "布兰特"
	cityMap["Hartford"] = "哈特福德"
	cityMap["Hobart"] = "霍巴特"
	cityMap["Kissimmee"] = "基西米"
	cityMap["Shizuishan"] = "石嘴山"
	cityMap["Tuscaloosa"] = "塔斯卡卢萨县"
	cityMap["Arctic Red River"] = "北极红河村"
	cityMap["Whyalla"] = "怀阿拉"
	cityMap["Cardiff"] = "加的夫"
	cityMap["Gimje"] = "金堤"
	cityMap["Merritt"] = "梅里特"
	cityMap["Carlisle"] = "卡莱尔"
	cityMap["Jasin"] = "野新"
	cityMap["Goose Bay"] = "古斯贝(哈皮瓦利-古斯贝的旧称)"
	cityMap["Oxnard"] = "奥克斯纳德"
	cityMap["Ranau"] = "拉瑙"
	cityMap["Broken Arrow"] = "布罗肯阿罗"
	cityMap["Conway"] = "康韦"
	cityMap["Gwangmyeong"] = "光明"
	cityMap["Washington County"] = "华盛顿县"
	cityMap["Canterbury"] = "坎特伯雷"

}
