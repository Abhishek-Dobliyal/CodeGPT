<template>
  <div class="flex items-center my-10 mx-5 sm:mt-16 sm:mx-7 space-x-5 justify-start text-left">
    <i :class="`${iconStyle} text-4xl`"></i>
    <p class="text-lg mx-2">{{ text }}</p>
  </div>
</template>

<script setup>
const props = defineProps({
  iconStyle: String,
  text: String,
});
</script>
