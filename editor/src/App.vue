<template>
  <Header></Header>
  <MainView></MainView>
</template>

<script>
import Header from './components/Header.vue'
import MainView from './views/MainView.vue'

export default {
  name: 'App',
  components: {
    Header,
    MainView
  }
}
</script>

<style>
html, body { padding: 0; margin: 0 }

body {
  display: flex;
  color: #444;
  font-size: 13px;
  font-family: Muli, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.vue-feather {
  vertical-align: middle;
}

#app {

  display: flex;
  flex-direction: column;

  width: 100vw;
  height: 100vh;
}

main {
  flex-grow: 1;
}

aside {
  flex-grow: 1;
}

</style>
