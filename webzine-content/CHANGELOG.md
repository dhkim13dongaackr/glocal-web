# D-insight 웹진 상세페이지 — devfive 제시안 vs 변경요청안

> **devfive 제시안**: [magazine/8/1](https://dev-five-git.github.io/donga-dongseo-glocal/magazine/8/1)
> **변경요청안**: [webzine-content](https://dhkim13dongaackr.github.io/glocal-web/webzine-content/)
> **날짜**: 2026-02-23

---

## 비교표

| 영역 | devfive 제시안 (magazine/8/1) | 변경요청안 (webzine-content) |
|------|------------------------------|---------------------------|
| **헤더 배경** | 고정 흰색 | 초기 투명 → 스크롤 시 솔리드 전환 (히스테리시스 적용) |
| **타이틀 흡수** | — | 스크롤 시 본문 타이틀 축소 → 헤더 중앙에 compact 타이틀 표시 |
| **NEWSROOM 링크** | 항상 표시 | 스크롤 시 fade-out, compact 타이틀과 교체 |
| **읽기 진행률** | — | 헤더 하단 틸색 프로그레스 바 (스크롤 0→100%) |
| **Vol. 표시** | VolDropdown 컴포넌트 | 우측 상단 캡슐 (`Vol. 03 / 2026. 02`) |
| **히어로 구조** | 배경 이미지 위 타이틀 오버레이 | 타이틀과 이미지를 독립 블록으로 분리 |
| **타이틀 폰트** | SUIT Variable | Noto Serif KR 900, `clamp(32px, 5vw, 56px)` |
| **카테고리 태그** | 텍스트 | 틸색 배경 캡슐형 |
| **히어로 스크롤** | 없음 | 이미지 전체화면 확장 + 줌아웃 + 비율 변화 (JS lerp 보간) |
| **타이틀 스크롤** | 없음 | `scale 1→0.5`, `opacity 1→0` 점진 축소 후 헤더 흡수 |
| **본문 폭** | 컨테이너 기준 | `max-w-[680px]` 중앙 정렬 |
| **섹션 구분** | 전용 타이포 클래스 | 틸색 accent bar (`scaleX 0→1` 애니메이션) + 소제목/중제목 분리 |
| **인라인 이미지** | 정적 배치 | 패럴렉스 효과 (`translateY -12%→2%` 스크롤 연동) |
| **본문 스타일** | 전용 클래스 | `text-justify`, `leading-[2]`, `font-medium` 16~18px |
| **추천 콘텐츠** | 2×2 카드 그리드 | 리스트형 (좌 썸네일 + 우 텍스트), hover 효과, staggered reveal |
| **Credit 섹션** | — | 기고자 크레딧 추가 ("글 박태원 / 사진 출처") |
| **스크롤 엔진** | IntersectionObserver만 | IntersectionObserver + rAF 루프 (타이틀·이미지·패럴렉스·진행률을 매 프레임 통합 처리) |
| **이미지 소스** | 자체 CDN (webp) | Unsplash (데모용) |
| **프레임워크** | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |

---

## 원본에서 유지한 요소

- 전체 색상 시스템 (틸 `#2ca39c` 계열)
- DEEP DIVE / SPOTLIGHT / DISCOVER 카테고리 체계
- 푸터 구조 (동아대/동서대 캠퍼스 정보 + 로고 링크)
- 반응형 레이아웃
- IntersectionObserver 기반 reveal 애니메이션
