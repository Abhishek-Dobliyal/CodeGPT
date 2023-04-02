<template>
  <div class="flex justify-center items-center">
    <div class="container">
      <select
        class="px-2 py-1 bg-black text-white rounded-md capitalize"
        @change="setProperty"
        v-model="selectedValue"
      >
        <option value="" disabled>
          {{ props.disabledOption }}
        </option>
        <option v-for="opt in options" :key="opt" :value="opt">
          {{ opt }}
        </option>
      </select>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useStore } from "vuex";

const store = useStore();

const selectedValue = ref("");

const props = defineProps({
  onChangeMethod: String,
  for: String,
  value: String,
  options: Array,
  disabledOption: String,
});

const setProperty = () => {
  let currPlaygroundOptions = store.getters.getPlaygroundOptions;
  for (const key in currPlaygroundOptions) {
    if (key === props.for) {
      currPlaygroundOptions[key][val] = selectedValue.value;
      break;
    }
  }
  store.commit("setPlaygroundOptions", currPlaygroundOptions);
};
</script>
