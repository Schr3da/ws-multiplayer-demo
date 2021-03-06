package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"time"
)

//ReadBytes Maps []byte to a provided struct
func ReadBytes(data []byte, dest interface{}) error {
	if err := json.Unmarshal(data, &dest); err != nil {
		CatchError("ReadBytes", err)
		return err
	}
	return nil
}

//DataToBytes Create
func DataToBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

//IsStringEmpty Check for emptyness
func IsStringEmpty(s string) (bool, error) {
	if s == "" || len(s) == 0 {
		err := errors.New("Provided string is empty")
		CatchError("generateID", err)
		return true, err
	}
	return false, nil
}

//NewError Creates a new Error
func NewError(s string) error {
	return errors.New(s)
}

//PrintLog Print message to console
func PrintLog(s string) {
	if *LogsEnabled == false {
		return
	}

	log.Println(s)
}

//GetRandomValue Returns a random number from a seed
func GetRandomValue(min int, max int) int {
	return rand.Intn(max-min) + min
}

//SetTimeout delays an function call
func SetTimeout(someFunc func(), milliseconds int) {
	timeout := time.Duration(milliseconds) * time.Millisecond
	time.AfterFunc(timeout, someFunc)
}

//ClonePlayerMap Deep Clone map
func ClonePlayerMap(toClone map[string]Player) map[string]Player {
	var clonedPlayerMap = make(map[string]Player)
	for k, v := range toClone {
		clonedPlayerMap[k] = Player{
			WSPlayerData: WSPlayerData{
				PseudoID:  v.PseudoID,
				Direction: v.Direction,
				Position:  v.Position,
				Plane:     v.Plane,
			},
			id:     v.id,
			health: v.health,
			mode:   v.mode,
		}
	}
	return clonedPlayerMap
}

//CloneWorld Deep Clone of world
func CloneWorld(toClone [MapTilesY][MapTilesX]string) [MapTilesY][MapTilesX]string {
	var clonedWorld [MapTilesY][MapTilesX]string
	for y := 0; y < MapTilesY; y++ {
		for x := 0; x < MapTilesX; x++ {
			clonedWorld[y][x] = toClone[y][x]
		}
	}
	return clonedWorld
}
