<template>
  <div class="grid sm:grid-cols-2 grid-cols-1 h-screen">
    <div class="container relative animate__animated animate__fadeInDown">
      <div
        class="flex justify-start mt-4 ml-4 text-xl align-center space-x-2 items-center"
      >
        <i class="fa-solid fa-code fa-bounce"></i>
        <span class="font-semibold">CodeGPT</span>
      </div>
      <div class="container text-center my-20 sm:my-40">
        <Header title="Code"></Header>
        <Headline :text="headline"></Headline>
      </div>
      <div class="flex justify-start absolute bottom-0 items-center">
        <span class="mx-3 mb-3 font-semibold"
          >Backed by <span class="text-teal-600">ChatGPT</span></span
        >
      </div>
    </div>
    <div
      class="text-center text-white bg-black animate__animated animate__fadeInUp"
    >
      <ul v-for="obj in infoButtonContent" :key="obj">
        <InfoText :text="obj.text" :iconStyle="obj.iconStyle"></InfoText>
      </ul>
      <button
        class="bg-black border-2 hover:bg-white hover:text-black text-white font-semibold px-4 py-2 rounded-lg mb-5 mt-3"
        @click="$router.push('/playground')"
      >
        Get Started
      </button>
    </div>
  </div>
</template>

<script setup>
import Header from "@/components/Header.vue";
import Headline from "@/components/Headline.vue";
import InfoText from "@/components/InfoText.vue";

const headline = "A utility to help organise and maintain your code.";
const infoButtonContent = {
  comment: {
    iconStyle: "fa-solid fa-comment",
    text: "Add comments to stay up-to-date with your code.",
  },
  refactor: {
    iconStyle: "fa-solid fa-code-compare",
    text: "Effortlessly refactor your code.",
  },
  prettify: {
    iconStyle: "fa-solid fa-wand-magic-sparkles",
    text: "Beautify those messy lines of code into structed ones.",
  },
  convert: {
    iconStyle: "fa-solid fa-arrows-rotate",
    text: "Convert your logic written in one language to another",
  },
  review: {
    iconStyle: "fa-solid fa-laptop-code",
    text: "Code reviews on the fly",
  },
};
</script>

<style scoped>
.left,
.right {
  height: 100vh;
  width: 100%;
}

@media screen and (min-width: 768px) {
  .left,
  .right {
    height: 100vh;
  }
}
</style>
