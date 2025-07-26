package leetcode

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

//go:embed user_info.json
var userinfo string

//go:embed contest_info.json
var contestinfo string

//go:embed profile_calendar.json
var profilecalendar string

type LeetCodeInfo struct {
	SiteRanking             int    `json:"siteRanking"`
	Rating                  int    `json:"rating"`
	GlobalRanking           int    `json:"globalRanking"`
	GlobalTotalParticipants int    `json:"globalTotalParticipants"`
	LocalRanking            int    `json:"localRanking"`
	LocalTotalParticipants  int    `json:"localTotalParticipants"`
	SubmissionCalendar      string `json:"submissionCalendar"`
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
	profilecalendar := getProfileCalendar()
	return LeetCodeInfo{
		SiteRanking:             userInfo.Data.UserProfilePublicProfile.SiteRanking,
		Rating:                  int(contestInfo.Data.UserContestRanking.Rating),
		GlobalRanking:           contestInfo.Data.UserContestRanking.GlobalRanking,
		GlobalTotalParticipants: contestInfo.Data.UserContestRanking.GlobalTotalParticipants,
		LocalRanking:            contestInfo.Data.UserContestRanking.LocalRanking,
		LocalTotalParticipants:  contestInfo.Data.UserContestRanking.LocalTotalParticipants,
		SubmissionCalendar:      profilecalendar.Data.UserCalendar.SubmissionCalendar,
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

type UserInfo struct {
	Data struct {
		UserProfilePublicProfile struct {
			SiteRanking int `json:"siteRanking"`
		} `json:"userProfilePublicProfile"`
	} `json:"data"`
}

func getUserInfo() UserInfo {
	resp, err := http.Post("https://leetcode.cn/graphql/", "application/json", strings.NewReader(userinfo))
	if err != nil {
		return UserInfo{}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}
	}
	var userInfo UserInfo
	err = json.Unmarshal(bytes, &userInfo)
	if err != nil {
		return UserInfo{}
	}
	return userInfo
}

type ProfileCalendarInfo struct {
	Data struct {
		UserCalendar struct {
			SubmissionCalendar string `json:"submissionCalendar"`
		} `json:"userCalendar"`
	} `json:"data"`
}

func getProfileCalendar() ProfileCalendarInfo {
	resp, err := http.Post("https://leetcode.cn/graphql/noj-go/", "application/json", strings.NewReader(profilecalendar))
	if err != nil {
		log.Fatalln("failed to get profile calendar:", err)
		return ProfileCalendarInfo{}
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("failed to read response body:", err)
		return ProfileCalendarInfo{}
	}
	var profileCalendar ProfileCalendarInfo
	err = json.Unmarshal(bytes, &profileCalendar)
	if err != nil {
		log.Fatalln("failed to unmarshal profile calendar:", err)
		return ProfileCalendarInfo{}
	}
	return profileCalendar
}
