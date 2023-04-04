<template>
  <div class="w-full px-4 sm:w-1/2 lg:w-1/4 mt-3">
    <div class="mb-14">
      <div class="group relative inline-block">
        <i class="fa-sharp fa-solid fa-circle-info text-xl"></i>
        <div
          class="text-white absolute right-full top-1/2 z-20 mr-3 -translate-y-1/2 whitespace-nowrap rounded border bg-black py-[6px] px-4 text-sm font-semibold opacity-0 group-hover:opacity-100"
        >
          <span
            class="absolute -right-1 top-1/2 -z-10 h-2 w-2 -translate-y-1/2 rotate-45 rounded-sm bg-black font-thin"
          ></span>
          {{ tip }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  tip: String,
});
</script>
