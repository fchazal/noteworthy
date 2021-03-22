<template>
	<div class="item" @click="getElements"><vue-feather type="box"></vue-feather>{{ name }}</div>
	<div class="children">
    <Book :key="item.id" v-for="item in elements.books" :name="item.name" :id="item.id"></Book>
  </div>
</template>

<script>
import Book from './Book.vue'

export default {
  components: {
    Book
  },

  props: {
		id: String,
    name: String,
  },
	
  data() { return {
      elements: {
				books: []
      }
    }
  },

  methods: {
    async getElements() {
			const res = await fetch(`http://localhost:3000/api/library/${this.id}`)
      this.elements = await res.json()
		}
	}
}
</script>

<style>
.item {
  padding: 0.5em;
  margin: 4px;
  margin-right: 0;
  background: rgba(0,0,0,0.02);
  border-radius: 2em 0 0 2em;
}
.item:hover {
  background: rgba(0,0,0,0.05);
}
.children {
  margin-left: 2em;
}
</style>
