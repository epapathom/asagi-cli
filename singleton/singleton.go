package singleton

import (
	"asagi/logic"
	"asagi/utils"
	"sync"
)

type Singleton struct{}

var once sync.Once

func (s *Singleton) Init() {
	once.Do(
		func() {
			utils.LoggerSingleton = new(utils.Logger)
			utils.LoggerSingleton.Init()

			logic.ECSSingleton = new(logic.ECS)
			logic.ECSSingleton.Init()

			logic.S3Singleton = new(logic.S3)
			logic.S3Singleton.Init()
		})
}
