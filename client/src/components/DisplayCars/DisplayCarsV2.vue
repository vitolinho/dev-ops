<script setup>
import { ref, onMounted } from "vue";

const cars = ref([]);
const loading = ref(true);
const error = ref(null);

const showModal = ref(false);

const showDeleteModal = ref(false);
const carToDelete = ref(null); // Stocke la voiture à supprimer

const showCarModal = ref(false);
const selectedCar = ref(null); // Stocke la voiture sélectionnée


// Champs du formulaire
const newCar = ref({
  name: "",
  price: null,
});

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

// Fonction pour afficher la modale
const openModal = () => {
  showModal.value = true;
};

// Fonction pour fermer la modale
const closeModal = () => {
  showModal.value = false;
};

// Fonction pour ajouter une nouvelle voiture (ici on l'ajoute à la liste en local)
const addCar = () => {
  if (newCar.value.name && newCar.value.price) {
    // Ajouter la voiture à la liste
    cars.value.push({ ...newCar.value, id: cars.value.length + 1 });
    // Réinitialiser le formulaire
    newCar.value.name = "";
    newCar.value.price = null;
    // Fermer la modale
    closeModal();
  } else {
    alert("Veuillez remplir tous les champs.");
  }
};

// Fonction pour ouvrir la modale de confirmation de suppression
const openDeleteModal = (car) => {
  carToDelete.value = car;
  showDeleteModal.value = true;
};

// Fonction pour fermer la modale de suppression
const closeDeleteModal = () => {
  showDeleteModal.value = false;
  carToDelete.value = null;
};

// Fonction pour supprimer une voiture
const deleteCar = () => {
  if (carToDelete.value) {
    cars.value = cars.value.filter((car) => car.id !== carToDelete.value.id);
    closeDeleteModal();
  }
};

const openCarModal = (car) => {
  selectedCar.value = car;
  showCarModal.value = true;
};

// Fonction pour fermer la modale de visualisation
const closeCarModal = () => {
  showCarModal.value = false;
  selectedCar.value = null;
};

// Appel automatique au montage du composant
onMounted(fetchCars);
</script>

<template>
  <h1>Porsche</h1>
    <button class="actions-button" @click="openModal">ADD CAR</button>
  <div v-if="showModal" class="modal">
    <div class="modal-content">
      <h2>Ajouter une voiture</h2>
      <form @submit.prevent="addCar">
        <div>
          <label for="carName">Nom de la voiture:</label>
          <input type="text" id="carName" v-model="newCar.name" placeholder="Nom de la voiture" />
        </div>
        <div>
          <label for="carPrice">Prix:</label>
          <input type="number" id="carPrice" v-model="newCar.price" placeholder="Prix de la voiture" />
        </div>
        <button type="submit">Ajouter la voiture</button>
        <button class="close" @click="closeModal">Annuler</button>
      </form>
    </div>
  </div>

  <!-- Modale de visualisation d'un véhicule -->
  <div v-if="showCarModal" class="modal">
    <div class="modal-content large">
      <img v-if="selectedCar" src="../../assets/car1.jpg" class="modal-car-image" />
      <h2 v-if="selectedCar" style="text-transform: uppercase;">{{ selectedCar.name }}</h2>
      <p v-if="selectedCar" class="modal-price">{{ selectedCar.price.toLocaleString() }} €</p>
      <button class="close-modal" @click="closeCarModal">Retour</button>
    </div>
  </div>

  <!-- Modale de confirmation de suppression -->
  <div v-if="showDeleteModal" class="modal">
    <div class="modal-content">
      <h2>Confirmation</h2>
      <p>Êtes-vous sûr de vouloir supprimer cette voiture ?</p>
      <div class="modal-actions">
        <button @click="deleteCar" class="confirm-button">Supprimer</button>
        <button @click="closeDeleteModal" class="close">Annuler</button>
      </div>
    </div>
  </div>

  <div class="container">
    <div v-if="loading" class="loading">Chargement...</div>
    <div v-if="error" class="error">{{ error }}</div>

    <ul v-if="!loading && !error" class="car-list">
      <li v-for="car in cars" :key="car.id" class="car-item">
        <img src="../../assets/car1.jpg" class="car-image" />
        <h3 class="car-name" style="text-transform: uppercase;">{{ car.name }}</h3>
        <p class="price">{{ car.price.toLocaleString() }} € </p>
        <div class="actions">
          <button class="view-button" @click="openCarModal(car)" >Voir le véhicule</button>
        </div>
        <button class="del-button" @click="openDeleteModal(car)">🗑️</button>
      </li>
    </ul>
  </div>
</template>

<style scoped>


#carName, #carPrice{
  width: 90%;
}

#carPrice{
  margin-bottom: 20px;
}

/* Contenu de la modale */
.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5); /* Fond semi-transparent */
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: #1e1e1e;
  padding: 20px;
  border-radius: 10px;
  width: 400px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.close {
  margin-top: 20px;
  cursor: pointer;
}

.confirm-button:hover {
  background-color: #9d0000;
}

/* Formulaire dans la modale */
form {
  display: flex;
  flex-direction: column;
}

form div {
  margin-bottom: 10px;
}

form input {
  padding: 10px;
  margin-top: 5px;
  width: 100%;
}

form button {
  padding: 10px;
  background-color: #7c7c7c;
  color: white;
  border-radius: 5px;
  cursor: pointer;
}


/* Modale plus grande pour visualisation */
.modal-content.large {
  width: 600px;
  height: 600px;
  display: grid;
}

.modal-car-image {
  width: 100%;
  height: 300px;
  object-fit: cover;
  border-radius: 10px;
}

.modal-price {
  font-size: 20px;
  font-weight: bold;
  margin-top: 10px;
}

.close-modal {
  position: relative;
  left: 240px;
  background: #7c7c7c;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  cursor: pointer;
  width: 20%;
}

.close-modal:hover{
  background-color: darkred;
}

/* Styles pour la liste des voitures */
.container {
  margin: 0px;
  padding: 0px;
  font-family: Arial, sans-serif;
}

.car-list {
  list-style: none;
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  padding: 0;
}

.car-item {
  width: calc(20% - 20px);
  box-sizing: border-box;
  background: #000000;
  border-radius: 20px;
  padding: 10px;
  text-align: center;
}

.car-image {
  width: 100%;
  height: 150px;
  object-fit: cover;
  border-radius: 5px;
}

.price {
  font-size: 16px;
  font-weight: bold;
}

.view-button {
  margin-top: 10px;
  padding: 5px 10px;
  background-color: #696969;
  color: white;
  border-radius: 5px;
  cursor: pointer;
}


.del-button {
  width: auto;
  height: 10%;
  margin: 15px;
}

.del-button:hover {
  background-color: darkred;
}


</style>