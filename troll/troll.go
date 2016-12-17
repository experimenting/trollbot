package troll

import (
	"bytes"
	"fmt"
	"strings"
	template "text/template"
)

// Troll Struct
type Troll struct {
	repo     FeedRepository
	varsRepo VarRepository
	tmpl     *template.Template
}

// Feed Struct
type Feed struct {
	Text string
	Tags []string
}

//NewTroll Contructor inject the repository
func NewTroll(r FeedRepository, vr VarRepository) *Troll {

	tr := &Troll{
		repo:     r,
		varsRepo: vr,
	}

	tr.tmpl = template.New("troll")
	tr.tmpl = tr.tmpl.Funcs(tr.initFuncMap())
	return tr
}

//Troll generate a troll text
func (t *Troll) Troll(context string, to []string) (tr string, err error) {
	var f Feed
	var tmp *template.Template

	f, err = t.repo.GetByContext(context)
	if err != nil {
		return "", fmt.Errorf("Error fetching %s parsing: %s", context, err)
	}
	fmt.Print(f, "\n")
	tmp, err = t.tmpl.Parse(f.Text)

	if err != nil {
		return "", fmt.Errorf("Error template parsing: %s", err)
	}
	var doc bytes.Buffer
	err = tmp.Execute(&doc, nil)
	if err != nil {
		return "", fmt.Errorf("Error template executing: %s", err)
	}

	fmt.Print(to)
	if len(to) > 1 {
		first := to[:1]
		others := to[1:]
		return fmt.Sprintf("Hey %s %s \n cc %s.", strings.Join(first, ","), doc.String(), strings.Join(others, ", ")), nil
	}

	if len(to) > 0 {
		return fmt.Sprintf("Hey %s %s", strings.Join(to, ","), doc.String()), nil
	}

	return doc.String(), nil
}

func (t *Troll) GetKeywords() []string {
	topics := t.varsRepo.GetAllTopics()
	tags := t.repo.GetAllTags()

	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	elements := append(topics, tags...)
	for v := range elements {
		if !encountered[elements[v]] {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}

	return result
}

func (t *Troll) getRandomUniqueVar(varName string) string {
	str, err := t.varsRepo.GetRandomUniqueVar(varName)

	if err != nil {
		panic(err)
	}

	return str
}

func (t *Troll) initFuncMap() template.FuncMap {
	return template.FuncMap{
		"lang":     func() string { return t.getRandomUniqueVar("language") },
		"language": func() string { return t.getRandomUniqueVar("language") },
		"vip":      func() string { return t.getRandomUniqueVar("vip") },
		"os":       func() string { return t.getRandomUniqueVar("os") },
		"buzzword": func() string { return t.getRandomUniqueVar("buzzword") },
		"ide":      func() string { return t.getRandomUniqueVar("ide") },
	}
}
