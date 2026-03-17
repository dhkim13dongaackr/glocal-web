# D-insight 웹진 메인 — devfive 제시안 vs 변경요청안

> **devfive 제시안**: [magazine](https://dev-five-git.github.io/donga-dongseo-glocal/magazine)
> **변경요청안**: [webzine02](https://dhkim13dongaackr.github.io/glocal-web/webzine02/)
> **날짜**: 2026-02-23

---

## 비교표

| 영역 | devfive 제시안 (magazine) | 변경요청안 (webzine02) |
|------|--------------------------|----------------------|
| **데스크톱 레이아웃** | 히어로 → 콘텐츠 순차 배치 (일반 스크롤) | Sticky Split: 좌측 커버 고정 + 우측 아티클 그리드, 스크롤에 따라 커버 너비 축소 등 연동 |
| **헤더** | D-insight 로고 + 네비게이션, 고정 헤더 | "동아·동서 글로컬 연합대학 웹진" 브랜딩 추가, 연합대학 홈 링크 추가, 스크롤 시 3단계 전환 (기본→반투명→compact 축소) |
| **폰트** | SUIT | SUIT Variable, Rakkas, Noto Serif, ELAND Nice 등 추가 |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- Glocal Now 섹션 (수평 스크롤 프레스 카드)
- Explore Other Issues 섹션 (과거 발행호)
- 카테고리 체계 (ALL / DEEP DIVE / SPOTLIGHT / DISCOVER)
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 모바일 반응형 레이아웃
- IntersectionObserver 기반 reveal 애니메이션
