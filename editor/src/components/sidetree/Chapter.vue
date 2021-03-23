<template>
	<div class="item" @click="getElements">
    <div class="name"><vue-feather type="bookmark" size="1em"></vue-feather>{{ name }}</div>
    <div class="children">
      <Note :key="item.id" v-for="item in elements.notes" :name="item.name" :id="item.id"></Note>
    </div>
  </div>
</template>

<script>
import Note from './Note.vue'

export default {
  components: {
		Note
	},

  props: {
		id: String,
    name: String,
  },
  
  data() { return {
      elements: {
				notes: []
			}
    }
  },

  methods: {
    async getElements() {
			const res = await fetch(`http://localhost:3000/api/chapter/${this.id}`)
      this.elements = await res.json()
		}
	}
}
</script>

<style scoped>

</style>
