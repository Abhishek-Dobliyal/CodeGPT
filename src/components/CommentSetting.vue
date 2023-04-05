<template>
  <div class="container my-4 animate__animated animate__fadeIn">
    <ToggleButton
      label="Comment on each line (if possible)"
      setting="comment-commentEveryLine"
    ></ToggleButton>
    <ToggleButton
      label="Concise comments"
      setting="comment-brief"
    ></ToggleButton>
  </div>
</template>

<script setup>
import ToggleButton from "@/components/ToggleButton.vue";
</script>
