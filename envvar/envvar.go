package envvar

import "gopkg.in/mgo.v2/bson"

// EnvVar represents an environmental variable for Runs and Jobs.
type EnvVar struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Key             string        `bson:"key" json:"key"`
	Value           string        `bson:"value" json:"value"`
	Type            string        `bson:"type" json:"type"`
	IsHiddenFromLog bool          `bson:"isHiddenFromLog" json:"isHiddenFromLog"`
	IsPrivate       bool          `bson:"isPrivate" json:"isPrivate"`
}
