# EVM-Docker
api and EVM

> api相关文件在目录`code/go-zero-api/crypto_api/`中

# 调用函数与验证数据

> 2023/4/2 卢恒润

## 运行命令

```shell
docker exec -it crypto_api_golang_1 bash
goctl api go -api ./secret_sharing.api -dir ./
go run secretsharing.go
```

## Secret_sharing

### Ss

```go
// ss
SecretSSReq {
	Name   string `json:"name"` // 1,2
	N      byte   `json:"n"`
	K      byte   `json:"k"`
	Secret string `json:"secret"`
	Shares string `json:"shares"`
}
SecretSSRes {
	Shares string `json:"shares"`
	Secret string `json:"secret"`
	}
```

#### Case1

> 调用函数
>
> ```go
> //给定参与者人数n，门限值k，秘密值secret，获得秘密分享结果shares
> shares := ss.Split(n, k, secret)
> ```
>
> 输入：`n，k，secret`
>
> 输出：秘密分享值`shares`

输入

```shell
curl -X GET http://localhost:8000/crypto/secretss -H 'Content-Type:application/json' -d '{"name":"1","n":6,"k":3,"secret":"AQIDBAQFBAVkyA==","shares":""}'
```

输出

```shell
{"shares":"WyJleUpZSWpveExDSlpJam9pWmpWelRXd3hZMkpXY1RGYVUxRTlQU0o5IiwiZXlKWUlqb3lMQ0paSWpvaVoxUnBTMWhwU0dWM1JrMWhWSGM5UFNKOSIsImV5SllJam96TENKWklqb2lMelpIUm5wWVRFRnJkbk51ZW1jOVBTSjkiLCJleUpZSWpvMExDSlpJam9pT0ZSd1owNHpVVnBFVGpKSmEzYzlQU0o5IiwiZXlKWUlqbzFMQ0paSWpvaWFqWk9kbkJEWTBoWWJsY3hSV2M5UFNKOSIsImV5SllJam8yTENKWklqb2lZMUZFY0dKV1NFTjVTWFl5UmtFOVBTSjkiXQ==","secret":""}
```

#### Case2

> 调用函数
>
> ```go
> //给定门限值k，秘密分享结果的子集{shares...}恢复原秘密secret1
> secret1 := ss.Combine(k, shares...)
> ```
>
> 输入：`k，shares`
>
> 输出：秘密值`secret1`

输入

```shell
curl -X GET http://localhost:8000/crypto/secretss -H 'Content-Type:application/json' -d '{"name":"2","n":6,"k":3,"secret":"AQIDBAQFBAVkyA==","shares":"WyJleUpZSWpveExDSlpJam9pWmpWelRXd3hZMkpXY1RGYVUxRTlQU0o5IiwiZXlKWUlqb3lMQ0paSWpvaVoxUnBTMWhwU0dWM1JrMWhWSGM5UFNKOSIsImV5SllJam96TENKWklqb2lMelpIUm5wWVRFRnJkbk51ZW1jOVBTSjkiLCJleUpZSWpvMExDSlpJam9pT0ZSd1owNHpVVnBFVGpKSmEzYzlQU0o5IiwiZXlKWUlqbzFMQ0paSWpvaWFqWk9kbkJEWTBoWWJsY3hSV2M5UFNKOSIsImV5SllJam8yTENKWklqb2lZMUZFY0dKV1NFTjVTWFl5UmtFOVBTSjkiXQ=="}'
```

输出

```shell
{"shares":"","secret":"AQIDBAQFBAVkyA=="}
```

可以看到，输出了一样得秘密值

### Vss

`api`文件

```go
// vss
SecretVssReq {
	Name      string `json:"name"` // 1 2 3 4
	Num       int    `json:"num"`
	T         int    `json:"t"`
	SecretNum int64  `json:"secretNum"`
	Secret    string `json:"secret"`
	Vs        string `json:"vs"`
	Shares    string `json:"shares"`
}
SecretVssRes {
	Secret  string `json:"secret"`
	Vs      string `json:"vs"`
	Shares  string `json:"shares"`
	Res     bool   `json:"res"`
	Secret1 string `json:"secret1"`
}

@handler VSS
get /crypto/secretvss (SecretVssReq) returns (SecretVssRes)
```

#### Case1

> 输入：`name = "1"`，`secretNum`
>
> 返回：`secret`
>
> 调用函数
>
> ```go
> //设置秘密值secret，
> secret := new(big.Int).SetInt64(1000)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretvss -H 'Content-Type:application/json' -d '{"name":"1","num":6,"t":3,"secretNum":1000,"secret":"1000","vs":"","shares":""}'
```

预期输出值

```shell
{"secret":"1000","vs":"","shares":"","res":false,"secret1":""}
```

#### Case2

> 输入：`name = "2"`，`num`，`t`，`secret`
>
> 返回：`Vs`，`Shares`
>
> 调用函数
>
> ```go
> //给定曲线ec,门限值t,秘密值secret，参与者num获得对秘密分享多项式的承诺vs和秘密分享结果share，以及错误信息err
> vs, shares, _ := vss.Create(ec, t, secret, num)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretvss -H 'Content-Type:application/json' -d '{"name":"2","num":6,"t":3,"secretNum":1000,"secret":"1000","vs":"","shares":""}'
```

预期返回值

