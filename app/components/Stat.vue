<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'

type Variant = 'subtle' | 'solid' | 'outline'
type Size = 'sm' | 'md' | 'lg' | 'xl'

const props = withDefaults(defineProps<{
  label: string
  value: number | string
  /** optional manual percentage diff (e.g. 5 = +5%). If omitted, component will compute trend automatically. */
  diff?: number | null
  /** override text for diff, e.g. "+12% vs. Vormonat" */
  diffLabel?: string
  prefix?: string
  suffix?: string
  variant?: Variant
  size?: Size
  loading?: boolean
  clickable?: boolean
  locale?: string
  /** format big numbers compactly (12,300 → 12.3K) */
  compact?: boolean
  /**
   * Zeitfenster (ms) für Trend-Berechnung. Beispiel: 60000 → Vergleich zum Wert vor 60s.
   * Wenn gesetzt, beobachtet die Komponente den Wert selbst und berechnet diff intern.
   */
  trendOffsetMs?: number | null
  /**
   * Wie oft soll der Trend neu bewertet werden (ms)? Nur relevant, wenn trendOffsetMs gesetzt ist.
   */
  pollMs?: number
  /**
   * Maximale Historie, die im Speicher behalten wird (ms). Älteres wird aus der History gelöscht.
   * Default = 10 × trendOffsetMs (mind. 5 Minuten)
   */
  historyLimitMs?: number
}>(), {
  variant: 'subtle',
  size: 'md',
  loading: false,
  clickable: false,
  compact: true,
  trendOffsetMs: 60_000, // 1 Minute
  pollMs: 1_000,         // jede Sekunde neu bewerten
})

const emit = defineEmits<{ (e:'click'): void }>()

// ---------- Formatting ----------
const isNumber = computed(() => typeof props.value === 'number' && !Number.isNaN(props.value))
const formattedValue = computed(() => {
  if (!isNumber.value) return String(props.value)
  try {
    const n = new Intl.NumberFormat(props.locale || undefined, {
      notation: props.compact ? 'compact' : 'standard',
      maximumFractionDigits: 1,
    }).format(props.value as number)
    return n
  } catch {
    return String(props.value)
  }
})

// ---------- Trend tracking (intern) ----------
const history = ref<Array<{ t: number; v: number }>>([])
const nowTs = ref(Date.now())
const timer = ref<number | null>(null)

const historyLimitMs = computed(() => {
  const min = 5 * 60_000 // mind. 5 Min
  const suggested = (props.trendOffsetMs ?? 60_000) * 10
  return Math.max(props.historyLimitMs ?? suggested, min)
})

function pushSample(ts: number, v: number) {
  if (typeof v !== 'number' || Number.isNaN(v)) return
  const last = history.value[history.value.length - 1]
  if (!last || last.v !== v) history.value.push({ t: ts, v })
  // Cleanup alte Einträge
  const cutoff = ts - historyLimitMs.value
  let i = 0
  while (i < history.value.length && history.value[i].t < cutoff) i++
  if (i > 0) history.value.splice(0, i)
}

// finde Referenzwert zum Zeitpunkt (now - offset)
function valueAt(ts: number): number | null {
  const arr = history.value
  if (arr.length === 0) return null
  // binäre Suche wäre overkill; lineare reicht hier (History ist kurz)
  let prev = arr[0]
  for (let i = 1; i < arr.length; i++) {
    const cur = arr[i]
    if (cur.t >= ts) {
      // lineare Interpolation zwischen prev und cur, falls möglich
      if (cur.t === prev.t) return cur.v
      const r = (ts - prev.t) / (cur.t - prev.t)
      return prev.v + (cur.v - prev.v) * Math.min(Math.max(r, 0), 1)
    }
    prev = cur
  }
  // Wenn ts hinter dem letzten Punkt liegt → letzten Wert
  return arr[arr.length - 1].v
}

const internalDiff = computed<number | null>(() => {
  if (!isNumber.value) return null
  if (!props.trendOffsetMs) return null
  const referenceTs = nowTs.value - props.trendOffsetMs
  const refVal = valueAt(referenceTs)
  const curVal = props.value as number
  if (refVal === null || refVal === 0) return null
  const pct = ((curVal - refVal) / refVal) * 100
  // Begrenze auf sinnvolle Nachkommastellen
  return Number.isFinite(pct) ? pct : null
})

