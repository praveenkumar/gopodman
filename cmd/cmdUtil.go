package cmd

import (
	"strings"
)

// ReposToMap parses the specified repotags and returns a map with repositories
// as keys and the corresponding arrays of tags as values.
func reposToMap(repotags []string) map[string][]string {
	// map format is repo -> tag
	repos := make(map[string][]string)
	for _, repo := range repotags {
		var repository, tag string
		if len(repo) > 0 {
			li := strings.LastIndex(repo, ":")
			repository = repo[0:li]
			tag = repo[li+1:]
		}
		repos[repository] = append(repos[repository], tag)
	}
	if len(repos) == 0 {
		repos["<none>"] = []string{"<none>"}
	}
	return repos
}

/*func shortID(id string) string {
	if len(id) > idTruncLength {
		return id[:idTruncLength]
	}
	return id
}*/
