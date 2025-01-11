package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gClearMap(m map[int]int) {
	for k := range m {
		delete(m, k)
	}
}

func getCommonFriendsCount(userId1, userId2 int, usersFriends [][]int) int {
	res := 0
	for _, user1FriendId := range usersFriends[userId1] {
		for _, user2FriendId := range usersFriends[userId2] {
			if user1FriendId == user2FriendId {
				res++
			}
		}
	}
	return res
}

func G() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var usersCount, pairsCount int
	_, _ = fmt.Fscan(in, &usersCount, &pairsCount)

	usersFriends := make([][]int, usersCount)
	for i := 0; i < usersCount; i++ {
		usersFriends[i] = make([]int, 0, 5)
	}

	var f1, f2 int
	for i := 0; i < pairsCount; i++ {
		_, _ = fmt.Fscan(in, &f1, &f2)
		f1--
		f2--
		usersFriends[f1] = append(usersFriends[f1], f2)
		usersFriends[f2] = append(usersFriends[f2], f1)
	}

	possibleFriendsMap := make(map[int]int)
	possibleFriends := make([]int, 0, 25)
	for userId := 0; userId < usersCount; userId++ {
		gClearMap(possibleFriendsMap)
		maxCommonFriendsCount := 0
		for _, friendId := range usersFriends[userId] {
			for _, grandFriendId := range usersFriends[friendId] {
				if grandFriendId == userId {
					continue
				}
				grandFriendIsFriendOfUser := false
				for _, friendId_ := range usersFriends[userId] {
					if friendId_ == grandFriendId {
						grandFriendIsFriendOfUser = true
						break
					}
				}
				if grandFriendIsFriendOfUser {
					continue
				}
				if _, ok := possibleFriendsMap[grandFriendId]; ok {
					continue
				}
				commonFriendsCount := getCommonFriendsCount(userId, grandFriendId, usersFriends)
				maxCommonFriendsCount = max(maxCommonFriendsCount, commonFriendsCount)
				possibleFriendsMap[grandFriendId] = commonFriendsCount
			}
		}

		if len(possibleFriendsMap) == 0 {
			_, _ = fmt.Fprintln(out, 0)
			continue
		}

		possibleFriends = possibleFriends[:0]
		for friendId, commonFriendsCount := range possibleFriendsMap {
			if commonFriendsCount == maxCommonFriendsCount {
				possibleFriends = append(possibleFriends, friendId+1)
			}
		}

		sort.Slice(possibleFriends, func(i, j int) bool {
			return possibleFriends[i] < possibleFriends[j]
		})
		for i, possibleFriendId := range possibleFriends {
			_, _ = fmt.Fprint(out, possibleFriendId)
			if i+1 != len(possibleFriends) {
				_, _ = fmt.Fprint(out, " ")
			}
		}
		_, _ = fmt.Fprintln(out)
	}
}

func main() {
	G()
}
