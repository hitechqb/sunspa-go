package sdk

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sunspa/common"
)

type service struct {
	db      *gorm.DB
	engine  *gin.Engine
	ginPort string
	logger  common.Logger
	redis   RedisProvider
}

func NewService() *service {
	return &service{}
}

type Service interface {
	MustGet(key string) interface{}
	Init(hasDatabase bool, hasRedis bool) error
	Close()
	Handle(router func(Service))
	EngineGin() *gin.Engine
	Start()
}

func (s *service) Start() {
	addr := fmt.Sprintf(":%s", s.ginPort)
	if err := s.engine.Run(addr); err != nil {
		log.Fatalln(err)
	}
}

func (s *service) MustGet(key string) interface{} {
	if key == common.KeyMainDB {
		return s.db
	}

	if key == common.KeyMainLogger {
		return s.logger
	}

	if key == common.KeyMainRedis {
		return s.redis
	}

	return nil
}

func (s *service) Init(hasDatabase bool, hasRedis bool) error {
	if err := loadEnv(); err != nil {
		return err
	}

	logger := common.NewLogger()
	s.logger = logger

	if hasDatabase {
		conn := NewConnection()
		db, err := connectDB(conn)
		if err != nil {
			logger.Errorln("😖 [Database] - cannot connect db, err: ", err.Error())
			return err
		}

		db.LogMode(true)
		s.db = db
	}

	if hasRedis {
		redis := NewRedisProvider(logger)
		if err := redis.Connect(); err != nil {
			logger.Errorln("😖 [Redis] - cannot connect to redis, err:", err.Error())
			return err
		}

		s.redis = redis
	}

	logger.Infof(`🎉 Application running in environment [%s]`, os.Getenv("ENV"))

	s.ginPort = os.Getenv("GIN_PORT")
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.engine = gin.Default()
	return nil
}

func (s *service) Close() {
	_ = s.db.Close()
}

func (s *service) Handle(router func(Service)) {
	router(s)
}

func (s *service) EngineGin() *gin.Engine { return s.engine }

func loadEnv() error {
	var err error
	switch os.Getenv("ENV") {
	case "production":
		err = godotenv.Load("prod.env")
		break
	case "staging":
		err = godotenv.Load("stg.env")
		break
	default:
		err = godotenv.Load("dev.env")
	}

	if err != nil {
		return err
	}
	return nil
}

func connectDB(conn *Connection) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", conn.DbUser, conn.DbPassWord, conn.DbHost, conn.DbPort, conn.DbName)
	return gorm.Open(conn.Dialect, args)
}
