# D-insight 웹진 상세페이지 변경사항

> 기준: [dev-five-git magazine/8/1](https://dev-five-git.github.io/donga-dongseo-glocal/magazine/8/1)
> 변경: [webzine-content](https://dhkim13dongaackr.github.io/glocal-web/webzine-content/)
> 날짜: 2026-02-23

---

## 1. 헤더

### 기존 (magazine/8/1)
- D-insight 로고 + 카테고리 네비게이션 (ALL / DEEP DIVE / SPOTLIGHT / DISCOVER)
- 고정 헤더, 배경 흰색
- 하단에 뒤로가기 화살표

### 변경 (webzine-content)
- **투명 → 솔리드 전환 헤더**: 초기 투명, 스크롤 시 `rgba(255,255,255,0.97)` + `backdrop-filter: blur(20px)` 적용
- **히스테리시스 적용**: 깜빡임 방지를 위해 solid 전환 임계값(0.5)과 해제 임계값(0.2) 분리
- **타이틀 헤더 흡수**: 스크롤 시 본문 타이틀이 축소되며 헤더 중앙에 compact 타이틀로 전환
  - NEWSROOM 링크는 fade-out, compact 타이틀이 fade-in
- **읽기 진행률 표시**: 헤더 하단에 틸색(#2ca39c) 프로그레스 바, 스크롤 위치에 따라 0→100% 채워짐
- **Vol. 캡슐 표시**: 우측 상단에 `Vol. 03 / 2026. 02` 캡슐 버튼

---

## 2. 히어로 섹션

### 기존 (magazine/8/1)
- 고정 높이 배너 (min 400px ~ max 650px)
- 배경 이미지 위에 카테고리 + 타이틀 + 설명 오버레이
- 스크롤 인터랙션 없음

### 변경 (webzine-content)
- **타이틀-이미지 분리 구조**: 타이틀이 이미지 위가 아닌 독립 블록으로 배치
- **Noto Serif KR 세리프 타이틀**: `clamp(32px, 5vw, 56px)`, font-weight 900
- **카테고리 태그**: 틸색 배경 캡슐형 (`bg-wz-primary-dark`)으로 변경 (기존: 텍스트만)
- **설명문 배경 밴드**: `bg-wz-primary/[0.05]` 연한 틸색 배경 밴드
- **히어로 이미지 스크롤 모프**: 스크롤에 따라:
  - `maxWidth: 800px → viewport 100%` (전체화면 확장)
  - `scale: 1.15 → 1.0` (줌아웃 효과)
  - `aspectRatio: 800/360 → viewW/420` (가로 비율 확장)
- **타이틀 스크롤 축소**: `scale: 1 → 0.5`, `opacity: 1 → 0` 점진 전환 후 헤더에 흡수

---

## 3. 본문 레이아웃

### 기존 (magazine/8/1)
- 전용 타이포 클래스 (`typo-webzineDetailBody`, `typo-webzineDetailSubTitle`)
- 본문 폭 제한 없음 (컨테이너 기준)
- 인포그래픽 이미지 + 캡션

### 변경 (webzine-content)
- **680px 중앙 정렬 본문**: `max-w-[680px] mx-auto` 가독성 최적화 폭
- **섹션 구분 체계**:
  - 틸색 accent bar (w-12~16, h-3px) + scaleX(0→1) 애니메이션
  - 소제목: `text-wz-primary`, tracking +1.28px (넓은 자간)
  - 중제목: Noto Serif KR 아닌 SUIT 고딕, font-extrabold 24px
- **패럴렉스 인라인 이미지**: `data-parallax` 속성, 스크롤에 따라 `translateY(-12% → 2%)` 이동
- **본문 스타일**: `text-justify`, `leading-[2]` (줄간격 2배), `font-medium` 16~18px
- **스페이서**: 단락 간 36px 간격 (`div.spacer`)

---

## 4. More Insights 섹션

### 기존 (magazine/8/1)
- 4개 아티클 카드, 그리드 레이아웃 (2×2)
- 썸네일 + 카테고리 + 제목

### 변경 (webzine-content)
- **리스트형 레이아웃**: 가로 배치 (좌: 128×90 썸네일, 우: 텍스트)
- **hover 인터랙션**: 배경색 `rgba(44,163,156,0.05)`, 썸네일 `scale(1.06)`
- **staggered reveal**: 0.08s 간격으로 순차 등장 애니메이션
- **Noto Serif KR 섹션 타이틀**: "More Insights" (기존 SUIT)

---

## 5. Credit 섹션 (신규)

- 기존 magazine에 없던 기고자 크레딧 표시
- "글 박태원 / 사진 출처 동아대학교, 동서대학교"
- Noto Serif KR 32px 섹션 타이틀

---

## 6. 스크롤 애니메이션 체계

### 기존 (magazine/8/1)
- 기본 IntersectionObserver reveal

### 변경 (webzine-content)
- **requestAnimationFrame 기반 메인 루프**: 타이틀 모프, 이미지 모프, 패럴렉스, 프로그레스 바를 단일 rAF 루프에서 처리
- **IntersectionObserver**: `.reveal`, `.accent-bar`, `.insight-card` 클래스별 등장 애니메이션
- **easeOutCubic 이징**: `1 - (1-t)^3` 수식 기반 부드러운 전환
- **reduced-motion 대응**: `prefers-reduced-motion: reduce` 미디어 쿼리로 애니메이션 비활성화

---

## 7. 기술적 비교

| 항목 | magazine/8/1 (기존) | webzine-content (변경) |
|------|---------------------|----------------------|
| 프레임워크 | Next.js + Tailwind (빌드) | Tailwind CDN (정적 HTML) |
| 타이틀 폰트 | SUIT Variable | Noto Serif KR 900 |
| 본문 폰트 | SUIT (전용 클래스) | SUIT Variable 500 |
| 스크롤 인터랙션 | 없음 | rAF 루프 (타이틀 모프 + 이미지 확장) |
| 헤더 전환 | 고정 | 투명 → 솔리드 (히스테리시스) |
| 읽기 진행률 | 없음 | 프로그레스 바 |
| 이미지 효과 | 없음 | 패럴렉스 스크롤 |
| 추천 콘텐츠 | 2×2 카드 그리드 | 리스트형 (썸네일 + 텍스트) |
| 접근성 | - | prefers-reduced-motion 대응 |
| 이미지 소스 | 자체 CDN (webp) | Unsplash (데모) |

---

## 8. 유지된 요소
- 전체 색상 시스템 (틸 #2ca39c 계열)
- 푸터 구조 (동아대/동서대 캠퍼스 정보)
- DEEP DIVE / SPOTLIGHT / DISCOVER 카테고리 체계
- 반응형 레이아웃 (768px 브레이크포인트)