onMounted(() => {
  // Startwerte erfassen
  const n = typeof props.value === 'number' ? (props.value as number) : NaN
  pushSample(Date.now(), n)
  if (props.trendOffsetMs) {
    timer.value = window.setInterval(() => {
      nowTs.value = Date.now()
      // zyklisch den aktuellen Wert nochmal einspeisen, damit Interpolation/Ref stimmt
      const v = typeof props.value === 'number' ? (props.value as number) : NaN
      pushSample(nowTs.value, v)
    }, props.pollMs)
  }
})

onUnmounted(() => {
  if (timer.value) window.clearInterval(timer.value)
})

watch(() => props.value, (nv) => {
  const ts = Date.now()
  const num = typeof nv === 'number' ? (nv as number) : NaN
  pushSample(ts, num)
})

// ---------- Anzeige / Ableitung ----------
const effectiveDiff = computed<number | null>(() => {
  // Priorität: manuell übergebenes diff > intern berechnetes diff
  if (props.diff !== undefined && props.diff !== null) return props.diff
  return internalDiff.value
})

const showTrend = computed(() => effectiveDiff.value !== null && effectiveDiff.value !== undefined)
const trendUp = computed(() => (effectiveDiff.value ?? 0) > 0)
const trendDown = computed(() => (effectiveDiff.value ?? 0) < 0)
const trendClass = computed(() => ({ up: trendUp.value, down: trendDown.value }))
const diffText = computed(() => {
  if (props.diffLabel) return props.diffLabel
  if (effectiveDiff.value === null || effectiveDiff.value === undefined) return ''
  const val = Number(effectiveDiff.value)
  const sign = val > 0 ? '+' : ''
  return `${sign}${val.toFixed(1)}%`
})
const trendAria = computed(() => trendUp.value ? 'Trend steigend' : trendDown.value ? 'Trend fallend' : 'Kein Trend')
</script>

<template>
  <div
    class="stat"
    :class="[size]"
    :data-variant="variant"
    :aria-busy="loading ? 'true' : 'false'"
    role="group"
    :aria-label="label"
  >
    <div class="leading" aria-hidden="true">
      <slot name="icon">
        <!-- Default icon (sparkline) -->
        <svg class="icon" viewBox="0 0 24 24" focusable="false">
          <path d="M3 17l6-8 3 4 4-6 5 10" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </slot>
    </div>

    <dl class="content">
      <div class="row value-row">
        <dt class="label">{{ label }}</dt>
        <dd v-if="!loading" class="value">
          <span v-if="prefix" class="affix">{{ prefix }}</span>
          <span class="val">{{ formattedValue }}</span>
          <span v-if="suffix" class="affix">{{ suffix }}</span>
        </dd>
        <dd v-else class="value skeleton" aria-hidden="true"></dd>
      </div>

      <div v-if="(showTrend && !loading) || $slots.meta" class="row meta">
        <dd v-if="showTrend && !loading" class="trend" :class="trendClass" :aria-label="trendAria">
          <svg v-if="trendUp" class="arrow" viewBox="0 0 24 24" aria-hidden="true">
            <path d="M4 14l6-6 6 6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
          <svg v-else-if="trendDown" class="arrow" viewBox="0 0 24 24" aria-hidden="true">
            <path d="M4 10l6 6 6-6" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
          </svg>
          <span class="diff">{{ diffText }}</span>
        </dd>
        <dd v-if="$slots.meta" class="extra"><slot name="meta" /></dd>
      </div>
    </dl>

    <button
      v-if="clickable"
      class="ghost"
      type="button"
      @click="emit('click')"
      :aria-label="`Details zu ${label}`"
    />
  </div>
</template>

<style scoped>
:root {}

