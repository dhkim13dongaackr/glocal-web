# D-insight 웹진 v2 — 변경사항

> **기준**: [dev-five-git 매거진](https://dev-five-git.github.io/donga-dongseo-glocal/magazine) (Next.js + Tailwind 빌드)
> **변경**: [webzine02](https://dhkim13dongaackr.github.io/glocal-web/webzine02/) (Tailwind CDN 정적 HTML)
> **날짜**: 2026-02-23

---

## 1. 커버/히어로 — 대각선 분할 커버로 전면 개편

**기존**: 단일 배경 이미지(`magazine-hero.webp`) + "Bridging Local to Global" 영문 타이틀. 고정 높이 배너, 스크롤 인터랙션 없음.

**변경**:

- **대각선 분할 커버** — 동아대 승학캠퍼스 + 동서대 주례캠퍼스 실사진 2장을 `clip-path: polygon(0 0, 58% 0, 38% 100%, 0 100%)`으로 사선 분할 배치
- **SVG 대각선 구분선** — 흰색 0.5px 사선 라인(`<line>`)으로 두 캠퍼스 시각 구분
- **캠퍼스 라벨** — 각 사진 위에 `backdrop-filter: blur(6px)` 캡슐형 라벨
- **커버 텍스트 밴드** — 사진 하단 반투명 흰색(`rgba(255,255,255,0.88)`) 밴드, 우측 정렬
  - COVER STORY 라벨 + 호별 타이틀 + 서브타이틀 (매 호마다 교체)
- **스크롤 연동 이미지 포지셔닝** — 스크롤 시 `object-position` JS 보간으로 각 대학 간판 방향 자동 이동

---

## 2. 데스크톱 스크롤 인터랙션 (신규)

**기존**: 일반 세로 스크롤. 히어로와 콘텐츠 그리드가 순차 배치.

**변경**:

- **Sticky Split 레이아웃** — 좌측 커버(`position: sticky`) + 우측 아티클 그리드
- **히어로 너비 애니메이션** — 스크롤에 따라 커버 `width: 100% → 48%` JS 보간, 우측 콘텐츠 fade-in
- **Compact 헤더 트리거** — 스크롤 95% 이상 도달 시 대형 헤더 → 소형 로고만 남는 축소 전환
- **우측 영역 독립 스크롤** — compact 모드 진입 시 우측 콘텐츠가 내부 스크롤(`overflow-y: auto`)로 전환. wheel 이벤트를 우측 영역으로 전달, 상·하단 도달 시 페이지 스크롤 복귀
- **커버 텍스트 투명도** — 스크롤에 따라 밴드 배경 `rgba(255,255,255, 0.88 → 1.0)` 점진 변화
- **ResizeObserver** — 콘텐츠 높이 변화 감지 → split-wrap `min-height` 자동 조정

---

## 3. 헤더 개선

**기존**: D-insight 로고 + 네비게이션. 고정 배경 흰색 헤더.

**변경**:

- **조직 브랜딩** — "동아·동서 글로컬 연합대학 웹진" 서브타이틀 + Rakkas 서체 D-insight 로고
- **홈페이지 링크** — 좌측 상단 "연합대학 홈" (홈 아이콘 + 텍스트)
- **NEWSROOM 링크** — 우측 NEWSROOM 텍스트 + 틸색 펄스 도트 애니메이션
- **Vol. 드롭다운** — 클릭 시 발행호 목록 표시 (translateY + scale 애니메이션, 현재 Vol.01 단일 항목)
- **Compact 모드** — 대형 헤더 → 네비게이션 바만 남는 축소 전환
  - CSS `max-height` + `opacity` 트랜지션
  - `--header-h` CSS 변수로 sticky 오프셋 자동 조정
- **헤더 상태**: 기본(흰색) → scrolled(반투명 + box-shadow) → compact(축소)
  - ※ CSS에 `hero-mode`(투명 헤더) 스타일이 정의되어 있으나, JS 토글 미연결로 현재 미작동

---

## 4. 아티클 그리드

**기존**: 카드형 그리드, 기본 IntersectionObserver reveal.

**변경**:

- **2컬럼 Flex 레이아웃** — 좌(`#colL`) / 우(`#colR`) 컬럼 분리, 카드별 aspect ratio 다양화 (289, 377, 414 등)로 Masonry 유사 효과 구현. 실제 CSS Masonry는 아님
- **데스크톱 Split 모드** — 히어로 옆 sticky 배치 시 pre-reveal 처리 (개별 등장 애니메이션 생략)
- **모바일** — 기존 방식 유지 (IntersectionObserver staggered reveal)
- **탭 필터링** — 카테고리 전환 시 `opacity + translateY + scale` 트랜지션 + staggered delay

---

## 5. 기술 비교

| 항목 | magazine (기존) | webzine02 (변경) |
|------|----------------|-----------------|
| 프레임워크 | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |
| 스크롤 인터랙션 | 없음 | rAF 기반 scroll controller |
| 커버 이미지 | 단일 배경 webp | 2장 대각선 분할 (`clip-path`) |
| 헤더 상태 | 고정 (1상태) | 기본 → scrolled → compact (3상태) |
| 콘텐츠 스크롤 | 페이지 스크롤 | compact 모드 시 우측 내부 독립 스크롤 |
| 커버 이미지 소스 | CDN webp (`magazine-hero.webp`) | 로컬 JPG (`campus-seunghak.jpg`, `campus-jurye.jpg`) |
| 아티클 이미지 소스 | 자체 CDN webp | 원본 CDN webp 그대로 사용 (`dev-five-git` 도메인) |
| CSS 변수 | `--banner-base-h` | `--header-h`, `--ease-out-expo`, `--ease-out-quint` |
| 폰트 | SUIT | SUIT Variable, Rakkas, Noto Serif, ELAND Nice, Abril Fatface, Elice DigitalBaeum |
| 드래그 스크롤 | 없음 | 캐러셀에 drag-to-scroll + momentum |

---

## 6. 유지된 기능

- Glocal Now 섹션 (수평 스크롤 프레스 카드)
- Explore Other Issues 섹션 (과거 발행호 캐러셀)
- 푸터 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 모바일 반응형 레이아웃 (768px 브레이크포인트)
- IntersectionObserver 기반 reveal 애니메이션
- 카테고리 체계 (ALL / DEEP DIVE / SPOTLIGHT / DISCOVER)
