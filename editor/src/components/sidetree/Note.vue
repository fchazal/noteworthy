<template>
	<div class="item" @click="getElements">
    <div class="name"><vue-feather type="clipboard" size="1em"></vue-feather>{{ name }}</div>
  </div>
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
