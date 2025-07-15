package crawler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const pokeURLTest = "https://sv-news.pokemon.co.jp/ko/json/list.json"

func TestPokeCrawl(t *testing.T) {
	// @@@ html 내용이 테스트 시점에 따라 바뀌므로 다른 테스트 방식 고려 필요 @@@
	// Test: Good url with 5 days duration
	result, err := PokeCrawl(pokeURLTest, 5)
	require.NoError(t, err)
	assert.Equal(t, 1, len(result))
	assert.Equal(t,
		`<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="robots" content="nofollow">
  <meta name="focus-ring-visibility" content="hidden">
  <meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0, user-scalable=no">
  <link rel="stylesheet" href="../../resources/css/news2.css">
  <link rel="stylesheet" href="../../resources/css/news-web2.css">
  <link rel="shortcut icon" href="../../resources/image/Favicon_SVNews.png">
  <link rel="apple-touch-icon" sizes="180x180" href="../../resources/image/Favicon_SVNews.png">
  <title>검은 결정 테라 레이드배틀에 최강의 짜랑고우거가 출현 중!</title>
  
  <script src='../../resources/js/jquery-3.2.1.min.js'></script>
  <script src="../../resources/js/news-web3.js"></script>
  <!-- Google tag (gtag.js) -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-SNKH7X9LBK"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'G-SNKH7X9LBK');
  </script>
</head>
<body>
<article class="main-contents">
  <div class="content-wrapper">
    <div class="notice">
      <div class="type">테라 레이드배틀</div>
      <div class="content-box">
        <div class="content-preview">
          <div class="title">검은 결정 테라 레이드배틀에 최강의 짜랑고우거가 출현 중!</div>
          <div class="date-text" data-date="1752192000"></div>
          <div class="image-wrapper">
            <img src="banner/8_1746175664_00411747199672.jpg">
          </div>
          <div class="text-wrapper">
            <div class="main-text">
              <b>최강의 짜랑고우거가 출현 중!</b><br />7월 14일(월) 8:59까지, ★7의 검은 결정 테라 레이드배틀에 <b>짜랑고우거</b>가 출현 중!<br />이번에 출현하는 짜랑고우거는 격투 테라스탈타입의 매우 강력한 포켓몬!<br />본 이벤트에 출현하는 최강의 짜랑고우거는 <b>「최강의증표」</b>를 가지고 있습니다.<br />트레이너들과 힘을 합쳐서 승리해 보시기 바랍니다!<br /><h4>※본 이벤트에서 출현하는 최강의 짜랑고우거는 하나의 저장 데이터에서 1마리만 잡을 수 있습니다. 잡은 뒤에도 기간 안에 검은 결정 테라 레이드배틀에 참가하면, 보상을 받을 수 있습니다.<br />※본 이벤트에서 출현하는 짜랑고우거는 경우에 따라, 추후 이벤트에서 다시 출현하거나 다른 방식으로 만날 수 있습니다.</h4><br /><h1>기간</h1>2025년 7월 11일(금) 9:00~7월 14일(월) 8:59<br />(다음 일정)2025년 7월 18일(금) 9:00~7월 21일(월) 8:59<br /><h4>두 개최 기간에 출현하는 「최강의 짜랑고우거」는 능력과 배운 기술이 같습니다. 앞선 기간에서 「최강의 짜랑고우거」를 잡은 경우, 「최강의 짜랑고우거」를 추가로 잡을 수 없습니다.</h4><br /><h1>출현 포켓몬</h1>★7: 최강의 짜랑고우거<br /><br /><h1>이벤트 테라 레이드배틀이란?</h1>이벤트 테라 레이드배틀은 시기별로 출현하는 포켓몬이 바뀝니다.<br />좀처럼 보기 어려운 테라스탈타입의 포켓몬을 만날 수도 있다고 합니다.<br />포켓포털 등의 정보를 확인해서 원하는 포켓몬을 찾았다면 도전해 보시기 바랍니다.<br /> <br /><h1>이벤트 테라 레이드배틀에 도전 시 주의사항</h1>・이벤트 테라 레이드배틀을 플레이하려면 아래 방법으로 최신 정보를 받아야 합니다.<br /><h3><b>「포켓포털」→「이상한 소포」→「포켓포털 뉴스를 받는다」</b><br />※Nintendo Switch Online(유료) 가입은 필요하지 않습니다.</h3>・검은 결정 테라 레이드배틀에 도전하려면 엔딩 후에 진행할 수 있는 이벤트를 클리어해야 합니다. 단, 멀티 플레이 중인 트레이너의 테라 레이드배틀에 참가하거나, 암호를 통해 지인의 테라 레이드배틀에 참가하는 경우, 이벤트를 클리어하지 않은 트레이너도 검은 결정 테라 레이드배틀에 참가할 수 있습니다.<br />・인터넷 통신을 이용해 다른 트레이너와 테라 레이드배틀을 플레이하려면 Nintendo Switch Online(유료)에 가입해야 합니다.<br />・이벤트 테라 레이드배틀은 팔데아지방에서만 발생합니다.
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="footer">
      ©2022 Pokémon.
      <br>
      ©1995-2022 Nintendo/Creatures Inc. /GAME FREAK inc.
    </div>
  </div>
</article>
</body>
</html>`,
		result[0].Body,
	)

	// Test: Good url with 15 days duration
	result, err = PokeCrawl(pokeURLTest, 15)
	require.NoError(t, err)
	// assert.Equal(t, 5, len(result))
	assert.Equal(t,
		`<!DOCTYPE html>
<html lang="ko">
<head>
  <meta charset="UTF-8">
  <meta name="robots" content="nofollow">
  <meta name="focus-ring-visibility" content="hidden">
  <meta name="viewport" content="width=device-width, height=device-height, initial-scale=1.0, user-scalable=no">
  <link rel="stylesheet" href="../../resources/css/news2.css">
  <link rel="stylesheet" href="../../resources/css/news-web2.css">
  <link rel="shortcut icon" href="../../resources/image/Favicon_SVNews.png">
  <link rel="apple-touch-icon" sizes="180x180" href="../../resources/image/Favicon_SVNews.png">
  <title>전기타입 포켓몬이 대량발생 중!</title>
  
  <script src='../../resources/js/jquery-3.2.1.min.js'></script>
  <script src="../../resources/js/news-web3.js"></script>
  <!-- Google tag (gtag.js) -->
  <script async src="https://www.googletagmanager.com/gtag/js?id=G-SNKH7X9LBK"></script>
  <script>
    window.dataLayer = window.dataLayer || [];
    function gtag(){dataLayer.push(arguments);}
    gtag('js', new Date());

    gtag('config', 'G-SNKH7X9LBK');
  </script>
</head>
<body>
<article class="main-contents">
  <div class="content-wrapper">
    <div class="notice">
      <div class="type">대량발생</div>
      <div class="content-box">
        <div class="content-preview">
          <div class="title">전기타입 포켓몬이 대량발생 중!</div>
          <div class="date-text" data-date="1751587200"></div>
          <div class="image-wrapper">
            <img src="banner/8_1746175997_73061749629447.jpg">
          </div>
          <div class="text-wrapper">
            <div class="main-text">
              <b>전기타입 포켓몬이 대량발생 중!</b><br />7월 7일(월) 8:59까지, 이벤트 대량발생으로 <b>데덴네, 빠모, 파치리스, 모르페코, 플러시, 마이농</b>이 출현 중!<br />본 특별한 대량발생에서는 색이 다른 포켓몬들을 평소보다 만나기 쉬워진다고 합니다.<br />팔데아지방과 북신의 고장, 블루베리 아카데미를 모험하며, 잔뜩 동료로 만들어 보시기 바랍니다!<br /><br /><h1>기간</h1>2025년 7월 4일(금) 9:00~7월 7일(월) 8:59<br /><br /><h1>출현 장소</h1>▼데덴네, 빠모 대량발생<br />・팔데아지방<br /><br />▼파치리스, 모르페코 대량발생<br />・북신의 고장<br /><br />▼플러시, 마이농 대량발생<br />・블루베리 아카데미<br /><br /><h1>이벤트 대량발생이란?</h1>이벤트 대량발생이 개최 중인 기간에는 대량발생에서 특정 포켓몬이 출현하기 쉬워집니다.<br />원하는 포켓몬의 대량발생을 놓치지 않도록 포켓포털 등의 정보를 확인해 보시기 바랍니다.<br /> <br /><h1>이벤트 대량발생 진행 시 주의사항</h1>・이벤트 대량발생을 플레이하려면, 아래 방법으로 최신 정보를 받아야 합니다.<br /><h3><b>「포켓포털」→「이상한 소포」→「포켓포털 뉴스를 받는다」</b><br />※Nintendo Switch Online(유료) 가입은 필요하지 않습니다.</h3>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="footer">
      ©2022 Pokémon.
      <br>
      ©1995-2022 Nintendo/Creatures Inc. /GAME FREAK inc.
    </div>
  </div>
</article>
</body>
</html>`,
		result[2].Body,
	)

	// Test: temp
	result, err = PokeCrawl(pokeURLTest, 100)
	for _, r := range result {
		fmt.Println(r.Title)
		fmt.Println(r.StAt)
	}
	require.Error(t, err)

}
