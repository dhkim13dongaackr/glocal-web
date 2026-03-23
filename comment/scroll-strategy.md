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

푸터 제외 전체 섹션 스냅. 8대 추진과제는 가로 스크롤 스냅.

### 2. Innovation (`/innovation`)

푸터 제외 섹션 스냅. 8대 추진과제 영역은 섹션 진입 시 타이틀에 먼저 스냅되고, 이후 스크롤하면 카드가 한 장씩 가로 스크롤 스냅으로 전환. 모바일에서의 가로 스냅 처리 방식은 추천 필요.

---

### 3. IACF 통합산단 (`/iacf`)

추진성과(ACHIEVEMENT)를 제외한 전 섹션 섹션 스냅. 현재 모바일에서 스크롤 시안이 없으므로 모바일 대응 필요.

- 히어로 SCROLL DOWN 버튼 여백을 혁신방향(innovation) 페이지와 동일하게 조정 필요
- 초기 3개 원형 모션(BM 진입 시)이 너무 동시에 나옴 — 순차적으로 딜레이를 줘서 순서감을 표현할 것

- 3대 수익모델(3 BM): 모바일에서 가로 스크롤 스냅 (한 카드씩)
- 3대 추진전략: Innovation 8대 과제와 같은 방식 (타이틀 스냅 → 가로 스크롤 스냅)
- 추진성과(ACHIEVEMENT): 자유 스크롤 (콘텐츠 밀도 높음)
- VISION VIDEO: 랜딩페이지와 같이 스크롤 시 전체 화면으로 전환되는 모션 필요
- PRESS RELEASE: 섹션 스냅
---

### 4. Field Campus (`/field-campus`)

인트로 배너, ABOUT, EDUCATION에 섹션 스냅. 6개 연합전공, JA교원은 콘텐츠가 길어 자유 스크롤. 페이지 하단에 구축현황(`/field-campus/status`)으로 유도하는 CTA 필요 — "Field 캠퍼스에 대해 더 알아보기" 등.

### 5. Field Campus Status (`/field-campus/status`)

배너, 성과에 섹션 스냅. 캠퍼스 네트워크는 2단 스냅 — 첫 스냅에 타이틀과 필터, 다음 스냅에 지도가 전체 화면으로 표시. 21개 거점 카드는 자유 스크롤. 대표 Field 캠퍼스 4개 카드는 스크롤에 따라 한 장씩 순차 노출.

---

### 6. Major 연합전공 (`/major`)

실용 정보 페이지. 특별한 모션 불필요. 장학금 숫자에 CountUp 애니메이션 추가.

### 7. Major/Curriculum 교과과정표 (`/major/curriculum`)

각 전공 시작점에 섹션 스냅 적용. 전공별 콘텐츠가 한 화면을 넘지만 시작점에만 스냅이 걸리므로 내부는 자유 스크롤. 탭바는 sticky 유지.

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
