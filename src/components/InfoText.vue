<template>
  <div
    class="flex items-center mt-8 my-7 mx-5 sm:mt-12 sm:mx-7 sm:my-7 space-x-5 justify-start text-left"
  >
    <i :class="`${iconStyle} text-4xl`"></i>
    <p class="text-lg mx-2 animate__fadeInRight">{{ text }}</p>
  </div>
</template>

<script setup>
const props = defineProps({
  iconStyle: String,
  text: String,
});
</script>
