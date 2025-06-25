<template>
  <div class="container">
    <h1>🔢 Hitung Entropi</h1>

    <div class="input-group">
      <label for="mode">Pilih Mode:</label>
      <select id="mode" v-model="selectedMode">
        <option value="text">Teks</option>
        <option value="csv">CSV</option>
        <option value="image">Gambar</option>
      </select>
    </div>

    <div v-if="selectedMode === 'text'">
      <label for="text">Masukkan Teks:</label>
      <textarea
        id="text"
        rows="5"
        v-model="inputText"
        placeholder="Ketik teks di sini..."
      ></textarea>
    </div>

    <div v-else>
      <label for="file">Unggah File:</label>
      <input type="file" id="file" @change="handleFileChange" />
    </div>

    <button @click="submit">🚀 Hitung</button>
    <button @click="downloadPDF" class="download">📄 Unduh PDF</button>

    <div v-if="entropy !== null" class="result-box">
      <p><strong>Total Karakter:</strong> {{ totalKarakter }}</p>
      <p
        :style="{
          color: entropy < 2 ? 'red' : entropy < 3.5 ? 'orange' : 'green',
        }"
      >
        {{
          entropy < 2
            ? "Plagiarisme Mungkin"
            : entropy < 3.5
            ? "Perlu Dicek"
            : "Teks Unik / Asli"
        }}
      </p>

      <client-only>
        <canvas id="charChart"></canvas>
      </client-only>

      <div class="frekuensi-grid">
        <h3>📊 Frekuensi Karakter:</h3>
        <div class="frekuensi-list">
          <div
            v-for="(val, key) in frekuensi"
            :key="key"
            class="frekuensi-item"
          >
            <span class="char">{{ key }}</span>
            <span class="count">{{ val }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, nextTick } from "vue";
import Chart from "chart.js/auto";

const selectedMode = ref("text");
const inputText = ref("");
const entropy = ref(null);
const frekuensi = ref({});
const totalKarakter = ref(0);
let chart = null;
const fileInput = ref(null);

function handleFileChange(event) {
  fileInput.value = event.target.files[0];
}

async function submit() {
  let res;
  if (selectedMode.value === "text") {
    res = await useFetch("http://localhost:8080/entropi/text", {
      method: "POST",
      body: { text: inputText.value },
    });
  } else {
    const formData = new FormData();
    formData.append("file", fileInput.value);
    const endpoint = selectedMode.value === "csv" ? "csv" : "image";
    res = await useFetch(`http://localhost:8080/entropi/${endpoint}`, {
      method: "POST",
      body: formData,
    });
  }

  const { data, error } = res;
  if (!error.value) {
    entropy.value = data.value.entropi;
    frekuensi.value = data.value.frekuensi;
    totalKarakter.value = data.value.total_karakter;
    renderChart();
  }
}

async function renderChart() {
  await nextTick();
  const canvas = document.getElementById("charChart");
  if (!canvas) return;
  if (chart && chart.canvas) chart.destroy();
  chart = new Chart(canvas, {
    type: "bar",
    data: {
      labels: Object.keys(frekuensi.value),
      datasets: [
        {
          label: "Frekuensi Karakter",
          data: Object.values(frekuensi.value),
          backgroundColor: "#2c3e50",
        },
      ],
    },
  });
}

function downloadPDF() {
  window.open("http://localhost:8080/entropi/pdf", "_blank");
}
</script>

<style scoped>
.container {
  padding: 2rem;
  max-width: 600px;
  margin: auto;
  font-family: "Segoe UI", sans-serif;
}

h1 {
  text-align: center;
  color: #ed2d56;
  margin-bottom: 1rem;
}

.input-group {
  margin-bottom: 1rem;
}
label {
  display: block;
  margin-bottom: 6px;
  font-weight: bold;
}

select,
textarea,
input[type="file"] {
  width: 100%;
  padding: 10px;
  margin-bottom: 1rem;
  border-radius: 6px;
  border: 1px solid #ccc;
  font-size: 1rem;
}

button {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #ed2d56;
  color: white;
  font-weight: bold;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  transition: 0.3s;
  margin-bottom: 10px;
}
button:hover {
  background-color: #c91d45;
}
.download {
  background-color: #2c3e50;
}
.download:hover {
  background-color: #1a252f;
}

.result-box {
  background-color: #f9f9f9;
  border-radius: 8px;
  padding: 1.5rem;
  margin-top: 1.5rem;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);
}

canvas {
  margin: 1rem 0;
  max-width: 100%;
}

.frekuensi-grid {
  margin-top: 1rem;
}
.frekuensi-list {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.frekuensi-item {
  background: #ed2d56;
  color: white;
  border-radius: 6px;
  padding: 0.5rem 1rem;
  display: flex;
  justify-content: space-between;
  min-width: 60px;
  font-weight: bold;
}
.char {
  margin-right: 0.5rem;
}
</style>
