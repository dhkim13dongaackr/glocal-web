# D-insight 웹진 상세페이지 — 변경사항

> **원본**: [dev-five-git magazine/8/1](https://dev-five-git.github.io/donga-dongseo-glocal/magazine/8/1)
> **시안**: [webzine-content](https://dhkim13dongaackr.github.io/glocal-web/webzine-content/)
> **날짜**: 2026-02-23

원본 상세페이지의 콘텐츠 구조와 디자인 시스템을 기반으로, 스크롤 인터랙션과 타이포그래피를 시안 수준으로 변형한 버전입니다.

---

## 주요 변경점

| 영역 | 변경 내용 | 구현 방식 |
|------|----------|----------|
| **헤더** | 투명→솔리드 전환 (스크롤 연동) | `rgba(255,255,255,0) → 0.97` + `backdrop-filter: blur(20px)` |
| | 깜빡임 방지 히스테리시스 | solid 진입 `0.5` / 해제 `0.2` (JS) |
| | 타이틀 헤더 흡수 | 스크롤 시 본문 타이틀 축소 → 헤더 compact 타이틀 전환. 고정 텍스트 + 클래스 토글 |
| | 읽기 진행률 바 | 헤더 하단 틸색(`#2ca39c`) 프로그레스 바, 0→100% |
| | Vol. 캡슐 표시 | 우측 상단 `Vol. 03 / 2026. 02` |
| **히어로** | 타이틀-이미지 분리 구조 | 타이틀이 독립 블록으로 배치 |
| | Noto Serif KR 세리프 타이틀 | `clamp(32px, 5vw, 56px)`, weight 900 |
| | 카테고리 캡슐형 태그 | 틸색 배경 `bg-wz-primary-dark` |
| | 설명문 배경 밴드 | `bg-wz-primary/[0.05]` 연한 틸색 |
| | 히어로 이미지 스크롤 모프 | `maxWidth: min(800,viewW-80) → 100%`, `scale: 1.15→1.0`, `aspectRatio` 확장 (JS lerp) |
| | 타이틀 스크롤 축소 → 헤더 흡수 | `scale: 1→0.5`, `opacity: 1→0` |
| **본문** | 680px 중앙 정렬 | `max-w-[680px] mx-auto` |
| | 섹션 accent bar 애니메이션 | 틸색 `w-12~16 h-3px` + `scaleX(0→1)` |
| | 패럴렉스 인라인 이미지 | `data-parallax`, `translateY(-12% → 2%)` JS 보간 |
| | 본문 타이포 | `text-justify`, `leading-[2]`, `font-medium` 16~18px |
| **추천** | More Insights 리스트형 레이아웃 | 좌 128×90 썸네일 + 우 텍스트, hover 배경 + 썸네일 scale |
| | staggered reveal | `transition-delay` 0.08s 간격 순차 등장 |
| **신규** | Credit 섹션 추가 | "글 박태원 / 사진 출처 동아대학교, 동서대학교", Noto Serif KR 타이틀 |

---

## 기술 비교

| 항목 | 원본 | 시안 |
|------|------|------|
| 프레임워크 | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |
| 타이틀 폰트 | SUIT Variable | Noto Serif KR 900 |
| 스크롤 인터랙션 | — | rAF 루프 (타이틀 모프 + 이미지 확장 + 패럴렉스) |
| 헤더 전환 | 고정 | 투명→솔리드 (히스테리시스) + 타이틀 흡수 |
| 읽기 진행률 | — | 틸색 프로그레스 바 |
| 이미지 효과 | — | 패럴렉스 (`translateY` 보간) |
| 추천 콘텐츠 | 2×2 카드 그리드 | 리스트형 (썸네일 + 텍스트) |
| 이미지 소스 | 자체 CDN (webp) | Unsplash (데모용) |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- DEEP DIVE / SPOTLIGHT / DISCOVER 카테고리 체계
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 반응형 레이아웃 (768px 브레이크포인트)
- IntersectionObserver 기반 reveal 애니메이션

---

## 참고

- `prefers-reduced-motion`: `.reveal`, `.accent-bar`, `.insight-card`, `.hero-main-title` CSS 전환만 즉시화. JS rAF 루프(스크롤 모프/패럴렉스)는 미대응
- CSS 주석에 "JS clone" 표현이 있었으나 실제로는 고정 텍스트 + 클래스 토글로 구현 (주석 수정 완료)
