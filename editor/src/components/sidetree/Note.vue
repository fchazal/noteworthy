<template>
	<div class="item" @click="getElements"><vue-feather type="file"></vue-feather>{{ name }}</div>
</template>

<script>
export default {
  props: {
		id: String,
    name: String,
  },
}
</script>

<style scoped>

</style>
