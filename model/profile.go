// @Time : 2018/5/26 17:40
// @Author : minigeek
package model

import "encoding/json"

type Profile struct {
	Name       string //姓名或者昵称
	Gender     string //性别
	Age        int    //年龄
	Height     int    //身高
	Weight     int    //体重
	Income     string //收入
	Marriage   string //婚姻状况
	Education  string //学历
	Occupation string //职业
	Hokou      string //户口
	Xinzuo     string //星座
	House      string //房子
	Car        string //车子
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	res, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(res, &profile)
	return profile, err
}
