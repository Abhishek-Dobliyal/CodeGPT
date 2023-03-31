<template>
  <div class="container mt-10 px-4">
    <nav class="bg-black rounded-lg my-3">
      <div class="max-w-screen-xl px-4 py-3 mx-auto md:px-6">
        <div class="flex items-center">
          <ul class="flex flex-row mt-0 mr-6 space-x-8 text-sm font-medium">
            <SelectField
              disabledOption="language"
              :options="['python', 'javascript', 'go']"
            ></SelectField>
            <SelectField
              disabledOption="theme"
              :options="['dracula', 'monokai', 'material']"
            ></SelectField>
          </ul>
        </div>
      </div>
    </nav>
    <textarea ref="textarea"> </textarea>
  </div>
</template>

<script setup>
import SelectField from "@/components/SelectField.vue";
import { useStore } from "vuex";
import { ref, onMounted, watch } from "vue";

import * as CodeMirror from "codemirror";
import "codemirror/lib/codemirror.css";

import "codemirror/theme/blackboard.css";
import "codemirror/theme/dracula.css";
import "codemirror/theme/monokai.css";
import "codemirror/theme/material.css";

import "codemirror/mode/python/python.js";
import "codemirror/mode/javascript/javascript.js";
import "codemirror/mode/go/go.js";
import "codemirror/addon/display/autorefresh.js";
import "codemirror/addon/edit/matchbrackets.js";
import "codemirror/addon/edit/closebrackets";
import "codemirror/addon/hint/show-hint";
import "codemirror/addon/hint/show-hint.css";
import "codemirror/addon/hint/anyword-hint";

const store = useStore();
const textarea = ref(null);

onMounted(() => {
  const editor = CodeMirror.fromTextArea(textarea.value, {
    lineNumbers: true,
    theme: "blackboard",
    mode: "go",
    lineWrapping: true,
    autoRefresh: true,
    matchBrackets: true,
    autoCloseBrackets: true,
    extraKeys: {
      "Shift-Space": "autocomplete",
    },
  });
  store.commit("setEditorElement", editor);
  const wrapperEle = editor.getWrapperElement();
  wrapperEle.classList.add("p-2");
  wrapperEle.classList.add("rounded-lg");
});

const log = () => {
  console.log(store.state.editor);
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
</style>
