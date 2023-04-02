<template>
  <div class="my-5 flex justify-start items-center space-x-5 ml-4">
    <span class="text-md">Convert To</span>
    <SelectFieldPlayground
      :options="['Python', 'C++', 'Go', 'Java', 'Javascript']"
      disabledOption="language"
    ></SelectFieldPlayground>
  </div>
</template>

<script setup>
import SelectFieldPlayground from "@/components/SelectFieldPlayground.vue";
</script>
