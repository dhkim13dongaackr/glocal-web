# D-insight 웹진 v2 — 변경사항

> **원본**: [dev-five-git 매거진](https://dev-five-git.github.io/donga-dongseo-glocal/magazine)
> **시안**: [webzine02](https://dhkim13dongaackr.github.io/glocal-web/webzine02/)
> **날짜**: 2026-02-23

원본 매거진의 구조와 디자인 시스템을 기반으로, 커버/스크롤 인터랙션 등을 시안 수준으로 변형한 버전입니다.

---

## 주요 변경점

| 영역 | 변경 내용 | 구현 방식 |
|------|----------|----------|
| **커버** | 두 캠퍼스 사진을 대각선 분할로 배치 | `clip-path: polygon()` + SVG 사선 구분선 |
| | 캠퍼스 라벨 (블러 캡슐형) | `backdrop-filter: blur(6px)` |
| | 커버 텍스트 밴드 (반투명 흰색, 우측 정렬) | `rgba(255,255,255,0.88)`, COVER STORY 라벨 + 호별 타이틀 (매 호 교체) |
| | 스크롤 시 이미지 포지셔닝 연동 | `object-position` JS 보간 |
| **스크롤** | 데스크톱 Sticky Split 레이아웃 | 좌측 커버 sticky + 우측 아티클 그리드 |
| | 히어로 너비 축소 애니메이션 | `width: 100% → 48%` rAF 보간 |
| | Compact 모드 시 우측 독립 스크롤 | `overflow-y: auto` + wheel 이벤트 전달 |
| | 커버 텍스트 투명도 점진 변화 | `rgba` alpha `0.88 → 1.0` |
| | 콘텐츠 높이 자동 추적 | `ResizeObserver` → `min-height` 조정 |
| **헤더** | 조직 브랜딩 서브타이틀 추가 | "동아·동서 글로컬 연합대학 웹진" |
| | Rakkas 서체 D-insight 로고 | 대형(52px) + compact용 소형(22px) |
| | 연합대학 홈 링크 | 좌측 상단 홈 아이콘 + 텍스트 |
| | NEWSROOM 링크 + 펄스 도트 | 틸색 도트 `dotPulse` 애니메이션 |
| | Vol. 드롭다운 | translateY + scale 애니메이션 (현재 Vol.01 단일 항목) |
| | Compact 모드 (스크롤 95% 시 축소) | `max-height` + `opacity` 트랜지션, `--header-h` CSS 변수 |
| **아티클** | 2컬럼 Flex 레이아웃 | 좌(`#colL`) / 우(`#colR`), aspect ratio 다양화로 Masonry 유사 효과 |
| | 탭 필터링 애니메이션 | `opacity + translateY + scale` + staggered delay |
| **캐러셀** | drag-to-scroll + momentum | Glocal Now / Explore Other Issues 섹션 |

---

## 기술 비교

| 항목 | 원본 | 시안 |
|------|------|------|
| 프레임워크 | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |
| 스크롤 인터랙션 | — | rAF 기반 scroll controller |
| 커버 | 단일 배경 | 2장 대각선 분할 (`clip-path`) |
| 헤더 | 고정 | 기본 → scrolled → compact (3상태) |
| 커버 이미지 | CDN webp | 로컬 JPG (호별 교체) |
| 아티클 이미지 | 자체 CDN webp | 원본 CDN webp 그대로 사용 |
| 폰트 | SUIT | SUIT Variable, Rakkas, Noto Serif, ELAND Nice 등 |
| 드래그 스크롤 | — | 캐러셀에 drag-to-scroll + momentum |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- Glocal Now 섹션 (수평 스크롤 프레스 카드)
- Explore Other Issues 섹션 (과거 발행호 캐러셀)
- 카테고리 체계 (ALL / DEEP DIVE / SPOTLIGHT / DISCOVER)
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 모바일 반응형 레이아웃 (768px 브레이크포인트)
- IntersectionObserver 기반 reveal 애니메이션

---

## 참고

- CSS에 `hero-mode`(투명 헤더) 스타일이 정의되어 있으나 JS 토글 미연결로 현재 미작동
- 폰트 `Abril Fatface`, `Elice DigitalBaeum`도 로드되지만 주요 UI에서는 미사용
