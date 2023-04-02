<template>
  <div class="container">
    <button
      :class="isActive ? btnActiveClass : btnInactiveClass"
      :value="value"
      @click="updateBtnState"
      :disabled="false"
    >
      <i :class="`${iconStyle} text-4xl`"></i>
      <span>{{ text }}</span>
    </button>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useStore } from "vuex";

const store = useStore();
const props = defineProps({
  text: String,
  iconStyle: String,
  value: String,
});

const btnInactiveClass =
  "border-2 border-white rounded-md hover:text-black hover:bg-white w-28 h-28 inline-flex justify-center items-center flex-col space-y-2";
const btnActiveClass =
  "border-2 border-none text-black rounded-md bg-white w-28 h-28 inline-flex justify-center items-center flex-col space-y-2";

const isActive = computed(() => {
  const currPlaygroundOptions = store.getters.getPlaygroundOptions;
  if (props.value in currPlaygroundOptions) {
    return currPlaygroundOptions[props.value].isBtnPressed;
  }
});

const isDisabled = computed(() => {
  const currPlaygroundOptions = store.getters.getPlaygroundOptions;
  if (props.value in currPlaygroundOptions) {
    return !currPlaygroundOptions[props.value].isBtnPressed;
  }
  return;
});

const updateBtnState = () => {
  let currPlaygroundOptions = store.getters.getPlaygroundOptions;

  for (const key in currPlaygroundOptions) {
    currPlaygroundOptions[key].isBtnPressed = key === props.value;
  }
  store.commit("setPlaygroundOptions", currPlaygroundOptions);
};
</script>
