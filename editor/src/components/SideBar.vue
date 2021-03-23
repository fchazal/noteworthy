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
  padding-left: .5em;
  max-width: 300px;
  background-color: #69e4;
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
  margin: auto .5em auto 0;
  color: #69e;
  font-size: 1.2em;
}

aside .item {
  margin-bottom: 5px;
  background: #fff4;
  border-radius: 5px 0 0 5px;
  overflow: hidden;
}
aside .item:hover {
  background: #fff6;
}

aside .item .name {
  display: flex;
  padding: 0.5em;
  padding-right: 0;
}

aside .item .children {
  margin-left: 5px;
}
</style>