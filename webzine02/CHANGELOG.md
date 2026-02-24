# D-insight 웹진 v2 변경사항

> 기준: [dev-five-git 매거진](https://dev-five-git.github.io/donga-dongseo-glocal/magazine)
> 변경: [webzine02](https://dhkim13dongaackr.github.io/glocal-web/webzine02/)
> 날짜: 2026-02-23

---

## 1. 커버/히어로 섹션 전면 개편

### 기존 (magazine)
- 단일 배경 이미지 + 텍스트 오버레이
- "Bridging Local to Global" 영문 타이틀
- 고정 높이 배너, 스크롤 인터랙션 없음

### 변경 (webzine02)
- **대각선 분할 커버**: 동아대 승학캠퍼스 + 동서대 주례캠퍼스 실사진 2장을 `clip-path: polygon()`으로 사선 분할 배치
- **SVG 대각선 구분선**: 흰색 0.5px 사선 라인으로 두 캠퍼스 구분
- **캠퍼스 라벨**: 각 사진 위에 `backdrop-filter: blur(6px)` 처리된 캡슐형 라벨 ("동아대 승학캠퍼스", "동서대 주례캠퍼스")
- **커버 텍스트 밴드**: 사진 하단에 반투명 흰색(`rgba(255,255,255,0.88)`) 밴드 + 우측 정렬 텍스트
  - COVER STORY 라벨
  - "4개에서 6개로, 연합전공이 더 넓어진다" 타이틀 (SUIT 고딕)
  - "휴먼메타케어와 B-Heritage가 합류한 동아·동서의 새 학기" 서브타이틀
- **스크롤 연동 이미지 포지셔닝**: 스크롤 시 동아대/동서대 사진이 각 대학 간판이 보이는 방향으로 자동 이동 (`object-position` 애니메이션)

---

## 2. 데스크톱 스크롤 인터랙션 (신규)

### 기존 (magazine)
- 일반 세로 스크롤 레이아웃
- 히어로와 콘텐츠 그리드가 순차 배치

### 변경 (webzine02)
- **Sticky Split 레이아웃**: 좌측 커버(sticky) + 우측 아티클 그리드
- **히어로 너비 애니메이션**: 스크롤에 따라 커버가 `100% → 48%`로 축소, 우측 콘텐츠가 fade-in
- **Compact 헤더 모드**: 스크롤 95% 이상 시 헤더 축소 (대형 D-insight 로고 → 소형 로고)
- **우측 영역 독립 스크롤**: compact 모드에서 우측 콘텐츠 영역이 내부 스크롤로 전환
  - 페이지 어디서든 wheel 이벤트로 우측 영역 스크롤 가능
  - 상단/하단 도달 시 자연스러운 페이지 스크롤 복귀
- **커버 텍스트 투명도 전환**: 스크롤에 따라 `rgba(255,255,255,0.88) → 1.0` 점진 변화
- **`ResizeObserver`**: 콘텐츠 높이 변화에 따라 split-wrap `min-height` 자동 조정

---

## 3. 헤더 개선

### 기존 (magazine)
- D-insight 로고 + 네비게이션
- 고정 헤더

### 변경 (webzine02)
- **조직 브랜딩 추가**: "동아·동서 글로컬 연합대학 웹진" 서브타이틀
- **홈페이지 링크**: 좌측 상단 "연합대학 홈" 링크 (홈 아이콘 + 텍스트)
- **Vol. 드롭다운**: 클릭 시 발행호 선택 드롭다운 (애니메이션 포함)
- **Compact 모드**: 스크롤 시 대형 헤더 → 네비게이션 바만 남는 축소 전환
  - CSS `max-height` + `opacity` 트랜지션으로 부드러운 접힘
  - `--header-h` CSS 변수로 sticky 요소 오프셋 자동 조정

---

## 4. 아티클 그리드

### 기존 (magazine)
- 카드형 그리드, 기본 IntersectionObserver reveal

### 변경 (webzine02)
- **2단 Masonry 스타일**: 좌/우 컬럼 분리, 카드 높이 비율 다양화 (289, 377, 414 등)
- **데스크톱 Split 모드**: 히어로 옆에 sticky 배치, pre-reveal 처리 (개별 애니메이션 생략)
- **모바일**: 기존 방식 유지 (IntersectionObserver 기반 staggered reveal)
- **탭 필터링 애니메이션**: 카테고리 전환 시 `opacity + translateY + scale` 트랜지션, staggered delay

---

## 5. 기술적 차이

| 항목 | magazine (기존) | webzine02 (변경) |
|------|----------------|-----------------|
| 스크롤 인터랙션 | 없음 | requestAnimationFrame 기반 scroll controller |
| 커버 이미지 | 단일 배경 | 2장 대각선 분할 (clip-path) |
| 헤더 모드 | 고정 | 3단계 (full → scrolled → compact) |
| 콘텐츠 스크롤 | 페이지 스크롤 | 내부 독립 스크롤 (compact 모드) |
| 이미지 최적화 | webp CDN | 로컬 JPG (campus-seunghak, campus-jurye) |
| CSS 변수 | `--banner-base-h` | `--header-h`, `--ease-out-expo`, `--ease-out-quint` |
| 폰트 | 기본 | SUIT Variable + Rakkas + Noto Serif + ELAND Nice |
| 드래그 스크롤 | 없음 | 캐러셀에 drag-to-scroll + momentum |

---

## 6. 유지된 기능
- Glocal Now 섹션 (수평 스크롤 프레스 카드)
- Explore Other Issues 섹션 (과거 발행호)
- 푸터 (동아대/동서대 캠퍼스 정보)
- 모바일 반응형 레이아웃
- IntersectionObserver 기반 reveal 애니메이션
