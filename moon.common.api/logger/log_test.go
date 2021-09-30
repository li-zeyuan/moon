package logger

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestDebug(t *testing.T) {
	l := NewLogger(bson.NewObjectId().Hex())
	l.Debug("lizeyuan")
}

func TestInfo(t *testing.T) {
	l := NewLogger(bson.NewObjectId().Hex())
	l.Info("lizeyuan")
}

func TestLogger_Debugf(t *testing.T) {
	l := NewLogger(bson.NewObjectId().Hex())
	l.Debugf("l am %s", "lizeyuan")
}
