<template>
<header>
	<h1>
		<vue-feather type="edit"></vue-feather>
		Note Editor
	</h1>
</header>
</template>

<script>
export default {
  setup() {
    
  },
}
</script>

<style scoped>
header {
	padding: 1em;

	color: #fff;
	background-color: #69e;
}

h1 {
	display: flex;
	align-items: center;

	margin: 0;
	padding: 0;
	font-weight: 400;
}

h1 .vue-feather { margin-right: .2em }
</style>