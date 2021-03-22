<template>
<main>
  <SideBar></SideBar>
	<Editor></Editor>
</main>
</template>

<script>
import SideBar from '@/components/SideBar.vue'
import Editor from '@/components/Editor.vue'

export default {
  components: {
    SideBar,
    Editor
  }
}
</script>


<style scoped>
main {
	display: flex;
	flex-direction: row;
}
</style>