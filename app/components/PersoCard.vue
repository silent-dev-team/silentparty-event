<script setup lang="ts">
const props = defineProps<{
  form: {
    firstName: string;
    lastName: string;
    street: string;
    housenumber: string;
    zipCode: string;
    place: string;
  }
}>()
const { firstName, lastName, street, housenumber, zipCode, place } = toRefs(props.form);
</script>

<template>
  <div class="perso-view">
    <section class="perso-card" aria-label="Personalausweis Vorschau">
      <div class="perso-bg"></div>
      <div class="perso-overlay"></div>

      <header class="perso-header">
        <h1>PERSONALAUSWEIS</h1>
      </header>

      <div class="perso-body">
        <div class="photo" aria-hidden="true"></div>

        <div class="fields">
          <div class="row">
            <span class="lbl">Name / Surname</span>
            <span class="val">{{ lastName }}</span>
          </div>
          <div class="row">
            <span class="lbl">Vorname / Given name</span>
            <span class="val">{{ firstName }}</span>
          </div>

        </div>
      </div>
    </section>
    <section class="perso-card" aria-label="Personalausweis Vorschau">
      <div class="perso-bg"></div>
      <div class="perso-overlay"></div>

      <header class="perso-header">
        <h1>Rückseite</h1>
      </header>

      <div class="perso-body">
        <div class="photo" style="opacity: 0;" aria-hidden="true"></div>
        <div class="fields">
          <div class="row muted">
            <span class="lbl">Anschrift / Address / Adresse</span>
            <span class="val">
              {{ zipCode }} {{ place }} <br>
              {{ street }} {{ housenumber }}
            </span>
          </div>
        </div>
      </div>

      <footer class="mrz" aria-hidden="true">
        <div class="mrz-line">
          &lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;
        </div>
        <div class="mrz-line">
          &lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;
        </div>
        <div class="mrz-line">
          {{ (lastName) }}&lt;&lt;{{ (firstName) }}&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;&lt;
        </div>
      </footer>
    </section>
  </div>
</template>


<style scoped>
.perso-view {
  display: flex;
  gap: 1.5rem;
  width: 100%;
  margin: 0 auto;
  padding: 1rem;
  scale: 1.5;
}

.perso-card {
  position: relative;
  /* aspect-ratio: 86 / 54;  */
  width: 100%;
  border-radius: 12px;
  overflow: hidden;
  box-shadow:
    0 2px 8px rgba(0,0,0,0.15),
    inset 0 0 0 1px rgba(0,0,0,0.08);
  background: #f6ecdf;
  isolation: isolate;
}

/* background look */
.perso-bg {
  position: absolute; inset: 0;
  background:
    radial-gradient(120% 80% at 10% 10%, rgba(255,255,255,0.9), rgba(255,255,255,0) 60%),
    radial-gradient(120% 80% at 90% 20%, rgba(255,255,255,0.7), rgba(255,255,255,0) 65%),
    linear-gradient(135deg, #f6ecdf 0%,  #d3e0d9 100%);
  mix-blend-mode: screen;
  z-index: 1;
}
.perso-overlay {
  position: absolute; inset: 0;
  background-image:
    repeating-linear-gradient( 0deg, rgba(0,0,0,0.03) 0 1px, transparent 1px 3px),
    repeating-linear-gradient(90deg, rgba(0,0,0,0.02) 0 1px, transparent 1px 3px);
  opacity: .5;
  z-index: 1;
  pointer-events: none;
}

/* header */
.perso-header {
  position: relative; z-index: 2;
  display: flex; align-items: center; gap: .5rem;
  padding: .75rem .9rem .25rem;
}
.perso-header h1 {
  margin: 0; font-size: .92rem; letter-spacing: .12em; font-weight: 800;
  color: #1a2f27; text-transform: uppercase;
}

/* body */
.perso-body {
  height: calc(100% - 2.1rem);
  position: relative; z-index: 2;
  display: grid;
  grid-template-columns: 110px 1fr;
  gap: .9rem;
  padding: .4rem .9rem .9rem;
  align-items: start;
}
.photo {
  width: 100%; aspect-ratio: 3/4;
  background:
    linear-gradient(135deg, rgba(0,0,0,0.06), rgba(0,0,0,0.02)),
    repeating-linear-gradient(45deg, rgba(0,0,0,0.05) 0 2px, transparent 2px 6px);
  border-radius: 6px;
  box-shadow: inset 0 0 0 1px rgba(0,0,0,.1);
}

.fields { display: grid; gap: .35rem; }
.row { display: grid; gap: .1rem; }
.lbl {
  font-size: .62rem; letter-spacing: .06em; text-transform: uppercase;
  color: #2b4a3d; opacity: .85;
}
.val {
  font-feature-settings: "ss01" 1, "tnum" 1, "lnum" 1;
  font-weight: 700;
  font-size: 1.05rem;
  color: #10251d;
  text-shadow: 0 1px 0 rgba(255,255,255,0.6);
}
.row.muted .val { font-weight: 600; font-size: .9rem; opacity: .8; }

/* fake MRZ */
.mrz {
  position: absolute; left: 0; right: 0; bottom: 0;
  z-index: 2;
  padding: .35rem .9rem .55rem;
  background: linear-gradient(to bottom, rgba(0,0,0,0), rgba(0,0,0,.03));
  border-top: 1px dashed rgba(0,0,0,0.08);
}
.mrz-line {
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", monospace;
  font-size: .75rem;
  letter-spacing: .08em;
  color: #0a1c16;
  opacity: .8;
  white-space: nowrap;
  overflow: hidden;
}
</style>
