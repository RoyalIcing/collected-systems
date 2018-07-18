package types

func optionalStrings(items []string) []*string {
	out := make([]*string, 0, len(items))
	for _, item := range items {
		out = append(out, &item)
	}
	return out
}
