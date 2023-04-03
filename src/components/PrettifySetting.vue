<template>
  <div class="container my-4">
    <ToggleButton
      label="Follow best practices"
      setting="prettify-useBestPractice"
    ></ToggleButton>
    <Input
      setting="prettify-indentSpace"
      lowerInfoText="Recommended size: 2 or 4"
      placeholderText="Enter indent size"
      :computedFunc="checkIndentSize"
    ></Input>
  </div>
</template>

<script setup>
import ToggleButton from "@/components/ToggleButton.vue";
import Input from "@/components/Input.vue";

const checkIndentSize = (num) => {
  return num % 2 !== 0;
};
</script>
