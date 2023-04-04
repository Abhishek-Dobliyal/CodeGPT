<template>
  <div class="container mt-10 px-4">
    <nav class="bg-black rounded-lg my-3">
      <div class="max-w-screen-xl px-4 py-3 mx-auto md:px-6">
        <div class="flex items-center">
          <ul class="flex flex-row mt-0 mr-6 space-x-8 text-sm font-medium">
            <SelectFieldEditor
              disabledOption="language"
              :options="['python', 'javascript', 'go', 'c++', 'java']"
              @change="changeLanguage"
            ></SelectFieldEditor>
            <button class="text-white text-xl ml-24" @click="toggleTheme">
              <i class="fa-solid fa-moon" v-if="!isDark"></i>
              <i class="fa-solid fa-sun" v-else></i>
            </button>
          </ul>
        </div>
      </div>
    </nav>
    <textarea ref="textarea"> </textarea>
  </div>
</template>

<script setup>
import SelectFieldEditor from "@/components/SelectFieldEditor.vue";
import { useStore } from "vuex";
import { ref, onMounted, reactive, computed } from "vue";

import * as CodeMirror from "codemirror";
import "codemirror/lib/codemirror.css";

import "codemirror/theme/ayu-dark.css";
import "codemirror/theme/base16-light.css";

import "codemirror/mode/python/python.js";
import "codemirror/mode/javascript/javascript.js";
import "codemirror/mode/go/go.js";
import "codemirror/mode/clike/clike.js";
import "codemirror/addon/display/autorefresh.js";
import "codemirror/addon/edit/matchbrackets.js";
import "codemirror/addon/edit/closebrackets.js";
import "codemirror/addon/selection/mark-selection.js";
import "codemirror/addon/hint/show-hint.js";
import "codemirror/addon/hint/anyword-hint.js";

const store = useStore();
const textarea = ref(null);
const isDark = ref(true);
let editor = reactive({});

onMounted(() => {
  editor = CodeMirror.fromTextArea(textarea.value, {
    lineNumbers: true,
    theme: "ayu-dark",
    mode: "",
    lineWrapping: true,
    autoRefresh: true,
    matchBrackets: true,
    autoCloseBrackets: true,
    extraKeys: {
      "Shift-Space": "autocomplete",
    },
    styleSelectedText: true,
  });

  const wrapperEle = editor.getWrapperElement();
  wrapperEle.classList.add("p-2");
  wrapperEle.classList.add("rounded-lg");
  wrapperEle.classList.add("border-2");
  wrapperEle.classList.add("border-black");

  editor.on("change", () => {
    store.getters.getEditorContent.text = editor.getValue();
  });

  store.commit("setEditorInstance", editor);
});

const toggleTheme = () => {
  isDark.value = !isDark.value;
  editor.doc.getAllMarks().forEach((m) => {
    m.className = isDark.value ? "highlight-dark" : "highlight-light";
  });
  editor.setOption("theme", isDark.value ? "ayu-dark" : "base16-light");
};

const changeLanguage = () => {
  let currLanguage = store.getters.getEditorContent.language;
  if (currLanguage === "c++") {
    currLanguage = "text/x-c++src";
  } else if (currLanguage === "java") {
    currLanguage = "text/x-java";
  }
  editor.setOption("mode", currLanguage);
};
</script>

<style>
.CodeMirror {
  font-family: "Inconsolata", monospace;
}
.CodeMirror-editor {
  height: 100%;
}
.CodeMirror-scroller {
  overflow: auto;
}

select {
  border-right: 10px solid black;
  outline: none;
}

.highlight-light {
  background-color: lightblue;
}
.highlight-dark {
  background-color: grey;
}
</style>
