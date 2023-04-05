<template>
  <div class="container font-semibold text-center py-2">
    <ul v-for="point in points" :key="point" class="list-disc list-inside">
      <li class="my-2">{{ point }}</li>
    </ul>
  </div>
</template>

<script setup>
const props = defineProps({
  points: Array,
});
</script>
