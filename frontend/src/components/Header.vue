<template>
  <div class="container flex justify-center items-center space-x-1 mx-auto">
    <VueWriter
      :array="headerArr"
      class="text-6xl font-semibold"
      :typeSpeed="80"
      :eraseSpeed="70"
      caret="underscore"
    ></VueWriter>
  </div>
</template>

<script setup>
import VueWriter from "vue-writer";
const headerArr = ["Comment", "Refactor", "Prettify", "Convert"];
</script>

<style scoped>
.is-typed span.cursor {
  display: inline-block;
  width: 3px;
  background-color: black;
  animation: blink 1s infinite;
}
</style>
