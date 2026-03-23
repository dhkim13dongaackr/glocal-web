# 모바일(375px) 전수검사 결과

> 검사일: 2026-03-23
> 대상: https://dev-five-git.github.io/donga-dongseo-glocal/ 전체 7개 페이지
> 기준: iPhone SE/13 mini (375px), Apple HIG 44px 터치 영역, WCAG 2.1 AA

---

## 결론

**레이아웃 깨짐은 없다.** 가로 오버플로우, 카드 넘침, grid 전환 모두 정상.
문제는 전부 **터치 영역 부족 + 작은 폰트** — 접근성 영역이다.

---

## 공통 문제 (전 페이지)

### 1. 햄버거 버튼 터치 영역 — 높음

모바일에서 가장 중요한 내비게이션 진입점인데 기준의 절반 이하.

| 항목 | 현재값 | 기준 |
|------|--------|------|
| 크기 | 20×20px | 최소 44×44px |
| padding | 0px | — |

```css
/* 수정안 */
.hamburger-btn {
  min-width: 44px;
  min-height: 44px;
  display: flex;
  align-items: center;
  justify-content: center;
}
```

### 2. 헤더 로고 링크 높이 — 중간

| 항목 | 현재값 | 기준 |
|------|--------|------|
| 크기 | 100×34px | 최소 44px 높이 |
| padding | 0px | — |

### 3. 푸터 전화번호 링크 — 중간

| 항목 | 현재값 | 기준 |
|------|--------|------|
| 크기 | 113×18px | 최소 44px 높이 |

```css
/* 수정안 */
a[href^="tel:"] {
  display: inline-block;
  padding: 12px 0;
}
```

### 4. 푸터 대학 로고 링크 — 낮음

| 항목 | 현재값 | 기준 |
|------|--------|------|
| 크기 | ~140×36px | 최소 44px 높이 |

### 5. `typo-tiny` 폰트 10px — 중간

전 페이지에서 `font-size: 10px` 사용. iOS Safari 자동 확대 대상이 되어 레이아웃이 깨질 수 있음.

```css
/* 수정안 */
.typo-tiny {
  font-size: 12px; /* 10px → 12px */
}
```

### 6. html/body overflow-x 미적용 — 낮음

현재 `html.overflowX: visible`, `body.overflowX: visible`. 실제 가로 스크롤은 발생하지 않지만 방어적 적용 권장.

```css
html { overflow-x: hidden; }
```

---

## 페이지별 고유 문제

### 메인 (`/`)

**문제 없음.** 카드 세로 스택 전환 정상, 영상 섹션/성과/FAQ 모두 정상 렌더링.

---

### Innovation (`/innovation`)

| 문제 | 심각도 | 설명 |
|------|:---:|------|
| 페이지 높이 20,980px | 높음 (UX) | 8대 추진과제 세로 나열로 극심한 스크롤 피로. 가로 전환 적용 시 해소 |

레이아웃 깨짐은 없으나, 8개 과제가 동일 구조로 끝없이 반복되어 모바일에서 스크롤 포기율이 높을 것으로 예상.

---

### IACF (`/iacf`)

| 문제 | 심각도 | 설명 |
|------|:---:|------|
| `typo-tiny` 10px × 7군데 | 중간 | "통합산단 총괄 연구 조직", 설립일('24.4, '25.5), 날짜('25.08, '28, '26.02), 구성 수(동아10+동서5) |
| 카로셀 탭 네비 aria-label 없음 | 낮음 | 7개 탭 스크롤 가능하나 접근성 라벨 부재 |

BM 카드(프랜차이징/마켓부스팅/스노우볼링)는 카로셀로 의도된 동작, 깨짐 없음.

---

### Field Campus (`/field-campus`)

| 문제 | 심각도 | 설명 |
|------|:---:|------|
| "교과과정 보기" 링크 24px 높이 × 6개 | 높음 | 6개 연합전공마다 반복. 터치 영역 44px 미달 |
| "수강신청 방법 보기" 링크 42px | 낮음 | 44px에 2px 부족 |
| fixed 헤더 + sticky 탭바 = 화면 19.5% 점유 | 참고 | 헤더 86px + 탭바 72px = 158px / 812px. 콘텐츠 가용 영역 654px |
| `typo-tiny` 10px × 1군데 | 중간 | "적용사례: 전력반도체 Field 교과목..." |

```css
/* "교과과정 보기" 링크 수정안 */
.curriculum-link {
  padding: 10px 0;
  min-height: 44px;
  display: inline-flex;
  align-items: center;
}
```

---

### Field Campus Status (`/field-campus/status`)

| 문제 | 심각도 | 설명 |
|------|:---:|------|
| 6개 서브내비 탭 스크롤 인디케이터 없음 | 중간 | 6개 탭 합산 ~460px > 335px. 스크롤바 숨김으로 뒤쪽 탭(성과/구축현황) 인지 불가 |
| 필터 탭 높이 28px | 중간 | 7개 필터(전체+6분야), `padding: 7px 16px`. 44px 미달 |
| `.dd-item-loc` 10px 폰트 | 중간 | 거점 위치 텍스트 |

```css
/* 서브내비 스크롤 힌트 수정안 */
.sub-nav-wrap::after {
  content: '';
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  width: 40px;
  background: linear-gradient(to right, transparent, white);
  pointer-events: none;
}
```

---

### Major (`/major`)

공통 문제 외 고유 문제 없음. 탭 콘텐츠 높이 가변은 스크롤스냅 전략에서 이미 "자유 스크롤"로 처리.

---

### Major/Curriculum (`/major/curriculum`)

| 문제 | 심각도 | 설명 |
|------|:---:|------|
| 6개 전공 탭 스크롤 인디케이터 없음 | 중간 | 탭 합산 폭 > 375px, 스크롤바 숨김 |
| 과목 카드 가독성 | 낮음 | 과목이 많은 모듈에서 텍스트 밀도 높음 |

---

## 깨짐 없음 확인 항목

| 항목 | 결과 |
|------|------|
| 가로 오버플로우 (scrollWidth > clientWidth) | 없음 — 전 페이지 375px 유지 |
| 카드/이미지 넘침 | 없음 — 카로셀은 의도적 overflow-x: auto |
| 모바일/데스크톱 분기 | 정상 — MobileBanner/DesktopBanner 분리 |
| grid → 1열 스택 전환 | 정상 — 미디어 쿼리 전환 동작 |
| z-index 충돌 | 없음 — 헤더(100/2000), 탭바(90/1500) 분리 |
| position fixed/sticky 겹침 | 없음 — JS 토글로 위치 보정 |

---

## 조치 우선순위

### P0 — 즉시 (접근성 기본, 웹어워드 감점 요소)

1. 햄버거 버튼 `min-width/height: 44px`
2. "교과과정 보기" 링크 `padding: 10px 0` (field-campus 6군데)
3. `typo-tiny` 10px → 12px 상향 (전 페이지)

### P1 — 권장 (완성도)

4. 푸터 TEL 링크 `padding: 12px 0`
5. 서브내비/필터 탭 우측 그라데이션 마스크 (스크롤 힌트)
6. 필터 탭 padding 증가 (높이 44px 확보)
7. `html { overflow-x: hidden }` 방어적 적용

### P2 — 참고

8. 헤더 로고 링크 높이 보완
9. Innovation 8대 과제 모바일 스크롤 피로 → 가로 전환으로 해소
10. Field Campus 고정 UI 점유율 19.5% → 스크롤 시 헤더 숨김으로 완화
