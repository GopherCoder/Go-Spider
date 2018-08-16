## 懂球帝 API



BasePath|
---|
https://sport-data.dongqiudi.com|





### 1. 数据组



#### 中超

###### 1. 积分

Method|URL|
---|---|
GET|soccer/biz/data/standing?season_id=10219&app=dqd&version=162&platform=android|


Query|Content|
---|---|
season_id|10219|
app|dqd|
version|162|
plaform|android|


响应：

```json
{
	"template": "person_ranking",
	"content": {
		"data": [{
			"count": "162",
			"currency": "EUR",
			"person_id": "50000116",
			"person_logo": "https://img.dongqiudi.com/soccer/data/logo/person/116.jpg",
			"person_name": "梅西",
			"rank": "1",
			"team_id": "50001756",
			"team_logo": "https://img.dongqiudi.com/data/pic/2017.png",
			"team_name": "巴塞罗那",
			"value": "162000000"
		}]
	}
}
```

###### 2. 球员榜

```json
{
	"template": "ranking_types",
	"content": {
		"data": [{
			"name": "射手榜",
			"type": "goals",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=goals\u0026season_id=10219"
		}, {
			"name": "助攻榜",
			"type": "assists",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=assists\u0026season_id=10219"
		}, {
			"name": "关键传球",
			"type": "key_passes",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=key_passes\u0026season_id=10219"
		}, {
			"name": "射门",
			"type": "shots",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=shots\u0026season_id=10219"
		}, {
			"name": "射正",
			"type": "shots_on_target",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=shots_on_target\u0026season_id=10219"
		}, {
			"name": "越位",
			"type": "offsides",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=offsides\u0026season_id=10219"
		}, {
			"name": "传球",
			"type": "passes",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=passes\u0026season_id=10219"
		}, {
			"name": "成功传球",
			"type": "success_passes",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=success_passes\u0026season_id=10219"
		}, {
			"name": "拦截",
			"type": "interceptions",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=interceptions\u0026season_id=10219"
		}, {
			"name": "抢断",
			"type": "tackles",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=tackles\u0026season_id=10219"
		}, {
			"name": "解围",
			"type": "clearances",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=clearances\u0026season_id=10219"
		}, {
			"name": "犯规",
			"type": "fouls",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=fouls\u0026season_id=10219"
		}, {
			"name": "被犯规",
			"type": "fouled",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=fouled\u0026season_id=10219"
		}, {
			"name": "红牌",
			"type": "red_cards",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=red_cards\u0026season_id=10219"
		}, {
			"name": "黄牌",
			"type": "yellow_cards",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=yellow_cards\u0026season_id=10219"
		}, {
			"name": "扑救",
			"type": "saves",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=saves\u0026season_id=10219"
		}, {
			"name": "出场次数",
			"type": "appearances",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=appearances\u0026season_id=10219"
		}, {
			"name": "出场时间",
			"type": "time_played",
			"url": "https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd\u0026version=162\u0026platform=android\u0026type=time_played\u0026season_id=10219"
		}]
	}
}

```



####### 3. 球队榜


中文|英文|URL
---|---|---|
进球|goals|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=goals&season_id=10219|
点球|penalty_goals|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=penalty_goals&season_id=10219|
射门|shots|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=shots&season_id=10219|
射正|shots_on_target|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=shots_on_target&season_id=10219|
传球|passes|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=passes&season_id=10219|
关键传球|key_passes|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=key_passes&season_id=10219|
成功传球|success_passes|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=success_passes&season_id=10219|
传球成功率|pass_accuracy|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=pass_accuracy&season_id=10219|
传中成功率|crossing_accuracy|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=crossing_accuracy&season_id=10219|
拦截|interceptions|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=interceptions&season_id=10219|
抢断|tackles_won|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=tackles_won&season_id=10219|
解围|clearances|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=clearances&season_id=10219|
犯规|fouls_conceded|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=fouls_conceded&season_id=10219|
被侵犯|fouls_won|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=fouls_won&season_id=10219|
红牌|red_cards|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=red_cards&season_id=10219|
黄牌|yellow_cards|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=yellow_cards&season_id=10219|
击中门框|hit_woodwork|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=hit_woodwork&season_id=10219|

####### 4. 赛程

Method|URL|
---|---|
GET|https://sport-data.dongqiudi.com/soccer/biz/data/schedule?season_id=10219&app=dqd&version=162&platform=android|


##### 英超


**积分**

Method|URL|
---|---|
GET|/soccer/biz/data/standing?season_id=10267&app=dqd&version=162&platform=android

Query|Content|
---|---|
season_id|10267|
app|dqd|
version|162|
plaform|android|


响应：

