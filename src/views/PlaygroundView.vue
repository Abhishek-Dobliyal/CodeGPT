<template>
  <div class="grid sm:grid-cols-2 grid-cols-1 h-screen">
    <div class="container relative mb-12">
      <div
        class="flex justify-start mt-4 ml-4 text-xl align-center space-x-2 items-center cursor-pointer"
        @click="$router.push('/')"
      >
        <i class="fa-solid fa-code fa-bounce"></i>
        <span class="font-semibold">CodeGPT</span>
      </div>
      <CodeArea></CodeArea>
      <div class="container absolute inset-y-0 right-0 w-16 mt-1">
        <Tooltip tip="Use ⇧ + Space for auto-hints."></Tooltip>
      </div>
      <div class="container text-center" v-if="showOptions">
        <RefactorSetting></RefactorSetting>
        <button
          @click="null"
          class="bg-black border hover:bg-white hover:text-black text-white font-semibold px-3 py-2 rounded-lg mt-3 hover:border-black hover:border-2"
        >
          Proceed
        </button>
      </div>
    </div>
    <div class="container text-white bg-black text-center mr-0">
      <div class="container text-center mt-10 mb-4">
        <span class="sm:text-lg text-md"
          >Select a card & then proceed to the Editor section</span
        >
      </div>
      <div
        class="grid grid-rows-2 grid-cols-3 gap-y-12 text-center items-baseline mt-10 mb-5 sm:mt-20"
      >
        <PlaygroundButton
          text="Refactor"
          iconStyle="fa-solid fa-code-compare"
          value="refactor"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Prettify"
          iconStyle="fa-solid fa-wand-magic-sparkles"
          value="prettify"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Convert"
          iconStyle="fa-solid fa-arrows-rotate"
          value="convert"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Comment"
          iconStyle="fa-solid fa-comment"
          value="comment"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Coming Soon"
          iconStyle="fa-solid fa-question"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Coming Soon"
          iconStyle="fa-solid fa-question"
        ></PlaygroundButton>
      </div>
      <div
        class="container mt-10 px-4 mb-3 flex justify-center items-center space-x-2 font-semibold sm:mt-14"
      >
        <span>Want to contribute or suggest more features? </span>
        <a href="https://github.com/Abhishek-Dobliyal">
          <i
            class="fa-brands fa-github text-2xl hover:shadow-lg hover:shadow-orange-400"
          ></i>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import CodeArea from "@/components/CodeArea.vue";
import Tooltip from "@/components/Tooltip.vue";
import PlaygroundButton from "@/components/PlaygroundButton.vue";
import RefactorSetting from "@/components/RefactorSetting.vue";
import { computed } from "vue";
import { useStore } from "vuex";

const store = useStore();
const showOptions = computed(() => {
  const currPlaygroundOptions = store.getters.getPlaygroundOptions;
  for (const key in currPlaygroundOptions) {
    if (currPlaygroundOptions[key].isBtnPressed) {
      return true;
    }
  }
  return false;
});
</script>

<style></style>
