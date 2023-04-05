<template>
  <div class="container font-semibold text-center py-2 mt-4">
    <ul v-for="p in points" :key="p" class="list-disc list-inside">
      <li class="my-3">
        {{ p.point }}
        <a
          class="underline cursor-pointer italic"
          :href="p.hyperlink"
          target="_blank"
        >
          {{ p.hyperlinkTxt }}
        </a>
      </li>
    </ul>
  </div>
</template>

<script setup>
const props = defineProps({
  points: Object,
});
</script>