```shell
{"secret":"","vs":"[{\"Curve\":\"P-384\",\"Coords\":[6095129518569365160805201704943709729922302445321289134209606244630599099111039677404723770482645868843443820854231,23155358023125676497813639640188377411255018691677149213323450617726661088102385806435697814607903413705247051556557]},{\"Curve\":\"P-384\",\"Coords\":[9252742986921333122164728049768051816151160809816718240308901778034953489680090582787150413994430149326237303212515,384666538185442772055320367442819651441376014409854116327600061028909746140464832084529611239894822082325337232522]},{\"Curve\":\"P-384\",\"Coords\":[27978123804038020421323030750577700053739880729352911416639808851678055214185759849515613443886398731257917696677535,12931803311834877901531513434202370747736487688449716979407543950629071225660168831723007451217318871623934072649707]}]","shares":"[{\"Threshold\":2,\"ID\":15342673715538555721950288420501968419632767140776945168177807018293298361718539557062669343357559609596322834133517,\"Share\":33828620878707718317029425202097979792738806094246966705291285753129620717926856852692053585955525523803417144093670},{\"Threshold\":2,\"ID\":12393854769127684920744859509091706310878538552554243622396176524255824735537837237456668576869624076782299795758843,\"Share\":28820141314378654571054889261853115456103905074906113236357519952694108543100730560126119081790115632056306600613682},{\"Threshold\":2,\"ID\":32607720364226890439380755247281238559260108966751734204313838039612000752943321529602017299445313348630301021502470,\"Share\":4935206293393797173767077840374175526292432960787174778759796911318211833743073607832208395373043869292412467170171},{\"Threshold\":2,\"ID\":28144425374960526547000645817573369064454309562717219549925947692817713137565066375196351928900202114096494350094776,\"Share\":19494138585219323094865626259476850057231242279534718754309989746771846759623124395427584403989008410883913627896537},{\"Threshold\":2,\"ID\":25963268517029240062053679774854269820353803024915733648435220991903826862085306743412344966483084559883274577980806,\"Share\":27503189255684262463330307774313495364836948292855224468546250872331559988680712116129610343617182817828223357790381},{\"Threshold\":2,\"ID\":16286581661753858593518467800054859376549650459262204904133562930937344544979841036400573772923625219547575662560858,\"Share\":1369833856913664830835741358264899866870040009467794760731470941951546222624590119920522013679655890161204989361711}]","res":false,"secret1":""}
```

#### Case3

> 输入：`name = "3"`，`t`，`vs`，`shares`
>
> 返回：`res`
>
> 调用函数
>
> ```go
> //给定曲线ec，门限值t，秘密分享多项式的承诺vs，对子秘密share[0]验证
> res := shares[0].Verify(ec, t, vs)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretvss -H 'Content-Type:application/json' -d '{"name":"3","num":6,"t":3,"secretNum":1000,"secret":"1000","vs":"[{\"Curve\":\"P-384\",\"Coords\":[6095129518569365160805201704943709729922302445321289134209606244630599099111039677404723770482645868843443820854231,23155358023125676497813639640188377411255018691677149213323450617726661088102385806435697814607903413705247051556557]},{\"Curve\":\"P-384\",\"Coords\":[9252742986921333122164728049768051816151160809816718240308901778034953489680090582787150413994430149326237303212515,384666538185442772055320367442819651441376014409854116327600061028909746140464832084529611239894822082325337232522]},{\"Curve\":\"P-384\",\"Coords\":[27978123804038020421323030750577700053739880729352911416639808851678055214185759849515613443886398731257917696677535,12931803311834877901531513434202370747736487688449716979407543950629071225660168831723007451217318871623934072649707]}]","shares":"[{\"Threshold\":2,\"ID\":15342673715538555721950288420501968419632767140776945168177807018293298361718539557062669343357559609596322834133517,\"Share\":33828620878707718317029425202097979792738806094246966705291285753129620717926856852692053585955525523803417144093670},{\"Threshold\":2,\"ID\":12393854769127684920744859509091706310878538552554243622396176524255824735537837237456668576869624076782299795758843,\"Share\":28820141314378654571054889261853115456103905074906113236357519952694108543100730560126119081790115632056306600613682},{\"Threshold\":2,\"ID\":32607720364226890439380755247281238559260108966751734204313838039612000752943321529602017299445313348630301021502470,\"Share\":4935206293393797173767077840374175526292432960787174778759796911318211833743073607832208395373043869292412467170171},{\"Threshold\":2,\"ID\":28144425374960526547000645817573369064454309562717219549925947692817713137565066375196351928900202114096494350094776,\"Share\":19494138585219323094865626259476850057231242279534718754309989746771846759623124395427584403989008410883913627896537},{\"Threshold\":2,\"ID\":25963268517029240062053679774854269820353803024915733648435220991903826862085306743412344966483084559883274577980806,\"Share\":27503189255684262463330307774313495364836948292855224468546250872331559988680712116129610343617182817828223357790381},{\"Threshold\":2,\"ID\":16286581661753858593518467800054859376549650459262204904133562930937344544979841036400573772923625219547575662560858,\"Share\":1369833856913664830835741358264899866870040009467794760731470941951546222624590119920522013679655890161204989361711}]"}'
```

预期输出值

```shell
{"secret":"","vs":"","shares":"","res":true,"secret1":""}
```

#### Case4

> 输入：`name = "4"`，`t`，`shares`
>
> 返回：`secret1`
>
> 调用函数
>
> ```go
> secret1, _ := subset.ReConstruct(ec)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretvss -H 'Content-Type:application/json' -d '{"name":"4","num":6,"t":3,"secretNum":1000,"secret":"","vs":"","shares":"[{\"Threshold\":2,\"ID\":15342673715538555721950288420501968419632767140776945168177807018293298361718539557062669343357559609596322834133517,\"Share\":33828620878707718317029425202097979792738806094246966705291285753129620717926856852692053585955525523803417144093670},{\"Threshold\":2,\"ID\":12393854769127684920744859509091706310878538552554243622396176524255824735537837237456668576869624076782299795758843,\"Share\":28820141314378654571054889261853115456103905074906113236357519952694108543100730560126119081790115632056306600613682},{\"Threshold\":2,\"ID\":32607720364226890439380755247281238559260108966751734204313838039612000752943321529602017299445313348630301021502470,\"Share\":4935206293393797173767077840374175526292432960787174778759796911318211833743073607832208395373043869292412467170171},{\"Threshold\":2,\"ID\":28144425374960526547000645817573369064454309562717219549925947692817713137565066375196351928900202114096494350094776,\"Share\":19494138585219323094865626259476850057231242279534718754309989746771846759623124395427584403989008410883913627896537},{\"Threshold\":2,\"ID\":25963268517029240062053679774854269820353803024915733648435220991903826862085306743412344966483084559883274577980806,\"Share\":27503189255684262463330307774313495364836948292855224468546250872331559988680712116129610343617182817828223357790381},{\"Threshold\":2,\"ID\":16286581661753858593518467800054859376549650459262204904133562930937344544979841036400573772923625219547575662560858,\"Share\":1369833856913664830835741358264899866870040009467794760731470941951546222624590119920522013679655890161204989361711}]"}'
```

预期输出值

```shell
{"secret":"","vs":"","shares":"","res":false,"secret1":"1000"}
```

可以发现和第一步设置的秘密值一样。

