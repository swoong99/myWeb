<template>
  <div class="my-research">
    <h1>Research</h1>
    <p>Here is a summary of my research work:</p>
    <ul>
      <li v-for="paper in researchPapers" :key="paper.id">
        <h2>{{ paper.title }}</h2>
        <p>{{ paper.description }}</p>
        <p>{{ paper.stack }}</p>
        <p>{{ paper.duration }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "MyResearch",
  data() {
    return {
      researchPapers: [],
    };
  },
  created() {
    axios
        .get("http://localhost:8080/api/research")
        .then((response) => {
          console.log("Received data:", response.data);
          this.researchPapers = response.data;
        })
        .catch((error) => {
          console.error("Error fetching research data:", error);
        });
  },
};
</script>

<style scoped>
.my-research {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  margin: 20px;
}

h1 {
  color: #2c3e50;
}

ul {
  padding-left: 20px;
}
</style>