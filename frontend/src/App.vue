<template>
  <div id="app">
    <h1>File Indexer</h1>
    <button @click="indexFiles">Index Files</button>
    <p>{{ message }}</p>
    <input v-model="query" placeholder="Search query" />
    <button @click="searchFiles">Search</button>
    <ul>
      <li v-for="file in files" :key="file">{{ file }}</li>
    </ul>
  </div>
</template>

<script>
import { ref } from 'vue';
import { IndexFiles, SearchFiles } from '../wailsjs/go/backend/App';

export default {
  name: 'App',
  setup() {
    const message = ref('');
    const query = ref('');
    const files = ref([]);

    const indexFiles = async () => {
      message.value = await IndexFiles();
    };

    const searchFiles = async () => {
      files.value = await SearchFiles(query.value);
    };

    return {
      message,
      query,
      files,
      indexFiles,
      searchFiles,
    };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>