### Pvss

```go
	// pvss
	SecretPvssReq {
		Name      string `json:"name"` // 1 2 3 4
		N         int    `json:"n"`
		T         int    `json:"t"`
		X         string `json:"x"`
		H         string `json:"h"`
		Secret    string `json:"secret"`
		EncShares string `json:"encShares"`
		PubPoly   string `json:"pubPoly"`
	}
	SecretPvssRes {
		H         string `json:"h"`
		PubX      string `json:"pubX"`
		Prix      string `json:"prix"`
		Secret    string `json:"secret"`
		EncShares string `json:"encShares"`
		PubPoly   string `json:"pubPoly"`
		Res       bool   `json:"res"`
		Recovered string `json:"recovered"`
	}
```

#### Case1

> 获取秘密分享者的公私钥对，并获得秘密分享值`secret`
>
> 输入：`name = "1"`，`n`
>
> 返回：公私钥`x`，`X`，秘密分享值`secret`，`H`

输入

```shell
curl -X GET http://localhost:8000/crypto/secretpvss -H 'Content-Type:application/json' -d '{"name":"1","n":5,"t":0,"x":"","h":"","secret":"","encShares":"","pubPoly":"","xpri":""}'
```

预期输出

```shell
{"h":"\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\"","pubX":"[\"yKDgKD1y5Au/H8QX+qKGH1X5dOqXXZmwTfew5nW7D2k=\",\"UYl2wbBsHj5sHPjCPoi/sibL1vB2aU2OhjtR19trUK0=\",\"WsQGB1R0bVBGMV5nMVeHP5yof8P3OMb6ab1xowRaSJI=\",\"542n1Zr7jgbWXm6SqakifyME1BEWlcZTusxPNDZuZtY=\",\"Mq7LPeigwpM62/xj2dAuz0Uy8mgOvJk7fGpr2QUKuEc=\"]","prix":"[\"WxmYGooLslzIItpj4XGeWw7+N9EGZGI7qD4URu4GMwE=\",\"/U23kHrzxYzuqMh/0MZGiE4iu6R4+Xn31BM+WUKPdAk=\",\"FAxU5dWZN3/9Drx8Fo2F5D5WUlUSBh1TT5IORj9zkQQ=\",\"vzCGmhq9zhP6b1plkPsbDrk6YfCrWuPyoSudS4Zvzgo=\",\"piZYmS8PO0h3w499XWiRgjRj7cKlB0dMopxN5fGXhQE=\"]","secret":"\"At5WnhG1/QN43sIzvP2QHNk5M0cY4tsSdPbusHeVFwo=\"","encShares":"","pubPoly":"","res":false,"recovered":""}
```

#### Case2

> **秘密分发**
>
> 输入：公钥`X`，秘密`secret`，门限值`t`，`H`
>
> 返回：分享列表`enShares`，多项式承诺信息`pubPoly`
>
> ```go
> // （1） 秘密分配（分发者）：给定曲线信息suite，生成元H，获得加密的秘密分享信息列表encShares和对多项式的承诺信息pubPoly
> encShares, pubPoly, _ := pvss.CreateEnShare(suite, H, X, secret, t)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretpvss -H 'Content-Type:application/json' -d '{"name":"2","n":5,"t":2,"x":"[\"yKDgKD1y5Au/H8QX+qKGH1X5dOqXXZmwTfew5nW7D2k=\",\"UYl2wbBsHj5sHPjCPoi/sibL1vB2aU2OhjtR19trUK0=\",\"WsQGB1R0bVBGMV5nMVeHP5yof8P3OMb6ab1xowRaSJI=\",\"542n1Zr7jgbWXm6SqakifyME1BEWlcZTusxPNDZuZtY=\",\"Mq7LPeigwpM62/xj2dAuz0Uy8mgOvJk7fGpr2QUKuEc=\"]","h":"\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\"","secret":"\"At5WnhG1/QN43sIzvP2QHNk5M0cY4tsSdPbusHeVFwo=\"","encShares":"","pubPoly":"","xpri":""}'
```

预期输出

```shell
{"h":"","pubX":"","prix":"","secret":"","encShares":"[{\"SI\":0,\"SV\":\"J1KV0fA5YqVD7JLhWd9N17QcjJZPDfWlVnbTnI8rtwE=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"c4Y8aZTSEIJtjpNe8T9eoDmcl+wTyflHmTm2H6JdGwg=\",\"VG\":\"UAsrtzaN3L68TJTV/slZn+4QeMlHz8F5NZV43E4PczE=\",\"VH\":\"gy/AIBfzM6N88gzgiR6oeF9Xtlg03KerwiBw6W69I8E=\"},{\"SI\":1,\"SV\":\"DBWF35v7eLIxdcsSF91PW6YtkSmzEu8r9l9K9pD/8MU=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"6KzLts3szj19/1Q2mtpa1L2VeWCqxuzWWquavVjt2gA=\",\"VG\":\"XGDWy7qsHA0m3Oz6Yt5NyUhU6u9wfCsEVTe6hNblHbg=\",\"VH\":\"CINSyi5TpAsrexQJYCXKGubnpct1wrCqPyBGNQpkTzo=\"},{\"SI\":2,\"SV\":\"oybPD6JVTz9SWh1InvBt4EUWlEYV6Urzmd8mwUMPerA=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"D/MNXZEXSsnEDeKNYKKAK44xUxoiq/U3NtEBmfMfFQQ=\",\"VG\":\"4ujF6QD1/eSV6g+Bq3agKEJvrGHchb2feWdOAq0Zoa8=\",\"VH\":\"ofyPkiJDDo5fiOGKOWp3685hvyTqGtF2+v/ZjuS75AA=\"},{\"SI\":3,\"SV\":\"tb5icOw4j04PlVSLfv2hco0YxFY+zAR3dXLbGmCMhE0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"RJce3KiomtFdzHQi+fMluS44r3/KAoSqBuFj4dQbBQY=\",\"VG\":\"ddkMdFi9uE1mUlC/SRcqFWfqR/YB3tjipWGsV4pc2Rs=\",\"VH\":\"hEBPKh3EGlfzZLOeCBoD4Xk8PrOh30nR+/l3QVQdaOc=\"},{\"SI\":4,\"SV\":\"VbRDB3gUnHUtvMUnnRi3CEOS4+bP6+Na+Ctxot/VLw0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"Bg3RZSPfK1CfShAgR4QRyICchjuwRrCskf4yYrgcAAk=\",\"VG\":\"w8ghol4yyHZhMs4BQsipRr2ek8UOyS1ddjgAC4aQ4k4=\",\"VH\":\"21MjeSLUaWL8dY49XsWlKeBJnOG9nSuvS3iJrbPnCA4=\"}]","pubPoly":"{\"g\":\"Ed25519\",\"b\":\"\\\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\\\"\",\"commits\":\"[\\\"CQtj9fXfAtYSKg4Fbp7BupID77AHFVSPT4Q/UEIiw6o=\\\",\\\"2OMPSa68W6bBC2GPGei2kAKJK4WyV9B8VG5CyGwmS+o=\\\"]\"}","res":false,"recovered":""}
```