```json
{
	"template": "team_point_ranking",
	"content": {
		"description": "英超联赛名次决定方法：\n（一）每队胜一场得3分，平一场得1分，负一场得0分；\n（二）英超联赛结束，积分多的队名次列前；\n（三）如果两队或两队以上积分相等，依下列顺序排列名次：\na）比较球队的联赛净胜球数；\nb）比较球队的联赛进球数；\nc）如果上述方法仍不能决定球队的排名且涉及到联赛冠军及降级名额，则两队于中立场进行一场附加赛决出胜负，其他则按有关队伍作相同排名。\n\n欧战&降级名额：\n（一）联赛前四直接参加下赛季欧冠联赛\n（二）第五名参加下赛季欧联杯\n（三）英格兰足总杯冠军可参加欧联杯，若其已获欧战资格，则名额让给联赛未获欧战资格中排名靠前的球队\n（四）英格兰联赛杯冠军可参加欧联杯第三轮资格赛，若其已获欧战资格，则名额让给联赛未获欧战资格中排名靠前的球队\n（五）若欧冠、欧联冠军都是英超球队，且两支球队都未通过联赛获得欧冠资格，则第四名参加欧联杯\n（六）排名最后三位的球队降入英冠联赛。\n",
		"rounds": [{
			"content": {
				"data": [{
					"goals_against": "0",
					"goals_pro": "4",
					"last_rank": "0",
					"matches_draw": "0",
					"matches_lost": "0",
					"matches_total": "1",
					"matches_won": "1",
					"points": "3",
					"rank": "1",
					"team_id": "50000516",
					"team_logo": "https://img.dongqiudi.com/data/pic/663.png",
					"team_name": "利物浦"
				}]
			}
		}]
	}
}

```

**球员榜**

类型|URL|
---|---|
射手榜|https://sport-data.dongqiudi.com/soccer/biz/data/person_ranking?app=dqd&version=162&platform=android&type=goals&season_id=10267|
助攻榜|assists|
关键传球|key_passes|
射门|shots|
射正|shots_on_target|
越位|offsides|
传球|passes|
成功传球|success_passes|
拦截|interception|
抢断|tackles|
解围|clearances|
犯规|fouls|
被犯规|fouled|
红牌|red_cards|
黄牌|yellow_cards|
扑救|saves|
出场次数|appearances|
出场时间|time_played|


**球队榜**

Method|URL|
---|---|
GET|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=goals&season_id=10267



类型|URL|
---|---|
关键传球|https://sport-data.dongqiudi.com/soccer/biz/data/team_ranking?app=dqd&version=162&platform=android&type=key_passes&season_id=10267|
成功传球|success_passes|
传球成功率|pass_accuracy|
传中成功率|crossing_accuracy|
拦截|interceptions|
抢断|tackles_won|
解围|clearances|
犯规|fouls_conceded|
被侵犯|fouls_won|
红牌|red_cards|
黄牌|yellow_cards|
击中门框|hit_woodwork|

**赛程**

Method|URL|
---|---|
GET|https://sport-data.dongqiudi.com/soccer/biz/data/schedule?season_id=10267&app=dqd&version=162&platform=android|

#### 亚运男足

**积分榜**

Method|URL|
---|---|
GET|/soccer/biz/data/standing?season_id=1815&app=dqd&version=162&platform=android|


**射手榜**

Method|URL|
---|---|
GET|/soccer/biz/data/person_ranking?season_id=1815&app=dqd&version=162&platform=android&type=goals|

**赛程**

Method|URL|
---|---|
GET|soccer/biz/data/schedule?season_id=1815&app=dqd&version=162&platform=android|

#### 亚运女足


**积分榜**

Method|URL|
---|---|
GET|/soccer/biz/data/standing?season_id=1816&app=dqd&version=162&platform=android|


**射手榜**

Method|URL|
---|---|
GET|/soccer/biz/data/person_ranking?season_id=1816&app=dqd&version=162&platform=android&type=goals|

**赛程**

Method|URL|
---|---|
GET|/soccer/biz/data/schedule?season_id=1816&app=dqd&version=162&platform=android|


#### 西甲

**积分**

URL|
---|
/soccer/biz/data/standing?season_id=10332&app=dqd&version=162&platform=android|

**球员榜**

类型|URL|
---|---|
射手榜|/soccer/biz/data/person_ranking?app=dqd&version=162&platform=android&type=goals&season_id=10332|
助攻榜|assists|
关键传球|key_passes|
射门|shots|
射正|success_shots|
越位|
传球|
成功传球|
拦截|
抢断|
解围|
犯规|
被犯规|
红牌|
黄牌|
扑救|
出场次数|
出场时间|

**球队榜**

**赛程**