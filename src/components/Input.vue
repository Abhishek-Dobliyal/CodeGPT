<template>
  <div class="container ml-4 mt-3 space-x-4 w-1/2">
    <div
      class="justify-center h-8 items-start flex-col text-left space-x-1 mt-5 mb-7"
    >
      <input
        type="text"
        inputmode="numeric"
        id="error"
        class="bg-black text-white placeholder:text-white text-sm rounded-lg block p-2.5 outline-none placeholder:text-xs"
        :placeholder="placeholderText"
        v-model="num"
        @change="updateVal"
      />
      <p class="text-sm mt-1" v-if="showInfo">
        <span class="font-medium">{{ lowerInfoText }}</span>
      </p>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from "vue";
import { useStore } from "vuex";

const store = useStore();

const props = defineProps({
  setting: String,
  lowerInfoText: String,
  placeholderText: String,
});

const num = ref();

const showInfo = computed(() => {
  if (num.value < 0) {
    num.value = 2;
  }
  if (num.value > 6) {
    num.value = 6;
  }
  return num.value % 2 !== 0 || !(num.value > 0 && num.value <= 6);
});

const updateVal = () => {
  let [option, key] = props.setting.split("-");
  let currPlaygroundOptions = store.getters.getPlaygroundOptions;
  currPlaygroundOptions[option][key] = parseInt(num.value);
  console.log(store.getters.getPlaygroundOptions);
};
</script>

<style scoped></style>
