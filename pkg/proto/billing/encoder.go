package billing

import (
	"github.com/ProtocolONE/payone-repository/tools"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type MgoExpire struct {
	Month string `bson:"month" json:"month"`
	Year  string `bson:"year" json:"year"`
}

type MgoSavedCard struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Account    string        `bson:"account" json:"account"`
	ProjectId  bson.ObjectId `bson:"project_id" json:"project_id"`
	MaskedPan  string        `bson:"masked_pan" json:"masked_pan"`
	Pan        string        `bson:"pan" json:"pan"`
	CardHolder string        `bson:"card_holder" json:"card_holder"`
	Expire     *MgoExpire    `bson:"expire" json:"expire"`
	IsActive   bool          `bson:"is_active" json:"is_active"`
	CreatedAt  time.Time     `bson:"created_at" json:"created_at"`
}

func (s *SavedCard) GetBSON() (interface{}, error) {
	st := &MgoSavedCard{
		Id:         tools.ByteToObjectId(s.Id),
		Account:    s.Account,
		ProjectId:  tools.ByteToObjectId(s.ProjectId),
		MaskedPan:  s.MaskedPan,
		Pan:        s.Pan,
		CardHolder: s.CardHolder,
		Expire: &MgoExpire{
			Month: s.Expire.Month,
			Year:  s.Expire.Year,
		},
		IsActive: s.IsActive,
	}

	t, err := ptypes.Timestamp(s.CreatedAt)

	if err != nil {
		return nil, err
	}

	st.CreatedAt = t

	return st, nil
}

func (s *SavedCard) SetBSON(raw bson.Raw) error {
	decoded := new(MgoSavedCard)
	err := raw.Unmarshal(decoded)

	if err != nil {
		return err
	}

	s.Id = tools.ObjectIdToByte(decoded.Id)
	s.Account = decoded.Account
	s.ProjectId = tools.ObjectIdToByte(decoded.ProjectId)
	s.MaskedPan = decoded.MaskedPan
	s.Pan = decoded.Pan
	s.CardHolder = decoded.CardHolder
	s.Expire = &CardExpire{
		Month: decoded.Expire.Month,
		Year:  decoded.Expire.Year,
	}
	s.IsActive = decoded.IsActive
	s.CreatedAt, err = ptypes.TimestampProto(decoded.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}
