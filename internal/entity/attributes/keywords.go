package attributes

type Keyword string

const (
	KeywordOutlaw Keyword = "Outlaw"
	KeywordLaw    Keyword = "Law"
)

var keywordConflicts = map[Keyword][]Keyword{
	KeywordOutlaw: {KeywordLaw},
	KeywordLaw:    {KeywordOutlaw},
}

type Keywords struct {
	defaultKeywords []Keyword
	keywords        []Keyword
}

func NewKeywords(keywords []Keyword) *Keywords {
	return &Keywords{
		defaultKeywords: keywords,
		keywords:        keywords,
	}
}

// AddKeyword adds a keyword to the Keywords list. if a keyword conflicts with existing keywords, it removes the conflicting ones first.
func (k *Keywords) AddKeyword(keyword Keyword) {
	if conflicts, exists := keywordConflicts[keyword]; exists {
		for _, conflict := range conflicts {
			k.removeKeyword(conflict)
		}
	}

	if !k.hasKeyword(keyword) {
		k.keywords = append(k.keywords, keyword)
	}
}

// removeKeyword removes a keyword from the Keywords list.
func (k *Keywords) removeKeyword(keyword Keyword) {
	for i, kw := range k.keywords {
		if kw == keyword {
			k.keywords = append(k.keywords[:i], k.keywords[i+1:]...)
			break
		}
	}
}

// hasKeyword checks if a keyword exists in the Keywords list.
func (k *Keywords) hasKeyword(keyword Keyword) bool {
	for _, kw := range k.keywords {
		if kw == keyword {
			return true
		}
	}
	return false
}

// GetKeywords returns the list of keywords.
func (k *Keywords) GetKeywords() []Keyword {
	return k.keywords
}

// isDefaultKeyword checks if a keyword exists in the default keywords list.
func (k *Keywords) isDefaultKeyword(keyword Keyword) bool {
	for _, kw := range k.defaultKeywords {
		if kw == keyword {
			return true
		}
	}
	return false
}

// RemoveKeyword removes a keyword from the Keywords list and restores any conflicted default keywords.
func (k *Keywords) RemoveKeyword(keyword Keyword) {
	k.removeKeyword(keyword)

	// Check if the removed keyword had conflicts, and restore any default keywords that were conflicted
	if conflicts, exists := keywordConflicts[keyword]; exists {
		for _, conflictedKeyword := range conflicts {
			// If the conflicted keyword was a default keyword and isn't already present, restore it
			if k.isDefaultKeyword(conflictedKeyword) && !k.hasKeyword(conflictedKeyword) {
				k.keywords = append(k.keywords, conflictedKeyword)
			}
		}
	}
}
