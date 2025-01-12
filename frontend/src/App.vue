<template>
  <div id="app">
    <header>
      <nav>
        <button @click="navigateTo('introduction')">Introduction</button>
        <button @click="navigateTo('projects')">Projects</button>
        <button @click="navigateTo('research')">Research</button>
      </nav>
    </header>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<script>
export default {
  name: "App",
  methods: {
    navigateTo(section) {
      this.$router.push(`/${section}`);
    },
  },
};
</script>

<style>
header {
  background-color: #333;
  color: white;
  padding: 1rem;
  text-align: center;
}

nav button {
  background-color: #444;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  margin: 0 0.5rem;
  cursor: pointer;
}

nav button:hover {
  background-color: #555;
}

main {
  padding: 2rem;
}
</style>