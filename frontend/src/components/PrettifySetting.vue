<template>
  <div class="container my-4 animate__animated animate__fadeIn">
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
import { useStore } from "vuex";

const store = useStore();

const currPlaygroundOptions = store.getters.getPlaygroundOptions;

const checkIndentSize = (num) => {
  if (!isNum(num)) {
    return true;
  }
  currPlaygroundOptions.prettify.indentSpace =
    parseInt(num) <= 4 && parseInt(num) >= 1 ? parseInt(num) : 2;
  console.log(store.getters.getPlaygroundOptions);
  return num % 2 !== 0;
};

const isNum = (n) => {
  return !isNaN(parseInt(n)) && isFinite(n);
};
</script>
