package main

import (
	"fmt"

	"github.com/murasame29/docker-bussiness-card/selfintro"
	"github.com/murasame29/docker-bussiness-card/utils"
)

func main() {
	intro := selfintro.NewSelfIntro()
	intro.AddTag(fmt.Sprintf("👋 %s名前%s", utils.FontColor(185, 80, 25), utils.Reset)).AddContent("I'm Murasame29.")
	intro.AddTag(fmt.Sprintf("😼 %sGitHub%s", utils.FontColor(35, 134, 54), utils.Reset)).AddContent("https://github.com/murasame29")
	intro.AddTag(fmt.Sprintf("📫 %sEmail%s", utils.FontColor(35, 134, 54), utils.Reset)).AddContent("oogiriminister@gmail.com")
	intro.AddTag(fmt.Sprintf("🐦 %sTwitter%s", utils.FontColor(100, 100, 255), utils.Reset)).AddContent("https://twitter.com/fresh_salmon256")
	intro.AddTag(fmt.Sprintf("🍴 %s好きな食べ物%s", utils.FontColor(100, 100, 255), utils.Reset)).AddContent("🍖 焼き肉、山葵、🍣 寿司")
	intro.AddTag(fmt.Sprintf("📜 %s保有資格%s", utils.FontColor(100, 100, 255), utils.Reset)).AddContent("DD2種,基本情報,陸特2級")
	fmt.Println(intro.PrintIntro())
}
