package database

import (
	"context"
	"os"
	"sync"
	"time"

	"goweb/log"
	"goweb/util/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var mongoOnce sync.Once
var mongoMutex sync.Mutex

var mongoDatabaseMap map[string]*mongo.Database

type MongoDB struct {
	collection *mongo.Collection
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func NewMongoDB(dbName string, collectionName string) *MongoDB {
	db := GetMongoDB(dbName)
	table := db.Collection(collectionName)
	return &MongoDB{collection: table}
}

func GetMongoDB(dbName string) (v *mongo.Database) {
	if dbName == "" {
		for _, v = range mongoDatabaseMap {
			return
		}
	}
	v = mongoDatabaseMap[dbName]
	return
}

// InsertOne insert one document
func (m MongoDB) InsertOne(document interface{}) (insertID interface{}) {
	insertResult, err := m.collection.InsertOne(getContext(), document)
	if err != nil {

		insertID = nil
		return
	}
	insertID = insertResult.InsertedID
	return
}

// InsertMany insert many document
func (m MongoDB) InsertMany(documents []interface{}, opts ...*options.InsertManyOptions) (insertIDs []interface{}) {
	insertResult, err := m.collection.InsertMany(getContext(), documents, opts...)
	if err != nil {
		log.Error(err)
	}
	insertIDs = insertResult.InsertedIDs
	return
}

// Find match doc
func (m MongoDB) Find(filter interface{}, results interface{}, opts ...*options.FindOptions) {
	cursor, err := m.collection.Find(getContext(), filter, opts...)
	if err != nil {
		log.Error(err)
		return
	}
	err = cursor.All(getContext(), results)
	if err != nil {
		log.Error(err)
	}
}

// FindOne match doc
func (m MongoDB) FindOne(filter interface{}, result interface{}, opts ...*options.FindOneOptions) {
	singleResult := m.collection.FindOne(getContext(), filter, opts...)
	if singleResult != nil {
		_ = singleResult.Decode(result)
	}
	return
}

// FindOneAndUpdate doc
func (m MongoDB) FindOneAndUpdate(filter interface{}, update interface{}, result interface{}, opts ...*options.FindOneAndUpdateOptions) {
	singleResult := m.collection.FindOneAndUpdate(getContext(), filter, update, opts...)
	if singleResult != nil {
		_ = singleResult.Decode(result)
	}
	return
}

// UpdateOne doc
func (m MongoDB) UpdateOne(filter interface{}, update interface{}) *mongo.UpdateResult {
	updateResult, err := m.collection.UpdateOne(getContext(), filter, update)
	if err != nil {
		log.Error(err)
	}
	return updateResult
}

// Delete One doc
func (m MongoDB) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) int64 {
	deleteResult, err := m.collection.DeleteOne(getContext(), filter, opts...)
	if err != nil {
		log.Error(err)
		return 0
	}
	return deleteResult.DeletedCount
}

// Delete Many doc, return delete count
func (m MongoDB) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) int64 {
	deleteResult, err := m.collection.DeleteMany(getContext(), filter, opts...)
	if err != nil {
		log.Error(err)
		return 0
	}
	return deleteResult.DeletedCount
}

func init() {
	mongoConfigs := config.ServerConfig().MongoDB
	mongoOnce.Do(func() {
		mongoMutex.Lock()
		defer mongoMutex.Unlock()

		mongoDatabaseMap = make(map[string]*mongo.Database, len(mongoConfigs))
		for _, mongoConfig := range mongoConfigs {
			var err error
			var client *mongo.Client

			opt := options.Client().ApplyURI(mongoConfig.Addr)
			opt.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
			opt.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
			opt.SetMaxPoolSize(200)                    //使用最大的连接数
			opt.SetReadPreference(readpref.Primary())  //使用指定节点
			opt.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
			wc := writeconcern.New(writeconcern.W(1))
			opt.SetWriteConcern(wc) //请求确认写操作传播到大多数mongod实例
			if client, err = mongo.Connect(getContext(), opt); err != nil {
				log.Fatalf("mongodb connect failed: %v", err)
				os.Exit(0)
			}

			//判断服务是否可用
			if err = client.Ping(getContext(), readpref.Primary()); err != nil {
				log.Fatalf("mongodb ping failed: %v", err)
				os.Exit(0)
			}
			// 保存mongodb database
			mongoDatabaseMap[mongoConfig.Database] = client.Database(mongoConfig.Database)
		}
	})
}
