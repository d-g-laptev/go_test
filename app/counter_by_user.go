package app

import "fmt"

type CounterByUserMap map[string]uint

func (c CounterByUserMap) incrementOrInit(userID string) {
	_, ok := c[userID]
	if !ok {
		c[userID] = 1
		return
	}
	c[userID]++
}

func (c CounterByUserMap) decrementOrDelete(userID string) {
	userCnt, ok := c[userID]
	if !ok {
		panic(fmt.Sprintf("User with name %s must be in map", userID))
	}
	if userCnt <= 1 {
		delete(c, userID)
		return
	}
	c[userID]--
}

func (c CounterByUserMap) getUniqueUsersCount() int {
	return len(c)
}
