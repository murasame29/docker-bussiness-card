# docker-bussiness-card

## 使い方
リポジトリをクローンしてください。
CIを利用したい場合は、GitHubのsecretを設定してください。
| secret | 説明 |
|:---|:---|
| DOCKER_USERNAME | DockerHubのユーザー名 |
| DOCKER_PASSWORD | DockerHubのパスワード |

CIのままの設定では、DockerHubのユーザ名/bussiness-card:latestとしてイメージが作成されます。
きにくわなければ、.github/workflows/build.ymlの32,33行目を変更してください。

## selfintroパッケージについて
selfintroパッケージは、枠線とか綺麗に描画したいがために作ったものです。

```go
package main
import (
    "github.com/murasame29/docker-bussiness-card/selfintro"
)

func main() {
    i := selfintro.NewSelfIntro()
    i.AddTag("Name").AddContent("I'm Murasame29.")
    fmt.Println(i.PrintIntro())
}
```
みたいな使い方をしてください。
issueとかprとか大歓迎です。

一応、宣伝しときます。

https://zenn.dev/murasame29/articles/cf02058c577acb

いいねとかふぉろーとかしてくれると嬉しいです。