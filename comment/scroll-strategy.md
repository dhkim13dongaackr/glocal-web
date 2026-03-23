# 글로컬 연합대학 스크롤 인터랙션 전략

> 작성일: 2026-03-23
> 대상: https://dev-five-git.github.io/donga-dongseo-glocal/

---

## 핵심 원칙

**"이 사이트는 보는 사이트다. 발신자의 메시지를 섹션 단위로 전달하는 프레젠테이션이다."**

- 스크롤스냅은 "제한"이 아니라 **"연출"** — 백화점이 동선을 설계하듯, 방문자의 시선을 의도적으로 가이드
- 모든 곳에 일괄 적용하지 않는다 — **프레젠테이션 구간은 스냅, 탐색 구간은 자유 스크롤**
- `mandatory` 대신 **`proximity`** — 자유 스크롤을 유지하면서 정지 시에만 가까운 섹션으로 정렬

---

## 인터랙션 패턴 정의

본 전략에서 사용하는 5가지 패턴:

| 패턴 | 설명 | 데모 |
|------|------|------|
| **섹션 스냅 (시작점 기준)** | 세로 스크롤 시 각 섹션 시작점에 자동 정렬 | [01-y-proximity](01-y-proximity.html) |
| **가로 전환 (횡스크롤)** | 세로 스크롤 시 콘텐츠가 가로로 슬라이드 전환 | [02-horizontal-scroll](02-horizontal-scroll.html) |
| **스티키 사이드바** | 왼쪽 네비 고정, 오른쪽 콘텐츠 자유 스크롤, 현재 위치 하이라이트 | [03-sticky-sidebar](03-sticky-sidebar.html) |
| **스크롤 연동 확대** | 스크롤 진행에 따라 영상이 작은 창에서 풀스크린으로 확대 | [04-video-expand](04-video-expand.html) |
| **2단 스냅** | 같은 콘텐츠를 2단계로 분리 (타이틀 → 본체 풀스크린) | [05-two-stage-snap](05-two-stage-snap.html) |

---

## 페이지별 상세 전략

### 1. 메인 (`/`)

**전략: 섹션 스냅 (시작점 기준) 전체 적용**

가장 순수한 프레젠테이션 페이지. 모든 섹션이 시각 중심, viewport에 맞는 구조.

| # | 섹션 | 스냅 | 추가 인터랙션 |
|---|------|:---:|------|
| 1 | 메인 배너 (비디오, "경계를 넘어, 내일을 열다") | **섹션 스냅** | — |
| 2 | 혁신전략 (4카드: Global Alliance, Field Campus 등) | **섹션 스냅** | — |
| 3 | Field 캠퍼스 (6탭 카드) | **섹션 스냅** | — |
| 4 | 핵심 성과 (수강생 2,363 / 매출 918억 / 유학생 3,898) | **섹션 스냅** | CountUp 애니메이션 |
| 5 | 글로컬 뉴스 (4카드) | **섹션 스냅** | — |
| 6 | FAQ (4항목 아코디언) | **섹션 스냅** | — |
| 7 | 푸터 | 제외 | — |

```css
html { scroll-snap-type: y proximity; scroll-behavior: smooth; }
section { scroll-snap-align: start; min-height: 100vh; }
footer { /* scroll-snap-align 없음 */ }
```

---

### 2. Innovation (`/innovation`)

**전략: 상단 섹션 스냅 (시작점 기준) + 8대 추진과제 가로 전환**

| # | 섹션 | 패턴 | 설명 |
|---|------|------|------|
| 1 | 히어로 배너 (SCROLL DOWN) | **섹션 스냅** | 첫 인상 |
| 2 | 글로컬대학30 프로젝트 (비전/목표/전략 I·II) | **섹션 스냅** | 정보 요약형, viewport에 맞음 |
| 3 | OUR VISION (3개 혁신방향 × 5 세부과제) | **섹션 스냅** | 카드형 구조 |
| 4 | **8대 추진과제** | **가로 전환** | 핵심 연출 구간 |
| 5 | 푸터 | 제외 | — |

#### 8대 추진과제 상세

