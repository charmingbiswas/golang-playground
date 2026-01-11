package dsa

import (
	"github.com/charmingbiswas/golang-stl/heap"
)

type Pair struct {
	timestamp int
	tweetId   int
}

type Twitter struct {
	timestamp  int
	userTweets map[int][]Pair
	followMap  map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		timestamp:  0,
		userTweets: make(map[int][]Pair),
		followMap:  make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if tweets, ok := this.userTweets[userId]; ok {
		tweets = append(tweets, Pair{timestamp: this.timestamp, tweetId: tweetId})
		this.userTweets[userId] = tweets
	} else {
		newTweets := make([]Pair, 0, 10)
		newTweets = append(newTweets, Pair{timestamp: this.timestamp, tweetId: tweetId})
		this.userTweets[userId] = newTweets
	}
	this.timestamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	results := make([]int, 0, 10)

	comparator := func(a, b Pair) bool {
		return a.timestamp > b.timestamp
	}

	maxHeap := heap.NewHeapWithFunc(comparator)

	// first check if user follows anyone
	if mp, ok := this.followMap[userId]; ok {
		for key := range mp {
			if tweets, ok := this.userTweets[key]; ok {
				for _, tweet := range tweets {
					maxHeap.Push(tweet)
				}
			}
		}
	}

	if tweets, ok := this.userTweets[userId]; ok {
		for _, tweet := range tweets {
			maxHeap.Push(tweet)
		}
	}
	for !maxHeap.IsEmpty() {
		if len(results) == 10 {
			break
		}
		results = append(results, maxHeap.Top().tweetId)
		maxHeap.Pop()
	}

	return results
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if mp, ok := this.followMap[followerId]; ok {
		mp[followeeId] = true
	} else {
		newMap := make(map[int]bool)
		newMap[followeeId] = true
		this.followMap[followerId] = newMap
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if mp, ok := this.followMap[followerId]; ok {
		delete(mp, followeeId)
	}
}
