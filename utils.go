package aig

import "go.mongodb.org/mongo-driver/bson"

const counterName string = "seq"
const idFieldName string = "id"

func (m *AI) getUpdate() bson.D {
	return bson.D{
		{"$inc", bson.D{
			{counterName, uint64(1)},
		}},
	}
}

func (m *AI) getFilter(name string) bson.M {
	return bson.M{
		idFieldName: name,
	}
}
