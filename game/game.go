//package game
package main

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
	"werewolf-cli/game/role"
)

type roles struct {
	roleType string
	count    int
}
type players struct {
	uid  uuid.UUID
	name string
	role role.Role
}
type gameStatus struct {
	maxPCount int `maxPlayerCount`
	jPCount   int `joinedPlayerCount`
}

var roleList []role.Role
var playList []players // Player List
var status gameStatus  // Game Status

func init() {
	rand.Seed(time.Now().Unix())
}

func InitGame() {
	// TODO: 创建游戏
	rL := []roles{
		{"werewolf", 4},
		{"villager", 4},
		{"seer", 1},
		{"witch", 1},
		{"hunter", 1},
		{"idiot", 1},
	}
	_ = createRoles(rL)
	status = gameStatus{len(roleList), 0}
}

func createRoles(rL []roles) bool {
	rF := new(role.RoleFactory)
	for _, r := range rL {
		if r.count > 0 {
			for i := 0; i < r.count; i++ {
				newRole := rF.CreateRole(r.roleType)
				roleList = append(roleList, newRole)
			}
		} else {
			break
		}
	}
	return true
}
func joinPlayer(name string, uid uuid.UUID) bool {
	if status.jPCount < status.maxPCount {
		newPlayer := players{uid, name, nil}
		playList = append(playList, newPlayer)
		status.jPCount++
		return true
	} else {
		return false
	}
}

func randomList(length int) ([]int, error) {
	if length <= 0 {
		return nil, errors.New("the size of the parameter length illegal")
	}
	var list []int
	for i := 0; i < length; i++ {
		list = append(list, i)
	}
	for i := len(list) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		list[i], list[num] = list[num], list[i]
	}
	return list, nil
}

func assignRole() bool {
	if status.jPCount == status.maxPCount {
		randList, _ := randomList(status.maxPCount)
		for i := 0; i < len(randList); i++ {
			playList[i].role = roleList[randList[i]]
		}
		return true
	} else {
		return false
	}
}

func night()  {

}
func day()  {

}
func main() {
	InitGame()

	for i := 0; i < len(roleList); i++ {
		name := "玩家" + strconv.Itoa(i+1)
		uid := uuid.NewV4()
		_ = joinPlayer(name, uid)
	}

	_ = assignRole()

	//player
	for _, p := range playList {
		fmt.Printf("%+v\n", p)
	}
	for _, r := range roleList {
		r.Kill()
		fmt.Printf("%+v\n", r)
	}
	//for _, r := range roleList {
	//	fmt.Printf("%+v\n", r)
	//}
	// player
	for _, p := range playList {
		p.role.Kill()
		fmt.Printf("%+v\n", p.role)
	}
	for _, r := range roleList {
		fmt.Printf("%+v\n", r)
	}
}
