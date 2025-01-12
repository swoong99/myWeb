<template>
  <div class="my-projects">
    <h1>Projects</h1>
    <p>Here are some of the projects I have worked on:</p>
    <ul>
      <li v-for="project in projects" :key="project.id">
        <h2>{{ project.title }}</h2>
        <p>{{ project.description }}</p>
        <p>{{ project.stack }}</p>
        <p>{{ project.duration }}</p>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "MyProjects",
  data() {
    return {
      projects: [],
    };
  },
  mounted() {
    axios
        .get("http://localhost:8080/api/projects")
        .then((response) => {
          console.log("Received data:", response.data);
          this.projects = response.data;
        })
        .catch((error) => {
          console.error("Error fetching projects data:", error);
        });
  },
};
</script>

<style scoped>
.my-projects {
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