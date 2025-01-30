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

const showEditModal = ref(false); // Affichage de la modale de modification
const carToEdit = ref(null); // Voiture en cours de modification
const editedCar = ref({ name: "", price: null }); // Valeurs modifiées

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

const addCar = async () => {
  try {
    const response = await fetch("http://localhost:8000/api/v1/cars", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: newCar.value.name, price: newCar.value.price })
    });

    if (!response.ok) {
      throw new Error("Erreur lors de l'ajout de la voiture.");
    }

    //const addedCar = await response.json();

    await fetchCars(); // Recharge toute la liste après ajout
    closeModal();
    newCar.value = { name: "", price: null };
  } catch (err) {
    error.value = err.message;
  }
};


const deleteCar = async () => {
  if (!carToDelete.value) return; // Vérifie qu'une voiture est bien sélectionnée

  try {
    const response = await fetch(`http://localhost:8000/api/v1/cars/${carToDelete.value.id}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      throw new Error("Erreur lors de la suppression de la voiture.");
    }

    await fetchCars(); // Recharge toute la liste après ajout
    closeDeleteModal(); // Ferme la modale après suppression

  } catch (err) {
    error.value = err.message;
  }
};

const updateCar = async () => {
  if (!carToEdit.value) return;

  try {
    const response = await fetch(`http://localhost:8000/api/v1/cars/${carToEdit.value.id}`, {
      method: "PUT", // Utilisation de la méthode PUT pour mettre à jour
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(editedCar.value),
    });

    if (!response.ok) {
      throw new Error("Erreur lors de la mise à jour de la voiture.");
    }

    await fetchCars(); // Recharge toute la liste après ajout

    closeEditModal(); // Fermer la modale après la mise à jour
  } catch (err) {
    error.value = err.message;
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

const openCarModal = (car) => {
  selectedCar.value = car;
  showCarModal.value = true;
};

// Fonction pour fermer la modale de visualisation
const closeCarModal = () => {
  showCarModal.value = false;
  selectedCar.value = null;
};

const openEditModal = (car) => {
  carToEdit.value = car; // Sauvegarde la voiture sélectionnée
  editedCar.value = { ...car }; // Copie des données existantes
  showEditModal.value = true; // Affiche la modale
};

const closeEditModal = () => {
  showEditModal.value = false;
  carToEdit.value = null;
  editedCar.value = { name: "", price: null }; // Réinitialisation du formulaire
};

// Appel automatique au montage du composant
onMounted(fetchCars);
</script>

<template>
  <h1>Porsche</h1>
    <button class="actions-button" @click="openModal">ADD CAR</button>

  <!-- Modale d'ajout d'un véhicule -->
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
      <img v-if="selectedCar" src="../../assets/car1.jpg" class="modal-car-image" alt="porsche car"/>
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
        <button @click="deleteCar" class="suppr-button">Supprimer</button>
        <button @click="closeDeleteModal" class="cancel-button">Annuler</button>
      </div>
    </div>
  </div>

  <!-- Modale de confirmation de modification -->
  <div v-if="showEditModal" class="modal">
    <div class="modal-content">
      <h2>Modifier la voiture</h2>
      <form @submit.prevent="updateCar">
        <div>
          <label for="carName">Nom de la voiture:</label>
          <input type="text" id="carName" v-model="editedCar.name" placeholder="Nom de la voiture" />
        </div>
        <div>
          <label for="carPrice">Prix:</label>
          <input type="number" id="carPrice" v-model="editedCar.price" placeholder="Prix de la voiture" />
        </div>
        <button type="submit">Enregistrer</button>
        <button class="close" @click="closeEditModal()">Annuler</button>
      </form>
    </div>
  </div>

  <div class="container">
    <div v-if="loading" class="loading">Chargement...</div>
    <div v-if="error" class="error">{{ error }}</div>

    <ul v-if="!loading && !error" class="car-list">
      <li v-for="car in cars" :key="car.id" class="car-item">
        <img src="../../assets/car1.jpg" class="car-image" alt="porsche car"/>
        <h3 class="car-name" style="text-transform: uppercase;">{{ car.name }}</h3>
        <p class="price">{{ car.price.toLocaleString() }} € </p>
        <div class="actions">
          <button class="view-button" @click="openCarModal(car)" >Voir le véhicule</button>
          <div class="actions2">
          <button class="del-button" @click="openDeleteModal(car)">🗑️</button>
          <button class="edit-button" @click="openEditModal(car)">✏️</button>
          </div>
        </div>

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

.suppr-button:hover {
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
  left: 230px;
  background: #7c7c7c;
  color: white;
  padding: 5px 10px;
  border-radius: 5px;
  cursor: pointer;
  width: 20%;
}

.close-modal:hover{
}

/* Styles pour la liste des voitures */
.container {
  margin: 0;
  padding: 0;
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
  text-align: left;
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


.actions2,button {
  margin: 10px;

}

.del-button:hover {
  background-color: darkred;
}

.actions{
  text-align: center;
}

</style>