- 세로 스크롤 → 가로 슬라이드 전환 (GSAP `pin` + `scrub`)
- 각 과제: 이미지(좌 40%) + 텍스트(우 60%) 배치
- "1/8" ~ "8/8" 인디케이터 하단 고정
- 카드 내 콘텐츠가 viewport보다 길면: 세로 스크롤 소진 후 다음 카드로 전환
- 상단 프로그레스 바로 전체 진행률 표시

```
[히어로] → snap
[글로컬30] → snap
[OUR VISION] → snap
[8대 추진과제 ──────────────────] → 가로 전환 (8장)
  ┌─────┐ ┌─────┐ ┌─────┐
  │ 01  │→│ 02  │→│ ... │→ ...08
  └─────┘ └─────┘ └─────┘
[푸터]
```

---

### 3. IACF 통합산단 (`/iacf`)

**전략: 가장 복합적 — 섹션 스냅 + 가로 전환 + 스티키 사이드바 + 영상 확대**

| # | 섹션 | 패턴 | 설명 |
|---|------|------|------|
| 1 | TECHBRIDGE 메인 (비교 다이어그램) | **섹션 스냅** | 첫 인상, 마우스 아이콘 |
| 2 | 3 BUSINESS MODELS (3카드: 프랜차이징/마켓부스팅/스노우볼링) | **섹션 스냅** | viewport에 맞는 카드형 |
| 3 | **3 CORE STRATEGIES** | **가로 전환** | 전략 01→02→03 |
| 4 | **ACHIEVEMENT** (조직/대외협력/기술사업화) | **스티키 사이드바** | 좌측 네비 고정 + 우측 자유 스크롤 |
| 5 | **VISION VIDEO** ("통합산단 2.0") | **스크롤 연동 확대** | 60%→풀스크린 |
| 6 | PRESS RELEASE (6개 뉴스 카드) | **섹션 스냅** | 뉴스 그리드 |
| 7 | 푸터 | 제외 | — |

#### 3대 추진전략 상세

- Innovation의 8대 과제와 동일 패턴, 3장
- 전략 01 차별성(Global Expansion) → 02 전문성(AI & Platform) → 03 수익성(Network Accelerator)
- 각 전략: 목표 + KEY INITIATIVES 3-4개 + 글로벌 파트너/브랜치

#### ACHIEVEMENT 상세

```
┌──────────────────────────────────┐
│                                  │
│  ACHIEVEMENT        │ 조직       │
│  추진성과            │ ├ 조직도   │
│                     │ ├ 체크리스트│
│  ● 조직             │            │
│  ○ 대외협력·네트워크  │ 대외협력   │  ← 우측만 스크롤
│  ○ 기술사업화·IP     │ ├ 6카드    │
│                     │            │
│  (좌측 sticky)      │ 기술사업화  │
│                     │ ├ 4카드    │
│                     │ ├ 특허차트  │
└──────────────────────────────────┘
```

- 왼쪽: 섹션 제목 + 3개 서브섹션 네비게이션 (sticky, `top: 80px`)
- 오른쪽: 조직→대외협력→기술사업화 순차 스크롤
- IntersectionObserver로 현재 보이는 서브섹션을 왼쪽 네비에 하이라이트

#### VISION VIDEO 상세

```
스크롤 0%:   영상 width: 60%, border-radius: 24px
스크롤 50%:  영상 width: 80%, border-radius: 12px
스크롤 100%: 영상 width: 100vw, height: 100vh, border-radius: 0
             타이틀/헤딩 fade-out
```

- 스크롤 거리: `300vh` (충분한 확대 시간)
- ACHIEVEMENT(정보 밀도 높음) → VIDEO(시각 임팩트)로 **리듬 전환 효과**

```
[TECHBRIDGE] → snap
[3 BM] → snap
[3 CORE STRATEGIES ─────────] → 가로 전환 (3장)
[ACHIEVEMENT ───────────────] → sticky sidebar + 자유 스크롤
[VISION VIDEO ──────────────] → 영상 확대 (300vh)
[PRESS RELEASE] → snap
[푸터]
```

---

### 4. Field Campus (`/field-campus`)

**전략: 인트로 배너만 스냅, 나머지 자유 스크롤**

