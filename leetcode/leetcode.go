package leetcode

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

//go:embed user_info.json
var userinfo string

//go:embed contest_info.json
var contestinfo string

type LeetCodeInfo struct {
	SiteRanking             int `json:"siteRanking"`
	Rating                  int `json:"rating"`
	GlobalRanking           int `json:"globalRanking"`
	GlobalTotalParticipants int `json:"globalTotalParticipants"`
	LocalRanking            int `json:"localRanking"`
	LocalTotalParticipants  int `json:"localTotalParticipants"`
}

var LeetcodeInfo LeetCodeInfo

func setLeetCodeInfo() {
	info, err := getInfo()
	if err != nil {
		fmt.Println("Error fetching LeetCode info:", err)
		return
	}
	LeetcodeInfo = info
}

func init() {
	setLeetCodeInfo()
	c := cron.New()
	c.AddFunc("0 0 * * * *", func() {
		setLeetCodeInfo()
	})
	c.Start()
}

func GetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, LeetcodeInfo)
}

func getInfo() (LeetCodeInfo, error) {
	userInfo := getUserInfo()
	contestInfo := getContestInfo()
	return LeetCodeInfo{
		SiteRanking:             userInfo.Data.UserProfilePublicProfile.SiteRanking,
		Rating:                  int(contestInfo.Data.UserContestRanking.Rating),
		GlobalRanking:           contestInfo.Data.UserContestRanking.GlobalRanking,
		GlobalTotalParticipants: contestInfo.Data.UserContestRanking.GlobalTotalParticipants,
		LocalRanking:            contestInfo.Data.UserContestRanking.LocalRanking,
		LocalTotalParticipants:  contestInfo.Data.UserContestRanking.LocalTotalParticipants,
	}, nil
}

type ContestInfo struct {
	Data struct {
		UserContestRanking struct {
			Rating                  float32 `json:"rating"`
			GlobalRanking           int     `json:"globalRanking"`
			GlobalTotalParticipants int     `json:"globalTotalParticipants"`
			LocalRanking            int     `json:"localRanking"`
			LocalTotalParticipants  int     `json:"localTotalParticipants"`
		} `json:"userContestRanking"`
	} `json:"data"`
}

func getContestInfo() ContestInfo {
	resp, err := http.Post("https://leetcode.cn/graphql/noj-go/", "application/json", strings.NewReader(contestinfo))
	if err != nil {
		return ContestInfo{}
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ContestInfo{}
	}
	var contestInfo ContestInfo
	err = json.Unmarshal(bytes, &contestInfo)
	if err != nil {
		return ContestInfo{}
	}
	return contestInfo
}

type UserInfoResp struct {
	Data struct {
		UserProfilePublicProfile struct {
			SiteRanking int `json:"siteRanking"`
		} `json:"userProfilePublicProfile"`
	} `json:"data"`
}

func getUserInfo() UserInfoResp {
	resp, err := http.Post("https://leetcode.cn/graphql/", "application/json", strings.NewReader(userinfo))
	if err != nil {
		return UserInfoResp{}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfoResp{}
	}
	var userInfo UserInfoResp
	err = json.Unmarshal(bytes, &userInfo)
	if err != nil {
		return UserInfoResp{}
	}
	return userInfo
}
