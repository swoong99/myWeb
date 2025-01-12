<template>
  <div class="my-introduction">
    <h1>Introduction</h1>
    <p><strong>Name:</strong> {{ introduction.name }}</p>
    <p><strong>소개:</strong> {{ introduction.bio }}</p>
    <p><strong>Education:</strong> {{ introduction.education }}</p>
    <p><strong>Skills:</strong></p>
    <ul>
      <li v-for="skill in introduction.skills" :key="skill">{{ skill }}</li>
    </ul>
    <p><strong>연락처:</strong> {{ introduction.contact }}</p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "MyIntroduction",
  data() {
    return {
      introduction: {},
    };
  },
  mounted() {
    axios
        .get("http://localhost:8080/api/introduction")
        .then((response) => {
          console.log("Received data:", response.data);
          this.introduction = response.data;
        })
        .catch((error) => {
          console.error("Failed to fetch introduction data:", error);
        });
  },
};
</script>

<style scoped>
.my-introduction {
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
