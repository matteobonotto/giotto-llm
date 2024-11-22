package types

type ArcJob struct {
	JobId     int    `json:"jobId"`
	JobName   string `json:"jobName"`
	JobCmd    string `json:"jobCmd"`
	JobConfig string `json:"jobConfig"`
	Started   int    `json:"started"`
	Done      int    `json:"done"`
	Failed    int    `json:"failed"`
}
