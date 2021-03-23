<template>
	<div class="item" @click="getElements">
    <div class="name"><vue-feather type="book" size="1em"></vue-feather>{{ name }}</div>
    <div class="children">
      <Chapter :key="item.id" v-for="item in elements.chapters" :name="item.name" :id="item.id"></Chapter>
      <Note :key="item.id" v-for="item in elements.notes" :name="item.name" :id="item.id"></Note>
    </div>
  </div>
</template>

<script>
import Chapter from './Chapter.vue'
import Note from './Note.vue'

export default {
  components: {
    Chapter,
		Note
	},
  props: {
		id: String,
    name: String,
  },
  data() { return {
      elements: {
				chapters: [],
				notes: []
			}
    }
  },

  methods: {
    async getElements() {
			const res = await fetch(`http://localhost:3000/api/book/${this.id}`)
      this.elements = await res.json()
		}
	}
}
</script>

<style scoped>

</style>
