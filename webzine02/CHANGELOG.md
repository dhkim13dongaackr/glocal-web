# D-insight 웹진 메인 — devfive 제시안 vs 변경요청안

> **devfive 제시안**: [magazine](https://dev-five-git.github.io/donga-dongseo-glocal/magazine)
> **변경요청안**: [webzine02](https://dhkim13dongaackr.github.io/glocal-web/webzine02/)
> **날짜**: 2026-02-23

---

## 비교표

| 영역 | devfive 제시안 (magazine) | 변경요청안 (webzine02) |
|------|--------------------------|----------------------|
| **커버 이미지** | 단일 배경 이미지 (`magazine-hero.webp`) | 두 캠퍼스 실사진 배치 + 캠퍼스 라벨 |
| **커버 텍스트** | "Bridging Local to Global" 영문 타이틀 | 반투명 밴드 위 COVER STORY 라벨 + 호별 한글 타이틀 (매 호 교체) |
| **데스크톱 레이아웃** | 히어로 → 콘텐츠 순차 배치 | Sticky Split (좌측 커버 고정 + 우측 아티클 그리드) |
| **스크롤 인터랙션** | 없음 | 스크롤에 따라 커버 너비 축소(`100%→48%`), 이미지 위치 연동, 텍스트 투명도 변화 등을 JS로 매 프레임 처리 |
| **헤더 브랜딩** | D-insight 로고 + 네비게이션 | "동아·동서 글로컬 연합대학 웹진" 서브타이틀 + Rakkas 서체 로고 |
| **헤더 상태** | 고정 | 기본 → scrolled(반투명) → compact(축소, 소형 로고) |
| **연합대학 홈 링크** | — | 좌측 상단 홈 아이콘 + 텍스트 |
| **NEWSROOM 링크** | 텍스트만 | 텍스트 + 틸색 펄스 도트 애니메이션 |
| **Vol. 드롭다운** | VolDropdown 컴포넌트 | 클릭 시 목록 표시 (현재 Vol.01 단일 항목) |
| **아티클 그리드** | 카드형 그리드 | 2컬럼 Flex, 카드별 높이 다양화 (Masonry 유사) |
| **탭 필터링** | 카테고리 전환 | 전환 시 `opacity + translateY + scale` 애니메이션 추가 |
| **Compact 모드 스크롤** | 페이지 스크롤 | 우측 콘텐츠 영역 내부 독립 스크롤 |
| **캐러셀 조작** | 수평 스크롤 | drag-to-scroll + momentum 추가 |
| **폰트** | SUIT | SUIT Variable, Rakkas, Noto Serif, ELAND Nice 등 추가 |
| **프레임워크** | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- Glocal Now 섹션 (수평 스크롤 프레스 카드)
- Explore Other Issues 섹션 (과거 발행호)
- 카테고리 체계 (ALL / DEEP DIVE / SPOTLIGHT / DISCOVER)
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 모바일 반응형 레이아웃
- IntersectionObserver 기반 reveal 애니메이션
