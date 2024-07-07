package thk

type ThkContributor struct {
	Login string               `json:"login"`
	Id    int                  `json:"id"`
	Total float64              `json:"total"`
	Repos []ThkContributorRepo `json:"repos"`
}

type ThkContributorRepo struct {
	Repo  string  `json:"repo"`
	Score float64 `json:"score"` // current repo's score
}

type tmpThk struct {
	repo         string
	score        float64
	contributors []*Contributor
}