#### Case3

> 验证
>
> 输入：分发列表`encShares`，多项式承诺`pubPoly`，公钥`X`，`H`
>
> 返回：验证结果`res`
>
> ```go
> //  (2) 验证（任何人）:给定分享者公钥X[1]，多项式的承诺信息pubPoly，加密的分享信息encShares[1]，生成元H，验证加密的秘密分享信息的一致性
> res := pvss.VerifyEncShare(suite, H, X[1], pubPoly, encShares[1])
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretpvss -H 'Content-Type:application/json' -d '{"name":"3","n":5,"t":2,"x":"[\"yKDgKD1y5Au/H8QX+qKGH1X5dOqXXZmwTfew5nW7D2k=\",\"UYl2wbBsHj5sHPjCPoi/sibL1vB2aU2OhjtR19trUK0=\",\"WsQGB1R0bVBGMV5nMVeHP5yof8P3OMb6ab1xowRaSJI=\",\"542n1Zr7jgbWXm6SqakifyME1BEWlcZTusxPNDZuZtY=\",\"Mq7LPeigwpM62/xj2dAuz0Uy8mgOvJk7fGpr2QUKuEc=\"]","h":"\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\"","secret":"\"At5WnhG1/QN43sIzvP2QHNk5M0cY4tsSdPbusHeVFwo=\"","encShares":"[{\"SI\":0,\"SV\":\"J1KV0fA5YqVD7JLhWd9N17QcjJZPDfWlVnbTnI8rtwE=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"c4Y8aZTSEIJtjpNe8T9eoDmcl+wTyflHmTm2H6JdGwg=\",\"VG\":\"UAsrtzaN3L68TJTV/slZn+4QeMlHz8F5NZV43E4PczE=\",\"VH\":\"gy/AIBfzM6N88gzgiR6oeF9Xtlg03KerwiBw6W69I8E=\"},{\"SI\":1,\"SV\":\"DBWF35v7eLIxdcsSF91PW6YtkSmzEu8r9l9K9pD/8MU=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"6KzLts3szj19/1Q2mtpa1L2VeWCqxuzWWquavVjt2gA=\",\"VG\":\"XGDWy7qsHA0m3Oz6Yt5NyUhU6u9wfCsEVTe6hNblHbg=\",\"VH\":\"CINSyi5TpAsrexQJYCXKGubnpct1wrCqPyBGNQpkTzo=\"},{\"SI\":2,\"SV\":\"oybPD6JVTz9SWh1InvBt4EUWlEYV6Urzmd8mwUMPerA=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"D/MNXZEXSsnEDeKNYKKAK44xUxoiq/U3NtEBmfMfFQQ=\",\"VG\":\"4ujF6QD1/eSV6g+Bq3agKEJvrGHchb2feWdOAq0Zoa8=\",\"VH\":\"ofyPkiJDDo5fiOGKOWp3685hvyTqGtF2+v/ZjuS75AA=\"},{\"SI\":3,\"SV\":\"tb5icOw4j04PlVSLfv2hco0YxFY+zAR3dXLbGmCMhE0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"RJce3KiomtFdzHQi+fMluS44r3/KAoSqBuFj4dQbBQY=\",\"VG\":\"ddkMdFi9uE1mUlC/SRcqFWfqR/YB3tjipWGsV4pc2Rs=\",\"VH\":\"hEBPKh3EGlfzZLOeCBoD4Xk8PrOh30nR+/l3QVQdaOc=\"},{\"SI\":4,\"SV\":\"VbRDB3gUnHUtvMUnnRi3CEOS4+bP6+Na+Ctxot/VLw0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"Bg3RZSPfK1CfShAgR4QRyICchjuwRrCskf4yYrgcAAk=\",\"VG\":\"w8ghol4yyHZhMs4BQsipRr2ek8UOyS1ddjgAC4aQ4k4=\",\"VH\":\"21MjeSLUaWL8dY49XsWlKeBJnOG9nSuvS3iJrbPnCA4=\"}]","pubPoly":"{\"g\":\"Ed25519\",\"b\":\"\\\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\\\"\",\"commits\":\"[\\\"CQtj9fXfAtYSKg4Fbp7BupID77AHFVSPT4Q/UEIiw6o=\\\",\\\"2OMPSa68W6bBC2GPGei2kAKJK4WyV9B8VG5CyGwmS+o=\\\"]\"}","xpri":"[\"WxmYGooLslzIItpj4XGeWw7+N9EGZGI7qD4URu4GMwE=\",\"/U23kHrzxYzuqMh/0MZGiE4iu6R4+Xn31BM+WUKPdAk=\",\"FAxU5dWZN3/9Drx8Fo2F5D5WUlUSBh1TT5IORj9zkQQ=\",\"vzCGmhq9zhP6b1plkPsbDrk6YfCrWuPyoSudS4Zvzgo=\",\"piZYmS8PO0h3w499XWiRgjRj7cKlB0dMopxN5fGXhQE=\"]"}'
```

预期输出

```shell
{"h":"","pubX":"","prix":"","secret":"","encShares":"","pubPoly":"","res":true,"recovered":""}
```

<img src="./%E9%AA%8C%E8%AF%81%E6%95%B0%E6%8D%AE.assets/image-20230329224233529.png" alt="image-20230329224233529" style="zoom:80%;" />

#### Case4

> 解密获得子秘密，恢复机密
>
> 输入：分发列表`encShares`，多项式承诺`pubPoly`，公钥`X`，`H`
>
> 返回：秘密`recovered`
>
> ```go
> // 秘密恢复结果recovered
> recovered, _ := pvss.RecoverSecret(suite, G, K, E, D, t, n)
> ```

