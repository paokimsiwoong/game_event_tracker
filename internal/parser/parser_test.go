package parser

import (
	"testing"
	"time"

	"github.com/paokimsiwoong/game_event_tracker/internal/crawler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPokeParse(t *testing.T) {
	// Test: Good result
	loc := time.FixedZone("KST", 9*60*60)
	input := []crawler.PokeSVResult{
		{
			Title:   "검은 결정 테라 레이드배틀에 최강의 짜랑고우거가 출현 중!",
			Kind:    "1",
			KindTxt: "테라 레이드배틀",
			Body: `<!DOCTYPE html>
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
			StAt:    1752192000,
			Success: true,
		},
		{
			Title:   "전기타입 포켓몬이 대량발생 중!",
			Kind:    "8",
			KindTxt: "대량발생",
			Body: `<!DOCTYPE html>
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
			StAt:    1751587200,
			Success: true,
		},
		{
			Title:   "「스칼렛 커버」, 「바이올렛 커버」 선물!",
			Kind:    "5",
			KindTxt: "이상한 소포",
			Body: `<!DOCTYPE html>
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
  <title>「스칼렛 커버」, 「바이올렛 커버」 선물!</title>
  
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
      <div class="type">이상한 소포</div>
      <div class="content-box">
        <div class="content-preview">
          <div class="title">「스칼렛 커버」, 「바이올렛 커버」 선물!</div>
          <div class="date-text" data-date="1740668400"></div>
          <div class="image-wrapper">
            <img src="banner/8_1733822913_39171736133593.jpg">
          </div>
          <div class="text-wrapper">
            <div class="main-text">
              2025년 2월 27일(목) 23:00부터, 2월 27일(목) 「Pokémon Day」를 기념해서 <b>「스칼렛 커버」, 「바이올렛 커버」</b>를 선물 중!<br />「포켓몬스터스칼렛」에서는 「스칼렛 커버」를,<br />「포켓몬스터바이올렛」에서는 「바이올렛 커버」를 받을 수 있습니다.<br /><br />메뉴에서 「포켓포털」→「이상한 소포」→「시리얼 코드/암호로 받는다」를 선택한 다음, 아래 암호를 입력하면 받을 수 있습니다.<br /><br /><b>▼스칼렛 커버</b><br />※「포켓몬스터스칼렛」에서만 받을 수 있습니다<br /><h3><b>SB00KC0VER</b></h3>※알파벳 O(오), I(아이), Z(제트)는 사용되지 않습니다.<br /><br /><b>▼바이올렛 커버</b><br />※「포켓몬스터바이올렛」에서만 받을 수 있습니다<br /><h3><b>VB00KC0VER</b></h3>※알파벳 O(오), I(아이), Z(제트)는 사용되지 않습니다.<br /><br /><h1>기간</h1>2025년 2월 27일(목) 23:00~<br /><br /><h1>주의사항</h1>※「스칼렛 커버」, 「바이올렛 커버」를 받으려면 인터넷에 접속할 수 있어야 합니다.<br />※인터넷에 접속하려면, 본체의 유저가 닌텐도 어카운트와 연동되어 있어야 합니다. (Nintendo Switch Online(유료) 가입은 필요하지 않습니다.)<br />※「스칼렛 커버」, 「바이올렛 커버」는 패키지 버전/다운로드 버전 모두 받을 수 있습니다.<br />※암호 입력을 10회 틀리면, 일시적으로 입력을 할 수 없게 됩니다. 6시간 이상 경과한 후에 다시 시도해 주시기 바랍니다.
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
			StAt:    1740668400,
			Success: true,
		},
		{
			Title:   "2025년 7월 시즌(시즌 32) 개최 중!",
			Kind:    "2",
			KindTxt: "랭크배틀",
			Body: `<!DOCTYPE html>
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
  <title>2025년 7월 시즌(시즌 32) 개최 중!</title>
  
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
      <div class="type">랭크배틀</div>
      <div class="content-box">
        <div class="content-preview">
          <div class="title">2025년 7월 시즌(시즌 32) 개최 중!</div>
          <div class="date-text" data-date="1751342400"></div>
          <div class="image-wrapper">
            <img src="banner/8_1748507153_46061748507486.jpg">
          </div>
          <div class="text-wrapper">
            <div class="main-text">
              랭크배틀 <b>2025년 7월 시즌(시즌 32)</b>이 개최 중!<br />게임 내 메뉴에서 <b>「포켓포털」→「배틀스타디움」→「랭크배틀」</b>의 순서로 선택하고 랭크배틀에 도전하자!<br /> <br /><h1>기간</h1>2025년 7월 1일(화)~8월 1일(금) 8:59<br /> <br /><h1>레귤레이션에 대해</h1>2025년 7월 시즌에서는 <b>레귤레이션 I</b>를 사용합니다.<br /><br /><h2>사용할 수 있는 포켓몬</h2>・팔데아도감 No.001〜400<br />・북신도감 No.001〜200<br />・블루베리도감 No.001〜242<br />상기 포켓몬과 아래 포켓몬이 참가할 수 있습니다.<br /><details><summary><b>기타 사용할 수 있는 포켓몬(펼치기)</b></summary><h4>라이츄(알로라의 모습)<br />나옹(알로라의 모습)<br />나옹(가라르의 모습)<br />페르시온(알로라의 모습)<br />가디(히스이의 모습)<br />윈디(히스이의 모습)<br />찌리리공(히스이의 모습)<br />붐볼(히스이의 모습)<br />또도가스(가라르의 모습)<br />프리져<br />프리져(가라르의 모습)<br />썬더<br />썬더(가라르의 모습)<br />파이어<br />파이어(가라르의 모습)<br />뮤츠<br />블레이범(히스이의 모습)<br />포푸니(히스이의 모습)<br />라이코<br />앤테이<br />스이쿤<br />루기아<br />칠색조<br />레지락<br />레지아이스<br />레지스틸<br />라티아스<br />라티오스<br />가이오가<br />그란돈<br />레쿠쟈<br />유크시<br />엠라이트<br />아그놈<br />디아루가<br />디아루가(오리진폼)<br />펄기아<br />펄기아(오리진폼)<br />히드런<br />레지기가스<br />기라티나(어나더폼)<br />기라티나(오리진폼)<br />크레세리아<br />대검귀(히스이의 모습)<br />드레디어(히스이의 모습)<br />조로아(히스이의 모습)<br />조로아크(히스이의 모습)<br />워글(히스이의 모습)<br />코바르온<br />테라키온<br />비리디온<br />토네로스(화신폼)<br />토네로스(영물폼)<br />볼트로스(화신폼)<br />볼트로스(영물폼)<br />레시라무<br />제크로무<br />랜드로스(화신폼)<br />랜드로스(영물폼)<br />큐레무(큐레무의 모습)<br />큐레무(화이트큐레무)<br />큐레무(블랙큐레무)<br />미끄네일(히스이의 모습)<br />미끄래곤(히스이의 모습)<br />크레베이스(히스이의 모습)<br />모크나이퍼(히스이의 모습)<br />코스모그<br />코스모움<br />솔가레오<br />루나아라<br />네크로즈마<br />네크로즈마(황혼의 갈기)<br />네크로즈마(새벽의 날개)<br />나이킹<br />자시안<br />자마젠타<br />무한다이노<br />치고마<br />우라오스(일격의 태세)<br />우라오스(연격의 태세)<br />레지에레키<br />레지드래고<br />블리자포스<br />레이스포스<br />버드렉스<br />버드렉스(백마 탄 모습)<br />버드렉스(흑마 탄 모습)<br />신비록<br />다투곰<br />포푸니크<br />러브로스(화신폼)<br />러브로스(영물폼)</h4></details>참가할 수 있는 포켓몬 중, 아래 <b>특별한 포켓몬은 2마리까지</b> 참가할 수 있습니다.<h4>뮤츠<br />루기아<br />칠색조<br />가이오가<br />그란돈<br />레쿠쟈<br />디아루가<br />디아루가(오리진폼)<br />펄기아<br />펄기아(오리진폼)<br />기라티나(어나더폼)<br />기라티나(오리진폼)<br />레시라무<br />제크로무<br />큐레무(큐레무의 모습)<br />큐레무(화이트큐레무)<br />큐레무(블랙큐레무)<br />코스모그<br />코스모움<br />솔가레오<br />루나아라<br />네크로즈마<br />네크로즈마(황혼의 갈기)<br />네크로즈마(새벽의 날개)<br />자시안<br />자마젠타<br />무한다이노<br />버드렉스<br />버드렉스(백마 탄 모습)<br />버드렉스(흑마 탄 모습)<br />코라이돈<br />미라이돈<br />테라파고스</h4><h4>※「포켓몬스터스칼렛・바이올렛」 및 「제로의 비보」에서 잡은 포켓몬, 알에서 태어난 포켓몬, 공식으로 선물 받은 포켓몬 및 「Pokémon HOME」에서 데려온 포켓몬만 참가할 수 있습니다.<br />※자세한 내용은 아래 <b>「사용할 수 있는 포켓몬」</b>을 확인해 주십시오.</h4><br /><a class="btn-link" href="https://battle-lp1.pokemon-home.com/scvi/regulation/arloqotrdhil6ik5466m/pokemon-list-ko-s">사용할 수 있는 포켓몬</a><br /><br /><h2>포켓몬의 레벨</h2>레벨 1~100의 포켓몬을 등록할 수 있습니다.<br />모든 포켓몬은 대전 시 자동으로 레벨 50이 됩니다.<br /><br /><h2>포켓몬의 지닌 물건</h2>포켓몬에게 「도구」를 지니게 할 수 있습니다.<br />단, 참가하는 포켓몬 중 두 마리 이상의 포켓몬에게 같은 「도구」를 지니게 할 수 없습니다.<br /><br /><h2>대전 시간</h2>종합 시간: 최대 20분<br />플레이어당 제한 시간: 최대 7분<br />1턴당 선택 시간: 45초<br />대전에 내보낼 포켓몬 선택 시간: 90초<br /><br /><h1>보상에 대해</h1>시즌이 종료되어 결과가 발표되면, 성적에 따른 보상을 받을 수 있습니다. 보상을 받으려면 시즌 중에 승패 결과가 보이는 대전에 <b>1회 이상</b> 참가해야 합니다.<br /><br /><h2>마스터볼급</h2>비전 매콤스파이스: 5개<br />문볼: 2개<br />금색병뚜껑: 1개<br />은색병뚜껑: 3개<br />특성패치: 1개<br />리그페이: 300,000LP<br /><br /><h2>하이퍼볼급</h2>비전 매콤스파이스: 2개<br />문볼: 1개<br />금색병뚜껑: 1개<br />은색병뚜껑: 1개<br />특성캡슐: 1개<br />리그페이: 150,000LP<br /><br /><h2>슈퍼볼급</h2>기술머신「테라버스트」: 1개<br />은색병뚜껑: 1개<br />특성캡슐: 1개<br />리그페이: 60,000LP<br /><br /><h2>몬스터볼급</h2>기술머신「테라버스트」: 1개<br />리그페이: 20,000LP<br /><br /><h2>비기너급</h2>리그페이: 10,000LP<br /><br /><h1>랭크배틀이란?</h1>전 세계의 포켓몬 트레이너와 대전해서 실력을 겨룰 수 있는 기능입니다.<br />랭크배틀에는 트레이너의 강함을 나타내는 랭크가 존재하며, 대전 결과에 따라 랭크가 변동됩니다. 대전은 자신의 랭크와 비슷한 트레이너와 진행하게 되어 치열한 배틀을 즐길 수 있습니다.<br /><br /><h1>주의사항 및 이용 약관</h1>・온라인 플레이를 이용하기 위해서는 Nintendo Switch Online(유료)에 가입해야 합니다.<br />・대전의 배틀 데이터와 성적은 Nintendo Co., Ltd.와 The Pokémon Company(및 그 자회사)에 제공됩니다.<br />・이하의 행위를 한 플레이어는 페널티가 부과될 수 있습니다. 또한, 이후 모든 인터넷 통신을 이용한 콘텐츠 및 이벤트에 참가할 수 없게 될 수 있습니다.<br />　1. 게임 소프트웨어의 리포트 등을 부정하게 개조하는 장치를 사용해서, 개조 코드에 의해 조작된 데이터가 포함된 포켓몬 및 개조 코드를 사용해서 만들어진 포켓몬을 사용한 경우<br />　2. 대전 중(대전 상대가 결정된 뒤 승패 결과가 서버에 송신 완료되기까지) 절단 횟수가 현저하게 많은 경우(인터넷 통신은 통신 상태가 안정된 환경에서 접속해 주십시오)<br />　3. 다른 참가자에게 피해를 주거나 불쾌한 이미지를 주는 행위<br />　4. 크래킹 등 대전이나 운영에 지장을 주는 행위<br />　5. 부정한 내용으로 등록하거나 타인으로 위장하여 등록하는 행위<br />　6. 개인이 복수의 닌텐도 어카운트를 이용해 본 랭크배틀에 참가하는 행위<br />　7. 개인이 복수의 기기나 저장 데이터를 이용하거나, 다른 사람과 연대하는 등의 방법을 통해 의도적으로 승패를 조작하는 행위<br />　8. 기타 Nintendo Co., Ltd. 및 The Pokémon Company(및 그 자회사)가 부적절하다고 판단하는 행위
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
			StAt:    1751342400,
			Success: true,
		},
		{
			Title:   "검은 레쿠쟈 강림!",
			Kind:    "1",
			KindTxt: "테라 레이드배틀",
			Body: `<!DOCTYPE html>
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
  <title>검은 레쿠쟈 강림!</title>
  
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
          <div class="title">검은 레쿠쟈 강림!</div>
          <div class="date-text" data-date="1734652800"></div>
          <div class="image-wrapper">
            <img src="banner/8_1728467685_39331730785407.jpg">
          </div>
          <div class="text-wrapper">
            <div class="main-text">
              <b>색이 다른 전설의 포켓몬 「레쿠쟈」가 이벤트 테라 레이드배틀에 등장!</b><br />2025년 1월 6일(월) 8:59까지, 색이 다른 레쿠쟈(드래곤 테라스탈타입)가 ★5 테라 레이드배틀 결정에서 출현 중!<br />본 이벤트에서 출현하는 색이 다른 레쿠쟈는 일반적인 플레이로는 만날 수 없습니다.<br />이번 기회에 도전해서 모험 동료로 만들어 보시기 바랍니다!<br /><h4>※본 이벤트에서 출현하는 색이 다른 레쿠쟈는 하나의 저장 데이터에서 1마리만 잡을 수 있습니다. 잡은 뒤에도 기간 안에 검은 결정 테라 레이드배틀에 참가하면, 보상을 받을 수 있습니다.<br />※본 이벤트에서 출현하는 색이 다른 레쿠쟈는 경우에 따라, 추후 이벤트에서 다시 출현하거나 다른 방식으로 만날 수 있습니다.</h4> <br /><h1>기간</h1>2024년 12월 20일(금) 9:00~2025년 1월 6일(월) 8:59<br /><br /><h1>출현 포켓몬</h1>★5: 색이 다른 레쿠쟈<br /> <br />또한, ★5의 테라 레이드배틀 결정에서 해피너스를 만나기 쉬워진다고 합니다!<br />출현하는 해피너스는 다양한 테라스탈타입을 가지고 있습니다.<br />승리하면 <b>다양한 타입의 테라피스</b>와 <b>경험사탕</b>을 평소보다<b>많이 획득할 수 있습니다!</b><br /><br /><h1>이벤트 테라 레이드배틀이란?</h1>이벤트 테라 레이드배틀은 시기별로 출현하는 포켓몬이 바뀝니다.<br />좀처럼 보기 어려운 테라스탈타입의 포켓몬을 만날 수도 있다고 합니다.<br />포켓포털 등의 정보를 확인해서 원하는 포켓몬을 찾았다면 도전해 보시기 바랍니다.<br /> <br /><h1>이벤트 테라 레이드배틀에 도전 시 주의사항</h1>・이벤트 테라 레이드배틀을 플레이하려면 아래 방법으로 최신 정보를 받아야 합니다.<br /><h3><b>「포켓포털」→「이상한 소포」→「포켓포털 뉴스를 받는다」</b><br />※Nintendo Switch Online(유료) 가입은 필요하지 않습니다.</h3>・인터넷 통신을 이용해 다른 트레이너와 테라 레이드배틀을 플레이하려면 Nintendo Switch Online(유료)에 가입해야 합니다.<br />・★5 테라 레이드배틀에 도전하려면 엔딩까지 진행해야 합니다. 단, 멀티 플레이 중인 트레이너의 테라 레이드배틀에 참가하거나, 암호를 통해 지인의 테라 레이드배틀에 참가하는 경우, 엔딩까지 진행하지 않은 트레이너도 ★5 테라 레이드배틀에 참가할 수 있습니다.<br />・이벤트 테라 레이드배틀은 팔데아지방에서만 발생합니다.
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
			StAt:    1734652800,
			Success: true,
		},
	}

	result, err := PokeParse(input)
	require.NoError(t, err)
	assert.Equal(t, 4, len(result))

	// 첫번쨰 테라레이드
	assert.Equal(t, 1, result[0].Kind)
	assert.Equal(t, time.Date(2025, time.July, 11, 9, 0, 0, 0, loc), result[0].StartsAt[0])
	assert.Equal(t, time.Date(2025, time.July, 18, 9, 0, 0, 0, loc), result[0].StartsAt[1])
	assert.Equal(t, time.Date(2025, time.July, 14, 8, 59, 0, 0, loc), result[0].EndsAt[0])
	assert.Equal(t, time.Date(2025, time.July, 21, 8, 59, 0, 0, loc), result[0].EndsAt[1])

	// 이상한 소포
	assert.Equal(t, 5, result[2].Kind)
	assert.Equal(t, 1, len(result[2].StartsAt))
	assert.Equal(t, 0, len(result[2].EndsAt))
	assert.Equal(t, time.Date(2025, time.February, 27, 23, 0, 0, 0, loc), result[2].StartsAt[0])

	// 연말~연초인 이벤트
	assert.Equal(t, 1, result[3].Kind) // @@@ 랭크배틀은 제외되므로 검은 레쿠쟈 인덱스는 4가 아니라 3
	assert.Equal(t, 1, len(result[3].StartsAt))
	assert.Equal(t, 1, len(result[3].EndsAt))
	assert.Equal(t, time.Date(2024, time.December, 20, 9, 0, 0, 0, loc), result[3].StartsAt[0])
	assert.Equal(t, time.Date(2025, time.January, 6, 8, 59, 0, 0, loc), result[3].EndsAt[0])
}
