# D-insight 웹진 상세페이지 — devfive 제시안 vs 변경요청안

> **devfive 제시안**: [magazine/8/1](https://dev-five-git.github.io/donga-dongseo-glocal/magazine/8/1)
> **변경요청안**: [webzine-content](https://dhkim13dongaackr.github.io/glocal-web/webzine-content/)
> **날짜**: 2026-02-23

---

## 비교표

| 영역 | devfive 제시안 (magazine/8/1) | 변경요청안 (webzine-content) |
|------|------------------------------|---------------------------|
| **헤더** | 고정 흰색 배경 | 초기 투명 → 스크롤 시 솔리드 전환, 읽기 진행률 바 추가, 스크롤 시 본문 타이틀이 헤더로 흡수 |
| **히어로 구조** | 배경 이미지 위 타이틀 오버레이 | 타이틀과 이미지를 독립 블록으로 분리, 스크롤 시 이미지 전체화면 확장 + 줌아웃 효과 |
| **본문 레이아웃** | 컨테이너 기준 폭 | 680px 중앙 정렬, 섹션별 accent bar 애니메이션, 패럴렉스 인라인 이미지 |
| **추천 콘텐츠** | 2×2 카드 그리드 | 리스트형 (좌 썸네일 + 우 텍스트) |
| **Credit 섹션** | — | 기고자 크레딧 추가 |
| **타이틀 폰트** | SUIT Variable | Noto Serif KR 900 |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- DEEP DIVE / SPOTLIGHT / DISCOVER 카테고리 체계
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 반응형 레이아웃
- IntersectionObserver 기반 reveal 애니메이션