输入

```shell
curl -X GET http://localhost:8000/crypto/secretpvss -H 'Content-Type:application/json' -d '{"name":"4","n":5,"t":2,"x":"[\"yKDgKD1y5Au/H8QX+qKGH1X5dOqXXZmwTfew5nW7D2k=\",\"UYl2wbBsHj5sHPjCPoi/sibL1vB2aU2OhjtR19trUK0=\",\"WsQGB1R0bVBGMV5nMVeHP5yof8P3OMb6ab1xowRaSJI=\",\"542n1Zr7jgbWXm6SqakifyME1BEWlcZTusxPNDZuZtY=\",\"Mq7LPeigwpM62/xj2dAuz0Uy8mgOvJk7fGpr2QUKuEc=\"]","h":"\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\"","secret":"\"At5WnhG1/QN43sIzvP2QHNk5M0cY4tsSdPbusHeVFwo=\"","encShares":"[{\"SI\":0,\"SV\":\"J1KV0fA5YqVD7JLhWd9N17QcjJZPDfWlVnbTnI8rtwE=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"c4Y8aZTSEIJtjpNe8T9eoDmcl+wTyflHmTm2H6JdGwg=\",\"VG\":\"UAsrtzaN3L68TJTV/slZn+4QeMlHz8F5NZV43E4PczE=\",\"VH\":\"gy/AIBfzM6N88gzgiR6oeF9Xtlg03KerwiBw6W69I8E=\"},{\"SI\":1,\"SV\":\"DBWF35v7eLIxdcsSF91PW6YtkSmzEu8r9l9K9pD/8MU=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"6KzLts3szj19/1Q2mtpa1L2VeWCqxuzWWquavVjt2gA=\",\"VG\":\"XGDWy7qsHA0m3Oz6Yt5NyUhU6u9wfCsEVTe6hNblHbg=\",\"VH\":\"CINSyi5TpAsrexQJYCXKGubnpct1wrCqPyBGNQpkTzo=\"},{\"SI\":2,\"SV\":\"oybPD6JVTz9SWh1InvBt4EUWlEYV6Urzmd8mwUMPerA=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"D/MNXZEXSsnEDeKNYKKAK44xUxoiq/U3NtEBmfMfFQQ=\",\"VG\":\"4ujF6QD1/eSV6g+Bq3agKEJvrGHchb2feWdOAq0Zoa8=\",\"VH\":\"ofyPkiJDDo5fiOGKOWp3685hvyTqGtF2+v/ZjuS75AA=\"},{\"SI\":3,\"SV\":\"tb5icOw4j04PlVSLfv2hco0YxFY+zAR3dXLbGmCMhE0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"RJce3KiomtFdzHQi+fMluS44r3/KAoSqBuFj4dQbBQY=\",\"VG\":\"ddkMdFi9uE1mUlC/SRcqFWfqR/YB3tjipWGsV4pc2Rs=\",\"VH\":\"hEBPKh3EGlfzZLOeCBoD4Xk8PrOh30nR+/l3QVQdaOc=\"},{\"SI\":4,\"SV\":\"VbRDB3gUnHUtvMUnnRi3CEOS4+bP6+Na+Ctxot/VLw0=\",\"PC\":\"NZIBSQA9IvrodMQ9oV3gzg0mipFrT6e9lO6Jto0KaAg=\",\"PR\":\"Bg3RZSPfK1CfShAgR4QRyICchjuwRrCskf4yYrgcAAk=\",\"VG\":\"w8ghol4yyHZhMs4BQsipRr2ek8UOyS1ddjgAC4aQ4k4=\",\"VH\":\"21MjeSLUaWL8dY49XsWlKeBJnOG9nSuvS3iJrbPnCA4=\"}]","pubPoly":"{\"g\":\"Ed25519\",\"b\":\"\\\"3LoA1Wv1EpiRTLlQUhRT/ag2/SD7LnHHWsEVFuf8Xsc=\\\"\",\"commits\":\"[\\\"CQtj9fXfAtYSKg4Fbp7BupID77AHFVSPT4Q/UEIiw6o=\\\",\\\"2OMPSa68W6bBC2GPGei2kAKJK4WyV9B8VG5CyGwmS+o=\\\"]\"}","xpri":"[\"WxmYGooLslzIItpj4XGeWw7+N9EGZGI7qD4URu4GMwE=\",\"/U23kHrzxYzuqMh/0MZGiE4iu6R4+Xn31BM+WUKPdAk=\",\"FAxU5dWZN3/9Drx8Fo2F5D5WUlUSBh1TT5IORj9zkQQ=\",\"vzCGmhq9zhP6b1plkPsbDrk6YfCrWuPyoSudS4Zvzgo=\",\"piZYmS8PO0h3w499XWiRgjRj7cKlB0dMopxN5fGXhQE=\"]"}'
```

函数出现报错：

<img src="./2023-3-29%20%E9%AA%8C%E8%AF%81%E6%95%B0%E6%8D%AE%20230752.assets/image-20230402121545911.png" alt="image-20230402121545911" style="zoom:80%;" />

也就是说输入数据后返回了空置，导致后面报错。在调用：

```go
a, _ = pvss.DecShare(suite, H, X[i], x[i], enShares[i], pubPoly)
```

返回空值。

## Key_exchange

### Steal

```shell
	// GenerateKeySt Req
	GenerateKeyStReq {
		Name      string `json:"name"` // name = req1, req2, req3
		Str1      string `json:"str1"`
		Str2      string `json:"str2"`
		Req_r     string `json:"reqr"`     // req2
		Req_pub1  string `json:"reqpub1"`  // req2
		Req_pub2  string `json:"reqpub2"`  // req2,req3
		Req_P     string `json:"reqP"`     // req3
		Req_R     string `json:"reqR"`     // req3
		Req_priv1 string `json:"reqPriv1"` // req3
		Req_priv2 string `json:"reqPriv2"` // req3
	}
	// GenerateKeySt Res
	GenerateKeyStRes {
		Priv1 string `json:"priv1"` // req1
		Pub1  string `json:"pub1"`  // req1
		Priv2 string `json:"priv2"` // req1
		Pub2  string `json:"pub2"`  // req1
		P     string `json:"p"`     // req2
		R     string `json:"r"`     // req2
		T     bool   `json:"t"`		// req3
		Priv  string `json:"priv"` // req3
	}
```

