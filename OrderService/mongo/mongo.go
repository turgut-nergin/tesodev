package mongo

import (
	"github.com/globalsign/mgo"
)

type Client struct {
	session *mgo.Session
}

// func GetMongoDB() *Client {
// 	host := "mongo-db:27017"
// 	dbName := "tesodev"
// 	session, err := mgo.Dial(host)
// 	if err != nil {
// 		panic("connection error.")
// 	}
// 	db := session.DB(dbName)
// 	return &Client{session: db}
// }

func NewClient(url string) *Client {
	session, err := mgo.Dial(url)
	if err != nil {
		panic("connection error!")
	}
	return &Client{session: session}
}

func (c *Client) NewSession() *mgo.Session {
	newSession := c.session.Copy()
	newSession.SetMode(mgo.Strong, true)
	return newSession
}
