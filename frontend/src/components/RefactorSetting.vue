<template>
  <div class="container my-4 animate__animated animate__fadeIn">
    <ToggleButton
      label="Refactor Into Multiple Function(s)"
      setting="refactor-splitCode"
    ></ToggleButton>
    <ToggleButton label="Add Tests" setting="refactor-addTests"></ToggleButton>
  </div>
</template>

<script setup>
import ToggleButton from "@/components/ToggleButton.vue";
</script>