#### Case1

> 调用函数
>
> ```go
> a, A, b, B := stealth_address.RecCalculateKeyPairs(str1, str2)
> ```
>
> 输入：`Str1，Str2，name = "1"`
>
> 输出：`Priv1 a，Pub1 A，Priv2 b，Pub2 B`

输入

```shell
curl -X GET http://localhost:8000/crypto/generatekeyst -H 'Content-Type:application/json' -d '{"name":"1","str1":"hello","str2":"you","reqr":"","reqpub1":"","reqpub2":"","reqP":"","reqR":"","reqPriv1":"","reqPriv2":""}'
```

返回

```shell
{"priv1":"20329878786436204988385760252021328656300425018755239228739303522659023427620","pub1":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","priv2":"20329878786436204988385760252021328656300425018755239228739303522659023427620","pub2":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","p":"","r":"","priv":""}
```

#### Case2

> 调用函数
>
> ```go
> P := stealth_address.SendCalculateObfuscateAddress(reqr, A, B)
> R := stealth_address.SendCalculatePublicKey(reqr))
> ```
>
> 输入：`reqr，reqpub1，reqpub2`
>
> 输出：`p，r`

输入

```shell
curl -X GET http://localhost:8000/crypto/generatekeyst -H 'Content-Type:application/json' -d '{"name":"2","str1":"hello","str2":"you","reqr":"1000","reqpub1":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","reqpub2":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","reqP":"","reqR":"","reqPriv1":"","reqPriv2":""}'
```

输出

```shell
{"priv1":"","pub1":"","priv2":"","pub2":"","p":"{\"X\":95479002247878751428610834670994115584124742924135938944738324937097549229600,\"Y\":12530620545552582535040653525052906683182466753764748726728834438188106319080}","r":"{\"X\":83667457367424346449913836100092076718726000981788703046900525431903421839106,\"Y\":69820075667727983197000401737468327286843439827068444235977723391007956708919}","priv":""}
```

#### Case3

> 调用函数
>
> ```go
> t := stealth_address.RecCalculateObfuscateAddress(P, R, a, B)
> 
> priv := stealth_address.RecCalculateAddressPrivatekey(a, b, R)
> ```
>
> 输入：`reqP，reqR，reqPriv1，reqPriv2`
>
> 输出：`t，priv`

输入

```shell
curl -X GET http://localhost:8000/crypto/generatekeyst -H 'Content-Type:application/json' -d '{"name":"3","str1":"hello","str2":"you","reqr":"1000","reqpub1":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","reqpub2":"{\"X\":53466807972227514710220285433514265750039844143986023941830406209187647784364,\"Y\":97578466278460447231433129414362077864850661904809116025127631511815465465044}","reqP":"{\"X\":95479002247878751428610834670994115584124742924135938944738324937097549229600,\"Y\":12530620545552582535040653525052906683182466753764748726728834438188106319080}","reqR":"{\"X\":83667457367424346449913836100092076718726000981788703046900525431903421839106,\"Y\":69820075667727983197000401737468327286843439827068444235977723391007956708919}","reqPriv1":"20329878786436204988385760252021328656300425018755239228739303522659023427620","reqPriv2":"20329878786436204988385760252021328656300425018755239228739303522659023427620"}'
```

输出

```shell
{"priv1":"","pub1":"","priv2":"","pub2":"","p":"","r":"","t":true,"priv":"35226410909133827193993112018665382520926747068382820998726431733653983940176"}
```

### Ecdh

```go
// GenerateKeyEcdh Req None
// GenerateKeyEcdh Res
GenerateKeyEcdh {
	SecretA string `json:"secretA"`
	SecretB string `json:"secretB"`
}
```

> 调用函数
>
> ```go
> shared1 := ecdh.CalculateNegotiationKey(publickey2, privatekey1)
> shared2 := ecdh.CalculateNegotiationKey(publickey1, privatekey2)
> ```

输入：无

```shell
curl -X GET http://localhost:8000/crypto/generatekeyecdh -H 'Content-Type:application/json'
```

输出：`secretA，secretB`

```shell
{"secretA":"\ufffdb\ufffd'\ufffd\u0010\u0010\t8\ufffd\ufffd\ufffd\ufffd\ufffd3\ufffd\ufffd\ufffd\ufffdRp\ufffd-\u00060;\u0002\ufffd\rhU\ufffd","secretB":"\ufffdb\ufffd'\ufffd\u0010\u0010\t8\ufffd\ufffd\ufffd\ufffd\ufffd3\ufffd\ufffd\ufffd\ufffdRp\ufffd-\u00060;\u0002\ufffd\rhU\ufffd"}
```

### Dh

```go
// GenerateKey Req None dh
// GenerateKey Res dh
GenerateKeyRes {
	SecretA string `json:"SecretA"`
	SecretB string `json:"SecretB"`
}
```

> 调用函数
>
> ```go
> //Alice计算会话密钥
> secretAlice := group.ComputeSecret(alicePrivate, bobPublic)
> //Bob计算会话密钥
> secretBob := group.ComputeSecret(bobPrivate, alicePublic)
> ```

输入：无

```shell
curl -X GET http://localhost:8000/crypto/generatekey -H 'Content-Type:application/json'
```

输出：`SecretA，SecretB`

