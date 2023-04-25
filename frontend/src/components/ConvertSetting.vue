<template>
  <div class="container text-left animate__animated animate__fadeIn">
    <div class="flex items-center justify-start mt-4 space-x-3 ml-4">
      <span class="text-md font-semibold">Convert To</span>
      <SelectFieldPlayground
        :options="['Python', 'C++', 'Go', 'Java', 'Javascript']"
        disabledOption="language"
        setting="convert-to"
      ></SelectFieldPlayground>
    </div>
    <Input
      setting="convert-start-end"
      placeholderText="Enter start-end line number"
      lowerInfoText="[start-end] line number"
      :computedFunc="highlightText"
    ></Input>
  </div>
</template>

<script setup>
import SelectFieldPlayground from "@/components/SelectFieldPlayground.vue";
import Input from "@/components/Input.vue";
import { useStore } from "vuex";

const store = useStore();
const editor = store.getters.getEditorInstance;
const numLines = editor.lineCount();
const editorContent = store.getters.getEditorContent;

const highlightText = (range) => {
  let digits = 0;
  for (const ch of range.split("-")) {
    if (isNum(ch)) digits++;
  }
  if (digits == 2) {
    let [start, end] = range.split("-").map((ele) => {
      return parseInt(ele);
    });
    console.log(start, end, !start, !end);
    if (!start || start >= end || start > numLines || !end) {
      clearHighlight();
      return true;
    }
    const marker = editor.markText(
      { line: start - 2 },
      { line: end - 1 },
      {
        className:
          editor.getOption("theme") === "ayu-dark"
            ? "highlight-dark"
            : "highlight-light",
      }
    );
    let lines = marker.lines;
    console.log(lines);
    for (let i = start - 1; i <= end - 1; i++) {
      if (lines[i] && "text" in lines[i])
        editorContent.highlightedText += lines[i].text + "\n";
    }
    console.log(store.state);
    return false;
  }
  clearHighlight();
  editorContent.highlightedText = "";
  return true;
};

const isNum = (n) => {
  return !isNaN(parseInt(n)) && isFinite(n);
};

const clearHighlight = () => {
  editor.doc.getAllMarks().forEach((marker) => marker.clear());
};
</script>
