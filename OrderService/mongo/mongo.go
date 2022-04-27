package mongo

import (
	"github.com/globalsign/mgo"
)

type Client struct {
	session *mgo.Session
}

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
