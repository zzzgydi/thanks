package thk

import "sort"

func tempToThkContributor(temp []*tmpThk) []*ThkContributor {
	total := 0.0 // all repos score

	for _, t := range temp {
		total += t.score
	}

	conMap := make(map[int]*ThkContributor)

	for _, t := range temp {
		for _, c := range t.contributors {
			if _, ok := conMap[c.Id]; !ok {
				conMap[c.Id] = &ThkContributor{
					Login: c.Login,
					Id:    c.Id,
				}
			}

			curScore := t.score / total * c.Score

			conMap[c.Id].Total += curScore
			conMap[c.Id].Repos = append(conMap[c.Id].Repos, ThkContributorRepo{
				Repo:  t.repo,
				Score: curScore,
			})
		}
	}

	result := make([]*ThkContributor, 0, len(conMap))
	for _, c := range conMap {
		result = append(result, c)
	}

	// sort by total score
	sort.Slice(result, func(i, j int) bool {
		return result[i].Total > result[j].Total
	})

	return result
}