```shell
{"SecretA":"342148992723789209576003990145122178507935795240990916792708173745468103361307411201373937317948455988886050619884667379844843858985356997990064711342851143785229870376178935141313482136340962345418517566455082573912362615430295092169453394296642418124722276188097960237096210128554313429564223472495076911477678409650868421227511690925589499639690431929009788014159396271269616996156953862404867572065296945693493604077416136223692826597333017454569409393537581966456939454950724438346974478108552543266505447897066401937787353128775658636671050545886283957254005994053382597994030205843794129478035548670169143120716183546839636620832206046130178840470912382088260040689396602749290725149279274843937698894336557963092719826592355136708260359035508112483559947061941595119585651015773217474124961000295759969227493901289642998828296716899454652536049446496906969419458622138816723825829435301821287090805463158656305975450920276963777923103017461201799125060572926044231032704796880489934603822037169658039101891501582036909397436250861123799518082274765578732293891061740875175280725140845569538266163576071003206509391159851507535336785090848512024036090894931823792446772092082764821222186551527833576365116844484278886443262212","SecretB":"342148992723789209576003990145122178507935795240990916792708173745468103361307411201373937317948455988886050619884667379844843858985356997990064711342851143785229870376178935141313482136340962345418517566455082573912362615430295092169453394296642418124722276188097960237096210128554313429564223472495076911477678409650868421227511690925589499639690431929009788014159396271269616996156953862404867572065296945693493604077416136223692826597333017454569409393537581966456939454950724438346974478108552543266505447897066401937787353128775658636671050545886283957254005994053382597994030205843794129478035548670169143120716183546839636620832206046130178840470912382088260040689396602749290725149279274843937698894336557963092719826592355136708260359035508112483559947061941595119585651015773217474124961000295759969227493901289642998828296716899454652536049446496906969419458622138816723825829435301821287090805463158656305975450920276963777923103017461201799125060572926044231032704796880489934603822037169658039101891501582036909397436250861123799518082274765578732293891061740875175280725140845569538266163576071003206509391159851507535336785090848512024036090894931823792446772092082764821222186551527833576365116844484278886443262212"}
```

## Vrf

```go
// pietrzak
vrfReq {
	Name    string `json:"name"`
	Message string `json:"message"`
	Vrf0    string `json:"vrf0"`
	Proof   string `json:"proof"`
}
vrfRes {
	Vrf0  string `json:"vrf0"`
	Proof string `json:"proof"`
	Res   bool   `json:"res"`
}
```

> 调用函数
>
> ```go
> // Case1	
> //生成公私钥对(priv,pub)
> priv, _ := vrf.Newprivatekey()
> pub, _ := vrf.GeneratePublickey(priv)
> 
> //利用消息message和私钥priv获得vrf结果vrf0以及证明proof
> message := []byte("heloo")
> vrf0, proof, _ := vrf.Prove(priv, message)
> 
> // Case2
> //利用公钥pub，消息message，证明proof 验证vrf结果vrf0的正确性
> res, _ := vrf.Verify(pub, message, vrf0, proof)
> fmt.Println(res)
> ```

### Case1

输入：`message`

```shell
curl -X GET http://localhost:8000/crypto/vrf -H 'Content-Type:application/json' -d '{"name":"1","message":"hello","vrf0":"","proof":""}'
```

输出：`vrf0，proof`

```shell
{"vrf0":"NCccjLl/BSpVGH6BCsPc6YPzUMOSSV+ZKisxD0m0a8GYXXxJSB5fJYkvIUxNrt43cQyuRYVA/VVdbgVPR/1Y3Q==","proof":"rnHIduCp8Em9YsVxyHE7M3mwfM2+4JLipQlmGmw7zuIgkXlNbIBz9H3THYFxreqEvUaeFVYabSHHuHh6uGxBnVU/z4sKqVbUFTl2V14Jwws=","res":false}
```

### Case2

输入：`message，vrf0，proof`

```shell
curl -X GET http://localhost:8000/crypto/vrf -H 'Content-Type:application/json' -d '{"name":"2","message":"hello","vrf0":"NCccjLl/BSpVGH6BCsPc6YPzUMOSSV+ZKisxD0m0a8GYXXxJSB5fJYkvIUxNrt43cQyuRYVA/VVdbgVPR/1Y3Q==","proof":"rnHIduCp8Em9YsVxyHE7M3mwfM2+4JLipQlmGmw7zuIgkXlNbIBz9H3THYFxreqEvUaeFVYabSHHuHh6uGxBnVU/z4sKqVbUFTl2V14Jwws="}'
```

输出：`res`

```shell
{"vrf0":"","proof":"","res":false}
```

> 因为**`pub`类型是私有结构体且没有提供获取数值方法，所以无法序列化为`json`进行传递**，实现时可以考虑存入数据库中，按照用户编号来一对一存储。
>
> 因此，每次输入后`pub`都会重新生成一个和之前不同的`pub`，所以此时验证结果为`false`。

## Vdf

```go
// pietrzak
pietrzakReq {
	Name       string `json:"name"`
	Challenge  int64  `json:"challenge"`
	Iterations int    `json:"iterations"`
	Proof      string `json:"proof"`
}
pietrzakRes {
	Out  string `json:"out"`
	Out2 bool   `json:"out2"`
}
```

#### Case1

输入：`name（调用哪个方法，1代表生成，2代表验证），challenge，iterations`

```shell
curl -X GET http://localhost:8000/crypto/vdfpietrzak -H 'Content-Type:application/json' -d '{"name":"1","challenge":170,"iterations":100,"proof":""}'
```

输出：`proof`

```shell
{"out":"005271e8f9ab2eb8a2906e851dfcb5542e4173f016b85e29d481a108dc82ed3b3f97937b7aa824801138d1771dea8dae2f6397e76a80613afda30f2c30a34b040baaafe76d5707d68689193e5d211833b372a6a4591abb88e2e7f2f5a5ec818b5707b86b8b2c495ca1581c179168509e3593f9a16879620a4dc4e907df452e8dd0ffc4f199825f54ec70472cc061f22eb54c48d6aa5af3ea375a392ac77294e2d955dde1d102ae2ace494293492d31cff21944a8bcb4608993065c9a00292e8d3f4604e7465b4eeefb494f5bea102db343bb61c5a15c7bdf288206885c130fa1f2d86bf5e4634fdc4216bc16ef7dac970b0ee46d69416f9a9acee651d158ac64915b\n","out2":false}
```

问题修复，原因：没有给二进制文件可执行权限

<img src="./api%E6%96%87%E6%A1%A3.assets/image-20230409104223340.png" alt="image-20230409104223340" style="zoom:80%;" />

#### Case2

输入：`name，challenge，iterations，proof`

```shell
curl -X GET http://localhost:8000/crypto/vdfpietrzak -H 'Content-Type:application/json' -d '{"name":"2","challenge":170,"iterations":100,"proof":"005271e8f9ab2eb8a2906e851dfcb5542e4173f016b85e29d481a108dc82ed3b3f97937b7aa824801138d1771dea8dae2f6397e76a80613afda30f2c30a34b040baaafe76d5707d68689193e5d211833b372a6a4591abb88e2e7f2f5a5ec818b5707b86b8b2c495ca1581c179168509e3593f9a16879620a4dc4e907df452e8dd0ffc4f199825f54ec70472cc061f22eb54c48d6aa5af3ea375a392ac77294e2d955dde1d102ae2ace494293492d31cff21944a8bcb4608993065c9a00292e8d3f4604e7465b4eeefb494f5bea102db343bb61c5a15c7bdf288206885c130fa1f2d86bf5e4634fdc4216bc16ef7dac970b0ee46d69416f9a9acee651d158ac64915b\n"}'
```

