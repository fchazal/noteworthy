<template>
<aside>
  <button @click="getElements()">REFRESH</button>
  <Library :key="item.id" v-for="item in elements.libraries" :name="item.name" :id="item.id"></Library>
  <Book :key="item.id" v-for="item in elements.books" :name="item.name" :id="item.id"></Book>
</aside>
</template>

<script>
import Library from './sidetree/Library.vue'
import Book from './sidetree/Book.vue'

export default {
  components: {
    Library,
    Book
  },
  data() { return {
      elements: {}
    }
  },
  beforeMount(){
    this.getElements();
  },
  methods: {
    async getElements(){
      const res = await fetch('http://localhost:3000/api/store/')
      this.elements = await res.json()
    }
  }
}
</script>

<style>
aside {
  max-width: 300px;
  border-right: 5px solid rgba(0,0,0,0.02);
}

aside button {
  display: block;
  outline: none;

  padding: .2em 1em;
  margin: 1em auto;
  border: none;
  border-radius: 1em;

  font: inherit;
  font-size: 0.8em;
  color: inherit;
}

aside .vue-feather {
  margin: auto .5em;
}
</style>