# D-insight 웹진 상세페이지 — 변경사항

> **기준**: [dev-five-git magazine/8/1](https://dev-five-git.github.io/donga-dongseo-glocal/magazine/8/1) (Next.js + Tailwind 빌드)
> **변경**: [webzine-content](https://dhkim13dongaackr.github.io/glocal-web/webzine-content/) (Tailwind CDN 정적 HTML)
> **날짜**: 2026-02-23

---

## 1. 헤더 — 투명→솔리드 전환 + 타이틀 흡수

**기존**: D-insight 로고 + 카테고리 네비게이션. 고정 흰색 배경 헤더. 하단 뒤로가기 화살표.

**변경**:

- **투명→솔리드 전환** — 초기 투명, 스크롤 시 `rgba(255,255,255,0.97)` + `backdrop-filter: blur(20px)`
- **히스테리시스** — 깜빡임 방지: solid 진입 임계값 `0.5` / 해제 임계값 `0.2` 분리 (JS 구현)
- **타이틀 헤더 흡수** — 스크롤 시 본문 타이틀이 축소되며 헤더 중앙에 compact 타이틀 표시
  - 구현: 고정 텍스트 요소를 미리 배치, 클래스 토글로 opacity 전환 (JS clone 아님)
  - NEWSROOM 링크 fade-out ↔ compact 타이틀 fade-in
- **읽기 진행률 바** — 헤더 하단 틸색(`#2ca39c`) 프로그레스 바, 스크롤 위치 기반 0→100%
- **Vol. 캡슐** — 우측 상단 `Vol. 03 / 2026. 02` 캡슐 버튼

---

## 2. 히어로 섹션 — 타이틀 분리 + 스크롤 모프

**기존**: 고정 높이 배너(400~650px). 배경 이미지 위 카테고리 + 타이틀 오버레이. 스크롤 인터랙션 없음.

**변경**:

- **타이틀-이미지 분리 구조** — 타이틀이 이미지 위가 아닌 독립 블록으로 배치
- **Noto Serif KR 세리프 타이틀** — `clamp(32px, 5vw, 56px)`, font-weight 900
- **카테고리 태그** — 틸색 배경 캡슐형(`bg-wz-primary-dark`) (기존: 텍스트만)
- **설명문 배경 밴드** — `bg-wz-primary/[0.05]` 연한 틸색 배경
- **히어로 이미지 스크롤 모프** — JS lerp 보간으로 스크롤에 따라:
  - `maxWidth: min(800, viewW-80)px → viewport 100%` (전체화면 확장)
  - `scale: 1.15 → 1.0` (줌아웃)
  - `aspectRatio: 800/360 → viewW/420` (가로 비율 확장)
- **타이틀 스크롤 축소** — `scale: 1→0.5`, `opacity: 1→0` 점진 전환 후 헤더에 흡수

---

## 3. 본문 레이아웃 — 680px 중앙정렬 + 패럴렉스

**기존**: 전용 타이포 클래스(`typo-webzineDetailBody` 등). 본문 폭 제한 없음(컨테이너 기준). 인포그래픽 이미지 + 캡션.

**변경**:

- **680px 중앙 정렬** — `max-w-[680px] mx-auto` 가독성 최적화
- **섹션 구분 체계**:
  - 틸색 accent bar (`w-12~16`, `h-3px`) + `scaleX(0→1)` 애니메이션
  - 소제목: `text-wz-primary`, `tracking +1.28px` (넓은 자간)
  - 중제목: SUIT 고딕 `font-extrabold 24px`
- **패럴렉스 인라인 이미지** — `data-parallax` 속성, 스크롤에 따라 `translateY(-12% → 2%)`
- **본문 스타일** — `text-justify`, `leading-[2]`(줄간격 2배), `font-medium` 16~18px
- **스페이서** — 단락 간 36px 간격(`div.spacer`)

---

## 4. More Insights 섹션 — 그리드→리스트 전환

**기존**: 4개 아티클 카드, 2×2 그리드. 썸네일 + 카테고리 + 제목.

**변경**:

- **리스트형 레이아웃** — 수직 flex 쌓기, 각 항목: 좌 128×90 썸네일 + 우 텍스트
- **hover 인터랙션** — 배경 `rgba(44,163,156,0.05)`, 썸네일 `scale(1.06)`
- **staggered reveal** — `transition-delay` 0.08s 간격 순차 등장
- **Noto Serif KR 섹션 타이틀** — "More Insights" (기존 SUIT)

---

## 5. Credit 섹션 (신규)

기존 magazine에 없던 기고자 크레딧 추가:
- "글 박태원 / 사진 출처 동아대학교, 동서대학교"
- Noto Serif KR 32px 섹션 타이틀

---

## 6. 스크롤 애니메이션 체계

**기존**: 기본 IntersectionObserver reveal.

**변경**:

- **rAF 메인 루프** — 타이틀 모프, 이미지 모프, 패럴렉스, 프로그레스 바를 단일 `requestAnimationFrame` 루프에서 처리
- **IntersectionObserver** — `.reveal`, `.accent-bar`, `.insight-card` 클래스별 등장 애니메이션
- **easeOutCubic 이징** — `1 - (1-t)³` 수식 기반 부드러운 전환
- **prefers-reduced-motion** — 일부 클래스(`.reveal`, `.accent-bar`, `.insight-card`, `.hero-main-title`)의 CSS 전환만 즉시화. 헤더 배경/프로그레스 바 transition 및 **JS rAF 루프는 미대응** (스크롤 모프/패럴렉스는 reduced-motion에서도 작동)

---

## 7. 기술 비교

| 항목 | magazine/8/1 (기존) | webzine-content (변경) |
|------|---------------------|----------------------|
| 프레임워크 | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |
| 타이틀 폰트 | SUIT Variable | Noto Serif KR 900 |
| 본문 폰트 | SUIT (전용 클래스) | SUIT Variable 500 |
| 스크롤 인터랙션 | 없음 | rAF 루프 (타이틀 모프 + 이미지 확장 + 패럴렉스) |
| 헤더 전환 | 고정 흰색 | 투명→솔리드 (히스테리시스 0.5/0.2) |
| 타이틀 흡수 | 없음 | 스크롤 시 본문 타이틀 → 헤더 compact 타이틀 전환 |
| 읽기 진행률 | 없음 | 틸색 프로그레스 바 |
| 이미지 효과 | 없음 | 패럴렉스 (`translateY` 보간) |
| 추천 콘텐츠 | 2×2 카드 그리드 | 리스트형 (썸네일 + 텍스트) |
| 접근성 | — | 일부 클래스 CSS 전환만 reduced-motion 대응 (JS 루프·헤더 transition 미대응) |
| 이미지 소스 | 자체 CDN (webp) | Unsplash (데모용) |

---

## 8. 유지된 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- DEEP DIVE / SPOTLIGHT / DISCOVER 카테고리 체계
- 반응형 레이아웃 (768px 브레이크포인트)
