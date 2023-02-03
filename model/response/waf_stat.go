package response

import "SamWaf/model"

type WafStat struct {
	AttackCountOfToday          int64 //今日攻击数量
	VisitCountOfToday           int64 //今天访问数量
	AttackCountOfYesterday      int64 //昨日攻击数量
	VisitCountOfYesterday       int64 //昨日访问数量
	AttackCountOfLastWeekToday  int64 //上周攻击数量
	VisitCountOfLastWeekToday   int64 //上周访问数量
	NormalIpCountOfToday        int64 //今日正常IP数量
	IllegalIpCountOfToday       int64 //今日非法IP数量
	NormalCountryCountOfToday   int64 //今日正常国家数量
	IllegalCountryCountOfToday  int64 //今日非法国家数量
	NormalProvinceCountOfToday  int64 //今日正常省份数量
	IllegalProvinceCountOfToday int64 //今日非法省份数量
	NormalCityCountOfToday      int64 //今日正常城市数量
	IllegalCityCountOfToday     int64 //今日非法城市数量
}
type WafStatRange struct {
	AttackCountOfRange map[int]int64 //区间攻击数量
	NormalCountOfRange map[int]int64 //区间正常数量
}
type WafCityStats struct {
	AttackCityOfRange map[string]int64 //区间攻击城市数量
	NormalCityOfRange map[string]int64 //区间正常城市数量
}
type WafIPStats struct {
	AttackIPOfRange []model.StatsIPCount //区间攻击IP数量
	NormalIPOfRange []model.StatsIPCount //区间正常IP数量
}