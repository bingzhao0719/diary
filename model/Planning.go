package model

type Planning struct {
	Id        int32
	Theme     string //规划主题
	Deadline  string //截止日期
	Period    string //周期
	CreatedAt string //规划创建时间
	Remarks   string //备注
	Status    int    //状态
}
