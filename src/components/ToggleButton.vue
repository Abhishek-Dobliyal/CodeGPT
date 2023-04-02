<template>
  <div
    class="flex items-center justify-start my-2 align-baseline py-2 ml-4 mr-0"
  >
    <label class="relative cursor-pointer">
      <input
        type="checkbox"
        :value="value"
        class="sr-only peer"
        @change="update"
      />
      <div
        class="w-11 h-6 bg-black rounded-full peer peer-checked:after:translate-x-full after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border after:rounded-full after:h-5 after:w-5 outline-none after:transition-allpeer-checked:bg-black border-2 border-black"
      ></div>
    </label>
    <span class="ml-3 text-md font-medium text-black">{{ label }}</span>
  </div>
</template>

<script setup>
import { useStore } from "vuex";

const store = useStore();

const props = defineProps({
  label: String,
  value: String,
});

const update = () => {
  let currPlaygroundOptions = store.getters.getPlaygroundOptions;
  let [option, toggle] = props.value.split("-");

  for (const key in currPlaygroundOptions) {
    if (key === option) {
      currPlaygroundOptions[key][toggle] = !currPlaygroundOptions[key][toggle];
      break;
    }
  }

  console.log(store.getters.getPlaygroundOptions[option]);
};
</script>
