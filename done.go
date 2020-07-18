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
}
