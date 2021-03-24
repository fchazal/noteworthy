<template>
<main id="editor"></main>
</template>

<script>
  import 'codemirror/lib/codemirror.css'; // Editor's Dependency Style
  import '@toast-ui/editor/dist/toastui-editor.css'
  import Editor from '@toast-ui/editor';

  export default {
    data: () => {
      return { editor: null }
    },
    mounted() {
      this.editor = new Editor({
        el: document.querySelector('main#editor'),
        initialEditType: 'markdown',
        previewStyle: 'vertical'
      })
    }
  };
</script>

<style>
main {
  display: flex;
  height: unset !important;
}

.tui-editor-defaultUI {
  border: none;
  flex-grow: 1;
}
</style>