<template>
  <div class="container my-4">
    <ToggleButton
      label="Follow best practices"
      setting="prettify-useBestPractice"
    ></ToggleButton>
    <Input
      setting="prettify-indentSpace"
      lowerInfoText="Possible size limit: 2, 4 or 6"
      placeholderText="Enter indent size"
    ></Input>
  </div>
</template>

<script setup>
import ToggleButton from "@/components/ToggleButton.vue";
import Input from "@/components/Input.vue";
</script>
