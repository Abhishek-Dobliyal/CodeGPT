<template>
  <div class="container">
    <button
      :class="fn ? btnActiveClass : btnInactiveClass"
      :value="value"
      @click="pressed"
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

const isDisabled = ref(false);
const props = defineProps({
  text: String,
  iconStyle: String,
  value: String,
});

const btnInactiveClass =
  "border-2 border-white rounded-md hover:text-black hover:bg-white w-28 h-28 inline-flex justify-center items-center flex-col space-y-2";
const btnActiveClass =
  "border-2 border-none text-black rounded-md bg-white w-28 h-28 inline-flex justify-center items-center flex-col space-y-2";
const isBtnPressed = ref(false);
const fn = computed(() => {
  return isBtnPressed.value;
});

const updateBtnState = () => {
  let currPlaygroundOptions = store.getters.getPlaygroundOptions;

  for (const key in currPlaygroundOptions) {
    if (key !== props.value) {
      currPlaygroundOptions[key].isBtnPressed = false;
      isBtnPressed.value = false;
      continue;
    }
    currPlaygroundOptions[key].isBtnPressed = true;
    isBtnPressed.value = true;
  }
  store.commit("setPlaygroundOptions", currPlaygroundOptions);
  console.log(store.getters.getPlaygroundOptions);
};

const pressed = () => {
  isBtnPressed.value = !isBtnPressed.value;
};
</script>
