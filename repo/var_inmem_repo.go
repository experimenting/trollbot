package repo

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// VarRepository is the structure for the In memory repository
type InMemoryVarRepository struct {
	Data map[string][]string
}

func NewInMemoryVarRepository() *InMemoryVarRepository {
	data := map[string][]string{
		"language": []string{"php", "nodejs"},
	}
	rand.Seed(time.Now().Unix())
	return &InMemoryVarRepository{data}
}

func NewInMemoryVarRepositoryFromYML(filename string) *InMemoryVarRepository {
	var err error
	var data map[string][]string
	data, err = loadDataFrom(filename)
	if err != nil {
		log.Fatalf("Error: loading from %s %v", filename, err)
	}
	rand.Seed(time.Now().Unix())
	return &InMemoryVarRepository{data}
}

func loadDataFrom(filename string) (data map[string][]string, err error) {
	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(filename)
	if err != nil {
		err = fmt.Errorf("error loading bytes from filename %s %v", os.Args[1], err)
		return
	}

	data = map[string][]string{}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		err = fmt.Errorf("error unmarshalling %s %v", yamlFile, err)
		return
	}

	return
}

func (r *InMemoryVarRepository) GetAllTopics() []string {
	topics := make([]string, len(r.Data))
	for k := range r.Data {
		topics = append(topics, k)
	}

	return topics
}

func (r *InMemoryVarRepository) GetRandomUniqueVar(varType string) (s string, err error) {

	if len(r.Data[varType]) <= 0 {
		err = fmt.Errorf("VarType '%s' not exists or empty", varType)
		return s, err
	}

	index := rand.Intn(len(r.Data[varType]))

	return r.Data[varType][index], nil
}
