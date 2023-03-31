import { createStore } from "vuex";

export default createStore({
  state: {
    editor: null,
    editorContent: {
      text: "",
      language: "",
      theme: "",
    },
  },
  getters: {
    getEditorContent(state) {
      return state.editorContent;
    },
  },
  mutations: {
    setEditorElement(state, payload) {
      state.editor = payload;
    },
    setEditorOtherProperty(state, payload) {
      state.editorContent[payload.property] = payload.val;
    },
    setEditorTextProperty(state, payload) {
      state.editorContent.text = payload;
    },
  },
  actions: {},
  modules: {},
});