.stat {
  --bg: var(--stat-bg, hsl(0 0% 100% / 1));
  --fg: var(--stat-fg, hsl(224 15% 15%));
  --muted: var(--stat-muted, hsl(224 6% 46%));
  --accent: var(--stat-accent, hsl(220 90% 56%));
  --up: hsl(142 72% 35%);
  --down: hsl(0 72% 45%);

  position: relative;
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 0.9rem 1rem;
  align-items: start;
  padding: 1rem 1.1rem;
  border-radius: 1rem;
  background: var(--bg);
  color: var(--fg);
  border: 1px solid hsl(220 14% 90%);
  box-shadow: 0 1px 2px hsl(220 3% 10% / 0.04);
}

/* Variants */
.stat[data-variant="solid"] {
  background: hsl(220 15% 96%);
  border-color: transparent;
}
.stat[data-variant="outline"] {
  background: transparent;
  border-style: dashed;
}
.stat[data-variant="subtle"] {
  background: var(--bg);
}

/* Sizes */
.stat.sm { padding: .7rem .85rem; border-radius: .75rem; }
.stat.md { padding: 1rem 1.1rem; }
.stat.lg { padding: 1.25rem 1.4rem; border-radius: 1.1rem; }
.stat.xl { padding: 1.5rem 1.8rem; border-radius: 1.25rem; gap: 1.2rem 1.4rem; grid-template-columns: auto 1fr; }

.leading { inline-size: 40px; block-size: 40px; display: grid; place-items: center; border-radius: .9rem; background: hsl(220 15% 96%); color: var(--accent); }
.stat.sm .leading { inline-size: 34px; block-size: 34px; border-radius: .7rem; }
.stat.lg .leading { inline-size: 48px; block-size: 48px; border-radius: 1rem; }
.stat.xl .leading { inline-size: 56px; block-size: 56px; border-radius: 1.1rem; }
.icon { width: 55%; height: 55%; }

.content { margin: 0; display: grid; gap: .35rem; }
.row { display: flex; align-items: baseline; gap: .5rem; flex-wrap: wrap; }

.label { margin: 0; font-size: .82rem; letter-spacing: .02em; color: var(--muted); font-weight: 600; }
.value { margin: 0 0 0 auto; font-variant-numeric: tabular-nums; font-weight: 700; font-size: 1.55rem; display: inline-flex; align-items: baseline; gap: .25rem; }
.stat.sm .value { font-size: 1.25rem; }
.stat.lg .value { font-size: 1.9rem; }
.stat.xl .value { font-size: 2.2rem; }
.val { line-height: 1; }
.affix { color: var(--muted); font-weight: 600; font-size: .95em; }

.meta { gap: .6rem; }
.trend { display: inline-flex; align-items: center; gap: .35rem; font-weight: 600; padding: .15rem .45rem; border-radius: .5rem; background: hsl(220 16% 94%); color: var(--muted); }
.trend.up { background: hsl(146 36% 94%); color: var(--up); }
.trend.down { background: hsl(0 36% 95%); color: var(--down); }
.arrow { width: 1em; height: 1em; }
.diff { font-variant-numeric: tabular-nums; }

.extra { color: var(--muted); }

/* Loading skeleton */
.skeleton { position: relative; min-inline-size: 8ch; min-block-size: 1.2em; border-radius: .4rem; background: linear-gradient(90deg, hsl(0 0% 90% / .6), hsl(0 0% 96%), hsl(0 0% 90% / .6)); background-size: 200% 100%; animation: shimmer 1.25s linear infinite; }
@keyframes shimmer { to { background-position-x: -200%; } }

/* Full-card clickable overlay */
.ghost { position: absolute; inset: 0; border: 0; background: none; cursor: pointer; border-radius: inherit; }
.ghost:focus-visible { outline: 2px solid var(--accent); outline-offset: 2px; }

/* Dark scheme */
@media (prefers-color-scheme: dark) {
  .stat { --bg: hsl(222 20% 10%); --fg: hsl(0 0% 98%); --muted: hsl(220 10% 65%); border-color: hsl(220 16% 18%); box-shadow: 0 1px 1px hsl(220 3% 10% / 0.3); }
  .leading { background: hsl(222 16% 16%); }
  .trend { background: hsl(220 16% 18%); }
  .trend.up { background: hsl(146 30% 18%); }
  .trend.down { background: hsl(0 50% 18%); }
  .skeleton { background: linear-gradient(90deg, hsl(220 10% 22%), hsl(220 10% 18%), hsl(220 10% 22%)); }
}
</style>