이 페이지는 "안내/소개" 성격. 서브 네비게이션(소개/교육방식/연합전공/JA교원)이 이미 존재하여 사용자가 직접 점프. 스냅보다 **읽기 흐름**이 중요.

| # | 섹션 | 스냅 | 이유 |
|---|------|:---:|------|
| 1 | 인트로 배너 ("無경계", 배경 이미지) | **섹션 스냅** | 깔끔한 첫 진입 |
| 2 | ABOUT (운영목적 3개 + 유형 4개) | 제외 | 읽기형 콘텐츠 |
| 3 | EDUCATION (4개 교육방식 카드) | 제외 | 자연스러운 흐름 |
| 4 | 6 JOINT MAJOR (6개 전공 카드) | 제외 | 탐색 영역 |
| 5 | JA교원 (설명 + 통계 + 3유형) | 제외 | 정보 밀도 높음 |
| 6 | 푸터 | 제외 | — |

**일관성 문제 없음** — 일관성의 기준은 "모든 페이지에 스냅"이 아니라 "페이지 성격에 맞는 처리". 프레젠테이션 페이지는 스냅, 안내 페이지는 자유.

---

### 5. Field Campus Status (`/field-campus/status`)

**전략: 2단 스냅 + 자유 탐색 + 성과 스냅**

| # | 섹션 | 패턴 | 설명 |
|---|------|------|------|
| 1 | 메인 배너 ("부산 전역, 캠퍼스가 되다" + 대표 카드) | **섹션 스냅** | 첫 인상 |
| 2 | 캠퍼스 네트워크 — 타이틀 | **2단 스냅 1단** | "21개 거점" 메시지 + 필터 UI |
| 3 | 캠퍼스 네트워크 — 지도 | **2단 스냅 2단** | 지도 풀스크린, "부산 전체에 깔려있다" 체감 |
| 4 | 21개 거점 카드 그리드 | 자유 스크롤 | 탐색 영역 |
| 5 | 성과 (21개소 / 29개 / 506명 / 15기관) | **섹션 스냅** | CountUp + 마무리 |
| 6 | 푸터 | 제외 | — |

#### 2단 스냅 상세

```
[1단 스냅]
┌─────────────────────────────┐
│  FIELD NETWORK              │
│  캠퍼스 네트워크              │
│  부산 전역 21개 산업 거점...  │
│  [전체21] [수소3] [반도체4]..│
│           ↓                 │
└─────────────────────────────┘

[2단 스냅]
┌─────────────────────────────┐
│  ┌───────────────────────┐  │
│  │    부산 지도 (21핀)     │  │
│  │    전체 화면           │  │
│  └───────────────────────┘  │
│  [범례: 에너지/휴먼/콘텐츠..]│
└─────────────────────────────┘
```

---

### 6. Major 연합전공 (`/major`)

**전략: 선택적 스냅 — snap→free→snap→free→snap**

| # | 섹션 | 스냅 | 이유 |
|---|------|:---:|------|
| 1 | 인트로 ("Field 연합전공이란?") | **섹션 스냅** | 정의, viewport에 맞음 |
| 2 | 운영현황 (6개 탭, MajorTabs) | 제외 | 탭 전환 시 높이 가변 → 스냅과 충돌 |
| 3 | 신청방법 (3열 + 3단계 절차) | **섹션 스냅** | 안내형, viewport에 맞음 |
| 4 | 이수체계 (로드맵 + 아코디언 FAQ) | 제외 | 아코디언 확장 시 높이 가변 |
| 5 | 혜택 (장학금 표 + 4카드) | **섹션 스냅** | 마무리, viewport에 맞음 |
| 6 | 푸터 | 제외 | — |

```css
/* 스냅 적용 섹션만 scroll-snap-align 부여 */
#intro { scroll-snap-align: start; }
/* #operation — 없음 */
#application { scroll-snap-align: start; }
/* #curriculum — 없음 */
#benefits { scroll-snap-align: start; }
```

자유 스크롤 구간을 지나다가 다음 스냅 포인트에 가까워지면 `proximity`가 자연스럽게 잡아줌.

---

### 7. Major/Curriculum 교과과정표 (`/major/curriculum`)

