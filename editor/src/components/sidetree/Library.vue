<template>
	<div class="item" @click="getElements">
    <div class="name"><vue-feather type="archive" size="1em"></vue-feather>{{ name }}</div>
    <div class="children">
      <Book :key="item.id" v-for="item in elements.books" :name="item.name" :id="item.id"></Book>
    </div>
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
</style>
