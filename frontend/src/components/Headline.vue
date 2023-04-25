<template>
  <div class="container mt-4">
    <p
      class="text-3xl font-normal mx-2"
    >
      {{ text }}
    </p>
  </div>
</template>

<script setup>
const props = defineProps({
  text: String,
});
</script>
