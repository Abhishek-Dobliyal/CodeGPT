<template>
  <div class="container my-4">
    <ToggleButton
      label="Refactor Into Multiple Function(s)"
      value="refactor-splitCode"
    ></ToggleButton>
    <ToggleButton label="Add Tests" value="refactor-addTests"></ToggleButton>
  </div>
</template>

<script setup>
import ToggleButton from "@/components/ToggleButton.vue";
</script>
