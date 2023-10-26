package selfintro

import (
	"fmt"

	"github.com/murasame29/docker-bussiness-card/utils"
)

type SelfIntro struct {
	tags        []string
	maxTagSize  int
	contents    []string
	maxContSize int
}

func NewSelfIntro() *SelfIntro {
	return &SelfIntro{}
}

func (s *SelfIntro) AddTag(tag string) *SelfIntro {
	s.tags = append(s.tags, tag)
	return s
}

func (s *SelfIntro) AddContent(content string) *SelfIntro {
	s.contents = append(s.contents, content)
	return s
}

func (s *SelfIntro) PrintIntro() (result string) {
	s.maxTagSize = utils.MaxLength(s.tags)
	s.maxContSize = utils.MaxLength(s.contents) + 2

	result += fmt.Sprintf("%s\n", utils.UpperLine(s.maxTagSize+s.maxContSize+3))
	for i := 0; i < len(s.tags); i++ {
		result += fmt.Sprintf("%s%s : %s%s\n", utils.Vl, utils.FillSpace(s.tags[i], s.maxTagSize), utils.FillSpace(s.contents[i], s.maxContSize), utils.Vl)
	}
	result += utils.LowerLine(s.maxTagSize + s.maxContSize + 3)
	return
}
