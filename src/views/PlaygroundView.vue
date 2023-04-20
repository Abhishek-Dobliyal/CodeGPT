<template>
  <div class="grid sm:grid-cols-2 grid-cols-1 h-screen">
    <div class="container mb-4 animate__animated animate__fadeInDown">
      <div
        class="flex justify-start mt-4 ml-4 text-xl align-center space-x-2 items-center cursor-pointer"
        @click="$router.push('/')"
      >
        <i class="fa-solid fa-code fa-bounce"></i>
        <span class="font-semibold">CodeGPT</span>
      </div>
      <CodeArea></CodeArea>
      <Bullets :points="bulletPoints" v-if="!showOptions"></Bullets>
      <div class="container text-center" v-if="showOptions">
        <RefactorSetting
          v-if="$store.state.playgroundOptions.refactor.isBtnPressed"
        ></RefactorSetting>
        <ConvertSetting
          v-else-if="$store.state.playgroundOptions.convert.isBtnPressed"
        ></ConvertSetting>
        <PrettifySetting
          v-else-if="$store.state.playgroundOptions.prettify.isBtnPressed"
        ></PrettifySetting>
        <CommentSetting
          v-else-if="$store.state.playgroundOptions.comment.isBtnPressed"
        ></CommentSetting>
        <button
          @click="null"
          class="bg-black border text-white font-semibold px-3 py-2 rounded-lg mt-7 hover:bg-gray-800"
        >
          Proceed
        </button>
      </div>
    </div>
    <div
      class="container text-white bg-black text-center mr-0 animate__animated animate__fadeInUp"
    >
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
          text="Review"
          iconStyle="fa-solid fa-laptop-code"
        ></PlaygroundButton>
        <PlaygroundButton
          text="Coming Soon"
          iconStyle="fa-solid fa-question"
        ></PlaygroundButton>
      </div>
      <div
        class="container mt-10 px-4 mb-3 flex justify-center items-center space-x-2 font-semibold sm:mt-14"
      >
        <span class="text-medium"
          >Want to contribute or suggest more features?
        </span>
        <a href="https://github.com/Abhishek-Dobliyal">
          <i
            class="fa-brands fa-github text-2xl hover:shadow-xl hover:shadow-teal-400"
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
import ConvertSetting from "@/components/ConvertSetting.vue";
import PrettifySetting from "@/components/PrettifySetting.vue";
import CommentSetting from "@/components/CommentSetting.vue";
import Bullets from "@/components/Bullets.vue";
import { computed, ref } from "vue";
import { useStore } from "vuex";

const store = useStore();

const bulletPoints = ref([
  {
    point: "Use ⇧ + ⎵ to display hints.",
    hyperlink: "",
    hyperlinkTxt: "",
  },
  {
    point: "Sublime Text key-bindings are the default.",
    hyperlink: "",
    hyperlinkTxt: "",
  },
  {
    point: "View the ",
    hyperlink: "https://rowannicholls.github.io/sublime_text/key_bindings.html",
    hyperlinkTxt: "ST4 Bindings",
  },
]);

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