输出：`out2`

```shell
{"out":"","out2":true}
```

### Wecolowski_rust

```go
// wrust
wrustReq {
	Name       string `json:"name"`
	Challenge  int64  `json:"challenge"`
	Iterations int    `json:"iterations"`
	Proof      string `json:"proof"`
}
wrustRes {
	Out  string `json:"out"`
	Out2 bool   `json:"out2"`
}
```

> 调用过程和`Pietrzak`类似，这里只展示输入输出数据，该程序也存在一些问题

#### Case1

输入

```shell
curl -X GET http://localhost:8000/crypto/vdfwrust -H 'Content-Type:application/json' -d '{"name":"1","challenge":170,"iterations":100,"proof":""}'
```

输出

```shell
{"out":"005271e8f9ab2eb8a2906e851dfcb5542e4173f016b85e29d481a108dc82ed3b3f97937b7aa824801138d1771dea8dae2f6397e76a80613afda30f2c30a34b040baaafe76d5707d68689193e5d211833b372a6a4591abb88e2e7f2f5a5ec818b5707b86b8b2c495ca1581c179168509e3593f9a16879620a4dc4e907df452e8dd0ffc4f199825f54ec70472cc061f22eb54c48d6aa5af3ea375a392ac77294e2d955dde1d102ae2ace494293492d31cff21944a8bcb4608993065c9a00292e8d3f4604e7465b4eeefb494f5bea102db343bb61c5a15c7bdf288206885c130fa1f2d86bf5e4634fdc4216bc16ef7dac970b0ee46d69416f9a9acee651d158ac64915b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001\n","out2":false}
```

#### Case2

输入

```shell
curl -X GET http://localhost:8000/crypto/vdfwrust -H 'Content-Type:application/json' -d '{"name":"2","challenge":170,"iterations":100,"proof":"005271e8f9ab2eb8a2906e851dfcb5542e4173f016b85e29d481a108dc82ed3b3f97937b7aa824801138d1771dea8dae2f6397e76a80613afda30f2c30a34b040baaafe76d5707d68689193e5d211833b372a6a4591abb88e2e7f2f5a5ec818b5707b86b8b2c495ca1581c179168509e3593f9a16879620a4dc4e907df452e8dd0ffc4f199825f54ec70472cc061f22eb54c48d6aa5af3ea375a392ac77294e2d955dde1d102ae2ace494293492d31cff21944a8bcb4608993065c9a00292e8d3f4604e7465b4eeefb494f5bea102db343bb61c5a15c7bdf288206885c130fa1f2d86bf5e4634fdc4216bc16ef7dac970b0ee46d69416f9a9acee651d158ac64915b000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001\n"}'
```

输出

```shell
{"out":"","out2":true}
```

## 

## zk-SNARKs问题说明

整体实现比较琐碎比较琐碎

### Merkle

因为`pk`，`vk`类型没有找到合适的导出方法，无法`json`序列化，现在还不知道怎么做

> `GPT4`：请注意，这个过程可能非常复杂，因为`groth16.ProvingKey`具有许多嵌套的数据结构。您需要确保为每个字段创建合适的JSON兼容类型，并正确地处理所有的转换。

### Cubic

只是实现了一个简单多项式`x**3 + x + 5 == y`的验证，现实实现中需要这个多项式足够复杂，而且系数可变，所以需要更高复杂度的多项式输入，但是这个多项式是如何生成的，怎么生成我也正在研究，作为毕设的一部分。作为替代，我换成了`Merkle`树证明，但是这里面的`Merkle`树的`pk`和`vk`不能传递，所以就停这里了。


> 要在 Linux 系统中将项目上传至 GitHub 仓库，请按照以下步骤操作：

安装 Git（如果尚未安装）：
打开终端并输入以下命令：
sudo apt-get update
sudo apt-get install git

在本地创建一个新的 Git 仓库，或者导航到现有项目的目录：
mkdir your_project_name
cd your_project_name

初始化新的 Git 仓库：
git init

将项目文件添加到新的 Git 仓库：
git add .
这将添加当前目录中的所有文件。如果只想添加特定文件，可以使用 git add file_name。

提交已添加的文件：
git commit -m "Initial commit"

在 GitHub 网站上创建一个新的仓库。记下新仓库的 URL，它应该类似于：
https://github.com/your_username/your_repository_name.git

在本地仓库中添加远程仓库：
git remote add origin https://github.com/your_username/your_repository_name.git

将本地仓库推送到 GitHub 远程仓库：
git push -u origin master
如果您的远程仓库已经有内容（例如，GitHub 自动生成的 README.md、.gitignore 或 LICENSE 文件），则需要先执行 git pull origin master 命令，将远程仓库内容合并到本地仓库。

您可以选择使用访问令牌。以下是创建并使用访问令牌的步骤：

登录到您的 GitHub 帐户。
点击右上角的头像，选择 "Settings"（设置）。
在左侧菜单中，选择 "Developer settings"（开发者设置）。
点击 "Personal access tokens"（个人访问令牌）。
点击 "Generate new token"（生成新令牌）。
为令牌设置一个描述性名称，选择适当的权限，然后点击 "Generate token"（生成令牌）。
在生成的页面上，复制生成的访问令牌。确保在离开页面之前将其复制，因为您无法再次查看令牌。

现在您可以使用访问令牌而不是密码进行身份验证。在 Git 中，您可以使用以下命令更改远程 URL（确保替换 your_username 和 your_token）：
git remote set-url origin https://your_username:your_token@github.com/19231224lhr/EVM-Docker.git
接下来，您应该能够正常推送代码到 GitHub 仓库了。

如果您更喜欢使用 SSH 密钥进行身份验证，可以参考这篇官方文档：Connecting to GitHub with SSH。
完成以上步骤后，您的项目应已成功上传至 GitHub 仓库。在将来的项目更改中，您只需要执行 git add、git commit 和 git push 命令即可将更改推送到远程仓库。
