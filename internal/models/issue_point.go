package models

type IssuePoint struct {
	Id          int64
	AddressId   int64
	IsAvailable bool
}

type IssuePointRetrieve struct {
	IssuePoint *IssuePoint
	Error      error
}
