<template>
  <div class="container text-left">
    <div class="flex items-center justify-start mt-4 space-x-3 ml-4">
      <span class="text-md font-semibold">Convert To</span>
      <SelectFieldPlayground
        :options="['Python', 'C++', 'Go', 'Java', 'Javascript']"
        disabledOption="language"
        setting="convert-to"
      ></SelectFieldPlayground>
    </div>
    <Input
      setting="convert-start-end"
      placeholderText="Enter start-end line number"
      lowerInfoText="[start-end] line number"
      :computedFunc="showInfo"
    ></Input>
  </div>
</template>

<script setup>
import SelectFieldPlayground from "@/components/SelectFieldPlayground.vue";
import Input from "@/components/Input.vue";
import { computed } from "vue";

const showInfo = (range) => {
  let hypen = 0,
    digits = 0;
  for (const ch of [...range]) {
    if (ch == "-") hypen++;
    if (isNum(ch)) digits++;
  }

  return hypen != 1 || digits != 2;
};

const isNum = (n) => {
  return !isNaN(parseInt(n)) && isFinite(n);
};
</script>
