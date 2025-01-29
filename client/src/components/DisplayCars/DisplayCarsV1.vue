<script setup>
import { ref, onMounted } from "vue";

const cars = ref([]);
const loading = ref(true);
const error = ref(null);

const fetchCars = async () => {
  try {
    const response = await fetch("http://localhost:8000/api/v1/cars");

    if (!response.ok) {
      throw new Error("Erreur lors de la récupération des voitures.");
    }

    cars.value = await response.json(); // Convertit la réponse en JSON
  } catch (err) {
    error.value = err.message;
  } finally {
    loading.value = false;
  }
};

// Appel automatique au montage du composant
onMounted(fetchCars);
</script>

<template>
  <div class="container">
    <h1>Cars </h1>

    <div v-if="loading" class="loading">Chargement...</div>
    <div v-if="error" class="error">{{ error }}</div>

    <ul v-if="!loading && !error" class="car-list">
      <li v-for="car in cars" :key="car.id">
        <strong>{{ car.name }}</strong> - 💰 {{ car.price.toLocaleString() }} €
      </li>
    </ul>
  </div>
</template>

<style scoped>
.container {
  max-width: 600px;
  margin: auto;
  padding: 20px;
  font-family: Arial, sans-serif;
}
.loading {
  color: blue;
  font-weight: bold;
}
.error {
  color: red;
  font-weight: bold;
}
.car-list {
  list-style: none;
  padding: 0;
}
.car-list li {
  background: #000000;
  margin: 10px 0;
  padding: 10px;
  border-radius: 5px;
}
</style>