**전략: 스냅 미적용**

- 탭 기반 전공 전환 (6개 탭) — 이미 탭이 전공 간 분리를 담당
- 각 탭 내부: 3-5개 모듈 × 4학기 교과목 카드 — 순수 참조/카탈로그
- 스냅을 걸 이유가 없음

단, **탭바는 sticky로 유지**하여 현재 어떤 전공을 보고 있는지 컨텍스트 유지.

```css
.tab-bar { position: sticky; top: 0; z-index: 100; }
```

---

## 전체 요약 매트릭스

| 페이지 | 섹션 스냅 (시작점 기준) | 가로 전환 | 스티키 사이드바 | 스크롤 연동 확대 | 2단 스냅 | 자유 스크롤 |
|--------|:---:|:---:|:---:|:---:|:---:|:---:|
| `/` 메인 | 전체 | — | — | — | — | 푸터 |
| `/innovation` | 상단 3섹션 | 8대 추진과제 | — | — | — | 푸터 |
| `/iacf` | TECH, BM, PRESS | 3대 전략 | ACHIEVEMENT | VIDEO | — | — |
| `/field-campus` | 인트로만 | — | — | — | — | 나머지 전체 |
| `/field-campus/status` | 배너, 성과 | — | — | — | 지도 | 21개 카드 |
| `/major` | 인트로, 신청, 혜택 | — | — | — | — | 운영현황, 이수체계 |
| `/major/curriculum` | — | — | — | — | — | 전체 |

---

## 기술 구현 방향

### CSS Scroll Snap

```css
html {
  scroll-snap-type: y proximity;
  scroll-behavior: smooth;
}

/* 스냅 대상 섹션에만 */
.snap-section {
  scroll-snap-align: start;
  min-height: 100vh;
}

/* 접근성: 모션 감소 설정 시 스냅 비활성화 */
@media (prefers-reduced-motion: reduce) {
  html { scroll-snap-type: none; }
}
```

### 가로 전환 (GSAP ScrollTrigger)

```javascript
gsap.to(track, {
  x: () => -(track.scrollWidth - window.innerWidth),
  ease: 'none',
  scrollTrigger: {
    trigger: '.horizontal-section',
    start: 'top top',
    end: () => '+=' + (track.scrollWidth - window.innerWidth),
    pin: true,
    scrub: 1,
  }
});
```

### 영상 확대 (GSAP ScrollTrigger)

```javascript
gsap.to(videoBox, {
  width: '100vw',
  maxWidth: '100vw',
  borderRadius: '0px',
  scrollTrigger: {
    trigger: '.video-section',
    start: 'top top',
    end: 'bottom bottom',
    scrub: 1,
  }
});
```

### 스티키 사이드바

```css
.sidebar {
  position: sticky;
  top: 80px;
  height: fit-content;
}
```

```javascript
// IntersectionObserver로 현재 서브섹션 감지
const observer = new IntersectionObserver(entries => {
  entries.forEach(e => {
    if (e.isIntersecting) {
      // 사이드바 네비 하이라이트 업데이트
    }
  });
}, { rootMargin: '-20% 0px -60% 0px' });
```

---

## 주의사항

1. **모바일**: 터치 스크롤에서 `proximity`가 과민하게 반응할 수 있음. 필요 시 `@media (max-width: 768px)`에서 완화
2. **고정 헤더**: `scroll-margin-top` 또는 `scroll-padding-top`으로 헤더 높이만큼 보정 필요
3. **동적 높이**: 아코디언/탭 콘텐츠가 있는 섹션에는 스냅 미적용 (높이 변화 시 스냅 포인트 어긋남)
4. **앵커 링크**: 서브 네비게이션의 `#anchor` 이동과 스냅이 충돌할 수 있음 — 스냅 적용 페이지에서는 `scrollIntoView`로 통일
5. **접근성**: `prefers-reduced-motion: reduce` 미디어 쿼리로 스냅 비활성화 제공
6. **성능**: 가로 전환 + 영상 확대는 GSAP 사용 — Next.js 환경에서는 Framer Motion `useScroll` + `useTransform`이 더 자연스러울 수 있음
