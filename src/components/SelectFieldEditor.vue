<template>
  <li>
    <div class="flex justify-center">
      <div class="container">
        <select
          class="px-2 py-1 bg-black text-white rounded-md capitalize"
          @change="setEditorProperty"
          v-model="selectedValue"
        >
          <option value="" disabled>{{ disabledOption }}</option>
          <option v-for="opt in options" :key="opt" :value="opt">
            {{ opt }}
          </option>
        </select>
      </div>
    </div>
  </li>
</template>

<script setup>
import { useStore } from "vuex";
import { ref } from "vue";

const store = useStore();
const props = defineProps({
  disabledOption: String,
  options: Array,
});

const selectedValue = ref("");
const setEditorProperty = () => {
  store.commit("setEditorAttrs", {
    property: props.disabledOption,
    val: selectedValue.value,
  });
};
</script>
