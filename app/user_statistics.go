package app

import (
	"container/list"
	"sync"
	"time"
)

var statsInterval = time.Second * 10

var userStats = UserStatistics{
	cntByUserMap: CounterByUserMap{},
	requestsList: list.List{},
	Mutex:        &sync.Mutex{},
}

func HandleRequest(userID string) {
	userStats.handleNewRequest(userID)
	userStats.cleanOldRequests()
}

func GetUniqueCount() int {
	userStats.cleanOldRequests()
	return userStats.cntByUserMap.getUniqueUsersCount()
}

type UserStatistics struct {
	cntByUserMap CounterByUserMap
	requestsList list.List
	*sync.Mutex
}

type RequestInfo struct {
	userID string
	time   time.Time
}

func (us *UserStatistics) handleNewRequest(userID string) {
	us.Lock()
	us.cntByUserMap.incrementOrInit(userID)

	newRequest := RequestInfo{
		userID: userID,
		time:   time.Now(),
	}
	us.requestsList.PushFront(newRequest)
	us.Unlock()
}

func (us *UserStatistics) cleanOldRequests() {
	for {
		us.Lock()
		latestRequestElement := us.requestsList.Back()
		if latestRequestElement == nil {
			us.Unlock()
			return
		}

		lastestRequestInfo := latestRequestElement.Value.(RequestInfo)

		if time.Since(lastestRequestInfo.time) < statsInterval {
			us.Unlock()
			return
		}

		us.requestsList.Remove(latestRequestElement)
		us.cntByUserMap.decrementOrDelete(lastestRequestInfo.userID)
		us.Unlock()
	}
}
