package models

type PartnerRequest struct {
	Direction string `json:"direction"`
	MemberID  int    `json:"member_id"`
}

type PartnerList struct {
	MemberSourceID   int    `json:"member_source_id" db:"member_id1"`
	MemberTargetID   int    `json:"member_target_id" db:"member_id2"`
	MemberTargetName string `json:"member_target_name" db:"member_target_name"`
	Status           int    `json:"status" db:"status"`
	CreatedAt        string `json:"created_at" db:"created_at"`
}
